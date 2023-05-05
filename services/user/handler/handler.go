package handler

import (
	"context"
	"flookybooky/services/user/ent"
	pb "flookybooky/services/user/proto"

	"github.com/jinzhu/copier"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	client ent.Client
}

func NewUserHandler(client ent.Client) (*UserHandler, error) {
	return &UserHandler{
		client: client,
	}, nil
}

func (h *UserHandler) PostUser(ctx context.Context, req *pb.PostUserRequest) (*pb.PostUserResponse, error) {
	err := h.client.User.Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetRole(req.Role).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	var res pb.PostUserResponse = pb.PostUserResponse{
		User: &pb.User{
			Username: req.Username,
			Password: req.Password,
			Role:     req.Role,
		},
	}
	return &res, nil
}

func (h *UserHandler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := h.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.GetUsersResponse{}
	err = copier.Copy(&res.Users, &users)
	if err != nil {
		return nil, err
	}
	return res, nil
}
