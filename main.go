package main

import (
	"log"
	"net/http"

	"streaming_ecuatorianos/handlers" // Asegúrate de que este sea el nombre de tu módulo
)

func main() {
	// Servir archivos estáticos (CSS, JS, imágenes)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rutas para las páginas HTML (manejadores que sirven templates)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/catalog", handlers.CatalogHandler)
	http.HandleFunc("/cart", handlers.CartHandler)
	http.HandleFunc("/login", handlers.LoginHandler)           // Manejador para la página HTML de login
	http.HandleFunc("/register", handlers.RegisterPageHandler) // Ruta para mostrar la página de registro
	http.HandleFunc("/movie-detail", handlers.MovieDetailHandler)
	http.HandleFunc("/player", handlers.PlayerHandler)

	// Rutas para la API (JSON - manejadores que devuelven datos JSON)

	// API de Películas
	http.HandleFunc("/api/peliculas", handlers.GetPeliculasAPI)
	http.HandleFunc("/api/peliculas/", handlers.GetPeliculaByIDAPI) // Maneja /api/peliculas/{id}
	http.HandleFunc("/api/peliculas/genero/", handlers.GetPeliculasByGeneroAPI)
	http.HandleFunc("/api/peliculas/year/", handlers.GetPeliculasByYearAPI)
	http.HandleFunc("/api/destacados", handlers.GetDestacadosAPI) // API para películas destacadas

	// --- LAS SIGUIENTES LÍNEAS FUERON LAS CAUSANTES DEL ERROR Y HAN SIDO ELIMINADAS O CORREGIDAS ---
	// La línea comentada de SearchPeliculasAPI era el error anterior, se elimina definitivamente.
	// Las líneas de 'estadisticas' y 'LoginAPI', 'RegisterAPI', 'AddTocartAPI', etc.
	// tenian typos o no corresponden a funciones existentes en handlers.go.
	// Se han reemplazado por las llamadas correctas a las funciones ya definidas.

	// Rutas de Autenticación
	http.HandleFunc("/api/auth/register", handlers.RegisterHandler) // Correcto: RegisterHandler
	http.HandleFunc("/api/auth/login", handlers.LoginAPIHandler)    // Correcto: LoginAPIHandler (el renombrado)
	http.HandleFunc("/api/auth/logout", handlers.LogoutHandler)     // Correcto: LogoutHandler

	// Rutas del Carrito de compras
	http.HandleFunc("/api/carrito", handlers.GetCart)                 // Correcto: GetCart
	http.HandleFunc("/api/carrito/agregar", handlers.AddToCart)       // Correcto: AddToCart
	http.HandleFunc("/api/carrito/eliminar", handlers.RemoveFromCart) // Correcto: RemoveFromCart
	http.HandleFunc("/api/carrito/limpiar", handlers.ClearCart)       // Correcto: ClearCart
	http.HandleFunc("/api/carrito/checkout", handlers.CheckoutCart)   // Correcto: CheckoutCart

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
