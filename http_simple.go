package main
import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/gorilla/mux" // acquired using $ go get github.com/gorilla/mux
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
    /*
    The example creates a basic router,
    adds the route / and assigns the Index handler to run when that endpoint is called.
    */
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/",Index)
    log.Fatal(http.ListenAndServe(":8080",router))
}
