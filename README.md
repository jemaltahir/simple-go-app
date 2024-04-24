# Simple Golang Web Server with Docker

This project provides a simple Golang web server Go application using Docker.

## Requirements

- Docker installed ([Get Docker](https://www.docker.com/))

## Project Structure

- Dockerfile: Defines the instructions for building the Docker image.
- go.mod: Declares project dependencies.
- main.go: The Go source code for the web server.

## Building the Image

Open your terminal and navigate to the project directory.

Run the following command:

```bash
docker build -t my-golang-web-app .
```

```bash
docker run -p 8080:8080 my-golang-web-app
```

Open `http://localhost:8080/` in your web browser. You should see "Hello, World!".

To access the endpoint with a name, use:

`http://localhost:8080/hello/john`