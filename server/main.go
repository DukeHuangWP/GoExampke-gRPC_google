package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	protobufExample "gRPCServerExample/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//需依據protobuf中間文件之Server interface 繼承後才可以設置gRPC服務
type GRPCService struct{}

//單向傳輸(Unary): Client 單向，Server 單向
func (this *GRPCService) UnaryTrans(_ context.Context, request *protobufExample.RequestData) (*protobufExample.ResponseData, error) {
	log.Printf("UnaryTrans 收到請求 : %v\n", request.Data)
	return &protobufExample.ResponseData{Data: "服務端回傳 UnaryTrans"}, nil
}

//Server 單向串流(Streaming-Server): Client 單向，Server 串流
func (this *GRPCService) ServerStream(request *protobufExample.RequestData, response protobufExample.ServiceExample_ServerStreamServer) error {
	log.Printf("ServerStream 收到請求 : %v\n", request.Data)
	var err error
	for index := 0; index < 3; index++ {
		err = response.Send(&protobufExample.ResponseData{
			Data: time.Now().Format(time.RFC3339),
		})

		if err != nil {
			log.Printf("ServerStream Err : %v\n", err)
			break
		}
	}
	return err
}

//Client 單向串流(Streaming-Client): Client 串流，Server 單向
func (this *GRPCService) ClientStream(request protobufExample.ServiceExample_ClientStreamServer) error {
	var err error
	for {
		response, err := request.Recv()
		if err == nil {
			log.Printf("recvClientStream 接收到訊息為  : %v\n", response)
			continue
		}

		if err == io.EOF {
			if err := request.SendAndClose(&protobufExample.ResponseData{Data: "ClientStream 服務端接收完畢!"}); err != nil {
				log.Printf("recvClientStream 接收完畢後關閉連線失敗 : %v\n", err)
			} else {
				log.Printf("recvClientStream 成功接收完畢\n")
			}
		} else {
			log.Printf("recvClientStream err  : %v\n", err)
		}
		break

	}
	return err
}

//雙向串流(Streaming-Bidirectional): Client 串流，Server 串流
func (this *GRPCService) AllStream(request protobufExample.ServiceExample_AllStreamServer) error {

	syncWG := sync.WaitGroup{}
	syncWG.Add(2) //並行接收與傳送

	var errSend error
	var errRecv error
	go func() {
		for index := 0; index < 3; index++ {
			errSend = request.Send(&protobufExample.RequestData{
				Data: time.Now().Format(time.RFC3339),
			})

			if errSend != nil {
				log.Printf("AllStream Err : %v\n", errSend)
				break
			}
		}
		syncWG.Done()
	}()

	go func() {
		for {
			response, errRecv := request.Recv()
			if errRecv == nil {
				log.Printf("AllStream 接收到訊息為  : %v\n", response)
				continue
			}

			if errRecv == io.EOF {
				log.Printf("AllStream 成功接收完畢\n")
			} else {
				log.Printf("AllStream err  : %v\n", errRecv)
			}
			break
		}
		syncWG.Done()
	}()

	syncWG.Wait()
	if errSend != nil || errRecv != nil {
		return fmt.Errorf("errSend>%v , errRecv>%v", errSend, errRecv)
	} else {
		return nil
	}

}

const (
	maxMsgSize = 1024 * 1021 * 4
)

func main() {

	credentials, err := credentials.NewServerTLSFromFile("../configs/cert.pem", "../configs/key.pem")
	if err != nil {
		log.Fatalf("%v failed to load server credentials: %v", credentials, err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//gRPCServer := grpc.NewServer() //無TLS加密
	gRPCServer := grpc.NewServer(grpc.Creds(credentials),
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
	)
	protobufExample.RegisterServiceExampleServer(gRPCServer, &GRPCService{})

	go func() {
		if err := gRPCServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("gRPC server 啟動...")

	//========================往下為關機後觸發func========================================{
	quitChan := make(chan os.Signal, 10) //Notify：系統訊號轉將發至channel
	signal.Notify(quitChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quitChan
	//永久等待收到系統發出訊號類型,若符合則往下執行
	close(quitChan)

	if err := listener.Close(); err != nil {
		log.Fatal("關閉tcp port過程中發生錯誤:", err)
	}
	gRPCServer.GracefulStop()

}
