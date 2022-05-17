package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "exec",
				Aliases:  []string{"e"},
				Required: true,
			},
			&cli.IntFlag{
				Name:    "interval",
				Aliases: []string{"i"},
				Value:   60,
				Usage:   "interval in seconds",
			},
		},
		Action: func(context *cli.Context) error {
			execStr := context.String("exec")
			interval := context.Int("interval")
			if context.Args().Len() == 0 {
				return fmt.Errorf("please specify at least 1 url for monitor")
			}
			for {
				serverIsUp := true
				for _, v := range context.Args().Slice() {
					log.Println("GET " + v)
					res, err := http.Get(v)
					if err != nil {
						serverIsUp = false
						cmd := exec.Command("bash", "-c", execStr)
						err := cmd.Run()
						if err != nil {
							log.Println(err.Error())
						}
						continue
					}
					if res.StatusCode >= 400 {
						serverIsUp = false
						cmd := exec.Command("bash", "-c", execStr)
						err := cmd.Run()
						if err != nil {
							log.Println(err.Error())
						}
						continue
					}
				}
				if serverIsUp {
					log.Println("server is UP !")
				} else {
					log.Println("server is DOWN !")
				}
				log.Printf("sleeping %d sec...\n", interval)
				time.Sleep(time.Second * time.Duration(interval))
			}
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
