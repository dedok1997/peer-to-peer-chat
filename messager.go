package main

import (
	"flag"
	ms "./messager"
)



func main() {
	name := flag.String("name", "user", "The user name")

	port := flag.Int("port", 10000, "The server port")
	serverAddr  := flag.String("server_addr", "", "The server address in the format of host:port")

	flag.Parse()
	if *serverAddr == ""{
		ms.RunServer(*name, *port)
	}else{
		ms.RunClient(*serverAddr, *name)
	}

}
