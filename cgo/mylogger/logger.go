package mylogger

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var mlogger *logger
var once sync.Once

// start Log
func GetInstance() *logger {
	once.Do(func() {
		mlogger = createLogger("logfile.log")
	})
	return mlogger
}

func createLogger(fname string) *logger {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		panic(err)
	}
	//defer file.Close()

	loggerr := logger{
		filename: fname,
		Logger:   log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	log.SetOutput(file)

	loggerr.Println("test")

	return &loggerr

	// return &logger{
	// 	filename: fname,
	// 	Logger:   log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	// }
}
