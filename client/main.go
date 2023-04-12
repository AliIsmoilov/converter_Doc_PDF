package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "app/converter"
)


func main() {

	conn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error to connect: %+v", err)
	}

	c := pb.NewConverterServiceClient(conn)

	resp, err := c.ConvertToPDF(context.Background(), &pb.ConverterReq{File: "Document25"})

	if err != nil{
		log.Println(err)
		return
	}

	fmt.Println(resp)

}