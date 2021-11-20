package main

import (
	"context"
	protobufExample "gRPCClientExample/protobuf"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//單向傳輸(Unary): Client 單向，Server 單向
func sendUnaryTrans(ctx context.Context, gRPCClient protobufExample.ServiceExampleClient) {
	response, err := gRPCClient.UnaryTrans(ctx, &protobufExample.RequestData{Data: "UnaryTrans 客戶端請求訊息!"})
	if err != nil {
		log.Printf("ConnectUnaryTrans err  : %v\n", err)
	} else {
		log.Printf("ConnectUnaryTrans get  : %v\n", response.Data)
	}
}

//Server 單向串流(Streaming-Server): Client 單向，Server 串流
func recvServerStream(ctx context.Context, gRPCClient protobufExample.ServiceExampleClient) {
	serverStream, err := gRPCClient.ServerStream(ctx, &protobufExample.RequestData{Data: "ServerStream 客戶端請求訊息!"})
	if err != nil {
		log.Printf("recvServerStream err  : %v\n", err)
	}

	for {
		response, err := serverStream.Recv()
		if err == nil {
			log.Printf("recvServerStream 接收到訊息為  : %v\n", response)
			continue
		}

		if err == io.EOF {
			if err := serverStream.CloseSend(); err != nil {
				log.Printf("recvServerStream 接收完畢後關閉連線失敗 : %v\n", err)
			} else {
				log.Printf("recvServerStream 成功接收完畢\n")
			}
		} else {
			log.Printf("recvServerStream err  : %v\n", err)
		}
		break
	}
}

//Client 單向串流(Streaming-Client): Client 串流，Server 單向
func sendClientStream(ctx context.Context, gRPCClient protobufExample.ServiceExampleClient) {
	clientStream, err := gRPCClient.ClientStream(ctx)
	if err != nil {
		log.Printf("sendClientStream err  : %v\n", err)
	}

	for index := 0; index < 3; index++ {
		err = clientStream.Send(&protobufExample.RequestData{
			Data: time.Now().Format(time.RFC3339),
		})

		if err != nil {
			log.Printf("sendClientStream Err : %v\n", err)
			break
		}
	}

	if response, err := clientStream.CloseAndRecv(); err != nil {
		log.Printf("sendClientStream 傳送完畢後關閉連線失敗 : %v\n", err)
	} else {
		log.Printf("sendClientStream 成功傳送完畢,並獲得服務端訊息 : %v\n", response)
	}

}

//雙向串流(Streaming-Bidirectional): Client 串流，Server 串流
func recvNsendAllStream(ctx context.Context, gRPCClient protobufExample.ServiceExampleClient) {
	allStream, err := gRPCClient.AllStream(context.Background())
	if err != nil {
		log.Printf("recvNsendAllStream err  : %v\n", err)
	}
	if allStream == nil {
		log.Printf("recvNsendAllStream is nill \n")
	}

	syncWG := sync.WaitGroup{}
	syncWG.Add(2) //並行接收與傳送
	go func() {
		for index := 0; index < 3; index++ {
			err = allStream.Send(&protobufExample.RequestData{
				Data: time.Now().Format(time.RFC3339),
			})

			if err != nil {
				log.Printf("recvNsendAllStream Err : %v\n", err)
				break
			}
		}
		if err := allStream.CloseSend(); err != nil {
			log.Printf("recvNsendAllStream 傳送完畢後關閉連線失敗 : %v\n", err)
		}
		syncWG.Done()
	}()

	go func() {
		for {
			response, err := allStream.Recv()
			if err == nil {
				log.Printf("recvNsendAllStream 接收到訊息為  : %v\n", response)
				continue
			}

			if err == io.EOF {
				log.Printf("recvNsendAllStream 成功接收完畢\n")
			} else {
				log.Printf("recvNsendAllStream err  : %v\n", err)
			}
			break
		}
		syncWG.Done()
	}()

	syncWG.Wait()
}

const (
	maxMsgSize = 1024 * 1021 * 4
)

func main() {

	credentials, err := credentials.NewClientTLSFromFile("../configs/cert.pem", "")
	if err != nil {
		log.Fatalf("%v failed to load server credentials: %v", credentials, err)
	}

	// conn, err := grpc.Dial(":8081", grpc.WithTimeout(2*time.Second),
	// 	grpc.WithTransportCredentials(credentials),
	// 	grpc.WithBlock(),
	// 	grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(maxMsgSize), grpc.MaxCallRecvMsgSize(maxMsgSize)),
	// )
	timeoutCTX, timeoutClose := context.WithTimeout(context.Background(), 2*time.Second)
	defer timeoutClose()
	conn, err := grpc.DialContext(timeoutCTX,
		":8081",
		grpc.WithTransportCredentials(credentials),
		grpc.WithBlock(), //grpc.DialContext使用timeout狀況下必須設置grpc.WithBlock(
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(maxMsgSize), grpc.MaxCallRecvMsgSize(maxMsgSize)),
	)
	if err != nil {
		log.Fatalf("grpc Dail err : %v\n", err)
	}
	defer conn.Close()
	gRPCClient := protobufExample.NewServiceExampleClient(conn)

	sendUnaryTrans(context.Background(), gRPCClient)
	recvServerStream(context.Background(), gRPCClient)
	sendClientStream(context.Background(), gRPCClient)
	recvNsendAllStream(context.Background(), gRPCClient)

}
