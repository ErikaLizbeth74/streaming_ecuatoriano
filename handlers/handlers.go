package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"streaming_ecuatorianos/models" // ¡Asegúrate de que este sea el nombre de tu módulo!
)

// sendJSONResponse es una función de ayuda para enviar respuestas JSON.
func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

// --- Datos en memoria (simulando una base de datos global para todos los manejadores) ---

var apiPeliculas = []models.Content{
	{
		ID: 1, Title: "Ratas, ratones y rateros", Genre: "Drama", Year: 1999, Length: 108,
		ImageURL: "/static/img/movies/ratas_ratones_y_rateros.jpg",
		VideoURL: "https://player.vimeo.com/video/381375898?badge=0&autopause=0&player_id=0&app_id=58479",
		IsFree:   true,
	},
	{
		ID: 2, Title: "Qué tan lejos", Genre: "Comedia dramática", Year: 2006, Length: 92,
		ImageURL: "/static/img/movies/que_tan_lejos.jpg",
		VideoURL: "https://www.youtube.com/embed/ymJE6_32_bQ?si=Ar24MiL5pqCuZ3qo", // Ejemplo de URL de YouTube (Rickroll)
		IsFree:   false,
	},
	{
		ID: 3, Title: "Alba", Genre: "Drama", Year: 2016, Length: 94,
		ImageURL: "/static/img/movies/alba.jpg",
		VideoURL: "https://www.youtube.com/embed/yMEMHutUhOY?si=tE_gfhyUMiURhuYu", // Ejemplo
		IsFree:   true,
	},
	{
		ID: 4, Title: "La Tigra", Genre: "Drama", Year: 1990, Length: 100,
		ImageURL: "/static/img/movies/la_tigra.jpg",
		VideoURL: "https://geo.dailymotion.com/player.html?video=x94rs4i", // Ejemplo
		IsFree:   false,
	},
	{
		ID: 5, Title: "En el nombre de la hija", Genre: "Drama", Year: 2011, Length: 100,
		ImageURL: "/static/img/movies/en_el_nombre_de_la_hija.jpg",
		VideoURL: "https://www.youtube.com/embed/DoXUeTEIjl8?si=vLWGUCe2XQH5FP8D", // Ejemplo
		IsFree:   true,
	},
	{
		ID: 6, Title: "Un secreto en la caja", Genre: "Documental", Year: 2016, Length: 70,
		ImageURL: "/static/img/movies/un_secreto_en_la_caja.jpg",
		VideoURL: "https://geo.dailymotion.com/player.html?video=x96f44y", // Ejemplo
		IsFree:   false,
	},
	{
		ID: 7, Title: "Con mi corazón en Yambo", Genre: "Documental", Year: 2011, Length: 120,
		ImageURL: "/static/img/movies/yambo.jpg",
		VideoURL: "https://www.youtube.com/embed/-A94fsVBE1c?si=FuEJXHSoIMLqSlSj", // Ejemplo
		IsFree:   true,
	},
}

var carritoEnMemoria = models.Carrito{
	Items: make([]models.ItemCarrito, 0),
	Total: 0.0,
}
var cartMutex sync.Mutex

var dummyUsers = []models.User{
	{ID: 1, Username: "usuario1", Password: "password1", Email: "usuario1@example.com"},
	{ID: 2, Username: "admin", Password: "adminpass", Email: "admin@example.com"},
	{ID: 3, Username: "ejemplo", Password: "123", Email: "ejemplo@ejemplo.com"},
}

// --- Manejadores de Páginas HTML (Serven los templates) ---

// funcMap es el mapa de funciones personalizadas para las plantillas.
var funcMap = template.FuncMap{
	"lower": strings.ToLower, // Registra la función "lower"
}

// HomeHandler sirve la página principal.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/home.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para HomeHandler: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// CatalogHandler sirve la página HTML del catálogo.
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/catalog.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para CatalogHandler: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", apiPeliculas)
}

// CartHandler sirve la página HTML del carrito.
func CartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/cart.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para CartHandler: %v", err)
		http.Error(w, "Internal Server Error al cargar la página del carrito", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// LoginHandler sirve la página HTML de inicio de sesión. (Nombre de función HTML)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/login.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para LoginHandler: %v", err)
		http.Error(w, "Internal Server Error al cargar la página de inicio de sesión", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// RegisterPageHandler sirve la página HTML de registro.
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/register.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para RegisterPageHandler: %v", err)
		http.Error(w, "Internal Server Error al cargar la página de registro", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// PlayerHandler sirve la página HTML para la reproducción simulada de películas.
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/player.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para PlayerHandler: %v", err)
		http.Error(w, "Internal Server Error al cargar la página del reproductor", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// MovieDetailHandler sirve la página HTML para los detalles de la película antes de la reproducción.
func MovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/movie_detail.html")
	if err != nil {
		log.Printf("Error al analizar plantillas para MovieDetailHandler: %v", err)
		http.Error(w, "Internal Server Error al cargar la página de detalles de la película", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// --- Servicios Web Principales (APIs de Contenido - JSON) ---

// GetPeliculasAPI lista todas las películas.
func GetPeliculasAPI(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, apiPeliculas, http.StatusOK)
}

// GetPeliculaByIDAPI obtiene detalles de la película por ID.
func GetPeliculaByIDAPI(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/peliculas/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJSONResponse(w, map[string]string{"error": "ID de película inválido"}, http.StatusBadRequest)
		return
	}

	for _, p := range apiPeliculas {
		if p.ID == id {
			sendJSONResponse(w, p, http.StatusOK)
			return
		}
	}
	sendJSONResponse(w, map[string]string{"message": "Película no encontrada"}, http.StatusNotFound)
}

// GetPeliculasByGeneroAPI filtra películas por género.
func GetPeliculasByGeneroAPI(w http.ResponseWriter, r *http.Request) {
	genero := strings.TrimPrefix(r.URL.Path, "/api/peliculas/genero/")
	var filteredPeliculas []models.Content
	for _, p := range apiPeliculas {
		if strings.EqualFold(p.Genre, genero) {
			filteredPeliculas = append(filteredPeliculas, p)
		}
	}

	if len(filteredPeliculas) == 0 {
		sendJSONResponse(w, map[string]string{"message": fmt.Sprintf("No se encontraron películas para el género '%s'", genero)}, http.StatusNotFound)
		return
	}
	sendJSONResponse(w, filteredPeliculas, http.StatusOK)
}

// GetPeliculasByYearAPI filtra películas por año.
func GetPeliculasByYearAPI(w http.ResponseWriter, r *http.Request) {
	yearStr := strings.TrimPrefix(r.URL.Path, "/api/peliculas/year/")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		sendJSONResponse(w, map[string]string{"error": "Año inválido"}, http.StatusBadRequest)
		return
	}

	var filteredPeliculas []models.Content
	for _, p := range apiPeliculas {
		if p.Year == year {
			filteredPeliculas = append(filteredPeliculas, p)
		}
	}

	if len(filteredPeliculas) == 0 {
		sendJSONResponse(w, map[string]string{"message": fmt.Sprintf("No se encontraron películas para el año '%d'", year)}, http.StatusNotFound)
		return
	}
	sendJSONResponse(w, filteredPeliculas, http.StatusOK)
}

// GetDestacadosAPI devuelve una lista de películas destacadas (ej., las primeras 3).
func GetDestacadosAPI(w http.ResponseWriter, r *http.Request) {
	featured := apiPeliculas
	if len(featured) > 3 { // Ajusta este número según necesites
		featured = apiPeliculas[:3]
	}
	sendJSONResponse(w, featured, http.StatusOK)
}

// --- Autenticación y Gestión de Usuarios ---

// Almacén de sesiones en memoria (simple, sin JWT por ahora)
var sessions = make(map[string]models.UserSession)
var sessionMutex sync.Mutex

// RegisterHandler maneja el registro de usuarios.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, map[string]string{"error": "Método no permitido"}, http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		sendJSONResponse(w, map[string]string{"error": "Datos inválidos"}, http.StatusBadRequest)
		return
	}

	// Validación básica
	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" {
		sendJSONResponse(w, map[string]string{"error": "Todos los campos son requeridos"}, http.StatusBadRequest)
		return
	}

	// Verificar si el nombre de usuario o el email ya existen
	for _, user := range dummyUsers {
		if user.Username == newUser.Username {
			sendJSONResponse(w, map[string]string{"error": "El nombre de usuario ya existe"}, http.StatusConflict)
			return
		}
		if user.Email == newUser.Email {
			sendJSONResponse(w, map[string]string{"error": "El correo electrónico ya está registrado"}, http.StatusConflict)
			return
		}
	}

	// Añadir nuevo usuario (en una app real, hashear la contraseña y almacenarla de forma segura)
	newUser.ID = len(dummyUsers) + 1 // Asignación simple de ID
	dummyUsers = append(dummyUsers, newUser)

	sendJSONResponse(w, map[string]string{"message": "Registro exitoso"}, http.StatusCreated)
}

// LoginAPIHandler maneja el inicio de sesión de usuarios (API POST).
func LoginAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, map[string]string{"error": "Método no permitido"}, http.StatusMethodNotAllowed)
		return
	}

	var loginCreds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginCreds); err != nil {
		sendJSONResponse(w, map[string]string{"error": "Datos inválidos"}, http.StatusBadRequest)
		return
	}

	var foundUser *models.User
	for i := range dummyUsers { // Usar índice para obtener un puntero
		if dummyUsers[i].Username == loginCreds.Username && dummyUsers[i].Password == loginCreds.Password {
			foundUser = &dummyUsers[i]
			break
		}
	}

	if foundUser == nil {
		sendJSONResponse(w, map[string]string{"error": "Credenciales inválidas"}, http.StatusUnauthorized)
		return
	}

	// Crear sesión
	sessionID := fmt.Sprintf("%d-%s", foundUser.ID, time.Now().Format("20060102150405"))
	sessionMutex.Lock()
	sessions[sessionID] = models.UserSession{UserID: foundUser.ID, Username: foundUser.Username, CreatedAt: time.Now()}
	sessionMutex.Unlock()

	// Establecer cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour), // Sesión válida por 24 horas
		HttpOnly: true,                           // No accesible vía JavaScript
		Path:     "/",                            // Accesible en todo el sitio
	})

	sendJSONResponse(w, map[string]string{"message": "Inicio de sesión exitoso", "session_id": sessionID}, http.StatusOK)
}

// LogoutHandler maneja el cierre de sesión de usuarios.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		sendJSONResponse(w, map[string]string{"message": "No hay sesión activa"}, http.StatusOK)
		return
	}

	sessionMutex.Lock()
	delete(sessions, cookie.Value)
	sessionMutex.Unlock()

	// Expirar la cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Unix(0, 0), // Establecer fecha de expiración en el pasado
		HttpOnly: true,
		Path:     "/",
	})

	sendJSONResponse(w, map[string]string{"message": "Sesión cerrada exitosamente"}, http.StatusOK)
}

// --- APIs del Carrito de Compras ---

// AddToCart añade un ítem al carrito en memoria.
func AddToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, map[string]string{"error": "Método no permitido"}, http.StatusMethodNotAllowed)
		return
	}

	var item models.ItemCarrito
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		sendJSONResponse(w, map[string]string{"error": "Datos de ítem inválidos"}, http.StatusBadRequest)
		return
	}

	cartMutex.Lock()
	defer cartMutex.Unlock()

	// Verificar si el ítem ya existe en el carrito, actualizar cantidad o añadir nuevo
	found := false
	for i, existingItem := range carritoEnMemoria.Items {
		if existingItem.PeliculaID == item.PeliculaID {
			carritoEnMemoria.Items[i].Cantidad++ // Incrementar cantidad
			found = true
			break
		}
	}
	if !found {
		item.Cantidad = 1 // Establecer cantidad inicial
		carritoEnMemoria.Items = append(carritoEnMemoria.Items, item)
	}

	carritoEnMemoria.Total += item.Precio
	sendJSONResponse(w, map[string]interface{}{"message": "Ítem añadido al carrito", "cart": carritoEnMemoria}, http.StatusOK)
}

// GetCart obtiene el estado actual del carrito en memoria.
func GetCart(w http.ResponseWriter, r *http.Request) {
	cartMutex.Lock()
	defer cartMutex.Unlock()
	sendJSONResponse(w, carritoEnMemoria, http.StatusOK)
}

// RemoveFromCart elimina un ítem del carrito en memoria.
func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, map[string]string{"error": "Método no permitido"}, http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		PeliculaID int `json:"pelicula_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		sendJSONResponse(w, map[string]string{"error": "Datos inválidos"}, http.StatusBadRequest)
		return
	}

	cartMutex.Lock()
	defer cartMutex.Unlock()

	newItems := []models.ItemCarrito{}
	itemRemoved := false
	for _, item := range carritoEnMemoria.Items {
		if item.PeliculaID == payload.PeliculaID {
			// Disminuir cantidad o eliminar completamente
			if item.Cantidad > 1 {
				item.Cantidad--
				carritoEnMemoria.Total -= item.Precio
				newItems = append(newItems, item)
			} else { // Cantidad es 1, así que lo elimina
				carritoEnMemoria.Total -= item.Precio
			}
			itemRemoved = true
		} else {
			newItems = append(newItems, item)
		}
	}
	carritoEnMemoria.Items = newItems

	if !itemRemoved {
		sendJSONResponse(w, map[string]string{"error": "Ítem no encontrado en el carrito"}, http.StatusNotFound)
		return
	}
	sendJSONResponse(w, map[string]interface{}{"message": "Ítem eliminado del carrito", "cart": carritoEnMemoria}, http.StatusOK)
}

// ClearCart vacía el carrito en memoria.
func ClearCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendJSONResponse(w, map[string]string{"error": "Método no permitido"}, http.StatusMethodNotAllowed)
		return
	}

	cartMutex.Lock()
	defer cartMutex.Unlock()

	carritoEnMemoria.Items = make([]models.ItemCarrito, 0)
	carritoEnMemoria.Total = 0.0
	sendJSONResponse(w, map[string]string{"message": "Carrito vaciado"}, http.StatusOK)
}

// CheckoutCart simula un proceso de pago.
func CheckoutCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, map[string]string{"error": "Método no permitido"}, http.StatusMethodNotAllowed)
		return
	}

	cartMutex.Lock()
	defer cartMutex.Unlock()

	if len(carritoEnMemoria.Items) == 0 {
		sendJSONResponse(w, map[string]string{"error": "El carrito está vacío"}, http.StatusBadRequest)
		return
	}

	// Simular creación de orden
	orderID := time.Now().UnixNano() // ID único para la orden

	// En una aplicación real, harías:
	// 1. Procesar el pago
	// 2. Almacenar la orden en una base de datos
	// 3. Asociar el contenido comprado con la cuenta del usuario
	// 4. Vaciar el carrito

	// Por ahora, solo vacía el carrito y devuelve éxito
	carritoEnMemoria.Items = make([]models.ItemCarrito, 0)
	carritoEnMemoria.Total = 0.0

	sendJSONResponse(w, map[string]interface{}{"message": "Compra realizada con éxito", "id": orderID}, http.StatusOK)
}
