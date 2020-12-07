package matrix

import(
  "fmt"

  "github.com/michaeldcanady/SliceTils/SliceTils"
)

type IntMatrix [][]int

func (I IntMatrix) String()string{
  var print string
  for _,slice := range I{
    print += fmt.Sprintf("%v\n",slice)
  }
  return print
}

func EnterMatrixInt(slice [][]string)StringMatrix{
  return StringMatrix(slice)
}

func (I *IntMatrix) AppendCol(slice[]int){
  *I = append(*I,slice)
}

func (I *IntMatrix) AppendRow(dat map[int]int){
  keySlice := keysInt(dat)
  if len(dat) != len((*I)){
    panic("All columns must be appended to equal lengths")
  }

  for _,key := range keySlice{
    (*I)[key] = append((*I)[key],dat[key])
  }
}

func (I *IntMatrix) RemoveCol(col int){
  S1 := *I
  S1 = append(S1[:col], S1[col+1:]...)
  *I = S1
}

func (I *IntMatrix) Any(value []int)bool{
  for _,slices := range (*I){
    for _,v := range slices{
      for _,v1 := range value{
        if v == v1{
          return true
        }
      }
    }
  }
  return false
}

func (I *IntMatrix) All(value []int)bool{
  contained := []int{}
  for _,v1 := range value{
    for _,slices := range (*I){
      for _,v := range slices{
        if v == v1{
          contained = append(contained,v1)
        }
      }
    }
  }
  if slicetils.Equal(contained,value){
    return true
  }
  return false
}

func (I *IntMatrix) RemoveRow(row int){
  var NewMatrix IntMatrix
  for _,col := range (*I){
    newcol := col
    newcol =  append(newcol[:row], newcol[row+1:]...)
    col = newcol
    NewMatrix = append(NewMatrix,col)

  }
  *I = NewMatrix
}

func keysInt(mymap map[int]int)[]int{
  keys := make([]int, 0, len(mymap))
  for k := range mymap {
    keys = append(keys, k)
  }
  return keys
}
