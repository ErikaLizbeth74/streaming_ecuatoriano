Nombre: Erika Vinocunga
Cine Ecuatoriano 🇪🇨
Una plataforma de streaming de películas y documentales ecuatorianos, construida con Go (Golang) en el backend y un frontend dinámico con HTML, CSS y JavaScript. Este proyecto demuestra una arquitectura web completa con persistencia de datos usando SQLite y comunicación JSON.
1. Descripción del Proyecto
"Cine Ecuatoriano" es una plataforma web que permite a los usuarios explorar, buscar, y acceder a un catálogo de películas y documentales producidos en Ecuador. El proyecto destaca por su backend eficiente en Go, la persistencia de datos con SQLite y una interfaz de usuario interactiva que se comunica completamente a través de JSON. Incluye funcionalidades de autenticación, gestión de carrito y un sistema simulado de alquiler de contenido premium.
2. Características
•	Catálogo Dinámico: Explora una colección de películas ecuatorianas con detalles completos.
•	Búsqueda y Filtrado: Filtra películas por título, género y año.
•	Detalles de Película: Páginas dedicadas con sinopsis, director, actores, duración y calificación.
•	Autenticación de Usuarios: Registro e inicio de sesión con gestión de sesiones.
•	Contenido Premium/Gratuito: Distinción clara entre contenido gratuito y de pago.
•	Sistema de Alquiler: Simulación de alquiler de películas premium con un carrito de compras.
•	Acceso Persistente: Las películas premium "compradas" se asocian al perfil del usuario y persisten en la base de datos.
•	Reproductor de Video: Página para la visualización de tráilers o contenido.
•	Modales Personalizados: Mensajes de usuario amigables y consistentes.
•	Serialización JSON: Toda la comunicación frontend-backend se realiza mediante JSON.
•	Persistencia de Datos: Usuarios, películas, compras y favoritos se almacenan en SQLite.
3. Tecnologías Utilizadas
•	Backend:
   Go (Golang): Versión 1.22.x 
   Gorilla Mux: Librería para enrutamiento de URLs (github.com/gorilla/mux).
   Go-SQLite3: Driver de SQLite para Go (github.com/mattn/go-sqlite3).
•	Frontend:
   HTML5: Estructura de la interfaz de usuario.
   CSS3: Estilos y diseño.
   Bulma CSS Framework: Framework CSS para un diseño responsive y moderno.
   JavaScript (Vanilla JS): Interactividad, llamadas a la API y renderizado dinámico.
•	Base de Datos
    SQLite: Base de datos ligera basada en archivos (streaming.db).
Librerías Go Extraídas 
Para el backend Go, las siguientes librerías se obtienen y gestionan automáticamente por Go Modules:
•	github.com/gorilla/mux
•	github.com/mattn/go-sqlite3

5. Estructura del Proyecto
streaming_ecuatorianos/
├── main.go               # Punto de entrada de la aplicación, inicializa DB y rutas.
├── go.mod                # Módulos Go y dependencias.
├── go.sum                # Checksums de dependencias.
├── models/
│   └── models.go         # Definiciones de estructuras de datos (Content, User, etc.) con etiquetas JSON.
├── handlers/
│   └── handlers.go       # Lógica de negocio, implementación de los servicios web (APIs).
├── templates/            # Archivos de plantillas HTML.
│   ├── base.html         # Plantilla base con estructura común (header, footer, navbar).
│   ├── home.html         # Página de inicio.
│   ├── catalog.html      # Página del catálogo de películas.
│   ├── movie_detail.html # Página de detalles de una película específica.
│   ├── player.html       # Página del reproductor de video.
│   ├── login.html        # Página de inicio de sesión.
│   └── register.html     # Página de registro de usuario.
└── static/               # Archivos estáticos (CSS, imágenes, JS del cliente).
    ├── css/
    │   └── style.css     # Estilos CSS personalizados.
    ├── img/
    │   ├── movies/       # Imágenes de miniaturas de películas.
    │   ├── portada_de_inicio/ # Imagen de portada de la página de inicio.
    │   └── placeholder.jpg # Imagen de reemplazo para miniaturas no encontradas.
    └── js/               # Archivos JavaScript del cliente (vacío, JS está incrustado en HTML).



