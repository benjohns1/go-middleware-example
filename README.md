# Go Middleware Example

A collection of basic HTTP middleware implementation examples in Go

1. [loopchain](https://github.com/benjohns1/go-middleware-example/tree/master/loopchain/main.go): An extremely simple middleware implementation using a for loop
1. [recursivechain](https://github.com/benjohns1/go-middleware-example/tree/master/recursivechain/main.go): Middleware implementation using recursion that allows for a short-circuiting the middleware chain
1. [finalhandler](https://github.com/benjohns1/go-middleware-example/tree/master/finalhandler/main.go): Variant on the recursive implementation, but using a final logic handler in the chain as the first chain parameter
1. [customstate](https://github.com/benjohns1/go-middleware-example/tree/master/customstate/main.go): Middleware implementation using recursion with a custom http.ResponseWriter implementation to manage custom state along the request chain
1. [apiexample](https://github.com/benjohns1/go-middleware-example/tree/master/apiexample): A slightly more organized example JSON API using previous middleware techniques
