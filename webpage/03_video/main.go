package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

var frame []byte
var stop bool

func main() {
	fmt.Println("main()")

	router := gin.Default()

	// Play button
	router.POST("/Play", func(c *gin.Context) {
		webcam, err := gocv.VideoCaptureDevice(0)
		if err != nil {
			c.String(http.StatusOK, "Can't open webcam")
			return
		}
		stop = false
		go stream(webcam)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Stop button
	router.POST("/Stop", func(c *gin.Context) {
		stop = true
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Video
	router.GET("/video", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
		data := ""
		for {
			data = "--frame\r\n Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"
			time.Sleep(100 * time.Millisecond)
			c.Writer.Write([]byte(data))
		}
	})

  router.LoadHTMLGlob("html/*")
	router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
	})

  router.Run(":8080")

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
