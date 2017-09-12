package custom_pack

//Person 结构体除了属性和名称，还有一个tag属性
//可以使用refleact拿到里面的值
type Person struct {
	name    string "Test1"
	age     int    "Test2"
	address string "Test3"
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) GetAddress() string {
	return p.address
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}
func (p *Person) SetAddress(address string) {
	p.address = address
}
