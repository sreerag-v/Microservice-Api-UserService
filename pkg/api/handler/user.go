package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	pb "github.com/sreerag_v/Micro-Api-Auth/pkg/api/proto"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/domain"
	services "github.com/sreerag_v/Micro-Api-Auth/pkg/usecase/interface"
	"gorm.io/gorm"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	pb.UnimplementedAuthServiceServer
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (usr *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.Users{
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	exist, err := usr.userUseCase.FindByEmail(ctx, user.Email)

	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			Id:     exist.Id,
			Error:  fmt.Sprint(errors.New("email already exists")),
		}, nil
	}

	//User Registeration

	user, err = usr.userUseCase.Register(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &pb.RegisterResponse{
				Status: http.StatusConflict,
				Error:  fmt.Sprint(errors.New("username already exists")),
			}, nil
		}
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  fmt.Sprint(errors.New("failed to register user")),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Id:     user.Id,
	}, nil
}

func (usr *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userLogin := domain.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := usr.userUseCase.Login(ctx, userLogin)

	if err != nil {
		return &pb.LoginResponse{
			AccessToken: "",
			Error:       err.Error(),
		}, err
	}

	return &pb.LoginResponse{
		AccessToken: token,
		Error:       "",
	}, nil
}

func (usr *UserHandler) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Uname: req.Uname,
	}, nil
}

func (usr *UserHandler) Validate(ctx context.Context, token *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	SignedToken := token.Token

	fmt.Println("Token :=",SignedToken)

	userId, err := usr.userUseCase.Validate(ctx, SignedToken)

	if err != nil {
		return &pb.ValidateResponse{
			UserId: "",
			Status: http.StatusUnauthorized,
		}, err
	}

	return &pb.ValidateResponse{
		UserId: userId,
		Status: http.StatusOK,
	}, nil

}

func (usr *UserHandler) GetUsers(ctx context.Context,req *pb.GetUsersRequest)(*pb.GetUsersResponse,error){
	users,err:=usr.userUseCase.GetUsers(ctx)

	if err != nil {
		return &pb.GetUsersResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  fmt.Sprint(errors.New("unable to fetch data")),
		}, errors.New(err.Error())
	}

	var ProtoUser []*pb.User

	for _, user:=range users{
		PbUser:=&pb.User{
			Id: user.Id,
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			UserName: user.UserName,
		}
		ProtoUser = append(ProtoUser,PbUser )
	}
	return &pb.GetUsersResponse{
		Status: http.StatusOK,
		User:   ProtoUser,
	}, nil
}