package service

import (
	"log"
	"context"
	"errors"

	"github.com/ughvj/goguild-ergo-sum/pb"
	"github.com/ughvj/goguild-ergo-sum/adapter"
)

type CommandService struct {}

func (s *CommandService) Receive(
	ctx context.Context,
	req *pb.ReceiveRequest,
) (
	*pb.ReceiveResponse,
	error,
){
	repo, err := adapter.NewCommandRepositoryRedis()
	if err != nil {
		log.Println("Cannot connect the redis")
		return nil, errors.New("An error occured")
	}

	res, err := repo.Get("command")
	if err != nil {
		log.Println("Cannot connect the redis")
		return nil, errors.New("Resource was not found")
	}

	log.Println("Response: " + res)
	return &pb.ReceiveResponse{
		Command: res,
	}, nil
}