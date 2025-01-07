package main

import (
	"fmt"
	"os"
	"io"
	"encoding/json"
	"project2/P"
)


func main(){
	Peeps, _ := loadData()

	out := make(map[string]any)
	avgAge, avgSalary := Peeps.SummaryStats()
	
	Peeps.Sort(Peeps.AgeSort())
	oldest, youngest := Peeps.People[0], Peeps.People[len(Peeps.People) - 1] 

	Peeps.Sort(Peeps.SalarySort())
	richest, poorest := Peeps.People[0], Peeps.People[len(Peeps.People) - 1] 

	out["Average Age"] = avgAge
	out["Youngest person name"] = Peeps.GetNameByAge(youngest.Age)
	out["Oldest person name"] = Peeps.GetNameByAge(oldest.Age)
	out["Average Salary"] = avgSalary
	out["Lowest Salary person name"] = Peeps.GetNameBySalary(poorest.Salary)
	out["Highest Salary person name"] = Peeps.GetNameBySalary(richest.Salary)
	out["Counts By Education"] = Peeps.CountByEducation()

	fmt.Println(out)
	saveData(out)
}

func saveData(m map[string]any) {
	file, err := os.Create("output.json")
	if err != nil {
		return
	}
	defer file.Close()

	data, err := json.Marshal(m)
	if err != nil {
		return
	}

	_, err = file.Write(data)
	if err != nil {
		return
	}
}

func loadData() (P.People, error) {
	var P P.People

	file, err := os.Open("data.json")
	if err != nil {
		return P, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return P, err
	}
	err = json.Unmarshal(data, &P)
	if err != nil {
		return P, err
	}

	return P, nil
}
