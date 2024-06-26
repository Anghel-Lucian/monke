package main

import (
	"fmt"
    "os"
	"os/user"

    "monke/repl"
)

func main() {
    user, err := user.Current();

    if err != nil {
        panic(err);
    }

    fmt.Printf("Hello %s. This is Monke\n", user.Username);
    repl.Start(os.Stdin, os.Stdout);
}

