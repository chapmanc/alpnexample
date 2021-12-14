Build protoset
```
cd proto
prototool descriptor-set --include-imports -o ~/test.protoset
cd ..
```
Run server
```
go run main.go
```
Grpcurl
```
grpcurl -vv -cert ./cert.pem -key key.pem -d '{"msg": "Chris"}' -insecure  -protoset ~/test.protoset localhost:7999 hello.HelloWorldService.GetHello
```

Curl h2
```
 curl  -k --cert cert.pem --key key.pem https://localhost:7999/h2
```
