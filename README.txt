``` server
protoc --go_out=plugins=grpc:./server/sensorpb --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative protos/sensor.proto
```

```client
protoc sensor.proto --js_out=import_style=commonjs,binary:./../js-client/src/sensorpb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./../js-client/src/sensorpb
```


--go-grpc_opt=paths=source_relative

protoc --go_opt=paths=source_relative —= go_out=plugins=grpc:./../server/sensorpb --go-grpc_out=. protos/sensor.proto