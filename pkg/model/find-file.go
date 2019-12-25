package main

import (
    "bufio"
    "fmt"
    "os"
)


// func ReadString () {
//     reader := bufio.NewReader(os.Stdin)
//     fmt.Print("Enter string: ")
//     for true {
//     	vrb, _ = reader.ReadString('\n')
//     }
//     fmt.Print("You live in " + city)
// }

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter : ")
    var vrb string
   	var model struct

   	model := model{vrb:vrb}
    vrb, _ = reader.ReadString('\n')
    fmt.Print("You : " + vrb)
}