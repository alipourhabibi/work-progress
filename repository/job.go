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

func NewJob(name string, description string, amount float64, time string) *job {
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

func (j *job) Modify(newName string) {
	dt := time.Now()
	j.Time = dt.Format("01-02-2006")
	jobs := []job{}
	data, err := os.ReadFile(files.WORK)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &jobs)

	for k, v := range jobs {
		if v.Name == j.Name && v.Time == j.Time {
			if j.Description != "" {
				jobs[k].Description = j.Description
			}
			if j.Amont != -1 {
				jobs[k].Amont = j.Amont
			}
			if newName != "" {
				jobs[k].Name = newName
			}
			break
		}
	}
	data, err = json.Marshal(jobs)
	if err != nil {
		panic(err)
	}
	os.WriteFile(files.WORK, data, 0644)
}

func (j *job) Delete() {
	jobs := []job{}
	data, err := os.ReadFile(files.WORK)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &jobs)

	newJobs := []job{}

	for _, v := range jobs {
		if v.Name != j.Name && v.Time != j.Time {
			newJobs = append(newJobs, v)
		}
	}
	data, err = json.Marshal(newJobs)
	if err != nil {
		panic(err)
	}
	os.WriteFile(files.WORK, data, 0644)
}
