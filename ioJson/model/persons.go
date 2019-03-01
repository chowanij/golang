package model

type Person struct {
	Id string
	IdType string
	Name string
	SecondName string
	Age int8
	Address string
	Tags []string

}

func (p Person) ShowBasicInfo() {
	println("id type: ", p.IdType, " id: ", p.Id, " name: ", p.Name, " second name: ", p.SecondName, " age: ", p.Age )
}