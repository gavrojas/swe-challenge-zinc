# Proyecto de Indexación y Visualización de Correos Electrónicos

## Descripción
Este proyecto tiene como objetivo indexar y visualizar una base de datos de correos electrónicos con el dataset de Enron Corp utilizando ZincSearch como motor de búsqueda. La aplicación permite a los usuarios buscar y visualizar correos electrónicos a través de una interfaz web simple.

## ¿Qué incluye el proyecto?

### Parte 1: Indexar Base de Datos de Correo Electrónico
Se indexaron los contenidos del dataset [Enron Mail Dataset](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) (423MB) en ZincSearch usando un programa en GO. 

### Parte 2: Profiling
Se realizó profilling al indexador siguiendo la documentación de GO [Profiling](https://go.dev/doc/diagnostics#profiling) y se generaron gráficos para cpu, memoria, y go routines. 

### Parte 3: Visualizador
Se creó una interfaz simple con Vue, y tailwind para visualizar y buscar contenidos. 

## Tecnologías Utilizadas
- **Lenguaje Backend:** Go
- **Base de Datos:** ZincSearch
- **API Router:** chi
- **Interfaz:** Vue 3, Pinia, Vue Router, Vuetify,
- **CSS:** Tailwind
- **Despliegue** Docker

## Estructura del Proyecto

El proyecto está dividido en cuatro partes: el frontend, el backend, el indexador y el entorno de Docker.

### Frontend (mail_app)
- **Interfaz de Usuario:** Desarrollada con Vue 3, permite a los usuarios buscar correos electrónicos de manera intuitiva.
- **Autenticación**: Sistema básico de registro y login con usuario y contraseña. 
- **Gestión de Sesión**: Manejo de tokens JWT
- **Búsqueda de correos**: Filtros por campos to, from, body y subject
- **Visualización de correos**: Resultados de los primeros correos que coinciden con los campos de búsqueda y con opción de ver más, opción de visualizar contenido del correo. 

<details>
<summary>Estructura de Directorios Frontend</summary>

#### mail_app

```
src/
├── components/
│   ├── EmailList.vue
│   ├── EmailView.vue
│   ├── FoldersSidebar.vue
│   ├── LoginForm.vue
│   ├── RegisterForm.vue
│   ├── SearchCard.vue
│   └── Snackbar.vue
├── router/
│   └── index.ts
├── services/
│   ├── auth.ts
│   ├── email.ts
│   └── utils.ts
├── stores/
│   ├── auth.ts
│   └── emails.ts
├── types/
│   └── index.ts
├── views/
│   ├── FolderView.vue
│   ├── HomeView.vue
│   ├── LoginView.vue
│   └── RegisterView.vue
├── assets/
│   ├── base.css
│   ├── logo.svg
│   └── login.svg
└── main.ts
```
</details>

### Backend (mail_api)
- **API:** Implementada en Go, maneja las solicitudes de búsqueda de emails y autenticación de usuarios.
- **Autenticación**:
  - Registro y login de usuarios
  - Generación y validación de JWT
- **Manejo de Errores:** Respuestas adecuadas para errores en la carga y búsqueda de correos.
- **Búsqueda de emails por campos específicos**

<details>
<summary>Estructura de Directorios Backend</summary>

#### mail_api

```
mail_api/
├── auth/
│   ├── handlers.go
│   └── router.go
├── emails/
│   ├── handler.go
│   └── router.go
├── models/
│   ├── emails.go
│   ├── users.go
│   └── zinc.go
├── shared/
│   ├── jwt.go
│   ├── middlewares.go
│   └── sessions.go
├── users/
│   ├── handler.go
│   └── router.go
├── zinc/
│   └── zinc.go
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
```
</details>

### Indexación (zinc)
- **Indexer:** Programa que toma la base de datos de correos y la indexa en ZincSearch para facilitar la búsqueda.

<details>
<summary>Estructura de Directorios Indexador</summary>

#### zinc

```
zinc/
├── config/
│   └── config.go
├── handlers/
│   └── indexer.go
├── models/
│   └── email.go
├── services/
│   └── zinc_service.go
├── utils/
│   └── helpers.go
├── go.mod
└── main.go
```
</details>

### Docker (mail_tools)
Configuración para el despliegue de la aplicación usando Docker Compose:
- ZincSearch para base de datos
- Servidor Go para la API
- Configuración de red y volúmenes

<details>
<summary>Estructura de Directorios Docker</summary>
mail_tools/

```
mail_tools/
├── docker-compose.yml
└── README.md
```
</details>

## Despliegue de manera local
En el directorio zinc:
- Ejecutar: ./get_data.sh 
  Esto traerá toda la data y la dejará descomprimida para trabajar con ella.

Teniendo Docker instalado y corriendo: 
- Ejecutar: docker compose build 
- Ejecutar: docker compose up zinc -d 
- Ejecutar: docker compose up api -d

Esto levantará:
- el servicio de zincSearch en http://localhost:4080/
- el servidor de Go en: http://localhost:8080/

En el directorio zinc ejecutar: 
- ./indexer -cpuprofile=cpu.prof -memprofile=mem.prof -goroutineprofile=goroutine.prof data/enron_mail_20110402 > logsIndexer.txt 2>&1 &
  Esto ejecutará el indexador en segundo plano,guardará los logs en un archivo llamado logsIndexer.txt y además guardará el profilling para cpu, memoria y goroutines 

Después de tener los registros en ZincSearch, dirigirse al directorio mail_app y ejecutar: 
- ./enron_mails 

Esto levantará el servicio de visualización (frontend) en http://localhost:5173/

-----------------

*Para mas información puedes escribirme a mi correo grojas9807@gmail.com o contactarme por [LinkedIn](https://www.linkedin.com/in/gavrojas-dev/)*