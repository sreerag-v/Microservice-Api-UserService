// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/sreerag_v/Micro-Api-Auth/pkg/api"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/api/handler"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/config"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/db"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/repository"
	"github.com/sreerag_v/Micro-Api-Auth/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
