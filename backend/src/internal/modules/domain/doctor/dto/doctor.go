package dto

type Doctor struct {
	Id         int
	TokenId    string
	Surname    string
	Speciality string
	Role       string
	PhotoPath  string
}

type Doctors []*Doctor

type CreateDoctor struct {
	Surname    string
	TokenId    string
	Speciality string
	Role       string
	PhotoPath  string
}

type UpdateDoctor struct {
	TokenId    string
	Surname    string
	Speciality string
	Role       string
	PhotoPath  string
}
