传统的分布式事务处理框架——saga

​	XA uses two-phase commit (2PC)，缺点：

​		部分数据库不支持

​		同步调用，降低系统的可用性



Saga

​	Maintain data consistency across services using a sequence of local transactions that are coordinated using asynchronous messaging.



​	业界有两种实现方式

	* Choreography-based sagas（分布式的）



	* Orchestration-based sagas



为什么都用消息队列来做消息传递？而不是普通的rpc，好像有点道理

​	没有那个超时问题；接口要保证幂等性

