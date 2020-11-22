package main

import (
	"fmt"
	"html/template"
	"net/http"

	"gocv.io/x/gocv"

	"time"
)

var frame []byte
var stop bool

func main() {
  fmt.Println("Start")
	/*
		if len(os.Args) < 2 {
			fmt.Println("path is empty,, open webcam")
			os.Exit(1)
		}
		fmt.Println("video path = " + os.Args[1])

		webcam, err := gocv.VideoCaptureFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer webcam.Close()
	*/

  // Set http server & handling functions
  // Play buttons
  http.HandleFunc("/Play", func(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Start to streaming webcam")
    webcam, err := gocv.VideoCaptureDevice(0)
    if err != nil {
      fmt.Fprint(res, "can't open device")
      return
    }

    stop = false
	  go stream(webcam)
    http.Redirect(res, req, "/", http.StatusMovedPermanently)
  })
  // Stop buttons
  http.HandleFunc("/Stop", func(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Stop to streaming webcam")
    stop = true
    http.Redirect(res, req, "/", http.StatusMovedPermanently)
  })
  // Video container
	http.HandleFunc("/video", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
		data := ""
		for {
			data = "--frame\r\n Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"
			time.Sleep(100 * time.Millisecond)
			res.Write([]byte(data))
		}
	})
  // Index
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(res, "index")
	})

	http.ListenAndServe(":8080", nil)
}


func stream(cap *gocv.VideoCapture) {
	img := gocv.NewMat()
	frameNum := 0
	for {
    if stop {
      return
    }

		//time.Sleep(10 * time.Millisecond)

		if !cap.Read(&img) {
			fmt.Println("video done")
			return
		}
		frameNum++
		if img.Empty() {
			continue
		}
		frame, _ = gocv.IMEncode(".jpg", img)
	}
}
