package file

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func GenerateStruct () {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\nEnter your model file path: ")
    sourcePath,err := reader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }

    input, err := ioutil.ReadFile(strings.TrimSpace(sourcePath))
    if err != nil {
        fmt.Println("ERR MODEL:", err)
        return
    }

    savePath := "/home/mehrdad/Documents/SELF-PROJECT/NATS.P/pkg/file/struct.go"
    err = ioutil.WriteFile(savePath, input, 0644)
    if err != nil {
        fmt.Println("ERR MODEL:", err)
        return
    }

    fmt.Println("Model added successfully")
}
