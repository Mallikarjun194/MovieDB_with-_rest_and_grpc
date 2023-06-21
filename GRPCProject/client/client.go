package main

import (
	pb "GRPCProject/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	fmt.Println("Client running..!")
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	c := pb.NewMovieServiceClient(conn)

	//CREATE Movie

	req := &pb.CreateMovieRequest{
		Movie: &pb.Movie{
			Title: "kgf",
			Genre: "thriller",
		},
	}

	res, err := c.CreateMovie(context.Background(), req)
	fmt.Println("response from post call", res)

	//GET Movie

	getReq := &pb.ReadMovieRequest{
		Id: res.GetMovie().GetId(),
	}
	getRes, err := c.GetMovie(context.Background(), getReq)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	fmt.Println("response from get call", getRes)

	//UPDATE Movie

	updateReq := &pb.UpdateMovieRequest{
		Movie: &pb.Movie{
			Id:    res.GetMovie().GetId(),
			Title: "anvcd",
			Genre: "thriller",
		},
	}
	updateRes, err := c.UpdateMovie(context.Background(), updateReq)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	fmt.Println("response from update call", updateRes)

	//LIST Movie

	getAllReq := &pb.ReadMoviesRequest{}
	getAllResp, err := c.GetMovies(context.Background(), getAllReq)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	fmt.Println("response from get all call", getAllResp)

	//DELETE Movie

	DeLReq := &pb.DeleteMovieRequest{Id: res.GetMovie().GetId()}
	delRes, err := c.DeleteMovie(context.Background(), DeLReq)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	fmt.Println("response from delete call", delRes)

	if err != nil {
		log.Fatalf("Error while making grpc call %v", err)
	}

}
