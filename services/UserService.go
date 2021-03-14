package services

import (
	"apitest/db"
	"apitest/models"
	"fmt"
)

func GetAllUser(b *[]models.User) (err error) {
	if err = db.GetDB().Preload("Posts.Comments").Preload("Roles.Permissions").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetAllIdDescUser(b *[]models.User) (err error) {
	if err = db.GetDB().Order("id desc").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *models.User) (err error) {
	if err = db.GetDB().Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *models.User, id string) (err error) {
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *models.User, id string) (err error) {
	fmt.Println(b)
	db.GetDB().Save(b)
	return nil
}

func DeleteUser(b *models.User, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(b)
	return nil
}
