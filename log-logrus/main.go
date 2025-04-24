package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//LOG - package
	fmt.Println("Testing Log")
	log.Println("Log") //by default prints to the standard error stream

	//LOG - to a file
	file, err := os.OpenFile("logs_output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Testing Log")

	//LOG - custom loggers
	var WarningLogger, InfoLogger, ErrorLogger *log.Logger

	WarningLogger = log.New(file, "WARINING: ", log.Ldate|log.Ltime|log.Lshortfile) //create custom logger type
	WarningLogger.Println("Testin WarningLogger")

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Println("Testin InfoLogger")

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger.Println("Testing ErrorLogger")

	//LOGRUS
}
