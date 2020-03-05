## gc





go version 1.12: non-generation, concurrent, tri-color, mark and sweep collector



要理解gc中的细节，还得需要理解gc内存分配的细节



**Marking**

* Marking Setup ——STW

  write barrier turn on

​	every application goroutine running must by stopped

* Marking

  和业务代码并行，但是会占用一些cpu资源

   Inspect the stacks to find root pointers to the heap

   Traverse the heap graph from these root pointers

   Mark values on the heap that are still in-use

  如果在marking阶段，有其他业务协程序分配cpu过快的话，会开启Mark Assist降低内存的分配速度

* Mark Termination ——STW

  write barrier turn off

**Sweeping**

​	Occurs when Groutines attempt to allcoate new heap memeory

​	The lantency of Sweeping is added to the cost of performing an allocaiotn, not GC



 **GC Percentage**

​	GC Percentage is 100 by default

​	Presents a ratio of how much new heap memory can be allocated before the next collection has to start

​	比较有意思(待确认)的一个点是gc可以提前开始（在内存未达到GC Percentage比例时），collector通过pacing algorithm决定gc什么时候开始





> ```
> GODEBUG=gctrace=1 ./app
> ```

```go
gc 1405 @6.068s 11%: 0.058+1.2+0.083 ms clock, 0.70+2.5/1.5/0+0.99 ms cpu, 7->11->6 MB, 10 MB goal, 12 P

// General
gc 1404     : The 1404 GC run since the program started
@6.068s     : Six seconds since the program started
11%         : Eleven percent of the available CPU so far has been spent in GC

// Wall-Clock
0.058ms     : STW        : Mark Start       - Write Barrier on
1.2ms       : Concurrent : Marking
0.083ms     : STW        : Mark Termination - Write Barrier off and clean up

// CPU Time
0.70ms      : STW        : Mark Start
2.5ms       : Concurrent : Mark - Assist Time (GC performed in line with allocation)
1.5ms       : Concurrent : Mark - Background GC time
0ms         : Concurrent : Mark - Idle GC time
0.99ms      : STW        : Mark Term

// Memory
7MB         : Heap memory in-use before the Marking started
11MB        : Heap memory in-use after the Marking finished
6MB         : Heap memory marked as live after the Marking finished
10MB        : Collection goal for heap memory in-use after Marking finished
```



参考资料：

[Garbage Collection Semantics视频](<https://www.youtube.com/watch?v=q4HoWwdZUHs>)

[garbage collector](<https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html>)





# #

[gc详细](https://blog.golang.org/ismmkeynote)



如果一个slice里面的key是*int，gc的代价会高

key是int gc的代价会低

具体原因是啥？ 如果是指针，gc的时候需要全部扫描一遍，决定哪些是可以回收的

> freecache gc overhead
>
> FreeCache avoids GC overhead by reducing the number of pointers. No matter how many entries stored in it, there are only 512 pointers. The data set is sharded into 256 segments by the hash value of the key. Each segment has only two pointers, one is the ring buffer that stores keys and values, the other one is the index slice which used to lookup for an entry. Each segment has its own lock, so it supports high concurrent access.
>
> 



[gc详细的资料](https://mp.weixin.qq.com/s/o2oMMh0PF5ZSoYD0XOBY2Q)



### Trip-color



为什么不直接用两色标记法：

[为什么不直接用两色标记法](https://malcolmyu.github.io/2019/07/07/Tri-Color-Marking/)

白色对象——未被回收器扫描到的对象。在回收开始阶段，所有对象都为白色，当回首结束后，白色对象均不可达

灰色对象——已被回收器扫描到的对象，但回收器需要对其中的一个或者多个指针进行扫描，因为他们可能还指向白色对象

黑色对象——已被回收器访问到的对象，其中所有字段都已被扫描，黑色对象中任何一个指针都不可能直接指向白色对象



垃圾回收可能会出现的问题：

	* 垃圾没有被回收
	* 有用的被回收了 （这个问题比较致命）



灰色对象是为了可以增量回收，可以保证gc的过程中，还可以修改引用关系

写屏障是为了保证黑色对象不再引用白色对象；写屏障相当于程序在写指针操作的地方插入一段代码，当用户程序修改内存引用时，如果有黑色对象引用了白色对象，则把相应白色对象移入灰色集合，如果是新分配的对象则直接放入黑色集合



开启写屏障的时候需要stw，关闭屏障的时候需要stw？ 为什么需要stw？

根节点是什么，为什么从这里开始扫描就可以垃圾回收t



