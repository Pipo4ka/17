package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/set-cookie", setCookieHandler)
	http.HandleFunc("/read-cookie", readCookieHandler)
	http.HandleFunc("/delete-cookie", deleteCookieHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Работа с Cookie на Go</title>
		</head>
		<body>
			<h1>Добро пожаловать!</h1>
			<form action="/set-cookie" method="POST">
				<label for="username">Введите ваше имя:</label>
				<input type="text" id="username" name="username" required>
				
				<label for="language">Выберите язык:</label>
				<select id="language" name="language" required>
					<option value="ru">Русский</option>
					<option value="en">Английский</option>
					<option value="fr">Французский</option>
				</select>
				
				<button type="submit">Сохранить Cookie</button>
			</form>
			<a href="/read-cookie">Читать Cookie</a> | 
			<a href="/delete-cookie">Удалить Cookie</a>
		</body>
		</html>
	`
	w.Write([]byte(html))
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		language := r.FormValue("language")

		// Устанавливаем cookie на 7 дней
		expiration := time.Now().Add(7 * 24 * time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   username,
			Expires: expiration,
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "language",
			Value:   language,
			Expires: expiration,
		})

		http.Redirect(w, r, "/read-cookie", http.StatusSeeOther)
	}
}

func readCookieHandler(w http.ResponseWriter, r *http.Request) {
	usernameCookie, err1 := r.Cookie("username")
	languageCookie, err2 := r.Cookie("language")

	if err1 != nil || err2 != nil {
		w.Write([]byte("Имя пользователя и язык не найдены. Пожалуйста, задайте их."))
		return
	}

	message := fmt.Sprintf("Здравствуйте, %s! Ваш выбранный язык: %s.", usernameCookie.Value, languageCookie.Value)
	w.Write([]byte(message))
}

func deleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем cookie с истекшим сроком действия
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "language",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})

	w.Write([]byte("Cookie успешно удалены!"))
}
