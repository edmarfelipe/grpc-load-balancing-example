# GRPC Load Balancing

This is a simple example of how to use the GRPC load balancing feature.

In this example we are using an client side load balancing, where the client is responsible for choosing the server, the list of servers is provided by a custom [name resolver](https://grpc.io/docs/guides/custom-name-resolution/).

## How to run

```shell
make run-server-1
make run-server-2
make run-client
```

## Reference
- [https://grpc.io/blog/grpc-load-balancing/#client-side-lb-options](https://grpc.io/blog/grpc-load-balancing/#client-side-lb-options)
- [https://grpc.io/docs/guides/custom-name-resolution/](https://grpc.io/docs/guides/custom-name-resolution/)