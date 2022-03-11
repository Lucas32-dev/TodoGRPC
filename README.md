# TodoGRPC

## What is GRPC?

Basically, in a GRPC application, the client can call directly methods on a remote server application, working like a local object.
How? You can define multiple services in a proto file, specifying the methods that can be called with their parameters and return type. The serve implements this interface an runs an GRPC server, and the client has a stub that provides the same methods. And this magic happens through Proto Requests :D

## Project application

This project is an simple TODO project, where the server holds an list of todo and their state, and client calls methods to manipulate it.

## Run project

Install Go 1.17

Set up the server

```bash
go run server/main.go
```

Run client

```bash
go run client/main.go
```

Feel free to play around with the client.  
You can check the proto file to see which methods are implemened in Todo service
