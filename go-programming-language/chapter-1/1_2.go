/*
1.2: Modifique o programa echo para exibir o índice e o valor de cada um de seus argumentos, um por linha.
package main

import (
        "fmt"
        "os"
        "strings"
)

func main() {
        fmt.Println(strings.Join(os.Args[1:], " "))
}
*/
package main

import (
        "fmt"
        "os"
)

func main() {
		for idx, arg := range os.Args {
			fmt.Println("idx: ", idx, " - arg:", arg)
		}
}
