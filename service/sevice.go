package service

import (
	"database/sql"

	"github.com/GURUAKASH-MUTHURAJAN/swagger/model"
)

func ListUser() []model.User {
	return model.Users
}
func GetUser(Name string)(*model.User,error){
	for _,user := range model.Users {
		if user.Name == Name{
			return &user,nil
		}
	}
	return nil,sql.ErrNoRows
}

func CreateUser(user model.User)error{
	model.Users=append(model.Users, user)
	return nil
}
