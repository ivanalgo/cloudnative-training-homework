package main

import "net/http"
import "log"
import "io"
import "fmt"
import "strings"
import "runtime"
import "net"
import "errors"

func main() {
  http.HandleFunc("/healthz", healthz)
  http.HandleFunc("/", indexz)
  err := http.ListenAndServe(":80", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func copyHeader(w http.ResponseWriter, req *http.Request) {
  for k,v := range req.Header {
    w.Header().Set(k, fmt.Sprintf(strings.Join(v, ",")))
  }

  w.Header().Set("VERSION", runtime.Version())
}

func clientIP(req *http.Request) (string, error) {
  ip := req.Header.Get("x-Real-IP")
  if net.ParseIP(ip) != nil {
    return ip, nil
  }

  ip = req.Header.Get("X-Forward-For")
  for _, i := range strings.Split(ip, ",") {
    if net.ParseIP(i) != nil {
      return i, nil
    }
  }

  ip, _, err := net.SplitHostPort(req.RemoteAddr)
  if err != nil {
    return "", err
  }

  if net.ParseIP(ip) != nil {
    return ip, nil
  }

  return "", errors.New("no valid ip found")
}

func logHttp(w http.ResponseWriter, req *http.Request, statusCode int) {
  ip, err := clientIP(req)
  if err != nil {
    fmt.Println("IP : ", err)
  } else {
    fmt.Println("IP : ", ip)
  }

  fmt.Println("HttpCode :", statusCode)
}

func healthz(w http.ResponseWriter, req *http.Request) {
  copyHeader(w, req)
  w.WriteHeader(200)
  logHttp(w, req, 200)
  io.WriteString(w, "ok")
}

func indexz(w http.ResponseWriter, req *http.Request) {
  copyHeader(w, req)
  w.WriteHeader(404)
  logHttp(w, req, 404)
}
