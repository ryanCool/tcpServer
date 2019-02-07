# honestbeeHomeTest
It's honestbee home test about tcp server implementation

You can search cat for adoption by color.

Just connect to server through TCP then type the cat color you like, and server will return all adoption info you need!!!



## Build server
 go build server/main.go
## Build client
 go build server/client.go
 
### Server
```
 go run server/main.go -port=3000 
```
TCP server default read timeout: 10 sec
Http server port : 8000
### Client
```
 go run client/main.go -port=3000
```

## Test
```
go test ./server/tcp/
```


### Statistics API
```
curl -X GET localhost:8000/stat
```

### Healthy check API
```
curl -X GET localhost:8000/healthy
```
