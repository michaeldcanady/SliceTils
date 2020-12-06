package slicetils

import(
  "fmt"
)


// Checks if two slices are equal
func Equal(a, b interface{}) bool {
  switch v := a.(type) {
  case []int:
    return equalInt(a.([]int),b.([]int))
  case []string:
    return equalStr(a.([]string),b.([]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return false
}

//func RemoveDuplicateValues(a interface{}){
//  switch v := a.(type) {
//  case []int:
//    return removeDuplicateValuesInt(a.([]int))
//  case []string:
//    return removeDuplicateValuesStr(a.([]string))
//  default:
//    fmt.Errorf("%v is not a valid type",v)
//  }
//  return ""
//}
