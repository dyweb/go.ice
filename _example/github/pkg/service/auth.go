package service

type AuthService interface {
	Register(name string, password string) error
	Login(name string, password string) error
}
