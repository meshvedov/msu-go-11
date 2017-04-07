package users

// если мы хотим тестить юзера через моки, то мы должны положить тесты в отдлельный файл
// потому что моки импортируют юзера и если мы будем импортировать моке, то получится циклическая зависимость

type User struct {
	ID   int
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(newName string) {
	u.Name = newName
}

type UserInterface interface {
	GetName() string
	SetName(string)
}
