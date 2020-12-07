package slicetils

import(
  "strconv"
  "sort"
  "math/rand"
)

var(
  Alpha = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
)

// Finds the lowest letter contained in the slice
func minStr(array []string)string{
  var min string = Alpha[25]
  for _, value := range array {
      if indexOfStr(min,Alpha) > indexOfStr(value,Alpha) {
          min = value
      }
  }
  return min
}

// Finds the Highest letter contained in the slice
func maxStr(array []string)string{
  var max string = Alpha[0]
  for _, value := range array {
      if indexOfStr(max,Alpha) < indexOfStr(value,Alpha) {
          max = value
      }
  }
  return max
}

// returns a slice of i size
func nthRandStr(i int)[]string{
  b := []string{}
  for j := 0; j < i; j++ {
    b = append(b, Alpha[rand.Intn(len(Alpha))])
  }
  return b
}

// gets the indext of an element in a slice
func indexOfStr(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1    //not found.
}

// Checks if the two slices are equal
func equalStr(a, b []string) bool {
  sort.Strings(b)
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

// Removes duplicate values from an string slice
func removeDupStr(intSlice *[]string){
    keys := make(map[string]bool)
    list := []string{}

    // If the key(values of the slice) is not equal
    // to the already present value in new slice (list)
    // then we append it. else we jump on another element.
    inSlice := *intSlice
    for _, entry := range inSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    *intSlice = inSlice
}

// Converts a 2D slice into a 1D slice
func twoToOneStr(input [][]string)[]string{
  NewSlice := []string{}
  for _,slice := range input{
    for _,s := range slice{
      NewSlice = append(NewSlice,s)
    }
  }
  return NewSlice
}

// Converts a string slice to an int slice
func stringToInt(Slice []string)[]int{
  NewSlice := []int{}
  for _,i := range Slice{
    new,_ := strconv.Atoi(i)
    NewSlice = append(NewSlice,new)
  }
  return NewSlice
}

// Returns if a slice contains a value and its index
func containStr(s []string, e string)(bool,int) {
    for i, a := range s {
        if a == e {
            return true,i
        }
    }
    return false,-1
}
