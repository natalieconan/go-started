package main

import ("fmt")

/**
type structName struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
**/

type User struct {
  name string
  age int
  email string
}

func testStruct() {
  user := User{
    name: "Nova",
    age: 21,
    email: "nova@gmail.com",
  }
  fmt.Println(user)
}

func main() {

}
