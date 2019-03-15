**[hashicorp raft](https://github.com/hashicorp/raft)**

​	两层 logStore FSM

​	日志处于commit 之后，才apply应用到FSM，client应该是在apply之后才返回

​	如果想要避免stales read，所有的读请求应该发送到leader



​	



