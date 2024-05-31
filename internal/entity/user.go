package entity

type Role string

type User struct {
	Email      string `json:"email"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	Role       Role   `json:"role"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	Id   string `json:"id"`
	Role Role   `json:"role"`
}
