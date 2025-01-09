package repository

import (
	"encoding/json"
	"io"
	"os"
	"project3/model"
	"fmt"
) 


var ServerStorage []model.Student

func LoadJson() error{
	f, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	json.Unmarshal(data, &ServerStorage)
	return nil
}

func SaveJson(l []model.Student) error{
	fmt.Println("Saving")
	fmt.Println(l)

	f, err := os.Create("data.json")
	if err != nil {
		return err
	} 
	defer f.Close()

	data, err := json.Marshal(l)
	if err != nil {
		return err 
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	
	return nil
}




