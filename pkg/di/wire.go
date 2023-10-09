//go:build wireinject
// +build wireinject

package di

import (
	"github.com/sreerag_v/Micro-Api-Auth/pkg/api/handler"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/config"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/db"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/repository"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/usecase"
	http "github.com/sreerag_v/Micro-Api-Auth/pkg/api"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, 
			repository.NewUserRepository, 
			usecase.NewUserUseCase,
			handler.NewUserHandler, 
			http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}