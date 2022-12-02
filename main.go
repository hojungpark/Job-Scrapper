package main

import "net/http"

var baseURL string = "https://ca.indeed.com/jobs?q=software+developer&vjk=36dbc4d045017359&limit=50"

func main(){
	pages := getPages()
}

func getPages() int{
	res, err := http.Get(baseURL)
	return 0
}