package main

import ("fmt")

/**
Input function:
- Scan() : default
- Scanlt()
- Scanf()
**/

/**
Output function:
- Print() : default
- Println() : with new line ('\n')
- Printf() : print with format
      - %v : value
      - %T : type
      - %b : base 2
      - %d : base 10
      - %+d : base 10 and show sign
      - %o : base 8
      - %x : base 16, lowercase
      - %X : base 16, uppercase
      - %.2f : precision(2)
      - %t : value of boolean
**/

func testScan() {
  var i, j int
  fmt.Scan(&i, &j)
  fmt.Println(i, j)
}

func testScanf() {
  var i, j int
  fmt.Scanf("%v %v", &i, &j)
  fmt.Println(i, j)
}

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

func testPrintPrecision() {
  var var1 float32 = 0.153432
  fmt.Printf("%.2f", var1)
}

func main() {
}
