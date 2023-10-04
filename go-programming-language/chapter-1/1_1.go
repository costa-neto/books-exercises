/*
1.1 Modifique o programa echo para exibir também os.Args[0], que é o nome do comando que o chamou.
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
        "strings"
)
func main() {
        fmt.Println(strings.Join(os.Args[0:], " "))
}