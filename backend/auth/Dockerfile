# Versión de golang a utilizar
FROM golang:1.22.4-alpine

# Definición del directorio de trabajo
WORKDIR /app

# Instalación de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia los archivos en el directorio de trabajo
COPY . ./

# Expone el puerto 3000
EXPOSE 4000

# Ejecuta proyecto
CMD ["go", "run", "main.go"]

