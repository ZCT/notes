##  ## 

go语言中的信号机制是什么样的

​	给一个程序中的一个线程发送消息，如何做到只要监听了的gorountine都可以收到消息

​	推测是runtime这里做了一层转发（还没有看到相关的代码和文章）