package user

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Role string

const (
	Admin Role = "Admin"
	Staff Role = "Staff"
)

type User struct {
	UserName string
	Password string
	Role     Role
}

func (User) TableName() string {
	return "users"
}
