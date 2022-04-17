package chart

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func BarChart(data []map[string]interface{}, port int) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Activities",
		Subtitle: "Based on hour",
	}))

	items := make([]opts.BarData, 0)
	xAxis := []string{}
	for _, v := range data {
		xAxis = append(xAxis, v["name"].(string))
		items = append(items, opts.BarData{Value: v["amount"].(float64)})
	}
	bar.SetXAxis(xAxis).AddSeries("day 1", items)

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		bar.Render(w)
	})
	fmt.Printf("Open brower and go to http://localhost:%d/ to see chart\n", port)
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, nil)
}
