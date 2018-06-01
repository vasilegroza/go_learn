# Package

Packages helps us to organize our code

* one folder many files    
    * folder name = package name
    * package name declaration in every file
    * package scope
        * you can access acros files within the same package
    * exported / unexported
        * aka public/private
        * lowercase  --> visible only from files within the package
        * capitalize --> visible outside the package
        
* to use 3rd packages    
  ```bash
  go get -u github.com/vasilegroza/go_learn 
  ```
  if the repo is private rungo
  ```sh 
  git config --global url."git@github.com:".insteadOf "https://github.com/"
  ```
  ```go
  import (  
	    "fmt"
	    "github.com/vasilegroza/go_learn/02_package/stringutil" // istalled custom package
  )
  ```

* go comands
    * go run
    * go build
    * go install
        - if you are in outside gopath
        ```sh
        export GOBIN=$GOPATH/bin
        ```
    * go clean
