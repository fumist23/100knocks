package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    fp, err := os.Open("hightemp.txt")
    if err != nil {
        log.Fatalln(err)
    }
    defer fp.Close()
    scanner := bufio.NewScanner(fp)
    lines := []string{}
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    for i := range lines {
        fmt.Println(lines[len(lines)-i-1])
    }
}