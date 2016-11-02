package main

import "os"
import "strings"
import "fmt"
import "log"
import "net/http"
//import "html"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

const defaultListen = ":8080"
const defaultDBDriver = "sqlite3"
const defaultDBDSN = "domains.db"

var db sql.DB

func defaultEnv(def string, envVar string) string {
    var actual = os.Getenv(envVar)
    if (actual == "") {
        return def
    }
    return actual
}

func dbInit() {
    db, err := sql.Open(defaultEnv(defaultDBDriver, "DBDRIVER"), defaultEnv(defaultDBDSN, "DBDSN"))
    if (err != nil) {
        log.Fatal(err)
    }
    defer db.Close()
}

func main() {
    dbInit()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        urlparts := strings.Split(r.URL.Path, "/")
        if (urlparts[1] == ".well_known") {
            var err error
            var domaindata string
            var fqdn string = urlparts[2]
            log.Print(fqdn)
            err = db.QueryRow("SELECT data FROM domaindata WHERE fqdn = ?;", fqdn).Scan(&domaindata)
            if err != nil {
                if err == sql.ErrNoRows {
                    domaindata = "NO DATA IN DATABASE, PANIC!"
                } else {
                    domaindata = "AARGH"
                }
            } else {
                fmt.Fprintf(w, domaindata)
            }
        } else {
            fmt.Fprintf(w, "Hello, You don't want to be here, go away.")
        }
    })
    
    log.Fatal(http.ListenAndServe(defaultEnv(":8080","LISTEN"), nil))
}
