package models

//database.AutoMigrate(&models.Comment{}) //add to Db.go
type Comments struct {
	Comments []Comments `json:"comments"`
}

type Comment struct {
	ModelDefault
	CommentForCreate
	Post Post `json:"post"`
}

type CommentForCreate struct {
	CommentForUpdate
}

type CommentForUpdate struct {
	Comment string `json:"comment"   `             // Comment comment
	PostID  string `json:"post_id" gorm:"index"  ` // Comment postid
}

func (m *Comment) TableName() string {
	return "comments"
}
