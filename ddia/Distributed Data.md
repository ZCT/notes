Replication：同一份数据存在不同的结点，如果其中某一个结点不可用了，服务还可用

Partitioning:  将一份大的数据分成子集，存储于不同的结点上

# ch5. Replication



* sync replication or async replication

* how to handle failed replicas

**handle node outages**

follower failure的时候，比较好处理(通过本地日志追上leader就可以)

leader failure需要考虑：

​	如果是async replication，可能要丢弃部分数据（因为最新的数据leader可能还没有同步给follower）

​	split brain(两个结点都宣称自己是leader)

​	检测leader是否dead的时间间隔（太长->恢复时间长，太短->不必要的failover)

**Implementations of Replication Logs**

* Statement-based replication(就是讲具体的每条命令同步过去）
* Write-ahead log replication(讲存储底层的日志同步过去，换一个新的存储的时候会有问题)
* Logic log replication(主要是将其与具体的存储解耦开来，最终的数据，比如insert，就是修改后的数据)



