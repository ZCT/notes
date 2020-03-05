``

```go
a := [...]{1,2.3,4,5,6,7.8}  //a是一个数组
b:= a  //b是一个slice

for i, v := range &a {
	a[3] = 100 
	if i == 3 {
		fmt.Println(v) // 100
	}
}

for i, v := range a {
	a[3] = 100
	if i == 3 {
		fmt.Println(v)  //4
	}
}

for i, v := range b {
	b[3] = 100 
	if i == 3 {
	   	fmt.Println(v)  //100
	}
}

原因是因为如果有取值（也就是有v) ，range会对a/b进行拷贝，如果a是数组 ，其中所有的元素都会被拷贝一遍，如果a是slice或者数组地址，则只会拷贝slice或者数组地址。

具体解释参看这个文档：https://www.obs.cr/2017/03/01/range-semantics-in-go/
```







