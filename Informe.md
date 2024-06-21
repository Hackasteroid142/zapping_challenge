# Informe de challenge

## Interpretación e implementación

El objetivo de la app es poder simular un servicio de streaming para usuarios registrados en la plataforma, y por lo tanto, que sean capaces de iniciar sesión en esta. Para esto se desarrollaron diversos modulos en el proyecto.

### Video-api

API que tiene como objetivo el disponibilizar el archivo m3u8 necesario para el streaming. Es así, que en la implementación se contempla una goroutine que permite la actualización del archivo m3u8 cada 10 segundos. De esta manera, si n usuarios llaman a la api dentro de esos 10 segundos, todos ellos verán el mismo archivo. 

En cada actualización del archivo, se verifica el atributo EXT-X-MEDIA-SEQUENCE para conocer que segment se debe agregar a continuación. Es decir, si se tiene un EXT-X-MEDIA-SEQUENCE 0, se interpreta que el archivo contiene los segment 0, 1 y 2, por lo que el segment 3 es el siguiente. Además, se elimina el primer segment del archivo y se incrementa el valor de EXT-X-MEDIA-SEQUENCE. 

En el caso borde, cuando se este por acabar los segment disponibles, lo que se hace es ir agregando los segment iniciales. Es decir, si al m3u8 se agrega el segment 63 (que es el último disponible), a la siguiente iteración se agregará el segment 0 y así sucesivamente con las siguientes iteraciones. Cabe destacar que el valor de EXT-X-MEDIA-SEQUENCE se reinicia a 0 solo cuando se vuelve al estado inicial de tener los segment 0, 1 y 2. Este enfoque se realiza debido a que se considera un streaming en vivo, por lo cual siempre se tiene que tener disponible el video para futuros usuarios. 


### Auth 

Es una API que tiene como principal objetivo la implementación objetivo el registro e inicio de sesión de los usuarios. 

Para el registro se crea un endpoint del tipo POST /users. Este recibe el nombre, email y contraseña para ser guardado en MongoDB. Cabe destacar que la contraseña es encriptada antes de ser guardada en base de datos.


```
# Body de solicitud POST /users
{
    "name": "Vicente",
    "email": "vicente@gmail.com",
    "contraseña": "1234"
}
```

El inicio de sesión de los usuarios se realiza a través del endpoint POST /logIn. Este captura el email y contraseña del usuario, donde este es buscado en DB para verificar su existencia y en caso de que esto se cumpla, se compara la contraseña de BD con la contraseña capturada. Si todo es correcto, se crea un token JWT que es retornado en la respuesta.

```
# Body de solicitud POST /logIn
{
    "email": "vicente@gmail.com",
    "contraseña": "1234"
}
```

### Streaming_api

Frontend que será presentado al usuario. Este contempla las siguientes vista.

- **Home**: Página principal que permite al usuario ingresar al inicio de sesión o al registro. Si el usuario ya inicio sesión el bóton de para iniciar sesión se cambia a uno para ingresar directamente al streaming. 

- **LogIn**: Página que presenta el formulario para incio de sesión y redirige al stream, por lo cual se debe ingresar email y contraseña. La página mostrará error en caso de que el usuario no exista. 

- **Register**: Página que permite el registro de un usuario y redirige a la página para iniciar sesión en caso de éxito. Muestra un mensaje de error en caso de usuario existente en BD. 

- **Live**: Página que muestra el streaming para usuarios que ya han iniciado sesión. 


## Organización

> Los tiempos reflejados a continuación no indican necesariamente que fue realizado de corrido, si no que se omiten tiempos donde se realizaron actividades fuera del challenge. 

- Primer día
    - **Objetivos cumplidos**: Investigación de conceptos y configuración de entorno local para desarrollo.
    - **Tiempo estimado**: 3 Hrs. 

- Segundo día 
    - **Objetivos cumplidos**:  
        - Creación de proyecto inicial de go para video-api
        - Creación de proyecto inicial de front
        - Desarrollo para envio de archivo m3u8 desde video-api
        - Implementación de video.js para la reproducción de archivos recibidos desde video-api
        - Conexión entre front y video-api para la reproducción de video.
        - Investigación de tecnologías para el manejo de archivos m3u8 en go, creación de api y reproducción de video en front.
    - **Tiempo estimado**: 6 hrs.

- Tercer día
    - **Objetivos cumplidos**:
        - Desarrollo para casos borde en envio de archivo m3u8 por parte de video-api.
        - Creación de endpoint para la creación de usuarios y guardado en MongoDB.
        - Creación de endpoint para la creación de token JWT a través del email y contraseña de usuario. 
        - Conexión de endpoints para registro e inicio de sesión de un usuario. 
        - Desarrollo de vista home, inicio de sesión, registro y streaming en front. 
        - Implementación de flujo en plataforma para usuario. 
    - **Tiempo estimado**: 8 hrs.

- Cuarto día
    - **Objetivos cumplidos**:
        - Implementación de rutas protegidas en front. 
        - Flujo para usuario logeado y no logeado. 
        - Arreglo de detalles funcionales y visuales en front. 
        - Implementación de funcionalidad extra.
        - Manejo de errores en inicio de sesión y registro. 
        - Dockerización de aplicación. 
    - **Tiempo estimado**: 5 hrs. 