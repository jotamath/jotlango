# HTTP Library

The `http` library provides HTTP server and client functionality for JotLang.

## Import

```jt
import "http"
```

## Classes

### Request

```jt
class Request {
    prop string Method
    prop string Path
    prop map<string, string> Headers
    prop map<string, string> Query
    prop any Body
}
```

Represents an HTTP request.

### Response

```jt
class Response {
    prop int Status
    prop map<string, string> Headers
    prop any Body

    fn json(data any) : void
    fn text(content string) : void
    fn html(content string) : void
}
```

Represents an HTTP response with methods for different content types.

### Server

```jt
class Server {
    prop int Port
    prop map<string, fn(Request):Response> Routes

    fn listen() : void
    fn get(path string, handler fn(Request):Response) : void
    fn post(path string, handler fn(Request):Response) : void
    fn put(path string, handler fn(Request):Response) : void
    fn delete(path string, handler fn(Request):Response) : void
}
```

HTTP server implementation with routing capabilities.

## Functions

### Client Functions

```jt
fn get(url string) : Response
```
Performs an HTTP GET request.

```jt
fn post(url string, body any) : Response
```
Performs an HTTP POST request.

```jt
fn put(url string, body any) : Response
```
Performs an HTTP PUT request.

```jt
fn delete(url string) : Response
```
Performs an HTTP DELETE request.

## Examples

### Basic Server

```jt
import "http"
import "io"

// Create server
server = new http.Server {
    Port = 8080,
    Routes = {}
}

// Add routes
server.get("/hello", fn(req http.Request) : http.Response {
    resp = new http.Response {
        Status = 200,
        Headers = {}
    }
    resp.text("Hello, World!")
    return resp
})

// Start server
io.println("Server started on port 8080")
call server.listen()
```

### REST API

```jt
import "http"

class UserApi {
    prop map<string, any> Users

    fn getUsers(req http.Request) : http.Response {
        resp = new http.Response {
            Status = 200,
            Headers = {}
        }
        resp.json(Users)
        return resp
    }

    fn createUser(req http.Request) : http.Response {
        user = req.Body
        Users[user.id] = user

        resp = new http.Response {
            Status = 201,
            Headers = {}
        }
        resp.json(user)
        return resp
    }
}

// Create API
api = new UserApi {
    Users = {}
}

// Setup server
server = new http.Server {
    Port = 8080,
    Routes = {}
}

server.get("/users", api.getUsers)
server.post("/users", api.createUser)

call server.listen()
```

### HTTP Client

```jt
import "http"
import "io"

// GET request
response = http.get("https://api.example.com/users")
if response.Status == 200 {
    io.println(response.Body)
}

// POST request
user = {
    "name" = "John Doe",
    "email" = "john@example.com"
}
response = http.post("https://api.example.com/users", user)
```

## Best Practices

1. Always set appropriate status codes
2. Include proper headers for content types
3. Handle errors gracefully
4. Use appropriate HTTP methods
5. Validate request data
6. Implement proper error responses
7. Use meaningful route paths
8. Document your API endpoints 