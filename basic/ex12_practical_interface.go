package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpenses(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	permEmp1 := Permanent{1, 5000, 20}
	permEmp2 := Permanent{2, 6000, 30}
	contractEmp1 := Contract{3, 3000}
	employees := []SalaryCalculator{permEmp1, permEmp2, contractEmp1}
	totalExpenses(employees)
}
