# Wikidocs - Backend
This is the backend to the [Wikidocs frontend](https://github.com/jon-ryan/wiki-demoserver-frontend). It is implemented in Go using the [fiber framework](https://github.com/gofiber/fiber). This repo is for demonstration and learning. 


# Clone and build
Clone the repo with:
```bash
git clone https://github.com/jon-ryan/wiki-demoserver.git
```
Make sure you have all the dependencies installed.
Then build with:
```bash
go build main.go
```

In order to use the server properly, place the compiled files from the [Wikidocs frontend](https://github.com/jon-ryan/wiki-demoserver-frontend) in the *static_files* directory. The server can be reached in the browser on
```
localhost:8080
```
