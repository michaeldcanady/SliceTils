package matrix

import(
  "fmt"
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

func (I *IntMatrix) Remove(col int){
  S1 := *I
  S1 = append(S1[:col], S1[col+1:]...)
  *I = S1
}

func keysInt(mymap map[int]int)[]int{
  keys := make([]int, 0, len(mymap))
  for k := range mymap {
    keys = append(keys, k)
  }
  return keys
}
