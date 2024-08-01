package main

import ("fmt")

/**
Arrays:
- var arrayName = [length]datatype{values} // length is defined
- var arrayName = [...]datatype{values} // length is inferred

len(): length of the array
**/

/**
Slice similar to arrays -> powerful and flexible
= vector (c++)

- slice := []datatype{values}

len() : length of current slice
cap() : capacity -> maximum of elements

create Slice
slice := make([]type, length, capacity)

**/

func testArray() {
  var array1 = [5]int{1, 2, 3, 4, 5}
  fmt.Println("Array1 : ", array1)

  var array2 = [...]int{1, 2, 3, 4}
  for i := 0; i < len(array2); i += 1 {
    fmt.Printf("%v : %v, %T", i, array2[i], array2[i])
  }
}

func testSlice() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("len myslice1 = %v\n", len(myslice1))
  fmt.Printf("cap myslice1 = %v\n", cap(myslice1))
}

func testModifySlice() {
  prices := []int{10, 20, 30}

  // before modify
  fmt.Printf("Before Modify : %v\n", prices[0])

  prices[0] = 20

  // after modify
  fmt.Printf("After Modify : %v\n", prices[0])
}

// slice =  append(slice, element1, element2, ...)
func testAppendElementsToSlice() {
  slice1 := []int{1, 2, 3, 4}
  fmt.Printf("slice1 = %v\n", slice1)
  fmt.Printf("length = %v\n", len(slice1))
  fmt.Printf("capacity = %v\n", cap(slice1))

  slice2 := append(slice1, 20, 21)
  fmt.Printf("slice2 = %v\n", slice2)
  fmt.Printf("length = %v\n", len(slice2))
  fmt.Printf("capacity = %v\n", cap(slice2))
}

// slice := append(slice1, slice2...)
// ... is neccessary
func testAppendSlices() {
  slice1 := []int{1, 2, 3, 4}
  slice2 := []int{5, 6, 7}

  slice := append(slice1, slice2...)
  fmt.Println(slice)
}

// modify length of slice
func testModifySliceLength() {
  arr1 := [6]int{1, 2, 3, 4, 5, 6}
  slice1 := arr1[1:5]
  fmt.Printf("slice before : %v %T\n", slice1, slice1)

  slice2 := arr1[3:5]
  fmt.Printf("slice after : %v %T\n", slice2, slice2)
}

// memory efficiency
/*
If array is large and need only a few of elements -> use copy() function
-> reduce the memory used

copy(dest, src)
*/

func testCopySlice() {
  arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Printf("array1 : %v\n", arr1)

  neededNumbers := arr1[:len(arr1) - 3]
  fmt.Printf("needed numbers : %v\n", neededNumbers)

  copyNumbers := make([]int, len(neededNumbers))
  copy(copyNumbers, neededNumbers)
  fmt.Printf("Copy Numbers : %v\n", copyNumbers)
}

func testMakeFunction() {
  var intSlice []int32 = make([]int32, 3, 10) // length, capacity
  fmt.Println(intSlice)
}

func main() {
}
