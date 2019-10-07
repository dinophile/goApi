# My first goAPI! Sniff, they grow up so fast!

## To use:

#### 1. Clone repo into your $GOPATH/src folder

#### 2. If not already installed run `go get github.com/pkg/errors`

#### 3. Change directory to $GOPATH/src/app/main and run `go run .` 

### Application runs on port 8080

## Routes:

### Health
* `curl "localhost:8080/health"`
* Returns: 
  - String: "You're good to...GO! :D"

### Echo
* `curl -F 'file=@matrix.csv' "localhost:8080/echo"`
* Returns []string:  
  - ```
    1,2,3
    4,5,6
    7,8,9
    ```

### Invert
* `curl -F 'file=@matrix.csv' "localhost:8080/invert"`
* Returns []string:  
  - ```
    1,4,7
    2,5,8
    3,6,9
    ```

### Flatten
* `curl -F 'file=@matrix.csv' "localhost:8080/flatten"`
* Returns string:  
  - `1,2,3,4,5,6,7,8,9`
   
### Sum
* `curl -F 'file=@matrix.csv' "localhost:8080/sum"`
* Returns string:  
  - `45`
  
### Multiply
* `curl -F 'file=@matrix.csv' "localhost:8080/multiply"`
* Returns string:  
  - `362880`
