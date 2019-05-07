

**golang.org/x/sync** 

| Singleflight | 重复的函数调用只会被执行一次，常用于防止缓存击穿 |
| ------------ | ------------------------------------------------ |
| errgroup     | 把子任务聚合在一起，如果一个失败，整体就失败     |



**cache**

* bigcache

  nogc 是因为预先分配了内存，只在map里存了int型的数据，value是指向数据块的offset

  特性：过期数据周期删除、数据分片、zero gc

  底层使用的是byteQueue，一种循环队列

  存在的问题：如果在短时期内重复设置同一个key，并且在初始化的时候没有设置hardmaxlimitedsize，会导致内存无限增长


[concurrency-limits](https://github.com/Netflix/concurrency-limits)

​	自适应的限流库

> ```java
> newLimit = currentLimit × gradient + queueSize
> gradient=(RTTnoload/RTTactual)
> ```
>
> 但是我没有弄明白怎么用最小的延迟去做分母，这样用户会<=1，而且大多数情况会<1，所以窗口不会一直增长）？



[gofail](https://github.com/etcd-io/gofail)

[failpoint](https://github.com/pingcap/failpoint)

​	在代码中模拟错误发生，而不干扰原来的逻辑

​	 [gofail的使用 ](https://github.com/lkk2003rty/notes/blob/master/gofail.md)

​		gofail的简单来说就是定义一个injection点，return这个语义具有欺骗性，其功能就是简单把值付给变量，然后执行后面的代码，注释中可以是任何代码。