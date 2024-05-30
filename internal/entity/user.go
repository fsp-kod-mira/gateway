package entity

type Role int32

type User struct {
	Email      string `json:"email"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	Id   string `json:"id"`
	Role Role   `json:"role"`
}
