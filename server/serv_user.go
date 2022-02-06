package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/rrm003/grpc/user"
)

const(
	addr=":8080"
)

var (
	userregistry []*pb.User
)

type UserManagementServer struct{
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, data *pb.NewUser) (*pb.User, error){
	id := int32(rand.Int())
	fmt.Println("Recived Data", data)
	newuser := &pb.User{Name: data.GetName(), Age: data.GetAge(), Id: id}
	userregistry = append(userregistry, newuser)
	return newuser,nil
}

func (s *UserManagementServer) 	GetAllUser(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error){
	fmt.Printf("UserList %+v\n",userregistry)
	return &pb.ListUsersResponse{List: userregistry},nil
}

func main(){
	log.Println("starting server...")
	lis,err:=net.Listen("tcp",addr)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("listening at :", addr)
	mux := runtime.NewServeMux()
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()
	
	pb.RegisterUserManagementHandlerServer(ctx, mux, &UserManagementServer{})
	
	if err=http.Serve(lis,mux);err!=nil{
		log.Fatal(err)
	}
}
