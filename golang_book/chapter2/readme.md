#Problems

* [x] What is whitespace?
    - whitespaces are characters like  **space**, **tab**, **new line** 
* [x] What is a comment? What are the two ways of writing a comment?
    - a comment is(are) line(s) of text that are ignored by compiler
    - in go there are two ways to define comments. with : **// text**, and **/\*\*/**
    ```go
    // this is single line comment
    /*
    This is 
    a multiline 
    comment
    */
    ```


* [x] Our program began with package main. What would the files in the fmt package begin with?
    ```go
    package fmt
    ```

* [x] We used the Println function defined in the fmt package. If we wanted to use the Exit function from the os package what would we need to do?
    - first we need to import **os** package
    - get signature of **os.Exit()**
    ```bash
    $ godoc os Exit
    use 'godoc cmd/os' for documentation on the os command

    func Exit(code int)
        Exit causes the current program to exit with the given status code.
        Conventionally, code zero indicates success, non-zero an error. The
        program terminates immediately; deferred functions are not run.
    ```
    - use this in your go file
    ```go
    import "os"
    os.Exit(0)

    ```

* [x] Modify the program we wrote so that instead of printing Hello World it prints Hello, my name is followed by your name.
    - check out [main.go](./main.go)