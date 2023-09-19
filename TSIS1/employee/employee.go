package employee

type Employee interface {
	GetPosition() string
	GetSalary() int
	GetAddress() string
	SetPosition(position string)
	SetSalary(salary int)
	SetAddress(address string)
}
