package main

import(
  "fmt"
  "image"
  "image/color"
  "os"
  "strconv"

  "gocv.io/x/gocv"
)

func main() {
  fmt.Println("running cv example ...")

  if len(os.Args) < 3 {
    fmt.Println("How to run:\n\tfacedetect [camID] [XML]")
    return
  }


  deviceID, _ := strconv.Atoi(os.Args[1])
  xmlFile := os.Args[2]

  webcamm, err := gocv.VideoCaptureDevice(int(deviceID))
  if err != nil {
    fmt.Println(err)
    return
  }
  defer webcamm.Close()

  window := gocv.NewWindow("Face Detect")
  defer window.Close()

  img := gocv.NewMat();
  defer img.Close()

  blue := color.RGBA{0,0,255,0}

  classifier := gocv.NewCascadeClassifier()
  defer classifier.Close()

  if !classifier.Load(xmlFile) {
    fmt.Printf("Error reading cacade file: %v\n", xmlFile)
    return
  }

  fmt.Println("start reading camera device : %v\n", deviceID)
  for {
    if ok := webcamm.Read(&img); !ok {
      fmt.Printf("can't read device %d\n", deviceID)
      return
    }
    if img.Empty() {
      continue
    }

    rects := classifier.DetectMultiScale(img)
    fmt.Printf("found %d faces\n", len(rects))

    for _, r := range rects {
      gocv.Rectangle(&img, r, blue, 3)

      size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
      pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2),r.Min.Y-2)
      gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
    }

    window.IMShow(img)
    if window.WaitKey(1) >= 0 {
      break
    }

  }
}
