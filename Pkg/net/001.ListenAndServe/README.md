ListenAndServe
===

## Handler Function
- The helloHandler function is defined to handle HTTP requests. It takes an http ResponseWriter and an *http.Request as parameters.
- Inside the handler, fmt.Fprintf writes "Hello, World!" to the response writer, which will be sent back to the client.


## http.ListenAndServe
- The http.ListenAndServe function in Go is used to start an HTTP server that listens on a specified address and handles incoming requests using a specified handler. Here’s the function signature:

## Function Signature
  ```
  func ListenAndServe(addr string, handler Handler) error

  ```
  - addr: The TCP network address to listen on, in the form of host:port.
  - handler: The HTTP handler to invoke for each incoming request. If handler is nil, 
    http.DefaultServeMux is used.

## Register Handler Function : 
- In the main function, http.HandleFunc("/hello", helloHandler) is called to register the helloHandler function for the /hello URL path.
- This means that any HTTP request to http://localhost:8080/hello will be handled by the helloHandler function.
  
## Starting the server : 
- http.ListenAndServe(":8080", nil) starts the HTTP server on port 8080.
- The first argument ":8080" specifies the network address to listen on. The empty host "" means all available interfaces.
- The second argument is nil, which means the server will use http.DefaultServeMux as the handler. Since we registered the /hello route with http.HandleFunc, it is added to http.DefaultServeMux

## Error Handling 
- http.ListenAndServe returns an error if it fails to start the server. The error is checked, and if not nil, it prints an error message.

### Running the Code : 
- To run the example, save the code to a file
- Run it using go run command
- You should see the message "Starting server on :8080". 
- Open a web browser and navigate to http://localhost:8080/hello. 
- You should see "Hello Gophers !!!!!!! " displayed in the browser.
```
```
