/*
1.3 Experimente medir a diferença de tempo de execução entre nossas versões
potencialmente ineficientes e a versão que usa strings.Join.
echo2
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

echo3
package main
import (
        "fmt"
        "os"
        "strings"
)

func main() {
        fmt.Println(strings.Join(os.Args[1:], " "))
}

RESULTS
/tmp/go-build638443159/b001/exe/1_3 a b c d e f g h i j k l
echo2 took 0.0000446 seconds
/tmp/go-build638443159/b001/exe/1_3 a b c d e f g h i j k l
echo3 took 0.0000121 seconds

*/
package main

import (
        "fmt"
        "os"
        "time"
)

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}


func main() {
	startEcho2 := time.Now()
	echo2(os.Args)
	fmt.Printf("echo2 took %.7f seconds\n", time.Since(startEcho2).Seconds())
	startEcho3 := time.Now()
	echo3(os.Args)
	fmt.Printf("echo3 took %.7f seconds\n", time.Since(startEcho3).Seconds())
}