# Middleware 

Middleware adalah entitas yang terhubung ke pemrosesan permintaan/respons server. Middleware, izinkan kami menerapkan berbagai fungsi seputar Permintaan HTTP masuk ke server Anda dan respons keluar.

## Example Third Party Middleware
* Negroni
* Echo
* Interpose
* Alice

## Middleware in Echo 
* Basic Auth
* Body Dump
* Body Limit
* CORS
* CSRF
* Casbin Auth
* Gzip
* JWT
* Key Auth
* Logger
* Method Override
* Proxy
* Recover
* Redirect
* Request ID
* Rewrite
* Secure
* Session
* Static
* Trailing Slash

## HTTPS Redirect
HTTPS redirect middleware redirects

### Example 
http://labstack.com will be redirected to https://labstack.com

```
e := echo.new()
e.Pre(middleware.HTTPSRedirect())
```

## CORS Middleware
What is CORS :
* CORS stands for Cross Origin Resource Sharing
* Enable resource sharing with different origin/domains
* Typically used in web applications that acess API with different domains

Common CORS Configuration :
* Acess-Control-Allow-Origin -> untuk menentukan domain/origin yang dapat mengirim request ke server
* Acess-Control-Allow-Headers -> untuk menentukan request header yang dapat digunakan ketika mengirim request ke server
* Acess-Control-Allow-Methods -> Untuk menentukan HTTP method yang dapat digunakan ketika mengirim request ke server
* Access-Control-Allow-Methods -> Untuk menentukan HTTP Method yang dapat digunakan ketika mengirim request ke server
* Access-Control-Max-Age -> Untuk menentukan durasi maksimum preflight request yang dapat dilakukan caching

Cors Implementation

```
Package main

import(
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main(){
    e := echo.New()
    // implement CORS middleware
    e.Use(middleware.CORS())

    e.GET("/",func (c echo.Context) error{
        return c.String(http.StatusOK,"hello")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
```

## Rate Limiter Middleware
What is Rate Limiter ?
* Limit the number of request to the server
* Prevent over fetching request to ensure the performance of the server is maintained
* Prevent security breaches like DDos,Brute Force Attack

Rate Limiter Implementation
```
Package main

import(
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main(){
    e := echo.new()
    e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

    e.Get("/",func (c echo.Context)error{
        return c.String(http.statusOK,"hello")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
```

## Log Middleware

Logger Middleware
* Logs the information HTTP request
* As a footprint
* Datasource for analytics

Step :
1. Create Folder Middleware
2. Create Log Middleware
3. Implement Middleware
4. Result

## Auth Middleware
Why Using Authentication ?
* User Identification
* Secure Data and information

Authentication Middleware
* Basic Authentication
* JSON Web Token

What is Basic Authentication?
* Basic Auth is an http authentication technique request, this method requires username and password information to be inserted in the request header.

```
Information Encoded Format :
'Authorization : Basic ' + base64encode('username:password')
Header Generate : 
Authorization : Basic dXN1cm5hbWU6cGFzc3dvcmQ=
```
