# url-shortner
A simple URL shortener built with Go, allowing users to convert long URLs into short, unique links and redirect users to the original URL when the short link is accessed. This project demonstrates the fundamentals of HTTP servers, request handling, URL redirection, and key generation in Go.

Features:

Serve a user-friendly form for URL input.
Generate short, random alphanumeric keys for URLs.
Store mappings in memory using a map.
Redirect users to the original URL using the shortened link.
Technologies Used:

Go's net/http package for server and routing.
Random key generation with math/rand.
Limitations:

URLs are stored in memory (data is lost on server restart).
Not suitable for production without additional improvements like persistent storage, HTTPS, and scalability enhancements.

To create a new go.mod file and initialize a Go module in your project folder, follow these steps:

**1. Navigate to Your Project Directory**
Open a terminal and navigate to your project directory where you want to initialize the Go module:
eg: cd /path/to/your/project

**2. Initialize the Go Module**
Run the following command to create a new go.mod file:
eg: go mod init <<module-name>>
