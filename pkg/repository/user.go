package repository

import (
	"context"
	"fmt"

	"github.com/sreerag_v/Micro-Api-Auth/pkg/domain"
	interfaces "github.com/sreerag_v/Micro-Api-Auth/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}


func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}


func (usr *userDatabase) FindByEmail(ctx context.Context,email string)(domain.Users,error){
	var user domain.Users

	err:=usr.DB.Where("email = ?",email).First(&user).Error

	if err!=nil{
		fmt.Println("err from repos find name", err)
	}
	return user,err
}

func (usr *userDatabase) Register(ctx context.Context,user domain.Users)(domain.Users,error){
	err := usr.DB.Save(&user).Error
	return user, err
}

func (usr *userDatabase) GetUsers(ctx context.Context)([]domain.Users,error){
	var users []domain.Users
	err := usr.DB.Find(&users).Error

	if err!=nil{
		return nil,err
	}

	return users,nil
}

func (usr *userDatabase) FindById(ctx context.Context,Uid uint)(domain.Users,error){
	var user domain.Users
	err := usr.DB.First(&user, Uid).Error

	return user, err
}

func (usr *userDatabase) DeleteUser(ctx context.Context,Uid int64)(error){
	user := &domain.Users{Id: Uid}
	err := usr.DB.Delete(user).Error
	return err
}





