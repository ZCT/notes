**[hashicorp raft](https://github.com/hashicorp/raft)**

​	两层 logStore FSM

​	日志处于commit 之后，才apply应用到FSM，client应该是在apply之后才返回

​	如果想要避免stales read，所有的读请求应该发送到leader(不支持follower这样，etcd可以），leader也要验证是否当前还是leader





**[etcd raft](https://github.com/etcd-io/etcd/tree/master/raft)**

Ready结构体是raft模块跟上层打交道的结构体数据

snapshot是针对业务storage的snapshot(?还是日志的snapshot）? snapshot是业务store的快照，然后再加上一些meta data。 防止log无限增长，快照怎么生成的

snapshot是业务和日志索引的存储

storage 提供的是日志持久化存储的interface

raftlog˙中的unstable用来保存还没持久化的数据

memoryStorage 是保存快照和一些raft要用到的状态信息的（实现了raft中的storage interface)，业务的数据需要自己制定存储来实现

store是用来业务存储数据的



触发快照操作，是在进行快照的时候，将数据进行压缩处理()



结点有两个storage，raftstorage是raft模块下的memoryStorage(用于raft在现场处理)

​				   storage是业务自定义的storage（存储快照和日志信息），这两者需要联动(快照处理的时候)

​				 commit过的日志存储在storage中

注意term和log term的区别

etcd raft example:

​	log 已经commited，在将log apply到状态机时，失败怎么办? 看example代码是通过commitC异步处理的



**Inflights**

这个是用来控制消息数量的，原理就是一个循环数组，最大消息个数不能超过数组的size，对输入元素的要求是index是递增的



references:

> [raft在etcd中的实现](http://blog.betacat.io/post/raft-implementation-in-etcd/)

>  [raft代码解析详细](https://www.codedump.info/post/20180922-etcd-raft/#msgvote-msgprevote%E6%B6%88%E6%81%AF%E4%BB%A5%E5%8F%8Amsgvoteresp-msgprevoteresp%E6%B6%88%E6%81%AF)

match is the index of the highest known matched entry

Next is the index of the first entry that will be replicated to the follower.  Leader puts entries from next to its latest one in next replication message.

node.go 应用层和raft共识算法的衔接