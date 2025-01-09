package models

/*Post repräsentiert einen Beitrag eines Nutzers*/
type Post struct {
	UserID   int    `json:"userId"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments []Comment
}

/*Comment repräsentiert einen Kommentar zu einem Beitrag. Jeder Post hat eine ID, welche mit der PostID der Comments korrespondiert.*/
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
