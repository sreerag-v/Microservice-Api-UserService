package interfaces

import (
	"context"

	"github.com/sreerag_v/Micro-Api-Auth/pkg/domain"
)


type UserRepository interface {
	FindByEmail(ctx context.Context,email string)(domain.Users,error)
	Register(ctx context.Context,user domain.Users)(domain.Users,error)

	GetUsers(ctx context.Context)([]domain.Users,error)
	FindById(ctx context.Context,Uid uint)(domain.Users,error)
	DeleteUser(ctx context.Context,Uid int64)(error)


}