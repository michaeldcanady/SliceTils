package SliceTools

import(
  "errors"
  "fmt"
  "os"
)

// Removes strings from a slice
func RemoveSliceStr(a []string, remove ...string)([]string,error){
  // Remove the element at index i from a.
  for _, s := range remove{
    index := SliceIndex(len(a), func(i int) bool{return a[i] == s})
    if index != -1{
      copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
      a[len(a)-1] = ""     // Erase last element (write zero value).
      a = a[:len(a)-1]     // Truncate slice.
      a = RemoveEmptyStr(a)
    }else{
      return a, errors.New(fmt.Sprintf("'%s' is not in the slice\n",s))
    }
  }
  return a, nil
}

// Removes integers from a slice
func RemoveSliceInt(a []int, remove ...int)([]int,error){
  // Remove the element at index i from a.
  for _, s := range remove{
    index := SliceIndex(len(a), func(i int) bool{return a[i] == s})
    if index != -1{
      copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
      a[len(a)-1] = 0     // Erase last element (write zero value).
      a = a[:len(a)-1]     // Truncate slice.
      a = RemoveEmptyInt(a)
    }else{
      return a, errors.New(fmt.Sprintf("'%v' is not in the slice\n",s))
    }
  }
  return a, nil
}

// Removes empty values from string slice
func RemoveEmptyStr(s []string) []string{
  var r []string
  for _, str := range s{
    if str != ""{
      r = append(r,str)
    }
  }
  return r
}

// Removes empty values from int slice
func RemoveEmptyInt(s []int) []int{
  var r []int
  for _, str := range s{
    if str != 0{
      r = append(r,str)
    }
  }
  return r
}

func SliceIndex(limit int, predicate func(i int) bool) int {
    if limit == 0{
      return -1
    }
    for i := 0; i < limit; i++ {
        if predicate(i) {
            return i
        }
    }
    return -1
}

//Breaks main slice into subslices string
func SliceSpliterStr(slice []string, slicecount int)([][]string,error){
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

//Breaks main slice into subslices int
func SliceSpliterInt(slice []string, slicecount int)([][]string, error){
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

// Checks if a slice contains a specific string
func containsStr(s []string, e string)(bool,error){
  for _, a := range s {
    if a == e {
      return true,nil
    }
  }
  return false,errors.New(fmt.Sprintf("%s not in slice",e))
}

//Checks if a slice contains a specific int
func containsInt(s []int, e int)(bool,error){
  for _, a := range s {
    if a == e {
      return true,nil
    }
  }
  return false,errors.New(fmt.Sprintf("%v not in slice",e))
}

func SliceSize(slice []string) int64{
  var totalSize int64
  for _,files := range slice{
    file,err := os.Stat(files); if err != nil{
      panic(err)
    }else{
      totalSize += file.Size()
    }
  }
  return totalSize
}
