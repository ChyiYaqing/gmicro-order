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

gRPC 双向TLS(mutual TLS/mTLS)认证

OpenSSL 生成自签名证书

```bash
openssl req -x509 \     # 公钥证书格式
  -sha256 \             # 使用SHA-256 摘要算法
  -newkey rsa:4096 \    # 生成私钥及其证书
  -days 365 \           # 有效期
  -keyout ca-key.pem \  # 私钥文件
  -out ca-cert.pem \    # 证书文件
  -subj "/C=CN          # /C 表示国家/地区
         /ST=ZHEJIANG   # /ST 省份信息
         /L=HANGZHOU    # /L 表示城市信息
         /O=CHYIDL      # /O 表示组织
         /OU=Software   # /OU 表示组织单位
         /CN=*.chyidl.com  # 标识域名
         /emailAddress=chyiyaqing@gmail.com" # 证书中添加身份信息
  -nodes # No DES 表示私钥不会被加密

➜ openssl req -x509 -sha256 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=CN/L=HANGZHOU/O=CHYIDL/OU=Software/CN=*.chyidl.com/emailAddress=chyiyaqing@gmail.com" -nodes
...+..........+.....+.+.....+.......+..............+...+..................+....+...+..+....+.........+..+...+.+.........+.....+......+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*.....................+...+...+....+...........+.+......+...+.....+.+......+..+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*..+.......+..............+......+...+....+..+......+...+......+...+..........+............+...........+.........................+..+............+......+....+...........+...+.......+............+............+...+.....+.+..+.......+..............................+.................+.+...+............+.........+.........+..............+.......+......+........+.........+......................+.....+.........+......+..........+...........+...+.......+............+......+........................+..+.+............+...............+..+....+..+............+......+.......+..+.......+......+...............+.................+...+..........+...+.....+............+.........+.............+...+..+......+....+..........................+.+.....+.............+...+.....+.......+.....+............+....+..+............+...+.......+..+.........+......+..........+......+..............+....+.........+..+...............+....+....................+.+...+...+..+............+.........+.+...+..................+.................+.+......+.....+....+..+............+.+...........+...+......+.+..............+.........+.......+...+...+..................+....................+.+.........+......+.....+.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
...........+..+...+...+.........+......+......+....+...........+......+............+......+.+...+......+..+.........+.+........+......+.+..+.......+......+..+.......+...+..+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*.+.+..+....+.....+...+...+............+.+......+.....+....+......+.....+....+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*..............+....+.....+..........+.....+...+....+...+......+..+..................+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
-----
```

```bash
➜ tree .         
.
├── ca-cert.pem   # 公钥
└── ca-key.pem

1 directory, 2 files
```

验证生成的CA证书

```bash
➜ openssl x509 -in ca-cert.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            49:0e:ee:97:ff:5c:08:86:e8:77:df:5d:c5:57:17:59:be:65:b6:a5
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = CN, L = HANGZHOU, O = CHYIDL, OU = Software, CN = *.chyidl.com, emailAddress = chyiyaqing@gmail.com
        Validity
            Not Before: Dec 19 10:40:03 2024 GMT
            Not After : Dec 19 10:40:03 2025 GMT
        Subject: C = CN, L = HANGZHOU, O = CHYIDL, OU = Software, CN = *.chyidl.com, emailAddress = chyiyaqing@gmail.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (4096 bit)
                Modulus:
                    00:e2:79:25:4c:49:2b:92:d1:3e:a6:e1:2c:f1:c9:
                    aa:e5:f1:f3:63:92:df:5e:67:7e:15:32:6b:91:51:
                    17:49:e7:ee:0e:ed:88:22:35:0d:4d:af:6e:5b:dc:
                    db:b4:b0:69:8e:0b:0c:a9:a8:c1:1e:94:92:83:48:
                    ae:59:07:fb:95:ed:b7:99:2c:dd:96:85:2b:49:a8:
                    98:4c:3d:aa:48:40:2b:ef:15:1b:05:2d:73:f8:da:
                    d1:12:58:ec:89:86:32:ff:c2:60:54:98:73:f4:4c:
                    99:0b:e1:64:dc:81:98:0b:be:6e:85:1c:70:8c:cb:
                    30:7d:e4:ae:95:e6:3f:cd:7a:c4:2c:84:ec:00:4a:
                    84:6a:0b:7a:e7:21:d7:4c:92:f6:e9:cf:43:62:44:
                    6b:5e:65:4e:6e:24:a2:14:e0:64:85:b5:78:de:0b:
                    f5:b8:13:62:1f:a0:8c:f6:09:5f:34:79:b5:6b:29:
                    28:eb:40:9d:dd:4a:97:af:87:e6:84:b6:19:18:81:
                    d4:18:bf:e6:56:77:02:50:93:d3:77:40:3b:b7:8c:
                    03:04:6f:dd:3f:4c:a6:0c:84:94:44:73:b2:e5:6a:
                    e6:a0:98:56:e9:79:c8:43:c1:ca:00:09:37:b5:03:
                    12:e6:d4:a4:45:6b:f5:5d:b1:9a:75:16:ec:34:bb:
                    a2:b5:df:3b:26:8d:88:98:b2:e8:4a:64:e7:f0:f3:
                    8d:51:21:c2:f3:51:b4:71:c7:bf:6f:56:46:51:a6:
                    7e:aa:92:78:47:a3:1a:1d:cf:8a:59:d1:31:43:ff:
                    f2:8c:de:bb:10:dc:f4:44:ea:bb:42:87:0e:5a:e6:
                    95:32:3e:e5:08:28:f9:54:e9:d1:d0:2b:69:72:c1:
                    f1:79:44:0c:d6:01:8a:71:b5:c8:80:c2:cd:46:1a:
                    c3:69:95:fe:16:8c:f1:3f:5b:5b:19:41:46:24:18:
                    b2:3b:9d:b7:d1:ff:8a:df:46:05:be:bd:79:e3:6e:
                    b4:c0:f7:9b:7c:3d:02:fd:e5:15:db:97:98:b2:6e:
                    3d:0b:a5:35:ad:91:8e:22:69:b9:39:64:40:90:08:
                    2e:0a:96:b1:52:66:25:f0:3f:3a:50:9d:92:16:8f:
                    fc:33:60:d0:b2:70:9d:b5:f4:5a:b3:36:23:57:fb:
                    05:6d:e4:d2:7e:dd:45:a0:47:07:6e:3d:ec:3c:b7:
                    e3:a6:7f:c7:cf:20:3c:86:e7:25:dd:65:d3:d3:9a:
                    3a:8e:5d:ed:d6:b1:84:1d:32:0a:89:25:1b:31:bd:
                    3f:b0:e3:05:38:8c:4e:64:4a:03:e8:33:65:75:db:
                    fc:c6:12:50:d5:14:71:ec:c9:91:42:f7:d1:e6:e0:
                    f1:b5:a7
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Key Identifier: 
                56:78:EF:86:51:EA:0E:BF:44:26:D8:8F:83:0C:99:20:D2:E2:89:90
            X509v3 Authority Key Identifier: 
                56:78:EF:86:51:EA:0E:BF:44:26:D8:8F:83:0C:99:20:D2:E2:89:90
            X509v3 Basic Constraints: critical
                CA:TRUE
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        6b:76:91:06:70:89:b7:d0:f2:7e:db:e0:d6:d6:e0:1d:be:66:
        98:86:fa:7a:33:0f:da:07:a3:87:cd:91:ff:9b:ed:33:bd:8c:
        9b:3b:15:93:a7:2b:f5:cd:43:57:89:ba:ce:09:be:f2:92:b5:
        07:9a:20:ff:8b:c6:45:2a:ab:ab:58:ec:e6:41:1f:57:c3:11:
        27:01:09:e0:02:e5:33:a7:77:5b:54:c3:78:b8:ea:6e:cb:2f:
        ef:6b:e4:3c:f9:76:15:66:f6:95:4c:c0:a9:6b:f3:9d:60:72:
        71:e7:07:6e:0f:72:7c:dd:b9:68:3b:fe:80:25:3d:62:77:c3:
        1e:e1:29:0f:75:c0:3d:d9:bc:8f:7c:ca:ca:7e:f0:35:de:4e:
        2f:16:1f:62:52:7f:63:6b:13:08:b5:3f:ce:78:ee:63:1f:0e:
        61:34:c6:a7:3b:86:a0:68:96:6d:c9:14:f6:33:95:cb:97:4f:
        18:9d:f9:02:f3:2b:f5:8f:2e:a0:ca:06:40:70:81:85:91:d2:
        54:51:cb:30:03:49:e9:cc:94:98:d2:f1:76:a6:0f:f6:dd:ef:
        79:89:b8:c3:b9:3a:d9:bd:be:e9:49:df:58:04:a4:f8:9b:07:
        d3:c5:b9:e7:96:6a:99:95:37:a1:71:89:50:ab:39:25:58:b9:
        87:8c:25:96:b9:50:4d:4c:f6:5c:6f:3f:c3:a4:8f:08:a1:76:
        0b:b4:ef:1e:79:e1:b7:26:e6:fe:03:61:cd:72:33:51:31:41:
        47:56:03:7b:4f:d3:c4:6d:1c:7c:a7:0c:29:77:97:0b:34:21:
        9f:9a:5a:c4:3f:eb:3f:fd:5d:3e:9b:ca:c7:12:0e:20:a7:92:
        2c:d2:43:86:f5:de:1b:1b:87:b5:a7:ad:83:43:78:dd:67:b3:
        2e:28:e7:b0:8d:bc:9d:a4:0a:eb:65:14:10:9f:37:ff:08:87:
        40:80:30:0c:ac:8e:fb:29:88:3a:2b:e8:a3:be:98:95:e7:6a:
        81:75:16:80:e8:6c:82:a0:12:d6:4e:d1:5c:cd:77:da:05:6f:
        9f:af:ec:54:cf:ee:3b:b6:74:7b:8c:70:93:08:d8:99:5d:e6:
        52:b4:dc:be:0a:8d:c3:4f:ca:c1:2b:5b:d2:49:7b:2c:a7:c3:
        83:26:05:fd:eb:f3:ce:51:1c:71:fb:66:1b:59:b2:d6:90:63:
        43:d2:79:14:58:e1:13:0e:59:6a:10:60:13:68:54:27:2c:83:
        00:4f:74:a2:ce:f8:b5:8c:1c:a1:d9:b0:d7:3a:30:aa:b9:d1:
        72:37:d2:08:53:e3:03:90:60:73:0a:85:84:3d:95:61:98:31:
        02:05:20:03:b0:a8:8f:66
```

生成server证书和私钥

```bash
➜ openssl req \  
-newkey rsa:4096 \
-keyout server-key.pem \ # 私钥位置
-out server-req.pem \  # 证书请求位置
-subj "/C=CN/L=HANGZHOU/O=CHYIDL/OU=Software/CN=*.chyidl.com/emailAddress=chyiyaqing@gmail.com" \
-nodes \
-sha256          
```

server-ext.cnf

```bash
➜ cat server-ext.cnf 
subjectAltName=DNS:*.chyidl.com,DNS:*.localhost,IP:0.0.0.0
```

使用CA-CERT.pem签署请求

```bash
➜ openssl req x509 \
-in server-req.pem \
-days 365  \
-CA ca-cert.pem \
-CAkey ca-key.pem \
-CAcreateserial \
-out server-cert.pem \
-extfile server-ext.cnf \
-sha256

Certificate request self-signature ok
subject=C = CN, L = HANGZHOU, O = CHYIDL, OU = Software, CN = *.chyidl.com, emailAddress = chyiyaqing@gmail.com
```