下面两段代码：

```go
func main() 
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("hello")
	}()
	i:=0
	go func() {
		fmt.Println("world")
	}()
	time.Sleep(time.Second)
}
```

代码反复运行输出结果一直都为：

> hello
>
> world

```go
func main() 
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("hello")
	}()
	i:=0
	go func() {
		fmt.Println("world")
	}()
	select{}
}
```

代码反复运行输出结果一直都为：

>world
>
>hello

​	在go语言的介绍中说，由于多个processor调度的原因，因此goroutine的执行顺序是随机的，并且跟关键字go出现的顺序没有关系。但是在代码的开头，我们已经设置了GOMAXPROCS=1，所以上述代码中的两个goroutine的执行顺序应该是固定的。但是比较奇怪的是，为什么这两段代码中的gorountine执行顺序会不一样。由于代码只在最后一行不一样，所以很容易怀疑是 time.Sleep和select，其中某个的底层实现原理导致gorountine的执行顺序发生了变化。

​	首先，我们来看go routine在processor中的执行顺序，其中go关键字主要是触发了以下函数逻辑：

```
newproc1(){
	runqput(_p_, newg, true){
			//将该goruntine放入runq数组尾部，根据最后一个参数是否为true，将next指针指向该goroutine
		}
	}
	
//代码位于runtime/proc.go文件中

```

![image-20190422015543137](https://github.com/zct/notes/blob/master/golang/img/image-20190422015543137.png?raw=true)

然后我们再来看，是如何调度gorountine运行的。调度的逻辑在schedule函数中，这里我们主要关心是如何从processor的run queue中取数据的，其中是通过函数runqget来获得的

![image-20190422020131164](https://github.com/zct/notes/blob/master/golang/img/image-20190422020131164.png?raw=true)

​	首先从next指针得到下一个要执行的goroutine地址，如果为空，则从runq数组中获取，从数组头部元素开始获取

​	回过头来再看我们原来的代码，程序下个运行的goroutine应该是最后一个创建的，然后再按照先进先出的顺序来执行，第二段代码的输出符合我们的预期

​	为什么第一段代码，在执行完time.Sleep后，goroutine的执行顺序跟预期的不一致呢？我们继续再来看看sleep函数里面做了什么，这是sleep中的函数调用逻辑：

```
time.Sleep 
	timeSleep (runtime/time.go)
		tb.addtimerLocked
			go timeProc() //创建一个goroutine

```

所以到这里，next指向的是timeProc，之前创建的goroutine则是按照先入先出的顺序来执行的



***结论***：

​	在processor的run queue中，gorountine是按照先进先出的顺序被调度执行的，但是最后一个创建的gorountine会被插入到下一个要执行的位置(next指针指向)



*相关阅读*



[time.Sleep sys call](https://github.com/golang/go/issues/25471)