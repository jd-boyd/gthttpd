package main

import (
  "log"
  "net/http"
  "flag"
  "fmt"
  "os"
  "path/filepath"
)

func main() {

  port := flag.Int("p", 8080, "http port")
  flag.Bool("D", false, "Don't daemonize (but we wouldn't anyway).")
  cur_dir, _ := os.Getwd()
  
  root_dir_arg := flag.String("d", cur_dir, "Root directory")

  flag.Parse()

  root_dir, err := filepath.Abs(*root_dir_arg)
  if err != nil {
    log.Fatal(err)
  }      

  fs := http.FileServer(http.Dir(root_dir))
  http.Handle("/", fs)
//  http.Handle("", http.NotFoundHandler())

  port_str := fmt.Sprintf(":%d", *port)

  log.Println("Will listen on", port_str, "and serve", root_dir)

  log.Println("Listening...")
  log.Fatal(http.ListenAndServe(port_str, nil))
}
