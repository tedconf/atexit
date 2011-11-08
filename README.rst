About
=====
Simple `atexit` implementation for Go_.

Note that you *have* to call `atexit.Exit` and not `os.Exit` to terminate your
program (that is, if you want the atexit handlers to execute).

Example usage
=============
::

    package main

    import (
        "atexit"
        "fmt"
    )

    func handler() {
        fmt.Println("Exiting")
    }

    func main() {
            atexit.Register(handler)
            atexit.Exit(0)
    }

Install
=======
::
    goinstall bitbucket.org/tebeka/atexit

Contact
=======
Home
    https://bitbucket.org/tebeka/atexit
Author
    Miki Tebeka <miki.tebeka@gmail.com>


.. _Go: http://golang.org
