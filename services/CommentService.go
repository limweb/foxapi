package services

import (
	"apitest/db"
	"apitest/models"
	"fmt"
) 
 
func GetAllComment(b *[]models.Comment) (err error) { 
	if err = db.GetDB().Find(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func GetAllIdDescComment(b *[]models.Comment) (err error) { 
	if err = db.GetDB().Order("id desc").Find(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func AddNewComment(b *models.Comment) (err error) { 
	if err = db.GetDB().Create(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func GetOneComment(b *models.Comment, id string) (err error) { 
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func PutOneComment(b *models.Comment, id string) (err error) { 
	fmt.Println(b) 
	db.GetDB().Save(b) 
	return nil 
} 
 
func DeleteComment(b *models.Comment, id string) (err error) { 
	db.GetDB().Where("id = ?", id).Delete(b) 
	return nil 
} 
	