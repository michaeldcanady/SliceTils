package slicetils

import(
  "fmt"
)

// Minimum := slicetils.Min([]int{}).([]int)
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

// Generates a random list of i size, type t, and if int or range r
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

// Converts a 2D slice to 1D
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

// Returns the index of a within b
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

func SubSlice(a interface{}, b interface{})(bool){
  switch v := a.(type) {
  case []int:
    return subsliceInt(a.([]int),b.([]int))
  case []string:
    return subsliceStr(a.([]string),b.([]string))
  default:
    fmt.Errorf("%v is not a valid type",v)
  }
  return false
}

// Removes all duplicate values from a
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

// converts slice a, as type t, to t1
// Figure out how to add designation type
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

// Checks if a contains elem b, returns true and index if it does, returns false and -1 if not.
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
