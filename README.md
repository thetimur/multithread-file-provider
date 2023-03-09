### Server/Client parameters:

```bash
server_exec cuncurrency_level
```

```bash
client_exec server_host server_port file_dest
```

### Examples:

Server:
```bash
timur@timur-lenovo:~/multithread-file-provider$ go run server/main.go 2
```

Client:

```bash
timur@timur-lenovo:~/multithread-file-provider$ go run client/main.go localhost 8080 a.txt
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: text/plain

Hello, world!
```