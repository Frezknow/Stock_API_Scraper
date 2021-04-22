package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	index, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(index),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
	scrapeForCSV()
}

//scrapeForCSV scrapes data into a CSV file
func scrapeForCSV() {
	link := "https://api.polygon.io/v2/reference/financials/AAPL?limit=30&type=Q&sort=calendarDate&apiKey=MBctIb6XJhtdvXZZRTasWYTbt2Qv0FX0"
	resp, err := http.Get(link)
	if err != nil {
		log.Panic(err)
	}
	createFile()
	log.Print(resp)
	fmt.Println(resp)
}

var path = "test.txt"

func createFile() {
	//check if file exist
	var _, err = os.Stat(path)
	//Create file if not exist
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Created SuccessFully", path)
}

func writeFile(resp) {
	//Open file using READ and WRITE permission
	var file, err = os.OpenFile(path, os.RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	//Write some text to file.
	_, err = file.WriteString("Hello \n")
	if isError(err) {
		return
	}
	//Save file changes
	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("File updated")

}
