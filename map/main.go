package main

import ("fmt")

/**
map[KEY_TYPE]VALUE_TYPE{}

**/

func testMap() {
  user := map[string]string{
    "name": "Nova",
    "email": "nova@gmail.com",
  }

  user["age"] = "20"

  fmt.Printf("Map : %v", user)
}

func testDeleteElementFromMap()  {
  user := map[string]string{
    "name": "Nova",
    "email": "nova@gmail.com",
  }

  // delete(mapName, key)
  delete(user, "name")
  fmt.Printf("Map : %v", user)
}

func testCheckSpecificElementsInMap() {
  user := map[string]string{
    "name": "Nova",
    "email": "nova@gmail.com",
  }

  val, ok := user["name"]
  fmt.Printf("Existed key in map : { val : %v, ok : %v }\n", val, ok)

  val, ok = user["password"]
  fmt.Printf("Non-existed key in map : { val : %v, ok : %v }\n", val, ok)
}

func testIterateOverMap() {
  user := map[string]string{
    "name": "Nova",
    "email": "nova@gmail.com",
  }

  for key, value := range user {
    fmt.Printf("key : %v, value : %v\n", key, value)
  }
}

func main() {
}
