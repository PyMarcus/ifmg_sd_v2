package controllers

import (
	"github.com/PyMarcus/distribuidos/dao"
	"github.com/PyMarcus/distribuidos/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers() []*models.User{
	user := dao.NewUser()
	return user.GetAllUsers()
}

func CreateUser(name, email string){
	user := models.NewUser(primitive.NewObjectID(), email, name)
	newUser := dao.NewUser()
	newUser.CreateUser(user)
}

func UpdateUser(name, email, oldName, oldEmail string){
	newUser := dao.NewUser()
	newUser.UpdateUser(oldName, oldEmail, name, email)
}

func DeleteUser(name, email string){
	newUser := dao.NewUser()
	newUser.DeleteUser(name, email)
}