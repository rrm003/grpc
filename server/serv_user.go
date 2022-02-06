package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/rrm003/grpc/user"
)

const(
	addr=":8080"
)

var (
	userregistry = map[int32] pb.User{}
)

type UserManagementServer struct{
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, data *pb.NewUser) (*pb.User, error){
	id := int32(len(userregistry)+1)
	log.Println("Recived Data", data)
	newuser := pb.User{
		Id: id,
		Fname: data.GetFname(), 
		Age: data.GetAge(), 
		City: data.GetCity(),
		Phone: data.GetPhone(),
		Height: data.GetHeight(),
		Married: data.GetMarried(),
	}
	userregistry[id] = newuser
	return &newuser, nil
}

func (s *UserManagementServer) 	GetAllUser(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error){
	log.Println("UserList :",userregistry)
	response := pb.ListUsersResponse{}
	for key, _ := range userregistry{
		user := userregistry[key]
		response.List = append(response.List, &user)
	}

	return &response,nil
}

func (s *UserManagementServer) 	GetUser(ctx context.Context, in *pb.UserRequest) (*pb.User, error){
	id := in.GetId()
	if val, ok := userregistry[id]; ok {
		log.Printf("User",val)
		return &val,nil
	}
	
	return nil, errors.New("no user found")
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
