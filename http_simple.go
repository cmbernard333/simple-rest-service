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

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w,"Welcome to the Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w,"Todo show:", todoId)
}

// assign all the functions mapped in the map 'm'
func SetupRouter(router *mux.Router, m map[string]func(http.ResponseWriter,*http.Request)) {
    for key, value := range m {
        router.HandleFunc(key,value)
    }
}

func main() {
    // map paths to handlers - make creates a reference to a map and assigns it to m
    m := make(map[string]func(http.ResponseWriter,*http.Request));
    m["/"] = Index
    m["/todos"] = TodoIndex
    m["/todos/{todoId}"] = TodoShow
    /*
    The example creates a router,
    adds the route / and assigns the Index handler to run when that endpoint is called.
    */
    router := mux.NewRouter().StrictSlash(true)
    SetupRouter(router,m)
    log.Fatal(http.ListenAndServe(":8080",router))
}
