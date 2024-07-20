package main

import ("fmt")

/**
Types
- int
- int8, int16, int32, int64
- float32
- string
- bool
**/

/**
var variableName type = value
**/

/**
variableName := value
**/

// global variable
var globalVariable string = "global Varable";

func staticAndDynamicTypeVariable() {
  // static type variable
  var studentFirstName string = "Nova"

  // dynamic type variable
  studentLastName := "Seele"

  var studentName string = studentFirstName + " " + studentLastName

  fmt.Println("Static Type Variable : ", studentFirstName)
  fmt.Println("Dynamic Type Variable : ", studentLastName)
  fmt.Println("Combined Variable : ", studentName)
}

func variableWithouInitialValue() {
  // variable declared without initial value
  var studentFullName string
  studentFullName = "Nova Seele"
  fmt.Println("Varirable Without Initial Value : ", studentFullName)
}

func printGlobalVariable() {
  fmt.Println("Global Variable : ", globalVariable)
}

func multipleVariablesDeclared() {
  // declare multiple variables
  var a, b, c, d, e = 1, 2, 3, 4, 5
  fmt.Println("Declare Multiple Variables : ", a, b, c, d, e)
}

func variableInsideBlock() {
  // declare in a block
  var (
    a int = 0
    b int = 1
    c string = "hello"
  )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

func printConstant() {
  // typed CONST
  const INF int = 1e9 + 5

  // untyped CONST
  const PI = 3.14
  fmt.Println(INF)
  fmt.Println(PI)
}

func main() {

}
