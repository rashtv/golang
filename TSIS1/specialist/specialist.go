package specialist

type Specialist struct {
	position string
	salary   int
	address  string
}

func (s *Specialist) GetPosition() string {
	return s.position
}

func (s *Specialist) GetSalary() int {
	return s.salary
}

func (s *Specialist) GetAddress() string {
	return s.address
}

func (s *Specialist) SetPosition(position string) {
	s.position = position
}

func (s *Specialist) SetSalary(salary int) {
	s.salary = salary
}

func (s *Specialist) SetAddress(address string) {
	s.address = address
}
