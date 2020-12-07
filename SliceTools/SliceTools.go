package SliceTools

import(
  "errors"
  "fmt"
  "os"
)

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
