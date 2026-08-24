package domainmaster

type User struct {
	Id       string
	Username string
	Password string
	KodeOpd  string
	Email    string
	IsActive bool //default true
}
