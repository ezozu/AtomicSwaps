// cmd/atomicswaps/main.go
package main

import (
"flag"
"log"
"os"

"atomicswaps/internal/atomicswaps"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := atomicswaps.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
