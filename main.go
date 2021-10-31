package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/urfave/cli/v2"
)

func performLoadTesting(method string, url string, ratePerSecond int, durationInSeconds int) {
	var rate = vegeta.ConstantPacer{Freq: ratePerSecond, Per: 1 * time.Second}
	var duration = time.Duration(durationInSeconds) * time.Second
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	var targeter vegeta.Targeter
	if method == "GET" {
		targeter = loadTargetForGetApi(url)
	} else if method == "POST" {
		targeter = loadTestPostApi(url)
	}
	for res := range attacker.Attack(targeter, rate, duration, "Load Test") {
		metrics.Add(res)
	}
	metrics.Close()

	prettyJSON, _ := json.MarshalIndent(metrics, "", "\t")
	fmt.Printf("%s", string(prettyJSON))
}

func loadTargetForGetApi(url string) func(target *vegeta.Target) error {
	return func(target *vegeta.Target) error {
		target.Method = "GET"
		target.URL = url
		target.Header = http.Header{
			"Content-type": []string{"application/json"},
		}

		return nil
	}
}

func loadTestPostApi(url string) func(target *vegeta.Target) error {
	return func(target *vegeta.Target) error {
		target.Method = "POST"
		target.URL = url
		target.Header = http.Header{
			"Content-type": []string{"application/json"},
		}
		// body in string format
		return nil
	}
}

func main() {
	app := &cli.App{
		Name:            "Load Test",
		Usage:           "Load testing utility with vegeta",
		HideHelpCommand: false,
		Commands: []*cli.Command{
			{
				Name: "ldt",
				Description: "A simple utility which used vegeta \n" +
					"make build \n" +
					"./out/ldt ldt -m GET -u <url>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "method",
						Value:   "GET",
						Usage:   "Http Method",
						Aliases: []string{"m"},
					},
					&cli.StringFlag{
						Name:    "url",
						Value:   "http://localhost:8080/",
						Usage:   "URL",
						Aliases: []string{"u"},
					}, &cli.StringFlag{
						Name:    "rate_per_second",
						Value:   "10",
						Usage:   "Hit rate per Second",
						Aliases: []string{"r"},
					}, &cli.StringFlag{
						Name:    "duration",
						Value:   "10",
						Usage:   "Duration of the test in seconds",
						Aliases: []string{"d"},
					},
				},
				Action: func(c *cli.Context) error {
					performLoadTesting(c.String("method"), c.String("url"), c.Int("rate_per_second"), c.Int("duration"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%s is failed : %+v", os.Args[1:], err)
	}
}
