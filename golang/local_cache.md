local_cache



[dgraph ristretto cache](https://github.com/dgraph-io/ristretto)



没有保存key，可能会有冲突，key的哈希碰撞（目前正在解决）

最近的set可能会丢弃

命中率比较高

Count-min sketch（bloom filter在统计方面的变形），基本原理就是对一个key做多个哈希，count分别加1。取元素的时候，从这多个哈希位置取最小的值

reset操作。每访问一次，有一个计数器会加一，直到计数器达到预设的值，所有的count都会减半

驱除时的原则，如果替换一个元素能够增加整体访问命中率就替换

[设计一个现代的缓存](http://biaobiaoqi.github.io/blog/2017/03/19/design-of-a-modern-cache/)

LFU和LRU的比较：	

>  总体来说LFU命中率较高，但是需要记录历史信息。LRU命中率会低一些，需要存储较多的数据才能勉强达到较高的命中率

