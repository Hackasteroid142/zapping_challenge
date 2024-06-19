# Versión de node a utilizar
FROM node:18.20.2-alpine

# Instalación de http server para servir el contenido
RUN npm install -g http-server

# Definición del directorio de trabajo
WORKDIR /app

# Copia de package.json, package-lock.json y yarn.lock si están disponibles
COPY package*.json ./
COPY yarn*.lock ./

# Instalación de dependencias
RUN yarn install

# Copia los archivos y carpetas en el directorio de trabajo (/app)
COPY . .

# Construye la aplicación para producción
RUN yarn build

# Expone el puerto 8080
EXPOSE 8080

# Ejecuta el proyecto
CMD ["http-server", "dist"]