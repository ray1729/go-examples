package main

import (
  "bufio"
  "fmt"
  "os"
  "crypto/sha256"
)

func hashSha256(data string) string {
  hash := sha256.Sum256([]byte(data))
  return fmt.Sprintf("%x", hash)
}

func main() {
  inch := make(chan string)
  ouch := make(chan string)

  go func() {
    for in := range inch {
      ouch<- hashSha256(in)
    }
  }()

  go func() {
    for ou := range ouch {
      fmt.Println(ou)
    }
  }()

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    inch<- scanner.Text()
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Error reading input: %v\n", err)
  }
  close(inch)
  close(ouch)
}
