package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/montanaflynn/stats"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

const (
	P99     = "p99"
	P97     = "p97"
	P95     = "p95"
	MAX     = "max"
	MIN     = "min"
	MEDIAN  = "median"
	AVERAGE = "average"
	STDDEV  = "stddev"
	SUM     = "sum"
	COUNT   = "count"
)

var RENDER_FORMATS = map[string]string{
	P99:     "%.4f",
	P97:     "%.4f",
	P95:     "%.4f",
	MAX:     "%.4f",
	MIN:     "%.4f",
	MEDIAN:  "%.4f",
	AVERAGE: "%.4f",
	STDDEV:  "%.4f",
	SUM:     "%.0f",
	COUNT:   "%.0f",
}

var DEFAULT_FIELDS []string = []string{COUNT, SUM, P99, P97, P95, MIN, MAX, AVERAGE, MEDIAN, STDDEV}

func NewStatistics(data []float64) map[string]float64 {
	var statistics map[string]float64 = make(map[string]float64)
	statistics[P99], _ = stats.Percentile(data, 99.0)
	statistics[P97], _ = stats.Percentile(data, 97.0)
	statistics[P95], _ = stats.Percentile(data, 95.0)
	statistics[MAX], _ = stats.Max(data)
	statistics[MIN], _ = stats.Min(data)
	statistics[MEDIAN], _ = stats.Median(data)
	statistics[AVERAGE], _ = stats.Mean(data)
	statistics[STDDEV], _ = stats.StandardDeviation(data)
	statistics[SUM], _ = stats.Sum(data)
	statistics[COUNT] = float64(len(data))
	return statistics
}

func Render(s map[string]float64, fields []string, valuesOnly bool) {
	if !valuesOnly {
		fmt.Println(strings.Join(fields, "\t"))
	}

	for i, field := range fields {
		if value, ok := s[field]; ok {
			fmt.Printf(RENDER_FORMATS[field], value)

			if i < len(fields)-1 {
				fmt.Printf("\t")
			}

		} else {
			panic("Invalid field in output list: " + field)
		}
	}
	fmt.Println()
}

func execute(c *cli.Context) {
	r, err := getInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	statistics := NewStatistics(r)

	output := c.StringSlice("output")
	if len(output) == 0 {
		Render(statistics, DEFAULT_FIELDS, c.Bool("values-only"))
	} else {
		Render(statistics, output, true)
	}
}

func getInput() ([]float64, error) {
	results := make([]float64, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var item string = strings.TrimSpace(scanner.Text())

		if item == "" {
			continue
		}

		f, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, errors.New("Invalid value provided, item must be a numeric value")
		}
		results = append(results, f)
	}

	return results, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "stats"
	app.Usage = "Outputs statistical information about line delimited numbers from stdin"
	app.Action = execute

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "values-only",
			Usage: "only output values (no header)",
		},
		cli.StringSliceFlag{
			Name:  "output",
			Usage: "(repeated) statistic to output (valid items are [count, sum, p99, p97, p95, min, max, avg, median, stddev])",
		},
	}

	app.Run(os.Args)
}
