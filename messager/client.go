package messager

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

var name string

func sendx(stream RouteGuide_SendMessageClient) {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		message := new(Message)
		s := string(msg)
		message.Message = s
		message.UserName = "user"
		err = stream.Send(message)
		if err != nil {
			return
		}
	}
}

func sendMessage(client RouteGuideClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.SendMessage(ctx)
	if err != nil {
		log.Fatalf("%v.Message(_) = _, %v: ", client, err)
	}
	go sendx(feature)
	for {
		in, err := feature.Recv()
		if err == io.EOF {
			//panic(err)
			return err
		} else
		if err != nil {
			return err
		} else {
			fmt.Printf("%s> %s", in.UserName, in.Message)
			fmt.Println()
		}
	}
}

func RunClient(serverAddr, userName string) {
	name = userName
	var opts []grpc.DialOption
	grpc.W
	opts = append(opts, grpc.WithStreamInterceptor())
	opts = append(opts, grpc.WithBlock())

	for {
		conn, err := grpc.Dial(serverAddr, opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		client := NewRouteGuideClient(conn)

		err = sendMessage(client)
		if err != nil{
			//panic(err)
		}
		err = conn.Close()
		if err != nil{
			//panic(err)
		}
	}

}
