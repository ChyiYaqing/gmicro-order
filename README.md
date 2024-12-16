# gmicro-order

```bash
➜ tree .       
.
├── cmd
│   └── main.go                                             ---- Dependency injection and application entry point
├── config
│   └── config.go                                           ---- Application Configuration
├── go.mod
├── internal
│   ├── adapters                                            ==== 适配器文件夹 - 此文件夹包含使用端口中定义的协议具体实现
│   │   ├── db
│   │   │   └── db.go                                       ---- DB specific implementation
│   │   └── grpc
│   │       ├── grpc.go                                     ---- gRPC adapter implementation
│   │       └── server.go
│   ├── application                                         ==== 应用程序文件夹 - 此文件夹包含微服务业务逻辑，它是引用业务实体的域模型和向其他模块公开核心功能的API的组合
│   │   └── core
│   │       ├── api
│   │       │   └── api.go                                  ---- Core functionalities
│   │       └── domain
│   │           └── order.go                                ---- Order domain model
│   └── ports                                               ==== 端口文件夹 - 此文件夹包含核心应用程序与第三方之间的协议信息
│       ├── api.go
│       └── db.go                                           ---- Interfaces that contain API and DB method signatures
└── README.md

12 directories, 11 files
```

运行Order服务

```bash
➜ SQLITE_DB=data/sqlite.db APPLICATION_PORT=8080 ENV=development go run cmd/main.go
2024/12/13 19:26:06 starting order service on port 8080 ...
```

使用grpcURl

```bash
➜ grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' -plaintext localhost:8080 Order/Create
{
  "orderId": "1"
}
```
