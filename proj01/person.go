package main

type Person struct {
	Id    int
	Name  string
	Email string
}

type People []Person

func (p *Person) Create() {
	db.Table("people").Create(&p)
}

func ListPeople() []string {
	names := make([]string, 0)

	db.Table("people").Model(Person{}).Order("id desc").Pluck("name", &names)

	return names
}
