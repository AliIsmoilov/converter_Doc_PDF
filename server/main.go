package main

import (
	pb "app/converter"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)


type server struct {
	pb.UnimplementedConverterServiceServer
}


func (s *server) ConvertToPDF(ctx context.Context, req *pb.ConverterReq) (*pb.ConverterResp, error) {

	filePath := "./documents/" + req.File + ".docx"
 	
	command := exec.Command("soffice", "--headless", "--convert-to", "pdf:writer_pdf_Export", filePath)
 	
	output, err := command.Output()
	if err != nil {
 		log.Fatal("Error when converting document")
 		return nil, err
	}

	return &pb.ConverterResp{File: string(output)}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil{
		log.Fatalf("Fail to listen : %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterConverterServiceServer(s, &server{})

	
	fmt.Println("Listen GRPC 8001...")

	err = s.Serve(lis)
	if err != nil{
		log.Fatalf("Fail to Serve %+v", err)
	}

}