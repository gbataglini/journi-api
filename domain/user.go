package domain

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserService interface {
	ListUsers() ([]User, error)
}

type UserStore interface {
	UserService
}
