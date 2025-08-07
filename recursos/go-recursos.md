# Empieza con Go: Librerías y Recursos para Todos los Niveles

Recopilación de `librerías`, `herramientas`, `frameworks` y recursos para desarrollar en Go. Desde frameworks web y herramientas CLI hasta utilidades para testing, bases de datos, APIs, blockchain, y más.

## 🌐 Frameworks Web y HTTP

- **[gin](https://gin-gonic.com/):** El framework web más popular en Go
- **[fiber](https://github.com/gofiber/fiber):** Framework web rápido inspirado en Express
- **[echo](https://github.com/labstack/echo):** Framework web minimalista y de alto rendimiento
- **[beego](https://github.com/beego/beego):** Framework web de código abierto y alto rendimiento para el lenguaje de programación Go
- **[go-chi/chi](https://github.com/go-chi/chi):** Enrutador ligero, idiomático y componible para crear servicios HTTP en Go
- **[gofr](https://github.com/gofr-dev/gofr):** Framework web moderno para desarrollo de microservicios
- **[ent](https://github.com/ent/ent):** Framework que facilita la creación y el mantenimiento de aplicaciones con grandes modelos de datos
- **[resty](https://github.com/go-resty/resty):** Cliente HTTP elegante
- **[imroc/req](https://github.com/imroc/req):** Cliente HTTP simple y potente
- **[rs/cors](https://github.com/rs/cors):** Middleware para CORS
- **[jub0bs/cors](https://github.com/jub0bs/cors):** Biblioteca de middleware CORS para Go
- **[gorilla/sessions](https://github.com/gorilla/sessions):** Proporciona sesiones de cookies y del sistema de archivos, así como la infraestructura necesaria para backends de sesión personalizados

## 🗄️ Bases de Datos, Migraciones y ORM

- **[jackc/pgx](https://github.com/jackc/pgx):** Driver nativo de PostgreSQL
- **[jmoiron/sqlx](https://github.com/jmoiron/sqlx):** Extensiones para database/sql
- **[go-redis](https://github.com/redis/go-redis):** Cliente para Redis
- **[gorm](https://github.com/go-gorm/gorm):** Biblioteca ORM para Go, diseñada para facilitar el trabajo de los desarrolladores
- **[uptrace/bun](https://github.com/uptrace/bun):** ORM Go basado en SQL
- **[golang-migrate/migrate](https://github.com/golang-migrate/migrate):** Migraciones de base de datos
- **[goose](https://github.com/pressly/goose):** Migraciones de base de datos
- **[jackc/tern](https://github.com/jackc/tern):** Migraciones de PostgreSQL
- **[sqlc](https://github.com/sqlc-dev/sqlc):** Generador de código SQL type-safe
- **[ariga/atlas](https://github.com/ariga/atlas):** Esquemas de base de datos como código

## 🔐 Autenticación y Autorización

- **[golang/oauth2](https://github.com/golang/oauth2):** Cliente OAuth2
- **[auth0/go-jwt-middleware](https://github.com/auth0/go-jwt-middleware):** Un middleware para el lenguaje de programación Go que comprueba los JWT en las solicitudes HTTP
- **[casbin](https://github.com/casbin/casbin):** Control de acceso y autorización
- **[markbates/goth](https://github.com/markbates/goth):** Autenticación multi-proveedor OAuth
- **[go-paseto](https://github.com/aidantwoods/go-paseto):** Implementación de tokens de seguridad independientes de la plataforma en Golang
- **[unrolled/secure](https://github.com/unrolled/secure):** Middleware HTTP para Go que facilita algunas mejoras rápidas en materia de seguridad

## ⚙️ CLI y Herramientas de Desarrollo

- **[cobra](https://github.com/spf13/cobra):** Framework para aplicaciones CLI
- **[viper](https://github.com/spf13/viper):** Manejo de configuración
- **[air](https://github.com/air-verse/air):** Live reload para desarrollo Go
- **[bubbletea](https://github.com/charmbracelet/bubbletea):** Framework para TUI (Terminal UI)
- **[goreleaser](https://github.com/goreleaser/goreleaser):** Automatización de releases
- **[gobuster](https://github.com/OJ/gobuster):** Herramienta para eliminar directorios/archivos, DNS y VHost
- **[tinygo](https://github.com/tinygo-org/tinygo):** Compilador Go para microcontroladores

## 🖥️ Aplicaciones Desktop y UI

- **[wails](https://github.com/wailsapp/wails):** Framework para apps desktop con web frontend
- **[fyne-io/fyne](https://github.com/fyne-io/fyne):** Kit de herramientas GUI multiplataforma en Go inspirado en `Material Design`
- **[charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss):** Estilos para terminales
- **[charmbracelet/huh](https://github.com/charmbracelet/huh):** Crear formularios y mensajes de terminal

## 🌍 Blockchain y Criptomonedas

- **[go-ethereum](https://github.com/ethereum/go-ethereum):** Implementación de Ethereum en Go

## 🐳 DevOps y Contenedores

- **[podman](https://github.com/containers/podman):** Manejo de contenedores
- **[slimtoolkit/slim](https://github.com/slimtoolkit/slim):** Optimización de imágenes de contenedores Docker

## 📄 Procesamiento de Documentos

- **[pdfcpu](https://github.com/pdfcpu/pdfcpu):** Procesador de PDFs
- **[johnfercher/maroto](https://github.com/johnfercher/maroto):** Una forma rápida de crear archivos PDF. Maroto se inspira en `Bootstrap` y utiliza `gofpdf`
- **[gotenberg](https://github.com/gotenberg/gotenberg):** Servicio de conversión de documentos

## 🕷️ Web Scraping y Crawling

- **[colly](https://github.com/gocolly/colly):** Framework de web scraping (rastreo y extracción de datos)
- **[kkdai/youtube](https://github.com/kkdai/youtube):** Descargar vídeos de YouTube en Go
- **[lux](https://github.com/iawia002/lux):** Biblioteca de descarga de vídeos rápida y sencilla y herramienta CLI escrita en Go

## 🧪 Testing y Validación

- **[testify](https://github.com/stretchr/testify):** Toolkit de testing
- **[validator](https://github.com/go-playground/validator):** Validación de estructuras y campos, incluyendo campos cruzados, estructuras cruzadas, maps, slices y arrays

## 📚 Utilidades y Estructuras de Datos

- **[emirpasic/gods](https://github.com/emirpasic/gods):** Estructuras de datos y algoritmos
- **[samber/Io](https://github.com/samber/lo):** Una biblioteca Go de estilo *Lodash* basada en Go 1.18
- **[ristretto](https://github.com/hypermodeinc/ristretto):** Caché de alto rendimiento y limitado en memoria
- **[gofrs/uuid](https://github.com/gofrs/uuid):** Generación de UUIDs
- **[fatih/color](https://github.com/fatih/color):** Colores para terminal
- **[joho/godotenv](https://github.com/joho/godotenv):** Carga de variables de entorno
- **[caarlos0/env](https://github.com/caarlos0/env):** Biblioteca sencilla y sin dependencias para analizar variables de entorno en estructuras
- **[idsulik/go-collections](https://github.com/idsulik/go-collections):** Proporciona implementaciones de estructuras de datos comunes
- **[go-pkgz/routegroup](https://github.com/go-pkgz/routegroup):** Biblioteca ligera para la agrupación de rutas y la integración de middleware con el estándar `http.ServeMux`
- **[sonirico/vago](https://github.com/sonirico/vago):** Kit de herramientas Go con utilidades genéricas para trabajar con slices, maps y primitivas de programación funcional, streams, bases de datos y mucho más

## 📡 APIs y Integraciones

- **[go-github](https://github.com/google/go-github):** Cliente de GitHub API
- **[discordgo](https://github.com/bwmarrin/discordgo):** Cliente de Discord
- **[stripe-go](https://github.com/stripe/stripe-go):** SDK de Stripe
- **[go-genai](https://github.com/googleapis/go-genai):** Google Gen AI Go SDK proporciona una interfaz para que los desarrolladores puedan integrar los modelos generativos de Google en sus aplicaciones Go
- **[anthropics/anthropic-sdk-go](https://github.com/anthropics/anthropic-sdk-go):** Acceso a las API del modelo de lenguaje de Anthropic, que prioriza la seguridad, a través de Go
- **[mercadopago/sdk-go](https://github.com/mercadopago/sdk-go):** SDK de MercadoPago
- **[modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk):** El SDK oficial de Go para servidores y clientes del Protocolo de contexto de modelo

## 📋 Documentación y Code Generation

- **[swaggo/swag](https://github.com/swaggo/swag):** Genera automáticamente documentación de API RESTful con Swagger 2.0
- **[bufbuild/buf](https://github.com/bufbuild/buf):** La mejor manera de trabajar con Protocol Buffers

## 📧 Comunicación y Mensajería

- **[go-mail](https://github.com/wneessen/go-mail):** Biblioteca fácil de usar para enviar correos electrónicos con Go
- **[watermill](https://github.com/ThreeDotsLabs/watermill):** Aplicaciones basadas en eventos de forma sencilla en Go
- **[twitchtv/twirp](https://github.com/twitchtv/twirp):** Framework RPC con definiciones de servicio protobuf

## 📖 Recursos de Aprendizaje

- **[100-go-mistakes](https://github.com/teivah/100-go-mistakes):** Guía de errores comunes en Go
- **[golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers):** Go para desarrolladores `Node.js`
- **[gofiber/recipes](https://github.com/gofiber/recipes):** Ejemplos de `Fiber`
- **[TheAlgorithms/Go](https://github.com/TheAlgorithms/Go):** Algoritmos y estructuras de datos implementados en Go para principiantes

## 🔧 Herramientas de Análisis

- **[timakin/bodyclose](https://github.com/timakin/bodyclose):** Comprueba si el cuerpo de la respuesta HTTP está cerrado y si no se bloquea la reutilización de la conexión TCP
- **[Oudwins/zog](https://github.com/Oudwins/zog):** Validación de esquemas simple inspirada en `Zod`
