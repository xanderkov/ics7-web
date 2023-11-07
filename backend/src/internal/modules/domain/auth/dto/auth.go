package dto

// Сервис для аунтефикации в телеграмм

type Auth struct {
	Id      int
	TokenId string
}

type Users []*Auth

type NewDoctor struct {
	TokenId    string
	Surname    string
	Speciality string
	Role       string
	PhotoPath  string
}

// NewAccount данные нового пользователя
type NewAccount struct {
	Login    string // Логин в системе
	Password string // Пароль
}

type Account struct {
	Id       int
	Login    string // Логин в системе
	Password string // Пароль
}

type Accounts []*Account

type CreateAccount struct {
	Login    string // Логин в системе
	Password string // Пароль
}

type UpdateAccount struct {
	Login    string // Логин в системе
	Password string // Пароль
}
