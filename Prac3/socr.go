package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
)

//получение IP хоста
func GetMyIP() net.IP {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

var urls = HashMap{}

func main() {

	http.HandleFunc("/", handleForm)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/short/", handleRedirect)

	fmt.Println("URL Shortener is running on :3333")
	http.ListenAndServe("0.0.0.0:3333", nil)

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	//главная страница
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Shortener</title>
		<style>
			body {
				display: flex;
				align-items: center;
				justify-content: center;
				height: 100vh;
				margin: 10px;
			}
			h2 {
			  margin-top: -15px;
			}

	
			form {
				text-align: centre;
			}
	
			input[type="url"] {
				width: 80%;
				padding: 0px;
				margin-bottom: 10px;
			}
	
			input[type="url"] {
				padding: 10px;
			}
		</style>
	</head>
	<body>
		<h2>Shortener</h2>
		<form method="post" action="/shorten">
			<input type="url" name="url" placeholder="Enter a URL" required>
			<br>
			<input type="submit" value="Shorten">
		</form>
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

	shortKey := generateShortKey()
	err := urls.Hadd(shortKey, originalURL)
	if err != nil {
		http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
		return
	}
	//urls[shortKey] = originalURL
	HostIp := GetMyIP()
	shortenedURL := fmt.Sprintf("/short/%s", shortKey)

	//Страница с результатом
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Shorted</title>
		</head>
		<body>
			<h2>Shorted</h2>
			<p>Original URL: `, originalURL, `</p>
			<p>Shorted URL: <a href="`, shortenedURL, `">`, fmt.Sprintf("%s:3333%s", HostIp, shortenedURL), `</a></p>
		</body>
		</html>
	`)
}

//
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	//получение оригинального url по ключу из хэштаблицы
	originalURL, err := urls.Hget(shortKey)
	if err != nil {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

//генерация короткого ключа на основе рандома длинны и символов
func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyLength := rand.Intn(6) + 1

	//rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
