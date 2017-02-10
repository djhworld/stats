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

type Statistics struct {
	p99     float64
	p97     float64
	p95     float64
	max     float64
	min     float64
	average float64
	median  float64
	stddev  float64
	sum     float64
	count   int
}

func NewStatistics(data []float64) *Statistics {
	statistics := new(Statistics)
	p99, _ := stats.Percentile(data, 99.0)
	p97, _ := stats.Percentile(data, 97.0)
	p95, _ := stats.Percentile(data, 95.0)
	max, _ := stats.Max(data)
	min, _ := stats.Min(data)
	median, _ := stats.Median(data)
	avg, _ := stats.Mean(data)
	stddev, _ := stats.StandardDeviation(data)
	sum, _ := stats.Sum(data)
	statistics.p99 = p99
	statistics.p97 = p97
	statistics.p95 = p95
	statistics.max = max
	statistics.min = min
	statistics.average = avg
	statistics.median = median
	statistics.stddev = stddev
	statistics.sum = sum
	statistics.count = len(data)
	return statistics
}

func (s *Statistics) Render(valuesOnly bool) {
	if !valuesOnly {
		fmt.Printf("count\t")
		fmt.Printf("sum\t")
		fmt.Printf("p99\t")
		fmt.Printf("p97\t")
		fmt.Printf("p95\t")
		fmt.Printf("min\t")
		fmt.Printf("max\t")
		fmt.Printf("avg\t")
		fmt.Printf("median\t")
		fmt.Printf("stddev\n")
	}
	fmt.Printf("%d\t", s.count)
	fmt.Printf("%.4f\t", s.sum)
	fmt.Printf("%.4f\t", s.p99)
	fmt.Printf("%.4f\t", s.p97)
	fmt.Printf("%.4f\t", s.p95)
	fmt.Printf("%.4f\t", s.min)
	fmt.Printf("%.4f\t", s.max)
	fmt.Printf("%.4f\t", s.average)
	fmt.Printf("%.4f\t", s.median)
	fmt.Printf("%.4f\n", s.stddev)
}

func main() {
	app := cli.NewApp()
	app.Name = "stats"
	app.Usage = "Outputs statistical information about line delimited numbers from stdin"
	app.Action = execute

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "values-only",
			Usage: "only output values in the order [count, sum, p99, p97, p95, min, max, avg, median, stddev]",
		},
	}

	app.Run(os.Args)
}

func execute(c *cli.Context) {
	r, err := getInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	statistics := NewStatistics(r)
	statistics.Render(c.Bool("values-only"))
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
