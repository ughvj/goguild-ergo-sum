### Launch

`$ docker-compose up -d`

### Register test data

`$ redis-cli`

`127.0.0.1:6379> set command testvalue`

### Send a request with gRPC client

`$ grpcurl -H "Authorization: Bearer hoge" -plaintext 127.0.0.1:50051 Command/Receive`

expected...

``` json
{
  "command": "testvalue"
}
```