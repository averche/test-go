```bash
❯ docker run --rm \
  -p 4770:4770 \
  -p 4771:4771 \
  -v "$(pwd)/echo.proto:/proto/echo.proto:ro" \
  bavix/gripmock /proto/echo.proto

```

```bash
grpcurl -plaintext -proto echo.proto \
  -d '{"name": "test"}' \
  localhost:4770 demo.Echo/Ping
```

Logs:

```bash
2025-08-06T17:03:33Z INF gRPC call completed grpc.code=Unknown grpc.component=server grpc.metadata={":authority":["localhost:4770"],"content-type":["application/grpc"],"user-agent":["grpcurl/dev-build (no version set) grpc-go/1.44.1-dev"]} grpc.method=Ping grpc.method_type=unary grpc.request.content={"name":"test"} grpc.service=demo.Echo grpc.time_ms=0.015583 peer.address=172.17.0.1:57282 protocol=grpc
```
