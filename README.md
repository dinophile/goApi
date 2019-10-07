# My first goAPI! Sniff, they grow up so fast!

So, let me preface this by saying ooh I wish I had not had a crazy week last week and could have spent more time before handing it to y'all! But I think it's a decent attempt to show most of the limits of my knowledge! 

I haven't worked with production golang at all so some of this was based solely on my experience working and building Node APIs! I figured out how to add tests and normally when I'm comfortable with a language to a degree I do try my hand at some TDD, but this was largely about implementing the main ask of the README! I'm going to keep going and play with testing some more since I have some experience using and writing tests this would be a great way to expand on that using go!

Also did not complete adding error handling but whoooo I love me some human readable, helpful errors! I also will likely keep learning about structuring go servers and their files! Everything in it's place!
 
Thanks for the opportunity to learn and play with some of my nacent skills at the very least! This was a load of fun!
Hope I can at least come by and chat with the team and maybe get some feedback to learn more!

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
  
  
