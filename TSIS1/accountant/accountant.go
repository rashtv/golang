package accountant

type Accountant struct {
	position string
	salary   int
	address  string
}

func (a *Accountant) CountSum(nums []int) int {
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func (a *Accountant) GetPosition() string {
	return a.position
}

func (a *Accountant) GetSalary() int {
	return a.salary
}

func (a *Accountant) GetAddress() string {
	return a.address
}

func (a *Accountant) SetPosition(position string) {
	a.position = position
}

func (a *Accountant) SetSalary(salary int) {
	a.salary = salary
}

func (a *Accountant) SetAddress(address string) {
	a.address = address
}
