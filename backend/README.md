




# Arquitectura hexagonal con diseño orientado a paquetes vertical

/backend
│
├── /cmd
│   └── /api
│       ├── /routes
│       │   ├── user.go
│       │   └── .go
│       └── main.go
│
├── /internal
│   ├── /config
│   │   └── /database
│   │       ├── connection.go    // conexión a la DB
│   │       └── migration.go     // migraciones de tablas
│   │
│   ├── /module
│   │   ├── /user
│   │        ├── /handler
│   │        │   └── handler.go       // controla lo que recibe HTTP
│   │        ├── /service
│   │        │   └── service.go       // lógica de negocio
│   │        ├── /repository
│   │        │   └── repository.go    // consultas SQL a la DB
│   │        └── /model
│   │            └── model.go         // definición de structs (modelo de datos)
│   │
│   └── vendor
│
└── go.mod



# Desarrollo del trabajo final del taller del go 

1. go mod init backend

2. instalar:
    go get -u github.com/gin-gonic/gin

    go get github.com/google/uuid

    go get github.com/go-resty/resty/v2

    go get -u go.uber.org/zap

    go get github.com/go-sql-driver/mysql

3. Vendorizar
    go mod vendor


