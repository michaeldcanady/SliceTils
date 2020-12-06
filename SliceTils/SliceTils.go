package slicetils

import(
  "fmt"
)

func Min(a interface{})interface{}{
  switch v := a.(type){
  case []int:
    return minInt(a.([]int))
  case []string:
    return minStr(a.([]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return false
}

//Returns an interface to allow for an infinite number of return types
//Return must be asserted as expected type to be used.
func Max(a interface{})interface{}{
  switch v := a.(type){
  case []int:
    return maxInt(a.([]int))
  case []string:
    return maxStr(a.([]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return false
}

// Checks if two slices are equal
func Equal(a, b interface{})bool{
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

func NthRand(i int,t string,r int)interface{}{
  switch t{
  case "int":
    return nthRandInt(i,r)
  case "string":
    return nthRandStr(i)
  default:
    fmt.Errorf("%v is not a valid type",t)
  }
  return false
}

func TwoToOne(a interface{})interface{}{
  switch v := a.(type) {
  case [][]int:
    return twoToOneInt(a.([][]int))
  case [][]string:
    return twoToOneStr(a.([][]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return false
}

func IndexOf(a interface{}, b interface{})(int){
  switch v := a.(type) {
  case []int:
    return indexOfInt(a.(int),b.([]int))
  case []string:
    return indexOfStr(a.(string),b.([]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return -1
}

func RemoveDup(a interface{})interface{}{
  switch v := a.(type) {
  case []int:
    return removeDupInt(a.([]int))
  case []string:
    return removeDupStr(a.([]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return -1
}

func Convert(a interface{},t string)interface{}{
  switch t{
  case "int":
    return intToString(a.([]int))
  case "string":
    return stringToInt(a.([]string))
  default:
    fmt.Errorf("%v is not a valid type",t)
  }
  return -1
}

func Contains(a,b interface{})(bool,int){
  switch v := a.(type) {
  case []int:
    return containInt(a.([]int),b.(int))
  case []string:
    return containStr(a.([]string),b.(string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return false,-1
}
