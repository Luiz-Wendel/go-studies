package users

type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	CreationDate string `json:"creationDate"`
	Age          uint   `json:"age"`
	Height       uint   `json:"height"`
	Active       bool   `json:"active"`
}
