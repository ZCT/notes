## Transactions

**Transactions**: Atomicity, Consistency, Isolation, Durability

## ACID

**Atomicity**

​	指的是当client发起几个写操作的时候，要么全部成功，要么全部都不成功，也许叫abortability更准确些。几个client同时操作时，一个client看到另一个client执行过程的中间状态，这种是在isolation中描述的。atomicity通过redo/undo log 来实现

**Consistency**

​	这个词在不同的地方有不同的含义。在数据库中，是指某些条件必须永远为真。In an accouting system, credits and debits across all accounts must always be balanced

**Isolation**

​	是指多个事务并行执行的时候，互不干涉。隔离有好几个级别，最高级别是serializable isolation

**Durability**

​	事务一旦提交，其中的数据就不会丢失。

**Single Object and multiple object operations **

​	Atomicity:use a log for cash recovery

​	Isolation:lock on each object

​	increment和CAS并不是事务，不能说做是轻量级事务

​	Multi Object transaction的存在是必要的，由于索引、外键这些因素

## weak Isolation Levels 

	### read committed 

主要用来解决dirty read 和dirty write

​	**定义：**

  * when reading from the database, you will only see data that has been commited

  * when writing to the database , you will only overwrite data that has been committed 

    解决办法**：

​	**no dirty read**

​		也可以通过加锁来做，但是这样实现性能会有问题。大部分数据库都会这样实现（MVCC的简化版）：

> for every object that is written , the database remembers both the old committed value and the new value set by the transaction that currently holds the write lock. While the transaction is ongoing, any other transactions that read the object are simply given the old value

​	**no dirty writes**

​		通过加锁来防止dirty write

### snapshot isolation

​	主要是破坏了事务中的consistency特性。假设A有500，B有500元，A在转账给B的场景。 有一个读事务(QueryA, QueryB)，和一个写事务(A-100, B+100)，如果这两个事务的执行顺序是这样的，QueryA->A-100->B+100->QueryB，这样虽然符合read committed 的标准，但是读事务看到的钱是900元，违反了consistency。虽然之后再执行一次读事务可以拿到正确的结果，但是在有些场景，比如备份和数据分析中，是不可容忍的。这种现象也叫做**nonrepeatable read and read skew**。**snapshot isolation **主要是用来解决这种问题

解决办法：MVCC

### prevent lost update

​	read commited 和 snapshot  isolation 都没办法解决lost update问题。lost update问题是由于read-modify-write cycle造成的，（比如对一个元素+1的操作，修改一个大json中的某个字段，多个人协同编辑文档，保存的时候把整个文档发送过去）。解决办法：

* atomic write operations（就是数据库提供一个接口，完成原来需要客户端read-modify-write三步完成的操作）

* explicit locking 

* Automatic detecting lost update (Mysql/InnoDB' repeatable read does not detect lost updates)

* Compare and set

  这里需要注意的是，locks and CAS 能够实现的假设条件是对象没有多个副本。在 replication 的场景下，这两种方法是不可行的

### serializable isolation 

​	之前说的dirty write 和 lost update 都是由于多个并行的客户端同时去写一个对象导致的，write skew 则是多个并行的客户端更细不同的对象导致的，可以说是lost update的泛化版。

​	write skew 的例子：

​	医院每天都会有一些值班的医生，值班的医生可以请假，但是至少要保证有一个医生，有个医生的数据表，医生有个状态代表自己是否值班。假设A和B两个医生今天值班，他们都想请假，A医生执行事务 select * from table where 值班=true，发现是2， 然后将自己的状态更改为不值班，B医生同时也执行相同的事务，他们在各自的事务中都符合条件，所以都请假了，但是这时医院就没有医生值班了。write skew 一般都是这样的模式：

1. 通过查询判断条件是否满足

2. 基于1的查询结果决定后面的写操作是否继续

3. 写操作的执行会影响1判断条件的结果

   解决办法：

   1. Actual Serial Execution 
   2. Two Phase Locking（关键点：如果一个transaction释放了它所持有的任意一个所，那它就不再不能后去任何锁，这个方法容易产生死锁，数据库必要要能够检测死锁，并且abort transaction），属于悲观锁的类型。这两的两阶段是指获取锁，释放锁的这两个阶段t
   3. Serializable Snapshot Isolation ：乐观锁的类型，怎么检测之前的查询结果条件被影响了，主要是这两种情形t：
      1. Detecting reads of a stale MVCC object version(uncommitted write occur before read)。在事务提交的时候，检查之前有没有ignore write 已经提交了，如果有的话，当前事务必须要abort
      2. Detecting writes that affect prior reads(the write occurs after read)

如何解决客户端超时问题，导致事情的二次发生——也就是exactly tonce语义