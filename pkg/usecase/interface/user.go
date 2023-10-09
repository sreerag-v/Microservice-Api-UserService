package interfaces

import (
	"context"

	"github.com/sreerag_v/Micro-Api-Auth/pkg/domain"
)
type UserUseCase interface{
	FindByEmail(ctx context.Context,email string)(domain.Users,error)
	Register(ctx context.Context,user domain.Users)(domain.Users,error)
	Login(ctx context.Context,user domain.UserLogin)(string ,error)
	Validate(ctx context.Context,token string)(string ,error)

	GetUsers(ctx context.Context)([]domain.Users,error)
}