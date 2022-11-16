/*
Copyright © 2022 NAME HERE <hernino25@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  "Trace the IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showData(ip)
			}

		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"location"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("unable to unmarshal response")
	}
	fmt.Println("DATA FOUND: ")
	//fmt.Println(data)
	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE :%s\nPOSTAL :%s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)

}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("unable to get response")
	}
	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("unable to read the response")
	}
	return responseByte
}
