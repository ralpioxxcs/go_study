package main

import (
	"fmt"
  "image"

	"gocv.io/x/gocv"
)

func open(id int) (vc* gocv.VideoCapture, err error) {
  vc, err = gocv.VideoCaptureDevice(id)
  return vc, err
}

func main() {

  fmt.Println("Video capturing")

  //webcam, _ := open(0)

  webcam, err := gocv.VideoCaptureDevice(0)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  defer webcam.Close()

  window := gocv.NewWindow("image")
  window.ResizeWindow(640, 480)
  //defer window.Close()

  img := gocv.NewMat()
  
  for {
    if ok := webcam.Read(&img); !ok {
      fmt.Sprintln("can't read camera device")
      return
    }
    if img.Empty() {
      continue
    }

    gocv.Resize(img, &img, image.Point{0, 0}, 0.5, 0.5, gocv.InterpolationArea)

    window.IMShow(img)
    if window.WaitKey(1) >= 0 {
      fmt.Println("image show break")
      break
    }
  }

}
