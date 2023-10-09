package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sreerag_v/Micro-Api-Auth/pkg/domain"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/jwt"
	interfaces "github.com/sreerag_v/Micro-Api-Auth/pkg/repository/interface"
	services "github.com/sreerag_v/Micro-Api-Auth/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (usr *userUseCase) FindByEmail(ctx context.Context, email string) (domain.Users, error) {
	user, err := usr.userRepo.FindByEmail(ctx, email)
	return user, err
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func (usr *userUseCase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {
	user.Password = HashPassword(user.Password)

	user, err := usr.userRepo.Register(ctx, user)

	return user, err
}

func (usr *userUseCase) Login(ctx context.Context, user domain.UserLogin) (string, error) {
	exist, err := usr.userRepo.FindByEmail(ctx, user.Email)

	if err != nil {
		fmt.Println("error found in finding ")
	}

	if exist.Id == 0 {
		return "", errors.New("user does not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("password incorrect")
	}

	token, err := jwt.GenerateJWT(uint(exist.Id))
	if err != nil {
		return "", errors.New("could not create token")
	}

	return token, nil
}

func (usr *userUseCase) Validate(ctx context.Context, token string) (string, error) {
	claims,err := jwt.ValidateToken(token)
	if err!=nil{
		return claims.Id, nil
	}
	return claims.Id, nil
}

func (usr *userUseCase)	GetUsers(ctx context.Context)([]domain.Users,error){
	users,err:=usr.userRepo.GetUsers(ctx)
	if err!=nil{
		return nil,err
	}
	return users,nil
}

