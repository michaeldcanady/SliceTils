package slicetils

import(
  "strconv"
  "sort"
  "math/rand"
)

// Gets the minimum value from a slice
func minInt(array []int) int {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min
}

// Gets the maxium value from a slice
func maxInt(array []int) int {
    var max int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
    }
    return max
}

func subsliceInt(s1 []int, s2 []int) bool {
    if len(s1) > len(s2) { return false }
    for _, e := range s1 {
      b,_ := containInt(s2,e)
        if !b{
            return false
        }
    }
    return true
}

// i is the number of values needed, r is the range of values needed
func nthRandInt(i int,r int)[]int{
  s := []int{}
  for a := 1; a <= i; a++ {
    s = append(s,rand.Intn(r))
  }
  return s
}

func equalInt(a, b []int) bool {
  sort.Ints(b)
  sort.Ints(a)
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

func indexOfInt(element int, data []int) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1    //not found.
}

// Removes duplicate values from an int slice
func removeDupInt(intSlice []int) []int {
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

// Converts a 2D slice into a 1D slice
func twoToOneInt(input [][]int)[]int{
  NewSlice := []int{}
  for _,slice := range input{
    for _,s := range slice{
      NewSlice = append(NewSlice,s)
    }
  }
  return NewSlice
}

// Converts a int slice to a string slice
func intToString(Slice []int)[]string{
  NewSlice := []string{}
  for _,i := range Slice{
    NewSlice = append(NewSlice,strconv.Itoa(i))
  }
  return NewSlice
}

// Returns if a slice contains a value and its index
func containInt(s []int, e int)(bool,int) {
    for i, a := range s {
        if a == e {
            return true,i
        }
    }
    return false,-1
}
