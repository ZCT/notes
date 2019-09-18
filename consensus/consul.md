## consul



* gossip 
* Serf  
* Vivaldi: A Decentralized Network Coordinate System



`gossip和raft协议怎么结合起来的`



consul服务发现





```go
docker network create --driver bridge --subnet 172.25.0.0/16 wordpress_net //创建一个名字为wordpress_net的network

docker network inspect wordpress_net //查看名字为wordpress_net的网络信息 

docker network ls  //查看当前有哪些network

docker rm -f  ${container_id}  //强制删除指定container_id

docker ps   //查看

docker exec -t node1 consul members  //docker上执行民命令
```

