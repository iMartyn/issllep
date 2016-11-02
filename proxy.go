package main

import "os"
//import "strings"
import "fmt"
import "log"
import "net/http"
import "html"
//import "database/sql"

const defaultListen = ":8080"
const defaultDBDriver = "SQLite"
const defaultDBDSN = "domains.db"

func defaultEnv(def string, envVar string) string {
    var actual = os.Getenv(envVar)
    if (actual == "") {
        return def
    }
    return actual
}

/*func dbInit() {
    var dbDriver string
    
    db, err := sql.Open(defaultEnv(defaultDBDriver, "DBDRIVER"), defaultEnv(defaultDBDSN, "DBDSN"))
    if err != nil (
        defer db.Close()
    )
}*/

func main() {
   // dbInit()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    
    log.Fatal(http.ListenAndServe(defaultEnv(":8080","LISTEN"), nil))
}
