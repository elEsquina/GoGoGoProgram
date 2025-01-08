package model 

type User struct {
	Id int
	Name string
	Email string
	Password string
}

func (u User) Authenticate(email, password string) bool {
	return u.Email == email && u.Password == password
}
