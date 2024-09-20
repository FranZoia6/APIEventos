# APIEventos


## Ejecución local

1. Clonar el proyecto:

```bash
 git clone https://github.com/FranZoia6/APIEventos.git
```

2. Redirígete a la carpeta

```bash
cd APIEventos
```

3. Iniciar el API y  front:

```bash
docker composer up --build
```

4. El frontend estará operativo en http://localhost:8000.

## Endpoints

### Consultar Evento por ID

- Consulta: GET http://localhost:8080/evento/{:id}

- Nos permite consultar un evento por su ID, en esta prueba hay 3 eventos cargados, el evento con ID 1 y 3 los vamos a poder recuperar sin problema, el evento con ID 2 no lo vamos a poder recuperar porque el estado de este es borrador. 


### Consultar Todos los Eventos

- Consulta: GET http://localhost:8080/eventos


- Nos permite recuperar todos los eventos publicados. 

### Login admin

- Consulta: POST http://localhost:8080/login

- Nos permite conseguir las credenciales necesarias para utilizar los endpoints a los que se necesita un token.

- Credenciales: admin: admin, password: adminpass.

### Crear Evento

- Consulta: POST http://localhost:8080/addEvent


- Nos permite enviar un JSON para poder crear un evento, para esta endpoints es necesario enviar un token porque solo la pueden resolver los administradores. 

- Ejemplo de JSON : 
```bash
  {
	"title": "Titulo",
	"descriptionShort": "Una descripción corta.",
	"descriptionLarge": "Una descripción un poco mas larga",
	"date": "una fecha con el siguiente formato  2024-09-15T10:12:00Z",
	"organizer": "El nombre de la organización",
	"place": "El lugar",
	"state": "Y el estado"
  }
```


### Crear Usuario

- Consulta: POST http://localhost:8080/register

- Se utiliza para crear un usuario el cual se puede suscribir a un evento.

### Suscribir a un Evento

- Consulta: POST http://localhost:8080/eventos/suscribir

- Se utiliza para suscribir a un evento, se pasa un ID del evento y el email de quien se quiere suscribir, una persona no se puede suscribir dos veces a un evento. 

### Actualizar Evento

- Consulta: PUT http://localhost:8080/updateEvent

- Se utiliza para actualizar un evento, en este caso no es necesario completar todos inputs, el ID es el único requerido y luego los que no se completan quedan igual como estaban antes 



