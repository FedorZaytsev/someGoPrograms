package main

import (
  "fmt"
)


func Partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int {
  var (
    i = low
    j = low
  )
  for j<high {
    if less(j,high) {
      swap(i,j)
      i++
    }
    j++
  }
  swap(i,high)
  return i
}

func quicksortRec(low, high int, less func(i, j int) bool, swap func(i, j int)) {
  if low<high {
    q:=Partition(low,high,less,swap)
    quicksortRec(low,q-1,less,swap)
    quicksortRec(q+1,high,less,swap)
  }
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
  quicksortRec(0,n-1,less,swap);
}

func main() {

  var count int

  fmt.Scanf("%d",&count)
  t := make([]int,count)

  for idx,_  := range t {
    fmt.Scanf("%d",&t[idx])
  }

  qsort(count,        func (i, j int) bool {
  					return t[i] < t[j]
				},
				func (i, j int) {
  					var q = t[i]
  					t[i] = t[j]
  					t[j] = q
				})
  

  for idx,_  := range t {
    fmt.Printf("%d ",t[idx])
  }
}
