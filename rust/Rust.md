##Rust



### The Rules of Ownership 

  *  Each value in Rust has a variable that's called its owner
  *  There can only be one owner at a time
  *  when the owner goes out of scope, the value will be dropped



>  对于在栈上分配的变量，赋值是copy

>  对于在堆上分配的变量，赋值是move

### The Rules of References

* At any given time , you can have either one mutable reference or any number of immutable references
* References must always be va







[Traits](file:///Users/tangzhongcheng/rust/book/book/ch10-02-traits.html)

[Generic Data Types](file:///Users/tangzhongcheng/rust/book/book/ch10-01-syntax.html)

这两章看了一遍，有些地方没有深究，后面还需要回过头来再看一遍



* As with generic type parameters, we need to declare generic lifetime parameters inside angle brackets between the function name and the parameter list.

> 为什么要这样，这样做是为了什么



##  lifetime

* each parameter that is a reference gets its own lifetime parameter
* if there is exactly one input lifetime parameter, that lifetime is assigned to all output lifetime parameters
* if there are multiple input lifetime parameters, but one of them is `&self` or `&mut self` because this is a method, the lifetime of `self` is assigned to all output lifetime parameters

