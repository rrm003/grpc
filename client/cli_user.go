package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rrm003/grpc/user"
	"google.golang.org/grpc"
)

const(
	address = "localhost:8080"
)

func main(){
	conn,err:=grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err!=nil{
		log.Fatal(err)
	}
		defer conn.Close()

	c:=pb.NewUserManagementClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	user,err:=c.CreateNewUser(ctx,&pb.NewUser{Name: "rrm",Age: 23})
	if err!=nil{
		log.Fatalf("Error creating user %v", err)
	}
	log.Printf("User Created %+v\n", user)

	list,err:=c.GetAllUser(ctx,&pb.ListUsersRequest{})
	if err!=nil{
		log.Fatalf("Error creating user %v", err)
	}
	log.Printf("User List %+v\n", list)
}
