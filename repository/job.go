package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alipourhabibi/work-progress/files"
)

type job struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewJob(name string, description string) *job {
	return &job{
		Name:        name,
		Description: description,
	}
}

func (j *job) Add() {
	// load jobs to a variable
	jobs := []job{}
	data, err := os.ReadFile(files.WORK)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &jobs)

	// check if name is unique
	unique := true
	for _, v := range jobs {
		if v.Name == j.Name {
			unique = false
		}
	}

	if !unique {
		fmt.Printf("[ERROR] name %s already exists\n", j.Name)
		return
	}

	jobs = append(jobs, *j)
	data, err = json.Marshal(jobs)
	if err != nil {
		panic(err)
	}
	os.WriteFile(files.WORK, data, 0644)
}
