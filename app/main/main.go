package main

import (
	"net/http"

	"github.com/dinophile/routes"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// So, let me preface this by saying ooh I wish I had not had a crazy week last week and could have spent more time
// before handing it to y'all! But I think it's a decent attempt to show most of the limits of my knowledge!
// I haven't worked with production golang at all so some of this was based solely on my experience working and
// building Node APIs!
// I figured out how to add tests and normally when I'm comfortable with a language to a degree I do try my hand at
// some TDD, but this was largely about implementing the main ask of the README!
// I'm going to keep going and play with testing some more since I have some experience using and writing tests this
// would be a great way to expand on that using go!
// Also did not complete adding error handling but whoooo I love me some human readable, helpful errors!
// I also will likely keep learning about structuring go servers and their files! Everything in it's place!
// Thanks for the opportunity to learn and play with some of my nacent skills at the very least! This was a load of fun!
// Hope I can at least come by and chat with the team and maybe get some feedback to learn more!

// main - abstracted out the new routes I've added
// TODO: learn about structuring GO APIs to further simplify the file and have routes in their own separate file
// Is there where you'd use the /internal folder? Hmmm...
func main() {
	http.HandleFunc("/echo", routes.Echo)
	http.HandleFunc("/invert", routes.Invert)
	http.HandleFunc("/flatten", routes.Flatten)
	http.HandleFunc("/sum", routes.Sum)
	http.HandleFunc("/multiply", routes.Multiply)
	http.HandleFunc("/health", routes.Health)
	http.ListenAndServe(":8080", nil)
}
