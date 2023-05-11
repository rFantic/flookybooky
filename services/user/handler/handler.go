package handler

import (
	"context"
	"flookybooky/internal/util"
	"flookybooky/pb"
	"flookybooky/services/user/ent"
	"flookybooky/services/user/ent/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}

	query := h.client.User.Create().
		SetUsername(req.Username).
		SetPassword(string(hash)).
		SetRole(user.Role(req.Role))
	if req.CustomerId != "" {
		query.SetCustomerID(req.CustomerId)
	}
	u, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	var postUser pb.User
	copier.Copy(&postUser, u)
	postUser.CustomerId = u.CustomerID
	var res pb.PostUserResponse = pb.PostUserResponse{
		User: &postUser,
	}
	return &res, nil
}
func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	entUser, err := h.client.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	var user pb.User
	copier.Copy(&user, entUser)
	user.CustomerId = entUser.CustomerID
	res := &pb.GetUserResponse{User: &user}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *UserHandler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	query := h.client.User.Query()
	query = query.Offset(int(req.Offset)).Limit(int(req.Limit))
	users, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.GetUsersResponse{}

	err = copier.Copy(&res.Users, &users)
	for i, c := range res.Users {
		c.CustomerId = users[i].CustomerID
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.client.User.Query().
		Where(user.Username(req.User.Username)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("wrong user name")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.User.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(util.Secretkey)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		JwtToken: tokenString,
	}, err
}
