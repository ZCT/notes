## Reliable 

* Hardware Faults
* SoftWare Errors
* Human Erros



## Scalability

* Describing Load

  twitter，两个主要的操作场景：发tweet（4.6k/s)，home timeline(300k/s)

  实现有两种方法：

  ​	1 发tweet的时候，插一条数据在数据库tweet表中，然后每个用户去查

  ​	2 发tweet的时候，将数据插在每个关注用户的timeline cache中

* Describing Performance

  P99,p999. For example, if the 95th percentile response time is 1.5 seconds, that means 95 out of 100 requests take less than 1.5 seconds, and 5 out of 100 requests take 1.5 seconds or more

* Approaches for Coping with Load

  scaling up(vertical scaling)

  scaling out(horizontal scaling)

An architecture that scales well for a particular application is built around assumptions of which operations will be common and which will be rare—the load parameters



## Maintainability

* Operability

  Make it easy for operations team to keep the system running smoothly

* Simplicity

  Make it easy for new engineers to understand  the system

* Evolvability

  Make it easy for engineers to make changes to the system in the future