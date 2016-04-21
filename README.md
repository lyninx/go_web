### Go web backend + REST api
#### Quick Start
1. `go get github.com/gorilla/mux`
2. `go get gopkg.in/mgo.v2`
3. `go run *.go`
4. `ab -c 500 -n 1000 http://localhost:8080/` (flood with requests)

##### About
```
The goal of this project was to build a basic web server, capable of connecting to a database, and serving pages, as well as page data through a REST api. Go makes for a great option for this due to the high quality http package built into the language which handles concurrency by spinning off goroutines for each new request to the server. It also makes for a good choice due to its simple type system, safety precautions, and easy modularization.
```
##### Comparison to other languages
```
Contrasting to using a language like NodeJS, Go likely has the upper hand when it comes to raw concurrency performance, allowing it to handle more requests with the same amount of CPU usage. Although Javascript has some benefits such as its popularity (which means for more libraries), as well as purely the number of language features, such as being able to use classes or functional programming.
Compared to Ruby, it seemed more difficult to be able to simply append html to a document, or build nested template components, although that's probably for the best to keep things clean. Go's template solution ended up working very well for me.
```

##### Some Observations
* slices are great.
* the language really tries to make sure that you're writing quality code. being compiled brings some confidence, too.
* no optional parameters or function overloading.
* performs well, on par with node for low number of requests.
* being a compiled language might make hot-reloading more difficult?
* the simplicity of the language made it fun to write in. "C, the good parts."

##### Moving forward
```
This opportunity also allowed me to compare some of the differences between having page content pre-rendered on the server, versus being pulled in from an API. Backend serving is more responsive, while front end fetching has the benefit of being reactive and interactive after the initial page load. I definitely will be looking into isomorphic web apps built with tools such as React and Vue to try and get the best of both worlds.
```

