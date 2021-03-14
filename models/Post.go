package models

//database.AutoMigrate(&models.Post{}) //add to Db.go
type Posts struct {
	Posts []Posts `json:"posts"`
}

type Post struct {
	ModelDefault
	PostForCreate
	Comments []Comment `json:"comments"`                         // HasMany
	UserID   string    `json:"user_id" gorm:"index" example:"" ` // Post useri
	User     User
}

type PostForCreate struct {
	PostForUpdate
}

type PostForUpdate struct {
	Title string `json:"title"  ` // Post title
	Desc  string `json:"desc"   ` // Post desc
}

func (m *Post) TableName() string {
	return "posts"
}
