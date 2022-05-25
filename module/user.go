package module

type User struct {
	Id       string
	Email    string
	Password string
}

type Refresh struct {
	Id       string
	Reftoken string
}

type ChangePWform struct {
	OldPW string
	NewPW string
}
