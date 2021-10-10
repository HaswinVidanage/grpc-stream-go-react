## GRPC Golang + React JS
generate server protoc
``` 
protoc --go_out=plugins=grpc:./server/sensorpb --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative protos/sensor.proto
```

generate client protoc
```
protoc protos/sensor.proto --js_out=import_style=commonjs,binary:././js-client/src/sensorpb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:././js-client/src/sensorpb
```

```
docker-compose -f docker-compose-macos.yml up --build
```