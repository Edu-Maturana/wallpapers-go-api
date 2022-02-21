package models

type Wallpaper struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Img   string `json:"img" bson:"img"`
	Date  string `json:"date" bson:"date"`
}
