

The arguments of a deferred function call or a goroutine function call are all evaluated at the moment when the function call is invoked.

- For a deferred function call, the invocation moment is the moment when it is pushed into the defer-call stack of its caller goroutine.
- For a goroutine function call, the invocation moment is the moment when the corresponding goroutine is created.



src/cmd/compile/internal/gc/ssa.go

src/cmd/compile/internal/gc/walk.go



```go
var err error

fmt.primtln(err== nil) //true


可以对一个interface申明变量,其underlytype为空

concretetype underlytype value
```

