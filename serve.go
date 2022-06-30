package main

import (
    "fmt"
    "net/http"
)

func main() {
        port := 3000;
        http.Handle("/", http.FileServer(http.Dir(".")));
        fmt.Println("Serving on port:", port);
        http.ListenAndServe(fmt.Sprintf(":%d", port), nil);
}
