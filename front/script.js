document.getElementById('consulta-evento-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const eventoID = document.getElementById('evento-id').value;
    const url = `http://localhost:8080/evento/${eventoID}`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('respuesta-evento').textContent = JSON.stringify(data, null, 2);
        })
        .catch(error => {
            document.getElementById('respuesta-evento').textContent = `Error al consultar el evento: ${error.message}`;
        });
});

document.getElementById('consulta-todos-eventos-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const url = `http://localhost:8080/eventos`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('respuesta-todos-eventos').textContent = JSON.stringify(data, null, 2);
        })
        .catch(error => {
            document.getElementById('respuesta-todos-eventos').textContent = `Error al consultar todos los eventos: ${error.message}`;
        });
});

document.getElementById('login-form').addEventListener('submit', function(event) {

    event.preventDefault(); 
    
    const url = `http://localhost:8080/login`;

    const adminname = document.getElementById('adminname').value;
    const password = document.getElementById('password').value;

    const loginData = {
        adminname: adminname,
        password: password
    };

   
    fetch(url, {
        method: 'POST', 
        headers: {
            'Content-Type': 'application/json' 
        },
        body: JSON.stringify(loginData) 
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('respuesta-token').textContent = JSON.stringify(data, null, 2);
    })
    .catch(error => {
        document.getElementById('respuesta-token').textContent = `Error al iniciar sesión: ${error.message}`;
    });
});

document.getElementById('crear-evento-form').addEventListener('submit', function(event) {
    event.preventDefault(); 
    
    const url = `http://localhost:8080/addEvent`;
    
    const eventData = document.getElementById('addEvent').value;
    const token = document.getElementById('token').value;

    let eventJson;
    try {
        eventJson = JSON.parse(eventData);
    } catch (error) {
        document.getElementById('respuesta-crear-evento').textContent = `Error: JSON inválido. Por favor, verifica el formato del evento.`;
        return;
    }

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(eventJson) 
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('respuesta-crear-evento').textContent = JSON.stringify(data, null, 2);
    })
    .catch(error => {
        document.getElementById('respuesta-crear-evento').textContent = `Error al crear el evento: ${error.message}`;
    });
});


document.getElementById('crear-usuario-form').addEventListener('submit', function(event) {
    event.preventDefault(); 

    const url = `http://localhost:8080/register`; 
    const email = document.getElementById('email').value; 

   
    const usuarioData = {
        email: email
    };

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(usuarioData) 
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('respuesta-crear-usuario').textContent = JSON.stringify(data, null, 2);
    })
    .catch(error => {
        document.getElementById('respuesta-crear-usuario').textContent = `Error al crear el usuario: ${error.message}`;
    });
});


document.getElementById('suscribir-evento-form').addEventListener('submit', function(event) {
    event.preventDefault(); 


    const url = `http://localhost:8080/eventos/suscribir`;

    const eventoID = document.getElementById('EventoID').value;
    const email = document.getElementById('email-suscribir').value;

    const suscripcionData = {
        eventoID: parseInt(eventoID, 10), 
        email: email
    };

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(suscripcionData) 
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('respuesta-suscribir-evento').textContent = JSON.stringify(data, null, 2);
    })
    .catch(error => {
        document.getElementById('respuesta-suscribir-evento').textContent = `Error al suscribir al evento: ${error.message}`;
    });
});



document.getElementById('update-evento-form').addEventListener('submit', function(event) {
    event.preventDefault(); 


    const eventoID = document.getElementById('evento-id-update').value;
    const title = document.getElementById('title').value;
    const description = document.getElementById('description').value;
    const descriptionLarge = document.getElementById('description-large').value;
    const date = document.getElementById('date').value.trim(); 
    const organizer = document.getElementById('organizer').value;
    const place = document.getElementById('place').value;
    const state = document.getElementById('state').value;
    const token = document.getElementById('token-update').value;


    let eventData;

    if (date !== "") {
        eventData = {
            id: parseInt(eventoID, 10),
            title: title,
            description_short: description,
            description_large: descriptionLarge,
            date: date, 
            organizer: organizer,
            place: place,
            state: state
        };
    } else {
        eventData = {
            id: parseInt(eventoID, 10),
            title: title,
            description_short: description,
            description_large: descriptionLarge,
            organizer: organizer,
            place: place,
            state: state
        };
    }

    fetch(`http://localhost:8080/updateEvent`, {
        method: 'PUT', 
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}` 
        },
        body: JSON.stringify(eventData) 
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('respuesta-actualizar-evento').textContent = JSON.stringify(data, null, 2);
    })
    .catch(error => {
        document.getElementById('respuesta-actualizar-evento').textContent = `Error al actualizar el evento: ${error.message}`;
    });
});