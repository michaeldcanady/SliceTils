package main

import(
  "fmt"

  "github.com/michaeldcanady/SliceTils/Matrix"
)

func main(){
  matr1 := [][]string{{"a","b","c"},{"d","e","f"},{"g","h","i"}}
  matr := matrix.EnterMatrix(matr1)
  fmt.Println(matr)
  matr.AppendCol([]string{"j","k","l"})
  fmt.Println(matr)
  matr.AppendRow(map[int]string{0:"w",1:"z",2:"y",3:"n"})
  fmt.Println(matr)
  matr.Remove(2)
  fmt.Println(matr)
}
