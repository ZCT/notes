

[linearizability vs serializability](http://www.bailis.org/blog/linearizability-versus-serializability/)

lineraizability 是分布式系统中的概念（在某一个时间点写入数据，后面时间点的读操作应该读到这个更新之后的数据）

serializability 是数据库中的概念（就像串行执行操作一样）

如果有两个操作，T1是更新数据A，T2是读取数据A，T1先于T2执行，在数据库调度中，如果T1先于T2执行，则符合linearizability，（满足线性？）。如果T2先于T1执行，则算是searilizability，不符合lineraizability



**Processing read-only queries**

保证无论在主还是从上读，都满足数据一致性（就是要读到最新已经commit的数据）

 * 和写一样，读也产生replication日志（性能不好）
 * bypassing the replication log

但是绕过replication log，会导致stale read，不符合linearizability，需要特殊的处理。比如，当前leader已经被分区了，剩下的集群已经选出了一个新的leader，并且提交了一些写的日志，如果旧的leader依旧可以处理读请求，则会返回stales read

*Leader*

1. If the leader has not yet marked an entry from its current term committed, it waits until it as done so. Raft handles this by having each leader commit a blank *no-op* entry into the log at the start of its term
2. leader save it current commit index in a local variable readIndex
3. issues a new round of heartbeats and waits for their acknowledgments from a majority of the cluster. Once these acknowledgments are received, the leader knows that there could not have existed a leader for a greater term at the moment it sent the heartbeats. Thus, the readIndex was, at the time, the largest commit index ever seen by any server in the cluster.
4. The Leader waits for its statemachine to advance as far as the readIndex. this is currnet enought to satisfy linearizbility
5. 处理读请求，返回给客户端

*Followers*

follower向leader询问当前readIndex，leader在返回readIndex给follower的时候，需要执行1-3步（因为leader需要check自己是否还是leader）。follower然后再执行4-5步



[etcd raft 如何实现linearizable Read](https://zhuanlan.zhihu.com/p/27869566)