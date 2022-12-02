package main

import (
	"log"
	"net/http"
)

var baseURL string = "https://ca.indeed.com/jobs?q=software+developer&vjk=36dbc4d045017359&limit=50"

func main(){
	pages := getPages()
}

func getPages() int{
	resp, err := http.Get(baseURL)

	checkErr(err)
	checkStatus(resp)
	if resp.StatusCode != 200{
		log.Fatalln("Request failed with status:", resp.Status)
	}
	return 0
}

func checkErr(err error){
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(resp *http.Response){
	if resp.StatusCode != 200{
		log.Fatalln("Request failed with status:", resp.Status)
	}
}