package model

type RegisterInput struct {
	Name         string
	Avatar       string
	Password     string
	UserSalt     string
	Sex          int
	Status       int
	Sign         string
	SecretAnswer string
}

type RegisterOutPut struct {
	Id uint
}

type LoginUserInput struct {
	Name     string
	Password string
}

type UpdatePasswordInput struct {
	Password     string
	UserSalt     string
	SecretAnswer string
}

type UpdatePasswordOutPut struct {
	Id uint
}
