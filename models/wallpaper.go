package models

type Wallpaper struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title" validate:"required"`
	Img   string `json:"img" bson:"img" validate:"required"`
	Date  string `json:"date" bson:"date"`
}
