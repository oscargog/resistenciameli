package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func CountWords(s string) int {
    return len(strings.Fields(s))
}

func GetLocation(posiciones string) {
 
    var posRebeldes string
    posRebeldes = posiciones
    
    switch posRebeldes {
    case "[-500, -200]":
        fmt.Print("Kenobi: [-500, -200]\n")
        fmt.Print("Ingresa el mensaje:")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        var space = " "
        var message = space +scanner.Text()
        del := " "
        var kenobiarr = strings.Split(message,del)
        fmt.Printf("%q\n", kenobiarr)


    case "[100, -100]":
        fmt.Print("Skywalker: [100, -100]\n")
        fmt.Print("Ingresa el mensaje:")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        var message = scanner.Text()
        del := " "
        var Skywalkerarr = strings.Split(message,del)
        Skywalkerarr[1] = " "
        fmt.Printf("%q\n", Skywalkerarr)


    case "[500, 100]":
        fmt.Print("Sato: [500, 100]\n")
        fmt.Print("Ingresa el mensaje:")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        var message = scanner.Text()
        del := " "
        var Satoarr = strings.Split(message,del)
        Satoarr[0] = " "
        Satoarr[1] = " "
        Satoarr[3] = " "
        fmt.Printf("%q\n", Satoarr)

    default:
        fmt.Print("Tus ojos pueden engañarte, no confíes en ellos - Obi-Wan \n")

        return
    }


}

func GetMessage (message ...[]string) (msg string){
    fmt.Print(message)
    return
}


func main() {

    fmt.Print("Ingresa las posiciones:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    posiciones := scanner.Text()
    GetLocation(posiciones)
}