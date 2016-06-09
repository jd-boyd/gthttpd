package main

import (
  "log"
  "net/http"
  "flag"
  "fmt"
  "os"
)

func main() {

  port := flag.Int("p", 80, "http port")
  flag.Bool("D", false, "Don't daemonize")
  cur_dir, _ := os.Getwd()
  
  root_dir := flag.String("d", cur_dir, "Root directory")

  flag.Parse()

  fs := http.FileServer(http.Dir(root_dir))
  http.Handle("/", fs)

  log.Println("Listening...")

  port_str := fmt.Sprintf(":%d", *port)
  log.Fatal(http.ListenAndServe(port_str, nil))
}
