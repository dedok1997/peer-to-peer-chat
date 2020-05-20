package messager

import (
	"bufio"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)



type routeGuideServer struct {
}

var user string

func send( stream RouteGuide_SendMessageServer){
	reader := bufio.NewReader(os.Stdin)
	for{
		msg, _, err := reader.ReadLine()
		if err != nil{
			panic(err)
		}
		message := new(Message)
		s := string(msg)
		message.Message = s
		message.UserName = user
		err = stream.Send(message)
		if err != nil{
			panic(err)
		}
	}
}

func (s *routeGuideServer) SendMessage(stream RouteGuide_SendMessageServer)  error {
	go send(stream)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s> %s", in.UserName, in.Message)
		fmt.Println()
	}
}



func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

func RunServer(userName string, port int) {
	user = userName
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
