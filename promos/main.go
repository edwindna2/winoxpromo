package main

import (
	"gihub.com/dna2/promos/cmd"
)

func main()  {
	cmd.Start()

	//Crontab
	/*c := cron.New()
	//c.AddFunc("0 50 07 * * *",func() {cmd.Start()})
	c.AddFunc("00 08 19 * * *",func() {
		log.Println("Start scrapping")
		//cmd.Start()
	})
	c.Start()
	forever := make(chan struct{})
	<-forever*/
}