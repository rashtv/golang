package main

import (
	"TSIS1/accountant"
	"TSIS1/manager"
	"fmt"
)

func main() {
	var mng = manager.Manager{}
	mng.SetPosition("Main")
	mng.SetSalary(500000)
	mng.SetAddress("8696 Bridge St.\nPensacola, FL 32503")
	fmt.Println(mng)

	var acc = accountant.Accountant{}
	acc.SetPosition("Main")
	acc.SetSalary(500000)
	acc.SetAddress("8696 Bridge St.\nPensacola, FL 32503")
	fmt.Println(acc.CountSum([]int{100, 200, 300, 400, 500}))

}
