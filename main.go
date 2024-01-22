package main

import (
	"fmt"
	"grpc_go/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	products := pb.Products{
		Pagination: &pb.Pagination{
			Total:       2,
			PerPage:     2,
			CurrentPage: 1,
			LastPage:    1,
		},
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "POLO SHIRT",
				Price: 70000,
				Stock: 3,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    1,
				Name:  "LEVIS SHIRT",
				Price: 130000,
				Stock: 7,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
		},
	}

	data, err := proto.Marshal(&products)
	if err != nil {
		log.Fatal("Err marshal: ", err)
	}

	fmt.Println("Data after marshal: ", data)

	var dataAfter pb.Products

	err = proto.Unmarshal(data, &dataAfter)
	if err != nil {
		log.Fatal("Err unmarshal: ", err)
	}

	fmt.Println("Data after unmarshal: ", dataAfter)
}
