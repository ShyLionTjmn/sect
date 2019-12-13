package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "regexp"
)

func main() {
  if len(os.Args) != 2 {
    log.Fatal("Usage: sect pattern")
  }

  r := regexp.MustCompile(os.Args[1])

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    //fmt.Printf("\t\"%s\"\n", line)
M:  if r.MatchString(line) {
      fmt.Println(line)
      spaces := 0
      for line[spaces] == ' ' { spaces++ }

      for scanner.Scan() {
        subline := scanner.Text()
        subspaces := 0
        for subline[subspaces] == ' ' && subspaces <= spaces { subspaces++ }
        if subspaces > spaces {
          fmt.Println(subline)
        } else {
          line = subline
          goto M
        }
      }
      break
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
