package main

import(
  "fmt"

  "github.com/michaeldcanady/SliceTils/Matrix"
)

func main(){
  matr1 := [][]string{{"a","b","c"},{"d","e","f"},{"g","h","i"}}

  test := []string{"a","b","w","a","d"}
  test2D := [][]string{{"e","z"},{"k","n"}}

  matr := matrix.EnterMatrix(matr1)
  fmt.Println(matr)
  matr.AppendCol([]string{"j","k","l"})
  fmt.Println(matr)
  matr.AppendRow(map[int]string{0:"w",1:"z",2:"y",3:"n"})
  fmt.Println(matr)
  matr.RemoveCol(2)
  fmt.Println(matr)
  matr.RemoveRow(2)
  fmt.Println(matr)
  fmt.Println(matr.Any(test))
  fmt.Println(matr.All(test))
  fmt.Println(matr.All(test2D))
  matr.Sub(test2D)
}
