package repository

import (
	"encoding/json"
	"os"
	"time"

	"github.com/alipourhabibi/work-progress/files"
)

type job struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amont       float64 `json:"amound"`
	Time        string  `json:"time"`
}

func NewJob(name string, description string, amount float64) *job {
	return &job{
		Name:        name,
		Description: description,
		Amont:       amount,
	}
}

func (j *job) Add() {
	dt := time.Now()
	j.Time = dt.Format("01-02-2006")
	// load jobs to a variable
	jobs := []job{}
	data, err := os.ReadFile(files.WORK)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &jobs)

	// check if already exists so add them up
	exists := false
	for k, v := range jobs {
		if v.Name == j.Name && v.Time == j.Time {
			jobs[k].Amont += j.Amont
			exists = true
		}
	}
	if !exists {
		jobs = append(jobs, *j)
	}
	data, err = json.Marshal(jobs)
	if err != nil {
		panic(err)
	}
	os.WriteFile(files.WORK, data, 0644)
}
