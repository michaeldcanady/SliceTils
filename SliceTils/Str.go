package slicetils

import(
  "strconv"
  "sort"
  "math/rand"
  "errors"
  "fmt"
)

var(
  Alpha = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
)

//Breaks main slice into subslices string
func sliceSpliterStr(slice []string, slicecount int)([][]string,error){
  var slices [][]string
  slicelen := len(slice)
  subSliceLen := (slicelen/slicecount)
  for{
    if subSliceLen < 1{
      difference := slicecount - slicelen
      slicecount = slicecount - difference
      subSliceLen = (slicelen/slicecount)
      return slices, errors.New(fmt.Sprintf("adjusts slicecount to %v due to, slice len < subslices.\n",slicecount))
    }else{
      break
    }
  }
  remainer := (slicelen%slicecount)
  n := 0
  for n < slicelen{
    var s []string
    t := n+subSliceLen
    if t != (slicelen-remainer){
      s = slice[n:t]
    }else{
      t = (t+remainer)
      s = slice[n:t]
    }
    slices = append(slices,s)
    n = t
  }
  return slices,nil
}

// Removes empty values from string slice
func removeEmptyStr(s []string) []string{
  var r []string
  for _, str := range s{
    if str != ""{
      r = append(r,str)
    }
  }
  return r
}

// Removes strings from a slice
func removeSliceStr(a []string, remove ...string)([]string,error){
  // Remove the element at index i from a.
  for _, s := range remove{
    index := IndexOf(s,a)
    if index != -1{
      copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
      a[len(a)-1] = ""     // Erase last element (write zero value).
      a = a[:len(a)-1]     // Truncate slice.
      a = removeEmptyStr(a)
    }else{
      return a, errors.New(fmt.Sprintf("'%s' is not in the slice\n",s))
    }
  }
  return a, nil
}

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

func subsliceStr(s1 []string, s2 []string) bool {
    if len(s1) > len(s2) { return false }
    for _, e := range s1 {
        b,_ := containStr(s2,e)
        if !b{
            return false
        }
    }
    return true
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
  sort.Strings(a)
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
func removeDupStr(intSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}

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

func removeMemberStr(a []string, b string)[]string{
  contains, index := containStr(a,b)
  if contains == false{
    panic(fmt.Sprint("that is not a value selection: %s,%s",a,b))
  }else{
    a[index] = ""
  }
  return removeEmptyStr(a)
}
