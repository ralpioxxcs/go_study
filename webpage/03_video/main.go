package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"html/template"
	"net/http"
	"os"
	"time"
)

var frame []byte

func main() {
	if len(os.Args) < 2 {
		fmt.Println("path is empty!")
		os.Exit(1)
	}
	fmt.Println("video path = " + os.Args[1])

	webcam, err := gocv.VideoCaptureFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer webcam.Close()

	fmt.Println("Video is streaming after 3seconds... (localhost:8080)")
	time.Sleep(1 * time.Second)

	go stream(webcam)

	http.HandleFunc("/video", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

		data := ""
		for {
			data = "--frame\r\n Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"
			time.Sleep(100 * time.Millisecond)
			res.Write([]byte(data))
		}
	})
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
		time.Sleep(100 * time.Millisecond)

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
