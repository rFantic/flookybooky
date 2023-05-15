package handler

import (
	"context"
	"flookybooky/internal/util"
	"flookybooky/pb"
	"flookybooky/services/user/ent"
	"flookybooky/services/user/ent/user"
	"flookybooky/services/user/internal"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
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

// func (h *UserHandler).Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)

func (h *UserHandler) PostUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, fmt.Errorf("generate hash: %w", err)
	}
	query := h.client.User.Create().
		SetUsername(req.Username).
		SetPassword(string(hash)).
		SetRole(user.Role(req.Role))
	if req.Customer != nil {
		query.SetCustomerID(req.Customer.Id)
	}
	userRes, err := query.Save(ctx)
	return internal.ParseUserEntToPb(userRes), err
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.UUID) (*pb.User, error) {
	_uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	userRes, err := h.client.User.Get(ctx, _uuid)
	return internal.ParseUserEntToPb(userRes), err
}

func (h *UserHandler) GetUsers(ctx context.Context, req *emptypb.Empty) (*pb.Users, error) {
	query := h.client.User.Query()
	// query = query.Offset(int(req.Offset)).Limit(int(req.Limit))
	usersRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseUsersEntToPb(usersRes), nil
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
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": expireTime,
	})
	tokenString, err := token.SignedString(util.Secretkey)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		JwtToken:   tokenString,
		ExpireTime: expireTime,
	}, err
}
