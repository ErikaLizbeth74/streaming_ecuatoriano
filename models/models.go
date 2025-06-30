package models

import "time" // ¡Asegúrate de que 'time' esté importado para 'time.Time'!

// Content representa una película, serie o documental.
type Content struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
	Length   int    `json:"length"` // Duración en minutos
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
	IsFree   bool   `json:"is_free"` // true si es gratuito, false si es premium
}

// ItemCarrito representa un elemento individual en el carrito de compras.
type ItemCarrito struct {
	PeliculaID int     `json:"pelicula_id"`
	Titulo     string  `json:"titulo"`
	ImageURL   string  `json:"image_url"`
	Precio     float64 `json:"precio"`
	Cantidad   int     `json:"cantidad"`
}

// Carrito representa la estructura completa del carrito de compras.
type Carrito struct {
	Items []ItemCarrito `json:"items"`
	Total float64       `json:"total"`
}

// User representa la estructura de un usuario registrado.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // En una app real, esto debería ser un hash
	Email    string `json:"email"`
}

// UserSession representa una sesión de usuario activa.
// ¡Verifica que esta estructura esté presente y bien escrita!
type UserSession struct {
	UserID    int
	Username  string
	CreatedAt time.Time // Asegúrate de que 'time.Time' esté bien escrito y 'time' importado
}
