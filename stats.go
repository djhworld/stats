package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
	"github.com/urfave/cli"
)

const (
	P99     = "p99"
	P97     = "p97"
	P95     = "p95"
	P90     = "p90"
	P75     = "p75"
	P50     = "p50"
	P25     = "p25"
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
	P90:     "%.4f",
	P75:     "%.4f",
	P50:     "%.4f",
	P25:     "%.4f",
	MAX:     "%.4f",
	MIN:     "%.4f",
	MEDIAN:  "%.4f",
	AVERAGE: "%.4f",
	STDDEV:  "%.4f",
	SUM:     "%.0f",
	COUNT:   "%.0f",
}

var DEFAULT_FIELDS []string = []string{COUNT, SUM, P99, P97, P95, P90, P75, P50, P25, MIN, MAX, AVERAGE, MEDIAN, STDDEV}

func NewStatistics(data []float64) map[string]float64 {
	var statistics map[string]float64 = make(map[string]float64)
	statistics[P99], _ = stats.Percentile(data, 99.0)
	statistics[P97], _ = stats.Percentile(data, 97.0)
	statistics[P95], _ = stats.Percentile(data, 95.0)
	statistics[P90], _ = stats.Percentile(data, 90.0)
	statistics[P75], _ = stats.Percentile(data, 75.0)
	statistics[P50], _ = stats.Percentile(data, 50.0)
	statistics[P25], _ = stats.Percentile(data, 25.0)
	statistics[MAX], _ = stats.Max(data)
	statistics[MIN], _ = stats.Min(data)
	statistics[MEDIAN], _ = stats.Median(data)
	statistics[AVERAGE], _ = stats.Mean(data)
	statistics[STDDEV], _ = stats.StandardDeviation(data)
	statistics[SUM], _ = stats.Sum(data)
	statistics[COUNT] = float64(len(data))
	return statistics
}

func Render(s map[string]float64, fields []string, valuesOnly bool, delimiter string) {
	for _, field := range fields {
		if value, ok := s[field]; ok {
			if !valuesOnly {
				fmt.Print(field + delimiter)
			}
			fmt.Printf(RENDER_FORMATS[field], value)
			fmt.Println()
		} else {
			panic("Invalid field in output list: " + field)
		}
	}
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
		Render(statistics, DEFAULT_FIELDS, c.Bool("values-only"), c.String("delimiter"))
	} else {
		Render(statistics, output, true, c.String("delimiter"))
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
			Usage: "statistic to output (valid items are [count, sum, p99, p97, p95, p90, p75, p50, p25, min, max, avg, median, stddev])",
		},
		cli.StringFlag{
			Name:  "delimiter",
			Usage: "output delimiter",
			Value: "\t",
		},
	}

	app.Run(os.Args)
}
