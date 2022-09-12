package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/alipourhabibi/work-progress/chart"
	"github.com/alipourhabibi/work-progress/files"
)

type job struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amont       float64 `json:"amount"`
	Time        string  `json:"time"`
}

func NewJob(name string, description string, amount float64, time string) *job {
	return &job{
		Name:        name,
		Description: description,
		Amont:       amount,
		Time:        time,
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

func (j *job) Get() {
	jobs := []job{}
	data, err := os.ReadFile(files.WORK)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &jobs)

	showJobs := []job{}

	if j.Name != "" && j.Time != "" {
		for _, v := range jobs {
			match, err := regexp.MatchString(j.Time, v.Time)
			if err != nil {
				panic(err)
			}
			if v.Name == j.Name && match {
				showJobs = append(showJobs, v)
			}
		}
	} else if j.Name == "" && j.Time != "" {
		for _, v := range jobs {
			match, err := regexp.MatchString(j.Time, v.Time)
			if err != nil {
				panic(err)
			}
			if match {
				showJobs = append(showJobs, v)
			}
		}
	} else if j.Name != "" && j.Time == "" {
		for _, v := range jobs {
			if v.Name == j.Name {
				showJobs = append(showJobs, v)
			}
		}

	} else {
		for _, v := range jobs {
			showJobs = append(showJobs, v)
		}
	}

	// show datas in table
	id := 1
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "AMOUNT"},
			{Align: simpletable.AlignCenter, Text: "TIME"},
		},
	}
	for _, row := range showJobs {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", id)},
			{Text: row.Name},
			{Text: row.Description},
			{Text: fmt.Sprintf("%f", row.Amont)},
			{Align: simpletable.AlignRight, Text: row.Time},
		}
		id += 1

		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func (j *job) Draw(chartName string, port int) {
	jobs := []job{}
	data, err := os.ReadFile(files.WORK)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &jobs)

	showJobs := []job{}

	if j.Name != "" && j.Time != "" {
		for _, v := range jobs {
			match, err := regexp.MatchString(j.Time, v.Time)
			if err != nil {
				panic(err)
			}
			if v.Name == j.Name && match {
				showJobs = append(showJobs, v)
			}
		}
	} else if j.Name == "" && j.Time != "" {
		for _, v := range jobs {
			match, err := regexp.MatchString(j.Time, v.Time)
			if err != nil {
				panic(err)
			}
			if match {
				showJobs = append(showJobs, v)
			}
		}
	} else if j.Name != "" && j.Time == "" {
		for _, v := range jobs {
			if v.Name == j.Name {
				showJobs = append(showJobs, v)
			}
		}

	}

	inInterface := []map[string]interface{}{}
	inRec, err := json.Marshal(showJobs)
	json.Unmarshal(inRec, &inInterface)

	switch chartName {
	case "bar":
		chart.BarChart(inInterface, port)
	}

}
