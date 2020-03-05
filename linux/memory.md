## 内存

RSS(常驻物理内存）是怎么计算的？

rss不包括文件缓存(page/cache)

> The amount of memory that *doesn’t* correspond to anything on disk: stacks, heaps, and anonymous memoryThe amount of memory that doesn’t correspond to anything on disk: stacks, heaps, and anonymous memory maps.maps  
>
> [docker 内存监控指标详解](https://docs.docker.com/config/containers/runmetrics/)
>
> rss = active_anon+inactive_anon - tmpfs
>
> 程序本身的二进制文件内存也包含在rss中
>
> 文件的读写I/O操作占用的内存是被计算在page/cache中的
>
> 

> 一般来说，业务进程使用的内存主要有以下几种情况：
>
> （1）用户空间的匿名映射页（Anonymous pages in User Mode address spaces），比如调用malloc分配的内存，以及使用MAP_ANONYMOUS的mmap；当系统内存不够时，内核可以将这部分内存交换出去；
>
> （2）用户空间的文件映射页（Mapped pages in User Mode address spaces），包含map file和map tmpfs；前者比如指定文件的mmap，后者比如IPC共享内存；当系统内存不够时，内核可以回收这些页，但回收之前可能需要与文件同步数据；
>
> （3）文件缓存（page in page cache of disk file）；发生在程序通过普通的read/write读写文件时，当系统内存不够时，内核可以回收这些页，但回收之前可能需要与文件/磁盘同步数据；
>
> （4）buffer pages，属于page cache；比如读取块设备文件。
>
> 

**大家（一般情况下也可以）普遍认为，buffers和cached所占用的内存空间是可以在内存压力较大的时候被释放当做空闲空间用的。**

貌似线上docker部署的服务在oom计算内存的时候，是包括了cache/buffer的

> [三年前的这个issue可以佐证这一点](https://github.com/moby/moby/issues/21759)
>
> page/cache一般是为了提高io的性能，但是它什么时候刷回磁盘呢，linux内核中有如下参数（跟docker无关，系统级别的行为）：
>
> ```sh
> $ sysctl -a | grep dirty
> 
> vm.dirty_background_bytes = 0
> 
> vm.dirty_background_ratio = 5
> 
> vm.dirty_bytes = 0
> 
> vm.dirty_expire_centisecs = 3000
> 
> vm.dirty_ratio = 10
> 
> vm.dirty_writeback_centisecs = 500
> ```
>
>  

一般来说，分配内存的时候，只是分配了虚拟内存空间，真正申请内存还是要等到实际使用的时候。

active_rss 和 inactive_rss 其中 active和inactive的是swap的特性要用到的，一般来说inactive是可以被swap的。不过业务线上系统一般都会禁止swap

[linux memory子系统介绍](https://github.com/digoal/blog/blob/master/201701/20170111_02.md)







