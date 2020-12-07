package matrix

import(
  "fmt"

  "github.com/michaeldcanady/SliceTils/SliceTils"
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

func (S *StringMatrix) RemoveCol(col int){
  S1 := *S
  S1 = append(S1[:col], S1[col+1:]...)
  *S = S1
}

func (S *StringMatrix) RemoveRow(row int){
  var NewMatrix StringMatrix
  for _,col := range (*S){
    newcol := col
    newcol =  append(newcol[:row], newcol[row+1:]...)
    col = newcol
    NewMatrix = append(NewMatrix,col)

  }
  *S = NewMatrix

}

func (S *StringMatrix) Any(value []string)bool{
  for _,slices := range (*S){
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

func (S *StringMatrix) All(test interface{})(bool,[][]int){
  value := []string{}
  switch test.(type){
  case [][]string:
    value = slicetils.TwoToOne(test).([]string)
  case []string:
    value = test.([]string)
  }
  contained := []string{}
  temp := [][]int{}
  for _,v1 := range value{
    for i,slices := range (*S){
      for i1,v := range slices{
        if v == v1{
          contained = append(contained,v1)
          temp = append(temp,[]int{i,i1})
        }
      }
    }
  }
  if slicetils.Equal(contained,value){
    return true,temp
  }
  return false,[][]int{}
}

func (S *StringMatrix) Sub(value [][]string)bool{
  
  return false
}

func (S *StringMatrix) SubB(value [][]string)bool{
  return false
}

func keysStr(mymap map[int]string)[]int{
  keys := make([]int, 0, len(mymap))
  for k := range mymap {
    keys = append(keys, k)
  }
  return keys
}
