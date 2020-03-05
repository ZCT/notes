



serializability 和 linearizability 的区别：

​	serializability 的意思是事务看起来就像串行执行一样

​	linearizability 是指数据看起来就像只有一个副本一样

linearizblility 并不能阻止write skew现象，除非加一些额外的手段。在serializbility的实现方法中：two-phase locking and  actual Serial execution are typically linearizable，但是serializability snapshot isolation 不是 linearizability。由于linearizability 必然包含casual consistency，所以linearizability 可以达到snapshot isolation 级别

Total order boradcast

	* Reliable delivery
	* Totally ordered delivery

Linearizability 方法





Lamport  timestamps 这个地方没有看懂，怎么定义全序的。相同counter情况下，node节点越大的，timestamp也就越大

[lamport and vector clock](https://lrita.github.io/2018/10/24/lamport-logical-clocks-vector-lock/)

[lamport and vector clock-wudi](https://www.cnblogs.com/foxmailed/archive/2012/01/11/2319854.html)



jian

朋友圈 casual consistency

Lineraizability 和 total order broadcast区别



服务发现不需要强一致的理由