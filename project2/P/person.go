package P

import (
	"sort"
)

type People struct {
	People []Person
}

type Person struct {
	Name string
	Age int
	Salary float64
	Education string
}

func (p *People) Sort(sorter func(i, j int) bool) {
	sort.Slice(p.People, sorter)
}

func (p *People) AgeSort() func(i, j int) bool {
	return func(i, j int) bool {
		return p.People[i].Age > p.People[j].Age
	}
}

func (p *People) SalarySort() func(i, j int) bool {
	return func(i, j int) bool {
		return p.People[i].Salary > p.People[j].Salary
	}
}

func (p *People) CountByEducation() map[string]int{
	peeps := p.People 
	m := make(map[string]int)

	for i := 0; i < len(peeps); i++ {
		_, exists := m[peeps[i].Education] 
		if exists{ 
			m[peeps[i].Education]++
		} else{
			m[peeps[i].Education] = 1
		}
	}

	return m
}

func (p *People) GetNameByAge(Age int) []string{
	var ret []string
	peeps := p.People 
	for i:= 0; i < len(peeps); i++{
		if peeps[i].Age == Age {
			ret = append(ret, peeps[i].Name)
		}
	}

	return ret
}

func (p *People) GetNameBySalary(Salary float64) []string{
	var ret []string
	peeps := p.People 
	for i:= 0; i < len(peeps); i++{
		if peeps[i].Salary == Salary {
			ret = append(ret, peeps[i].Name)
		}
	}

	return ret
}

func (p *People) SummaryStats() (float64, float64){

	peeps := p.People 
	var avgAge, avgSalary float64;

	if len(peeps) == 0 {
		return 0.0, 0.0
	}

	for i := 0; i < len(peeps); i++ {
		avgAge += float64(peeps[i].Age)
		avgSalary += float64(peeps[i].Salary)
	} 
	
	l := float64(len(peeps))
	return avgAge/l, avgSalary/l
}

