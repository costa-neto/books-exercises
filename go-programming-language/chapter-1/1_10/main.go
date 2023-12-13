// 1.10: Modifique fetchAll para exibir sua saída em um arquivo
// para que ela possa ser examinada.
// Fetchall fetches URLs in parallel and reports their times and sizes.
//package main
//
//import (
//  "fmt"
//  "io"
//  "io/ioutil"
//	"net/http"
//	"os"
//	"time"
//)
//func main() {
//  start := time.Now()
//  ch := make(chan string)
//  for _, url := range os.Args[1:] {
//    go fetch(url, ch) // start a goroutine
//  }
//  for range os.Args[1:] {
//    fmt.Println(<-ch) // receive from channel ch
//  }
//  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
//}
//
//func fetch(url string, ch chan<- string) {
//  start := time.Now()
//  resp, err := http.Get(url)
//  if err != nil {
//    ch <- fmt.Sprint(err) // send to channel ch
//    return
//}
//
//  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
//  resp.Body.Close() // don't leak resources
//  if err != nil {
//    ch <- fmt.Sprintf("while reading %s: %v", url, err)
//    return
//  }
//  secs := time.Since(start).Seconds()
//  ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
//}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	var output string
  start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		output += fmt.Sprintf("\n" + <-ch) // receive from channel ch
	}
	output += fmt.Sprintf("\n%.2fs elapsed\n", time.Since(start).Seconds())
  _ = os.WriteFile("output.txt", []byte(output), 0666)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
