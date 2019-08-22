## gc



go version 1.12: non-generation, concurrent, tri-color, mark and sweep collector



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