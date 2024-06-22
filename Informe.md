# Informe de challenge

## Interpretación e implementación

El objetivo de la app es poder simular un servicio de streaming para usuarios registrados en la plataforma, y, por lo tanto, que sean capaces de iniciar sesión en esta. Para esto se desarrollaron diversos módulos en el proyecto, los cuales serán explicados a continuación.

### Video-api

API que tiene como objetivo el disponibilizar el archivo m3u8 necesario para el streaming. Es así, que en la implementación se contempla una goroutine que permite la actualización del archivo m3u8 cada 10 segundos. De esta manera, si n usuarios llaman a la api dentro de esos 10 segundos, todos ellos verán el mismo archivo. 

En cada actualización del archivo, se verifica el atributo `EXT-X-MEDIA-SEQUENCE` para conocer qué segment se debe agregar a continuación. Es decir, si se tiene un `EXT-X-MEDIA-SEQUENCE` 0, se interpreta que el archivo contiene los segment 0, 1 y 2, por lo que el segment 3 es el siguiente. Además, se elimina el primer segment del archivo y se incrementa el valor de `EXT-X-MEDIA-SEQUENCE`. 

Para el primer llamado se define una variable global con un valor booleano true. De esta manera se puede identificar cuándo se debe enviar el archivo inicial m3u8 que contiene el EXT-X-MEDIA-SEQUENCE 0 y los segmentos 0, 1 y 2. Luego de esto, la variable global se cambia a false para que no se dé el caso inicial. 

En el caso borde, cuando se estén por acabar los segment disponibles, lo que se hace es ir agregando los segment iniciales. Es decir, si al m3u8 se agrega el segment 63 (que es el último disponible), a la siguiente iteración se agregará el segment 0 y así sucesivamente con las siguientes iteraciones. Cabe destacar que el valor de `EXT-X-MEDIA-SEQUENCE` se reinicia a 0 solo cuando se vuelve al estado inicial de tener los segment 0, 1 y 2. Este enfoque se realiza debido a que se considera un streaming en vivo, por lo cual siempre se tiene que tener disponible el video para futuros usuarios. 

Aunque en el caso borde se vuelve al caso inicial, no es necesario inicializar la variable global en `true`. Esto se debe a que el `EXT-X-MEDIA-SEQUENCE` se utiliza como indicador para determinar qué segmento agregar. Por lo tanto, a este valor se le debe sumar un valor distinto en cada caso. En detalle, sería lo siguiente:

- Si es el primer llamado, el EXT-X-MEDIA-SEQUENCE tiene valor 0, por lo que los segmentos serían aquellos con los valores  `EXT-X-MEDIA-SEQUENCE`,  `EXT-X-MEDIA-SEQUENCE`+1 y `EXT-X-MEDIA-SEQUENCE`+2.
- En el caso límite, dado que el máximo valor permitido es 63, desde el número 61 se deben volver a agregar los segmentos desde el inicio. Por lo tanto, el segmento a agregar se calcula como `EXT-X-MEDIA-SEQUENCE` - 61. Por ejemplo, si `EXT-X-MEDIA-SEQUENCE` es 62, se agregaría el `segment1.ts`.
- Todos los demás casos considera el agregar el segmento con valor `EXT-X-MEDIA-SEQUENCE`+3, ya que por ejemplo si se tiene un `EXT-X-MEDIA-SEQUENCE`EXT-X-MEDIA-SEQUENCE igual a 3 significa que se debe agregar el segment6.ts.  

### Auth 

Es una API que tiene como principal objetivo la implementación del registro e inicio de sesión de los usuarios. 

Para el registro se crea un endpoint del tipo POST /users. Este recibe el nombre, email y contraseña para ser guardado en MongoDB. Cabe destacar que la contraseña es encriptada antes de ser guardada en base de datos para asegurar las credenciales del usuario.

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
