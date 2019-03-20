

**basic raft algorithm**

* If two entries in different logs have the same index and term, then they store the same command.
* If two entries in different logs have the same index and term, then the logs are identical in all preceding entries

​	不准 Committing entries from preivious term，防止已经提交的日志被覆盖，只允许commit当前term的日志，如果当前term的日志已经commitd，自然而然的以前的日志也已经被commited了

​	选举时，log要up-to-date（最后一条日志的term是否更大，相同的情况下，日志索引是否大），这样可以保证选出来的leader，包含所有commited的日志，因为commited，代表日志已经在大多数结点上已经持久化了，如果日志没有包含已经commited过的日志，不会获得大部分的投票

[linearizability vs serializability](http://www.bailis.org/blog/linearizability-versus-serializability/)

lineraizability 是分布式系统中的概念（在某一个时间点写入数据，后面时间点的读操作应该读到这个更新之后的数据）

serializability 是数据库中的概念（就像串行执行操作一样）

如果有两个操作，T1是更新数据A，T2是读取数据A，T1先于T2执行，在数据库调度中，如果T1先于T2执行，则符合linearizability，（满足线性？）。如果T2先于T1执行，则算是searilizability，不符合lineraizability

对于linearizability的概念还是有点模糊，为什么这个例子里就不是符合linearizability

[strong consistency models](https://aphyr.com/posts/313-strong-consistency-models)

>  We call this consistency model *linearizability*; because although operations are concurrent, and take time, there is some place–or the appearance of a place–where every operation happens in a nice linear order.

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

> [etcd raft 如何实现linearizable Read](https://zhuanlan.zhihu.com/p/27869566)

**linearizability**

In linearizability,  each operation appears to execute instantaneously, exactly once, at some point between its invocation and its response

*avoid duplicate request*

基本思想就是server保存客户端的请求历史，用来去除client的重复请求。

给每个client分配一个id，client给每个命令分配一个serial number，server端记录一个session。为了防止并发的情况，记录client的所有请求，并且定期淘汰，客户端需要includes the lowest sequence number for which it has not yet received a response

