## Preface

consensus protocol 

* paxos(multi-paxos, raft, zab)

* pow



golang raft implementation

* [etcd-raft](https://github.com/etcd-io/etcd)
* [consul-raft](https://github.com/hashicorp/raft)

##  etcd-raft

​	raft作为一个库嵌入在etcd中，通过提供接口来屏蔽网络、存储等模块实现，交给上层应用者来实现提供

**raft node status**

| 状态      | 描述                                             |
| --------- | ------------------------------------------------ |
| Candidate | 候选人状态，节点切换到这个状态时，意味着一次选举 |
| Follower  | 跟随者状态，节点切换到这个状态时，意味着选举结束 |
| Leader    | 领导者状态，所有数据提交都必须先提交到leader上   |

**伪代码过程：**

```go
// HTTP server
HttpServer主循环:
  接收用户提交的数据：
    如果是PUT请求：
      将数据写入到proposeC中
    如果是POST请求：
      将配置变更数据写入到confChangeC中
// raft Node
raftNode结构体主循环：
  如果proposeC中有数据写入：
    调用node.Propose向raft库提交数据
  如果confChangeC中有数据写入：
    调用node.Node.ProposeConfChange向raft库提交配置变更数据
  如果tick定时器到期：
    调用node.Tick函数进行raft库的定时操作
  如果node.Ready()函数返回的Ready结构体channel有数据变更：
    依次处理Ready结构体中各成员数据
    处理完毕之后调用node.Advance函数进行收尾处理
```

**Raft库主要结构体：**

| struct/interface   | 文件         | 作用                                   |
| ------------------ | ------------ | -------------------------------------- |
| Node(interface)    | node.go      | 提供raft库与应用交互的接口             |
| node               | node.go      | 实现了Node interface                   |
| Config             | raft.go      | raft相关配置参数                       |
| raft               | raft.go      | raft 算法的实现                        |
| ReadState/readOnly | read_only.go | linearizability read                   |
| rafgLog            | log.go       | raft日志操作                           |
| Progress           | progress.go  | 用于leader中保存每个follower的状态信息 |
| Storage(interface) | stoage.go    | 持久化存储接口                         |
| MemoryStorage      | storage.go   | raft中状态转换要用到的storage信息      |

**Ready结构体：**

| Field            | Type         | 作用                                                         |
| ---------------- | ------------ | ------------------------------------------------------------ |
| SoftState        | SoftState    | 经常变动，且不需要持久化中的数据，包括：集群leader、结点当前的状态 |
| HardState        | HardState    | 需要写入持久化存储中，包括节点当前term、Vote、Commit(跟日志同一个存储) |
| ReadStates       | []ReasStates | 读一致性的数据                                               |
| Entries          | []pb.Entry   | 在向其他node发送消息之前需要先写入持久化存储的数据(一般是leader结点有这个数据) |
| Snapshot         | pb.Snapshot  | 需要写入持久化存储中的快照数据                               |
| CommittedEntries | []pb.Entry   | 需要输入到状态机的数据，之前已经保存到持久化存储中了         |
| Messages         | []pb.Message | 在entries被写入持久化存储中以后，需要发送出去的数据(leader和follower都有) |

**RaftLog结构体**：

```go
type raftLog struct {
	// storage contains all stable entries since the last snapshot.
	storage Storage  //用于保存最后一次snapshot之后提交的数据，已经持久化了

	// unstable contains all unstable entries and snapshot.
	// they will be saved into storage.
	unstable unstable  //用于保存还没有持久化的数据和快照，这些数据最终都会落到storage中

	// committed is the highest log position that is known to be in
	// stable storage on a quorum of nodes.
	committed uint64  /
	// applied is the highest log position that the application has
	// been instructed to apply to its state machine.
	// Invariant: applied <= committed
	applied uint64

	logger Logger
}
```

**消息类型：**

| MessageType                   | 描述                                                         |
| ----------------------------- | ------------------------------------------------------------ |
| MsgHup                        | 用于触发选举（tickElection)，通常只在本节点                  |
| MsgBeat                       | 内部消息，触发leader节点定期向follower节点发送心跳的消息     |
| MsgProp                       | 业务层propose数据(candidate、follower、leader处理逻辑各不相同) |
| MsgApp                        | 用于leader向集群中其他节点同步数据的消息                     |
| MsgSnap                       | 用于leader向集群中其他节点同步数据的消息                     |
| MsgAppResp                    | Follower节点针对leader节点的MsgApp消息的应答消息             |
| MsgVote/MsgVoteResp           | 选举的时候给其他节点发送的投票请求消息/应答                  |
| MspPreVote/MsgPreVoteResp     | 预先投票消息，防止disruptive service                         |
| MsgHeartbeat/MsgHeartbeatResp | leader向follower发送心跳的消息/应答(探活、commit成员、线性一致性读) |
| MsgUnreachable                | 业务层向raft库汇报某个节点当前已不可达(仅leader才处理这类消息k) |
| MsgSnapStatus                 | 用于向raft库汇报某个节点当前接受快照的状态                   |
| MsgIndex/MsgReadIndexResp     | linearizability read 相关消息                                |

**Leader view progress**



![image-20190319205615000](https://raw.githubusercontent.com/zct/notes/master/consensus/progress.png)



### Optimize

**如何避免分区的时候 disruptive serice**

考虑到一种情况：当出现网络分区的时候，A、B、C、D、E五个节点被划分成了两个网络分区，A、B、C组成的分区和D、E组成的分区，其中的D节点，如果在选举超时到来时，都没有收到来自leader节点A的消息（因为网络已经分区），那么D节点认为需要开始一次新的选举了。

正常的情况下，节点D应该把自己的任期号term递增1，然后发起一次新的选举。由于网络分区的存在，节点D肯定不会获得超过半数以上的的投票，因为A、B、C三个节点组成的分区不会收到它的消息，这会导致节点D不停的由于选举超时而开始一次新的选举，而每次选举又会递增任期号。

在网络分区还没恢复的情况下，这样做问题不大。但是当网络分区恢复时，由于节点D的任期号大于当前leader节点的任期号，这会导致集群进行一次新的选举，即使节点D肯定不会获得选举成功的情况下（因为节点D的日志落后当前集群太多，不能赢得选举成功）。

为了避免这种无意义的选举流程，节点可以有一种PreVote的状态，在这种状态下，想要参与选举的节点会首先连接集群的其他节点，只有在超过半数以上的节点连接成功时，才能真正发起一次新的选举

增加Prevote阶段，term不加一，在得到多数节点的响应后，才能发起正式投票请求

**如何保证linearizability read**

由于所有的leader和follower都能处理客户端的读请求，所以存在可能造成返回读出的旧数据的情况：

- leader和follower之间存在状态差，因为follower总是由leader同步过去的，可能会返回同步之前的数据。
- 如果发生了网络分区，某个leader实际上已经被隔离出了集群之外，但是该leader并不知道，如果还继续响应客户端的读请求，也可能会返回旧的数据。

因此，在接收到客户端的读请求时，需要保证返回的数据都是当前最新的

*Leader*

1. If the leader has not yet marked an entry from its current term committed, it waits until it as done so. Raft handles this by having each leader commit a blank *no-op* entry into the log at the start of its term
2. leader save it current commit index in a local variable readIndex
3. issues a new round of heartbeats and waits for their acknowledgments from a majority of the cluster. Once these acknowledgments are received, the leader knows that there could not have existed a leader for a greater term at the moment it sent the heartbeats. Thus, the readIndex was, at the time, the largest commit index ever seen by any server in the cluster.
4. The Leader waits for its statemachine to advance as far as the readIndex. this is currnet enought to satisfy linearizbility
5. 处理读请求，返回给客户端

*Follower*

Follower向leader询问当前readIndex，leader在返回readIndex给follower的时候，需要执行1-3步（因为leader需要check自己是否还是leader）。follower然后再执行4-5步

*具体实现：*

```
type ReadState struct {  //用于保存读请求到来时的节点状态
  Index uint64  			//接收到改读请求时，当前节点的commit索引
  RequestCtx []byte   //客户端读请求的唯一标识
}
```

1. server收到客户端的读请求，此时会调用raft.ReadIndex函数发起一个MsgReadIndex的请求，带上的参数是客户端读请求的唯一标识（此时可以对照前面分析的MsgReadIndex及其对应应答消息的格式）。

2. follower将向leader直接转发MsgReadIndex消息，而leader收到不论是本节点还是由其他server发来的MsgReadIndex消息，其处理都是：

   a. 首先如果该leader在成为新的leader之后没有提交过任何值，那么会直接返回不做处理。

   b. 调用r.readOnly.addRequest(r.raftLog.committed, m)保存该MsgreadIndex请求到来时的commit索引。

   c. r.bcastHeartbeatWithCtx(m.Entries[0].Data)，向集群中所有其他节点广播一个心跳消息MsgHeartbeat，并且在其中带上该读请求的唯一标识。

   d. follower在收到leader发送过来的MsgHeartbeat，将应答MsgHeartbeatResp消息，并且如果MsgHeartbeat消息中有ctx数据，MsgHeartbeatResp消息将原样返回这个ctx数据。

   e. leader在接收到MsgHeartbeatResp消息后，如果其中有ctx字段，说明该MsgHeartbeatResp消息对应的MsgHeartbeat消息，是收到ReadIndex时leader消息为了确认自己还是集群leader发送的心跳消息。首先会调用r.readOnly.recvAck(m)函数，根据消息中的ctx字段，到全局的pendingReadIndex中查找是否有保存该ctx的带处理的readIndex请求，如果有就在acks map中记录下该follower已经进行了应答。

   f. 当ack数量超过了集群半数时，意味着该leader仍然还是集群的leader，此时调用r.readOnly.advance(m)函数，将该readIndex之前的所有readI

   ndex请求都认为是已经成功进行确认的了，所有成功确认的readIndex请求，将会加入到readStates数组中，同时leader也会向follower发送MsgReadIndexResp。

   g. follower收到MsgReadIndexResp消息时，同样也会更新自己的readStates数组信息。

   h. readStates数组的信息，将做为ready结构体的信息更新给上层的raft协议库的使用者（应用层等待apply日志大于commitIndex之后，才返回给client读响应）。



### References

[raft-thesis paper](https://ramcloud.stanford.edu/~ongaro/thesis.pdf)

