package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"encoding/json"
	"time"
)


type Log struct {
	IP   string
	URL    string
	Time  string
}

//получение IP хоста
func GetMyIP() net.IP {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func sendCommandToDatabase(mainCommand, key, value string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		return "", fmt.Errorf("failed to connect to the database server: %s", err)
	}
	defer conn.Close()
	switch mainCommand{
	case "HPUSH":
			command := fmt.Sprintf("%s %s %s", mainCommand, key, value)
			_, err = conn.Write([]byte(command))
	case "HGET":
		command := fmt.Sprintf("%s %s", mainCommand, key)
		_, err = conn.Write([]byte(command))
	case "QPUSH":
		command:=fmt.Sprintf("%s %s",mainCommand,key)
		_, err = conn.Write([]byte(command))
	}
	responseBuf := make([]byte, 1024)
	n, err := conn.Read(responseBuf)
	if err != nil {
		return "", fmt.Errorf("failed to read response from the database server: %s", err)
	}

	return string(responseBuf[:n]), nil
}




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
	_,err := sendCommandToDatabase("HPUSH", shortKey, originalURL)
	if err != nil {
		http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
		return
	}
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

	originalURL, err := sendCommandToDatabase("HGET", shortKey, "")
	originalURL = strings.TrimPrefix(originalURL, shortKey+" ")
	fmt.Println(originalURL)
	if err != nil {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	currentTime := time.Now().Format("2006-01-02")
	IP := r.RemoteAddr

	log := Log{
		IP:   IP,
		URL:  originalURL,
		Time: currentTime,
	}

	logJSON, err := json.Marshal(log)
	if err != nil {
		http.Error(w, "Failed to create log entry", http.StatusInternalServerError)
		return
	}

	_, err = sendCommandToDatabase("QPUSH", string(logJSON),"")
	if err != nil {
		http.Error(w, "Failed to save log entry", http.StatusInternalServerError)
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
