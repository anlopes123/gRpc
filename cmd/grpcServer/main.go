package main

import (
	"github.com/anlopes123/gRpc/internal/database"
	"github.com/anlopes123/gRpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/mattn/go-sqlite3"
)

func main(){
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err:= grpcServer.Serve(lis); err!=nil {
		panic(err)
	}
	
}