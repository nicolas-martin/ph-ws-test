# Features

- Buffered write channel that will save data in a temporary channel in the client structure.
once server is up we immediately send a data
```go
	conn := WebSocketClient{
		sendBuf: make(chan []byte, 10),
	}
```
```go
for {
        ws := conn.Connect()
        if ws == nil {
                return
        }
        [...]
}
```
- Read and Write routines call checks if the ws != nil and calls `Connects` if it is. Uses RWMutex to assure only one client is created and read
- `Connect` loops for ever until it can connect and then unblocks the read/write 
```go
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}
```
- Fix memory leak and orphans goroutines by adding context and cancel functions to goroutines with infinity loops. The context has 50ms to write data before dropping it.
```go

	conn.ctx, conn.ctxCancel = context.WithCancel(context.Background())


```
Then every goroutine will check that context and exit if it's done
```go

        select {
        case <-conn.ctx.Done():
                return nil

```
