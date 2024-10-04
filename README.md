## Go Web Application Example

This project demonstrates a simple web application in Go that serves HTML content using dynamic templates. The application listens on port 8080 and serves a web page containing a list of films, displaying their titles and directors using Go's html/template package. It showcases the use of data-driven templates, dynamic content rendering, and basic HTTP server functionality.

# Getting Started

- Prerequisites
  Go installed on your machine (version 1.16+ recommended).
  An HTML file named index.html in the same directory as the Go source code.

  - Project Structure
    /project-root
    ├── main.go # Go source code for the web server
    ├── index.html # HTML template file to render

# Code Overview

1. Film struct: The Film struct is used to represent film data with fields for the title and director.
2. HTML Template Rendering: The code dynamically injects the film data into the HTML template using Go's html/template package. This allows the web page to be updated with film information when the server receives a request.
3. Serving Requests: The http.HandleFunc function is used to define a handler for the root URL (/), and the http.ListenAndServe function starts the server listening on port 8080.
