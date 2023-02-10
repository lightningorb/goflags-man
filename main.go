/*
* @Author: w
* @Date:   2023-02-10 16:20:07
* @Last Modified by:   w
* @Last Modified time: 2023-02-10 16:21:38
*/
package main

import (
    "fmt"
    "os"

    "github.com/jessevdk/go-flags"
)

type Options struct {
    Message string `short:"m" long:"message" description:"The greeting message" default:"Hello, World!"`
    GenerateManPage bool `short:"g" long:"generate-man-page" description:"Generate man page"`
}

func main() {
    var opts Options
    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()
    if err != nil {
        panic(err)
    }

    if opts.GenerateManPage {
        f, err := os.Create("program.1")
        if err != nil {
            panic(err)
        }
        defer f.Close()

        parser.WriteManPage(f)

        fmt.Println("Man page generated as program.1")
    } else {
        fmt.Println("Hello " + opts.Message)
    }
}
