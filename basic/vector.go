package main 

import "fmt"

// x y z values of 3D vector
type vector struct {
  X,Y,Z float64
}

func (a vector) Add(b vector) vector {
  return vector{
    X: a.X + b.X,
    Y: a.Y + b.Y,
    Z: a.Z + b.Z,
  }
}

func (a vector) Sub(b vector) vector {
  return vector{
    X: a.X - b.X,
    Y: a.Y - b.Y,
    Z: a.Z - b.Z,
  }
}

func (a vector) MultiplyScalar(s float64) vector {
  return vector{
    X: a.X * s,
    Y: a.Y * s,
    Z: a.Z * s,
  }
}

func (a vector) Dot(b vector) float64{
  return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func (a vector) Cross(b vector) vector {
  return vector{
    X: a.Y*b.Z - a.Z*b.Y,
    Y: a.Z*b.X - a.X*b.Z,
    Z: a.X*b.Y - a.Y*b.X,
  }
}

func main() {
  fmt.Println("main")

  result := vector{1,0,0}.Cross(vector{0,1,0})

  fmt.Println(result)
}


