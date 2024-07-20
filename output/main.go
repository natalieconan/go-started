package main

import ("fmt")

/**
Output function:
- Print() : default
- Println() : with new line ('\n')
- Printf() : print with format
      - %v : value
      - %T : type
**/

func testPrint() {
  var var1 int = 0
  var var2 int = 1
  fmt.Print(var1)
  fmt.Print(var2)
}

func testPrintln() {
  var var1 int = 0
  var var2 int = 1
  fmt.Println(var1)
  fmt.Println(var2)
}

func testPrintf() {
  var var1 int = 0
  var var2 string = "var2"
  fmt.Printf("%v %T\n%v %T", var1, var1, var2, var2)
}

func main() {

}
