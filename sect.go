package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "regexp"
)

//test comment

func main() {

  if len(os.Args) != 2 {
    log.Fatal("Usage: sect pattern")
  }

  r := regexp.MustCompile(os.Args[1])

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    last_line_printed := false
    //fmt.Printf("\t\"%s\"\n", line)
M:  if r.MatchString(line) {
      if !last_line_printed {
        fmt.Println(line)
      }
      last_line_printed = false
      spaces := 0
      for line[spaces] == ' ' || line[spaces] == '\t' { spaces++ }

      for scanner.Scan() {
        subline := scanner.Text()
        subspaces := 0
        for (subline[subspaces] == ' ' || subline[subspaces] == '\t') && subspaces <= spaces { subspaces++ }
        if subspaces > spaces {
          fmt.Println(subline)
        } else {
          last_char := line[len(line)-1]
          first_char := subline[subspaces]
          if (last_char == '(' && first_char == ')') ||
             (last_char == '{' && first_char == '}') ||
             (last_char == '[' && first_char == ']') ||
             false {
            //if
            fmt.Println(subline)
            last_line_printed = true
          }
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
