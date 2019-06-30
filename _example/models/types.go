package models

type RawBook struct {
	Authors []string `json:"authors"`
	Chapter *float32 `json:"chapter"`
	Series string `json:"series"`
	Title string `json:"title"`
	Volume *int32 `json:"volume"`
}
type Book struct {
	 RawBook
	UserBook *UserBook `json:"user_book"`
}
type RawUser struct {
	Name *string `json:"name"`
	Password *string `json:"password"`
}
type User struct {
	 RawUser
	UserBook *UserBook `json:"user_book"`
}
type RawUserBook struct {
	CurrentPage *int32 `json:"current_page"`
	Rating *float32 `json:"rating"`
}
type UserBook struct {
	 RawUserBook
	Book *Book `json:"book"`
	User *User `json:"user"`
}
