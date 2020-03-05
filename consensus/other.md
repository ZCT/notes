分布式系统对 fault tolerance 的一般解决方案是 state machine replication



分布式系统常见的做法：

​	主从全同步

​	多数派



paxos

![image-20190920004615896](/Users/tangzhongcheng/Library/Application Support/typora-user-images/image-20190920004615896.png)

​	client

​	proposer

​	acceptor

​	learner





选leader很重要的一点是保证有序？



Mulit paxos

 leader 唯一的proposer