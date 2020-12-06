package slicetils

import(
  "strconv"
  "fmt"
)

// Gets the minimum value from a slice
func Min(array []int) int {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min
}

// Gets the maxium value from a slice
func Max(array []int) int {
    var max int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
    }
    return max
}

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

func equalStr(a, b []string) string {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func equalInt(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

// Removes duplicate values from an int slice
func removeDuplicateValues(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}

    // If the key(values of the slice) is not equal
    // to the already present value in new slice (list)
    // then we append it. else we jump on another element.
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

// Converts a int slice to a string slice
func IntToString(Slice []int)[]string{
  NewSlice := []string{}
  for _,i := range Slice{
    NewSlice = append(NewSlice,strconv.Itoa(i))
  }
  return NewSlice
}

// Returns if a slice contains a value and its index
func Contains(s []string, e string)(bool,int) {
    for i, a := range s {
        if a == e {
            return true,i
        }
    }
    return false,-1
}
