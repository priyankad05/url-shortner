package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urls = make(map[string]string)

func main() {
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/short/", handleRedirect)

	fmt.Println("URL Shortener is running on :3030")
	http.ListenAndServe(":3030", nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	// Serve the HTML form
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>URL Shortener</title>
</head>
<body>
    <header style="text-align: center;">
        <h1 style="font-family: Arial, sans-serif;">URL Shortener</h1>
        <p><strong>Simplify your links in seconds!</strong></p>
    </header>
    <main style="display: flex; justify-content: center; margin-top: 20px;">
        <form method="post" action="/shorten" style="text-align: center;">
            <label for="url" style="font-family: Verdana, sans-serif; font-size: 1.1em;"><strong>Enter a URL:</strong></label><br><br>
            <input 
                type="url" 
                id="url" 
                name="url" 
                placeholder="https://example.com" 
                required 
                style="padding: 8px; width: 300px; font-family: Arial, sans-serif;"
            ><br><br>
            <input 
                type="submit" 
                value="Shorten" 
                style="padding: 8px 16px; font-family: Arial, sans-serif; font-weight: bold; cursor: pointer;"
            >
        </form>
    </main>
    <footer style="text-align: center; margin-top: 40px; font-family: Arial, sans-serif; font-size: 0.9em; color: gray;">
        <p>&copy; 2024 URL Shortener. All rights reserved.</p>
    </footer>
</body>
</html>

	`)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate a unique shortened key for the original URL
	shortKey := generateShortKey()
	urls[shortKey] = originalURL

	// Construct the full shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:3030/short/%s", shortKey)

	// Serve the result page
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>URL Shortener</title>
</head>
<body>
    <header style="text-align: center;">
        <h1 style="font-family: Arial, sans-serif;">URL Shortener</h1>
    </header>
    <main style="margin: 20px auto; max-width: 600px; font-family: Arial, sans-serif;">
        <p><strong>Original URL:</strong> <span style="font-family: Courier New, monospace;">`, originalURL, `</span></p>
        <p>
            <strong>Shortened URL:</strong> 
            <a 
                href="`, shortenedURL, `" 
                target="_blank" 
                style="color: blue; text-decoration: underline;"
            >
                `, shortenedURL, `
            </a>
        </p>
    </main>
</body>
</html>

	`)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := urls[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// Redirect the user to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey { //0 to keylen -1
		shortKey[i] = charset[rand.Intn(len(charset))] // between 0 to len(charset)-1
	}
	return string(shortKey) //from byte array

}
