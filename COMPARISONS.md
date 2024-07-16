### Comparion of Go with NodeJs

##### Equivalent of `npm init` in Go

-   It creates a new `go.mod` file in the current directory.
-   The `go.mod` file contains information about the module, its dependencies, and the Go version.
-   `go mod init` is used to initialize a new module.

```bash
go mod init <module_name>
```

##### Equivalent of `npm run start` in Go

```bash
go run <main_file.go>
```

-   `go run` is used to compile and run a Go program.
-   It compiles the program and executes it.

##### Equivalent of `npm install <package_name>` in Go

-   `go get` is used to download and install packages from remote repositories.
-   It does not handle versioning.
-   This command fetches the package and its dependencies (if any)
-   `go get` is not a package manager.

```bash
go get <package_name>
```

##### Equivalent of `package.json` in Go
- `go.mod`

```bash
go.mod file
```

-   It contains information about the module, its dependencies, and the Go version.

##### Equivalent of `npm install` in Go

-   `go mod tidy` is used to add missing and remove unused modules.
-   It updates the go.mod file to use the latest version of the dependencies.

```bash
go mod tidy
```

##### Equivalent of `JSON.stringify()` in Go

-   `json.Marshal()` is used to convert a Go data structure to a JSON string.

```go
import "encoding/json"

func main() {
    data := map[string]interface{}{
        "name": "John Doe",
        "age": 30,
    }

    jsonString, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(jsonString))
}
```

##### Equivalent of `JSON.parse()` in Go
-   `json.Unmarshal()` is used to convert a JSON string to a Go data structure.

```go
import "encoding/json"

func main() {
    jsonString := `{"name":"John Doe","age":30}`

    var data map[string]interface{}
    err := json.Unmarshal([]byte(jsonString), &data)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(data)
}
```

##### Equivalent of `nodemon` in Go

```bash
go install github.com/cosmtrek/air@latest
```

-   `air` is a live reload tool for Go applications.
-   It watches for file changes and automatically rebuilds and restarts the application.
-   It is similar to `nodemon` in the Node.js ecosystem.
-   There are other tools like `fresh` which can also be used for live reloading in Go.

##### Equivalent of `dotenv` in Go

```bash
go get github.com/joho/godotenv
```

-   `godotenv` is a Go package that loads environment variables from a `.env` file.
-   It is similar to `dotenv` in the Node.js ecosystem.
-   It allows developers to store sensitive information like API keys, database URIs, etc., in a `.env` file and load them into the application.

##### Code Example of Using `godotenv`

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  MONGODB_URI := os.Getenv("MONGODB_URI")
}
```

-   In this example, we load environment variables from a `.env` file using `godotenv.Load(".env")`.
-   We can then access the environment variables using `os.Getenv("MONGODB_URI")`.

### Gin

##### Equivalent of `Express.js` in Go

-   Gin is a web framework written in Go.
-   It is similar to Express in the Node.js ecosystem.
-   It provides a simple and elegant way to build web applications in Go.
-   It is fast, lightweight, and easy to use.
-   It is a popular choice for building web applications in Go.
-   It is used by many companies and projects, including Google, Uber, and Docker.
-   It is a good choice for building RESTful APIs, web applications, and microservices.
-   It provides a similar API to Express.js, making it easy for developers familiar with Express.js to transition to Go.
-   Other popular web frameworks in Go include `Gin` and `Echo`.

```bash
go get github.com/gin-gonic/gin
```

##### Equivalent of `Express.js Middleware` in Go

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Define a middleware
    r.Use(func(c *gin.Context) {
        // Middleware logic
        c.Next()
    })

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    r.Run(":3000")
}
```

- In this example, we define a middleware function that uses c.Next() to call the next handler.
- The middleware function can execute some logic before or after the c.Next() call.

##### Equivalent of `Express.js Route Handling` in Go

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    r.Run(":3000")
}
```

- In this example, we define a route handler function that sends a response back to the client using c.String().

##### Equivalent of `Express.js Error Handling` in Go

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    r.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "Page not found")
    })

    r.Run(":3000")
}

```
- In this example, r.NoRoute() is used to handle 404 errors (routes that are not found).
- This is similar to defining a catch-all route in Express.js.

### Fiber

##### Equivalent of `Express.js` in Go

-   `Fiber` is a web framework for Go that is inspired by Express.js.
-   It is fast, lightweight, and easy to use.
-   It provides a similar API to Express.js, making it easy for developers familiar with Express.js to transition to Go.
-   Other popular web frameworks in Go include `Gin` and `Echo`.

```bash
go get github.com/gofiber/fiber/v2
```

##### Equivalent of `Express.js Middleware` in Go

-   In this example, we define a middleware function that takes the `next` handler as an argument.
-   The middleware function wraps the `next` handler and executes some logic before calling the `next` handler.

```go
func main() {
    app := fiber.New()

    app.Use(middleware)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Middleware logic

        // .
        // .
        // .

        next.ServeHTTP(w, r)
    })
}
```


##### Equivalent of `Express.js Route Handling` in Go

```go
func main() {
    app := fiber.New()

    app.Get("/", helloHandler)

    app.Listen(":3000")
}

func helloHandler(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
}
```

-   In this example, we define a route handler function `helloHandler` that sends a response back to the client.
-   We then register the route handler with the `app.Get()` method, specifying the route path and the handler function.
