➜  ph-ws-test go run panic/main.go
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Hello
2021/09/27 14:55:05 Received: Trouble
2021/09/27 14:55:05 Error in receive: read tcp 127.0.0.1:63140->127.0.0.1:8080: use of closed network connection
panic: concurrent write to websocket connection

goroutine 1 [running]:
github.com/gorilla/websocket.(*messageWriter).flushFrame(0xc0000783c0, 0x1048901, 0x0, 0x0, 0x0, 0x1496f18, 0x30)
        /Users/nmartin/go/pkg/mod/github.com/gorilla/websocket@v1.4.2/conn.go:610 +0x62e
github.com/gorilla/websocket.(*messageWriter).Close(0xc0000783c0, 0x0, 0xc0000b3d98)
        /Users/nmartin/go/pkg/mod/github.com/gorilla/websocket@v1.4.2/conn.go:724 +0x65
github.com/gorilla/websocket.(*Conn).beginMessage(0xc00007c160, 0xc0002020f0, 0x1, 0x2, 0x2)
        /Users/nmartin/go/pkg/mod/github.com/gorilla/websocket@v1.4.2/conn.go:473 +0x262
github.com/gorilla/websocket.(*Conn).NextWriter(0xc00007c160, 0x1, 0xc0001001e0, 0xc0001001e0, 0x1496f18, 0xc000000208)
        /Users/nmartin/go/pkg/mod/github.com/gorilla/websocket@v1.4.2/conn.go:513 +0x53
github.com/gorilla/websocket.(*Conn).WriteMessage(0xc00007c160, 0x1, 0xc0002000a7, 0x5, 0x5, 0x1, 0x1)
        /Users/nmartin/go/pkg/mod/github.com/gorilla/websocket@v1.4.2/conn.go:766 +0x6e
main.main()
        /Users/nmartin/dev/ph-ws-test/panic/main.go:59 +0x2e5
exit status 2
