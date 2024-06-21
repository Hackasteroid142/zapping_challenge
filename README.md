# zapping_challenge

## Aplicación

Este proyecto contempla la implementación de un servicio de streaming. Este contiene formularios para el registro e inicio de sesión de usuarios, así como también una vista que permite ver el stream. 

## Tecnologías 

Para el desarrollo las tecnologías utilizadas y su versión son las siguientes.

- Mongodb: 7.0.12
- Go: 1.22.4
- VueJS: 5.0.8
- nodeJS: 18.20.2
- Docker: 26.1.4

## Organización

El proyecto contempla tres modulos principales y se organizan de la siguiente manera:

- `backend/auth/`: Desarrollado en Go, permite la autenticación y registro de los usuarios en la plataforma.
- `backend/video-api/`: Desarrollado en Go, permite el consumo de los distintos archivos para realizar el streaming de video.
- `streaming_frontend/`: Desarrallado en VueJs, es la página que será mostrada en el navegador. 

## Ejecución

El proyecto contempla un archivo docker-compose.yml por lo que el ejecutar este archivo a través del comando 

```
docker compose up
```

Este levanta imagenes para los siguientes servicios:

- `auth`: Ejecutado en el puerto 4000
- `video-api`: Ejecutado en el puerto 3000
- `streaming_frontend`: Ejecutado en el puerto 8080.
- `mongodb`: Ejecutado en el puerto 27017

Si se quiere detener y eliminar los servicios se debe ejecutar el siguiente comando

```
docker compose down
```
