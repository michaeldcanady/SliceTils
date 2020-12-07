package matrix

import(
  "fmt"
)

type StringMatrix [][]string

func (S StringMatrix) String()string{
  var print string
  for _,slice := range S{
    print += fmt.Sprintf("%v\n",slice)
  }
  return print
}

func EnterMatrix(slice [][]string)StringMatrix{
  return StringMatrix(slice)
}

func (S *StringMatrix) AppendCol(slice[]string){
  *S = append(*S,slice)
}

func (S *StringMatrix) AppendRow(dat map[int]string){
  keySlice := keysStr(dat)
  if len(dat) != len((*S)){
    panic("All columns must be appended to equal lengths")
  }

  for _,key := range keySlice{
    (*S)[key] = append((*S)[key],dat[key])
  }
}

func (S *StringMatrix) Remove(col int){
  S1 := *S
  S1 = append(S1[:col], S1[col+1:]...)
  *S = S1
}

func keysStr(mymap map[int]string)[]int{
  keys := make([]int, 0, len(mymap))
  for k := range mymap {
    keys = append(keys, k)
  }
  return keys
}
