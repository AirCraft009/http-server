 http-server - a basic implementation of a http-server
-
### Usage:

**Setup**
  - create a new server with `httpServer := server.NewServer(port int)`
  - if you want the base server to be able to serve Get requests to specific files
    - call `httpServer.AddFileSystem(path_to_folder)` relative or absoloute
    - now all subfolders and files will have specific routes.
    - make sure that you don't have duplicate files because they will override each other
    - you can override all premade with custom ones for your needs
  - now call `httpServer.ListenAndServe()`


**Using Builtin Handlers**
  - there are some builtin handlers like the base `Http404Handler()` 
    - it pulls 404.html from your sourcefolder if it doesn't exist it will only show a message

**Custom Handlers**
  - configuring custom handlers is easy they leave you in full control of the response
  - most important headers and the first line are already prewritten
    - startline; Content-lenght; Connection
  - you can still overwrite them if you want to.
  - all the data sent should be in the `Response.Body []byte`


**Adding Handlers**
  - adding is simple just call `httpServer.Handle(method, route, string handler Handler)`
