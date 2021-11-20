
A golang example of gRPCtutorials.

## 前言

gRPC效能約為RESTful API 兩倍 : https://medium.com/skyler-record/grpc-vs-rest-%E6%95%88%E8%83%BD%E6%B8%AC%E8%A9%A6-859f062f6ce3
官參考資料 : https://developers.google.com/protocol-buffers/docs/gotutorial

 
gRPC具備四種傳輸方式:

```
-   單向傳輸(Unary): Client 單向，Server 單向
-   Server 單向串流(Streaming-Server): Client 單向，Server 串流
-   Client 單向串流(Streaming-Client): Client 串流，Server 單向
-   雙向串流(Streaming-Bidirectional): Client 串流，Server 串流
```
## 檔案樹說明(File tree)
 ```bash
.
├─client #客戶端串接範例
├─configs #設定檔(TLS加密金鑰)
├─protobuf #gRPC傳輸中間文件(protobuf)範例
├─server #服務端範例
└─tools
    ├─Creat_TLS_keys #golang TLS加密金鑰腳本
    └─protobuf #protobuf中間文件生成工具
```

## 前置設定(Setup)
1. 安裝Golang相依套件 ``go get -a github.com/golang/protobuf/protoc-gen-go``
2. 依照``goroot``路徑和網域名稱``localhost``, 暫時關閉gomod版控``SET GO111MODULE="off"``, 生成TLS加密金鑰 `go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost`

## 建置步驟(Build)
1. 先定義prutobuf:``tools/protobuf/*.proto``, 執行生成腳本``tools/protobuf/creat_protobuf_example(golang).sh``
2. 將``*pb.go`` 文件放自新目錄內獨立成一個``package``使用

