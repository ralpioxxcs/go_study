package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"net/http"
	//"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tsak/concurrent-csv-writer"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

/*
				       [main]

				     [getpages]
  [getpage] [getpage] [getpage]
	[][][][]  [][][][]   [][][][]

*/

func main() {
	var jobs []extractedJob

	c := make(chan []extractedJob)

	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < 1; i++ {
		go getPage(i, c)
		// jobs = append(jobs, extractedJobs...) // meaning merge two arrays -> [] + [] => [ []+[] ]
		// jobs = append(jobs, extractedJobs) // meaning merge two arrays -> [] + []
	}

	for i := 0; i < 1; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	// !! code challenge - change writejobs func to go routines
	// -> change csv writer supporting concurrency!!
	writeJobs(jobs)

	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob

	c := make(chan extractedJob)

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c //receive
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())

	// sends to channel
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}

func writeJobs(jobs []extractedJob) {

	//file, err := os.Create("jobs.csv")
	//checkErr(err)

	//w := csv.NewWriter(file)
	//defer w.Flush()

	//headers := []string{"ID", "Title", "Location", "Salary", "Summary"}

	//wErr := w.Write(headers)
	//checkErr(wErr)

	//for _, job := range jobs {
	//  jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
	//  jwErr := w.Write(jobSlice)
	//  checkErr(jwErr)
	//}

	//###################################################
	// go routines
	//###################################################

	w, _ := ccsv.NewCsvWriter("jobs.csv")
	defer w.Close()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	done := make(chan bool)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		go writeFile(w, jobSlice, done)
	}

	for i := 0; i < len(jobs); i++ {
		<-done
	}

}

func writeFile(w *ccsv.CsvWriter, jobSlice []string, done chan<- bool) {
	jwErr := w.Write(jobSlice)
	checkErr(jwErr)
	done <- true
}

func cleanString(str string) string {
	// hello    f.  1
	// "hello","f","1"
	// hello f 1
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
