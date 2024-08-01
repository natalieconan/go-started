package main

import ("fmt")

/**
type structName struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...

  - there is no private in golang
  - lowercase -> protected
  - uppercase -> public
}
**/

type User struct {
  name string
  age int
  email string
}

func (user User) sayHello() string {
  return fmt.Sprintf("Hello! my name is %s", user.name)
}

func testStruct() {
  user := User{
    name: "Nova",
    age: 21,
    email: "nova@gmail.com",
  }
  fmt.Println("user : ", user)

  message := user.sayHello()
  fmt.Println("messsage : ", message)
}

func main() {
}
