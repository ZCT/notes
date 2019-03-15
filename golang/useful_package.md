

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

  