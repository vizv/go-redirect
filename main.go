package main

import (
    "io/ioutil"
    "net/http"
    "strings"
    "html"
    "fmt"
    "os"
)

var aliasesMap map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
    from := r.URL.String()
    to   := aliasesMap[from]

    if to == "" {
        message := fmt.Sprintf("The path [%s] was not found!",
                               html.EscapeString(from))
        http.Error(w, message, http.StatusNotFound)
    } else {
        w.Header().Add("Location", to)

        w.WriteHeader(http.StatusFound)
    }
}

func main() {
    // load the aliases from file
    aliasesFile := os.Getenv("ALIASES_FILE")
    if aliasesFile == "" {
        fmt.Fprintln(os.Stderr, "FATAL: Aliases file cannot be empty!")
        os.Exit(1);
    }

    aliases, err := ioutil.ReadFile(aliasesFile)
    if err != nil {
        fmt.Fprintln(os.Stderr, "FATAL: Cannot open aliases file!")
        os.Exit(1);
    }

    maxWidth := 0
    aliasesMap = make(map[string]string)
    for _, alias := range strings.Split(string(aliases), "\n") {
        // skip empty lines and comments
        if alias == "" || alias[0] == '#' { continue }

        // ignore malformed lines
        if strings.Index(alias, ":") == -1 {
            fmt.Fprintf(os.Stderr,
                        "Warning: Skip malformed configuration: [%s].\n",
                        alias)

            continue
        }

        pair := strings.SplitN(alias, ":", 2)
        from := "/" + strings.TrimSpace(pair[0])
        to   := strings.TrimSpace(pair[1])
        aliasesMap[from] = to

        if (len(from) > maxWidth) {
            maxWidth = len(from)
        }
    }

    fmt.Println("INFO: Aliases map loaded...")
    for from, to := range aliasesMap {
        fmt.Printf("%[2]*[1]s -> %[3]s\n", from, maxWidth, to);
    }

    http.HandleFunc("/", handler)

    fmt.Println("INFO: Server up and running...")
    http.ListenAndServe(":8080", nil)
}
