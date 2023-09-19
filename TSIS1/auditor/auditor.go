package auditor

type Auditor struct {
	position string
	salary   int
	address  string
}

func (a *Auditor) GetPosition() string {
	return a.position
}

func (a *Auditor) GetSalary() int {
	return a.salary
}

func (a *Auditor) GetAddress() string {
	return a.address
}

func (a *Auditor) SetPosition(position string) {
	a.position = position
}

func (a *Auditor) SetSalary(salary int) {
	a.salary = salary
}

func (a *Auditor) SetAddress(address string) {
	a.address = address
}
