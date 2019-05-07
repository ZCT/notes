

The arguments of a deferred function call or a goroutine function call are all evaluated at the moment when the function call is invoked.

- For a deferred function call, the invocation moment is the moment when it is pushed into the defer-call stack of its caller goroutine.
- For a goroutine function call, the invocation moment is the moment when the corresponding goroutine is created.