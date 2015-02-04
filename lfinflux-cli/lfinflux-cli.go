package main

import (
	//"bufio"
	//"os"
	//"fmt"
	"github.com/influxdb/influxdb/client"
	"log"
)

func main() {
	cli, err := client.NewClient(&client.ClientConfig{
		Host:     "127.0.0.1:8086",
		Database: "test",
		Username: "root",
		Password: "root",
	})

	if err != nil {
		log.Fatal(err)
	}

	series := &client.Series{
		Name:    "test",
		Columns: []string{"time", "value"},
		Points: [][]interface{}{
			[]interface{}{1394398326, 7},
		},
	}

	if err := cli.WriteSeriesWithTimePrecision([]*client.Series{series}, "s"); err != nil {
		log.Fatal(err)
	}

	/*
	   scanner := bufio.NewScanner(os.Stdin)

	   for scanner.Scan() {
	       fmt.Println(scanner.Text())
	   }

	   if err := scanner.Err(); err != nil {
	       log.Fatal(err)
	   }
	*/
}
