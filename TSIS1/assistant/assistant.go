package assistant

type Assistant struct {
	position string
	salary   int
	address  string
}

func (a *Assistant) GetPosition() string {
	return a.position
}

func (a *Assistant) GetSalary() int {
	return a.salary
}

func (a *Assistant) GetAddress() string {
	return a.address
}

func (a *Assistant) SetPosition(position string) {
	a.position = position
}

func (a *Assistant) SetSalary(salary int) {
	a.salary = salary
}

func (a *Assistant) SetAddress(address string) {
	a.address = address
}
