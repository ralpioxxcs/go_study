package main

import (
	"fmt"
	"time"
  "strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

var frame []byte

type camera struct {
  device *gocv.VideoCapture
  isStop bool
}
type framesize struct {
  width string  `form:"width"`
  height string `form:"hegith"`
}
var cam camera

func main() {
	router := gin.Default()

	// Play button
	router.POST("/Play", func(c *gin.Context) {
    var err error 
    cam.device, err = gocv.VideoCaptureDevice(0)
		if err != nil {
			c.String(http.StatusOK, "Can't open webcam")
			return
		}
		cam.isStop = false

    fmt.Printf("Frame width  : %f\n", cam.device.Get(gocv.VideoCaptureFrameWidth))
    fmt.Printf("Frame height : %f\n", cam.device.Get(gocv.VideoCaptureFrameHeight))

		go stream(cam.device)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Stop button
	router.POST("/Stop", func(c *gin.Context) {
		cam.isStop = true
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
  // Video - property
  router.GET("/property/brightness", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
  })
  // Video - property (frame size)
  router.GET("/property/framesize", func(c *gin.Context) {
    w, _ := strconv.ParseFloat( c.Query("width"), 64)
    h , _ := strconv.ParseFloat( c.Query("height"), 64)
    fmt.Printf("width : %f\n", w)
    fmt.Printf("height : %f\n", h)
    cam.device.Set(gocv.VideoCaptureFrameWidth, w)
    cam.device.Set(gocv.VideoCaptureFrameHeight, h)
		c.Redirect(http.StatusMovedPermanently, "/")
  })

  // Index
  router.LoadHTMLGlob("public/*")
	router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
	})


  router.Run(":8080")
}

func stream(cap *gocv.VideoCapture) {
	img := gocv.NewMat()
	frameNum := 0
	for {
		if cam.isStop {
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
