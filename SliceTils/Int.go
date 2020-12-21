package slicetils

import(
  "strconv"
  "sort"
  "math/rand"
  "errors"
  "fmt"
)

//Breaks main slice into subslices int
func sliceSpliterInt(slice []int, slicecount int)([][]int, error){
  var slices [][]int
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
    var s []int
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

// Removes empty values from int slice
func removeEmptyInt(s []int) []int{
  var r []int
  for _, str := range s{
    if str != 0{
      r = append(r,str)
    }
  }
  return r
}

// Removes integers from a slice
func removeSliceInt(a []int, remove ...int)([]int,error){
  // Remove the element at index i from a.
  for _, s := range remove{
    index := IndexOf(s,a)
    if index != -1{
      copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
      a[len(a)-1] = 0     // Erase last element (write zero value).
      a = a[:len(a)-1]     // Truncate slice.
      a = removeEmptyInt(a)
    }else{
      return a, errors.New(fmt.Sprintf("'%v' is not in the slice\n",s))
    }
  }
  return a, nil
}

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

func removeMemberInt(a []int, b int)[]int{
  contains, index := containInt(a,b)
  if contains == false{
    panic("that is not a value selection")
  }else{
    a[index] = 0
  }
  return removeEmptyInt(a)
}
