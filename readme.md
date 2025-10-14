 http-server - a basic implementation of a http-server
-
### Usage:

**Setup**
  - create a new server with `httpServer := server.NewServer(port int)`
  - if you want the base server to be able to serve Get requests to specific files
    - call `httpServer.AddFileSystem(path_to_folder)` relative or absoloute
    - or call `httpServer.AddFileSystemWithHandler(path, handler)` to use your own handler
    - now all subfolders and files will have specific routes.
    - make sure that you don't have duplicate files because they will override each other
    - you can override all premade with custom ones for your needs
  - now call `httpServer.ListenAndServe()`


**Using Builtin Handlers**
  - there are some builtin handlers like the base `Http404Handler()` 
    - it pulls 404.html from your sourcefolder if it doesn't exist it will only show a message
  - there is also the base `StreamHandler` it streams any data in the filesystem
    - it only works with GET req.
    - you can also use a custom handler with `AddFileSystemWithHandler`

**Overwriting the 404 Basehandler**

- the concept is just like mentioned above.
- but keep in mind the 404-Handler has to set all it's headers like 
  - content lenght
  - connection
  - content-type

**Custom Handlers**
  - configuring custom handlers is easy they leave you in full control of the response
  - most important headers and the first line are already prewritten
    - startline; Content-lenght; Connection
  - you can still overwrite them if you want to.
  - all the data sent should be in the `Response.Body []byte`#
    - best set with `response.SetBody([]byte"boring text or other data")` 

```Go 
func Homehandler(req *parser.Request) (res *server.Response) {
	res = server.NewResponse(req)
	res.StatusCode = http.StatusOK
	res.AddHeader("Content-Type", "text/html; charset=utf-8")
	res.AddHeader("Connection", "keep-alive")
	res.SetBody([]byte("Hello World"))
	return res
}
```



**Adding Handlers**
  - adding is simple just call `httpServer.Handle(method, route, string handler Handler)`
