package services

import (
    // "database/sql"
    "time"
    "fmt"
    "apieventos/models"
)

const layout = "2006-01-02 15:04"

var date1, _ = time.Parse(layout, "2024-09-15 10:12")
var date2, _ = time.Parse(layout, "2024-10-10 15:30")
var date3, _ = time.Parse(layout, "2024-11-05 07:45")

var events = []models.Event{
    {
        ID:               1,
        Title:            "Concierto de Jazz",
        DescriptionShort: "Un evento de jazz en vivo.",
        DescriptionLarge: "Únete a nosotros para una noche de jazz en vivo con músicos locales y nacionales. ¡No te lo pierdas!",
        Date:             date1,
        Organizer:        "Jazz Nights",
        Place:            "Sala de Conciertos ABC",
        State:            "publicada",
    },
    {
        ID:               2,
        Title:            "Feria de Libros",
        DescriptionShort: "Una feria dedicada a la venta de libros.",
        DescriptionLarge: "Explora una amplia variedad de libros de todos los géneros en nuestra feria anual.",
        Date:             date2,
        Organizer:        "Biblioteca Central",
        Place:            "Centro de Convenciones XYZ",
        State:            "borrador",
    },
    {
        ID:               3,
        Title:            "Maratón de la Ciudad",
        DescriptionShort: "Una maratón para todos los corredores.",
        DescriptionLarge: "Participa en el maratón anual de la ciudad.",
        Date:             date3,
        Organizer:        "Club de Corredores Locales",
        Place:            "Parque Central",
        State:            "publicada",
    },
}

func GetEventsPublished() []models.Event {
    var publishedEvents []models.Event
    for _, event := range events {
        if event.State == "publicada" {
            publishedEvents = append(publishedEvents, event)
        }
    }
    return publishedEvents
}


func GetEvents() []models.Event {
    return events
}


func GetEventoByID(id int) *models.Event {
    for _, event := range events {
        if event.ID == id && event.State == "publicada" {
            return &event
        }
    }
    return nil
}


func SetEvent(event models.Event) {
    event.ID = len(events) + 1 
    events = append(events, event)
}

func UpdateEvent(event models.Event) error {
    for i, existingEvent := range events {
        if existingEvent.ID == event.ID {
            if event.Title != "" {
                events[i].Title = event.Title
            }
            if event.DescriptionShort != "" {
                events[i].DescriptionShort = event.DescriptionShort
            }
            if event.DescriptionLarge != "" {
                events[i].DescriptionLarge = event.DescriptionLarge
            }
            if !event.Date.IsZero() {
                events[i].Date = event.Date
            }
            if event.Organizer != "" {
                events[i].Organizer = event.Organizer
            }
            if event.Place != "" {
                events[i].Place = event.Place
            }
            if event.State != "" {
                events[i].State = event.State
            }
            return nil
        }
    }
    return fmt.Errorf("El evento con ID %d no existe", event.ID)
}


var userEvents []models.UserEvent

func SuscribirUsuarioAEvento(eventoID int, usuarioID int) error {

    var event *models.Event
    for i := range events {
        if events[i].ID == eventoID && events[i].State == "publicada" {
            event = &events[i]
            break
        }
    }

    if event == nil {
        return fmt.Errorf("no se encontró el evento o no está publicado", eventoID)
    }

    today := time.Now().Truncate(24 * time.Hour)
    if event.Date.Before(today) {
        return fmt.Errorf("la fecha del evento ya pasó, no es posible suscribirse")
    }


    userEvents = append(userEvents, models.UserEvent{EventoID: eventoID, UserID: usuarioID})
    fmt.Printf("Usuario con ID %d suscrito al evento ID %d\n", usuarioID, eventoID)
    return nil
}




// Manejo para postgreSQL

// // Obtener todos los eventos
// func GetEventsPublished() []models.Event {
//     db := GetDB()

//     rows, err := db.Query("SELECT id, title, description_short, description_large, date, organizer, place, state FROM eventos where state = 'publicado'")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer rows.Close()

//     var events []models.Event
//     for rows.Next() {
//         var event models.Event
//         err := rows.Scan(&event.ID, &event.Title, &event.DescriptionShort, &event.DescriptionLarge, &event.Date, &event.Organizer, &event.Place, &event.State)
//         if err != nil {
//             log.Fatal(err)
//         }
//         events = append(events, event)
//     }

//     return events
// }


// func GetEvents() []models.Event {
//     db := GetDB()

//     rows, err := db.Query("SELECT id, title, description_short, description_large, date, organizer, place, state FROM eventos")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer rows.Close()

//     var events []models.Event
//     for rows.Next() {
//         var event models.Event
//         err := rows.Scan(&event.ID, &event.Title, &event.DescriptionShort, &event.DescriptionLarge, &event.Date, &event.Organizer, &event.Place, &event.State)
//         if err != nil {
//             log.Fatal(err)
//         }
//         events = append(events, event)
//     }

//     return events
// }


// func GetEventoByID(id int) *models.Event {
//     db := GetDB()

//     var event models.Event
//     err := db.QueryRow("SELECT id, title, description_short, description_large, date, organizer, place, state FROM eventos WHERE id=$1 AND state='publicado'", id).Scan(
//         &event.ID, &event.Title, &event.DescriptionShort, &event.DescriptionLarge, &event.Date, &event.Organizer, &event.Place, &event.State)
    
//     if err == sql.ErrNoRows {
//         return nil
//     } else if err != nil {
//         log.Fatal(err)
//     }

//     return &event
// }

// func SetEvent(event models.Event) {
//     db := GetDB()

//     _, err := db.Exec("INSERT INTO eventos (title, description_short, description_large, date, organizer, place, state) VALUES ($1, $2, $3, $4, $5, $6, $7)",
//         event.Title, event.DescriptionShort, event.DescriptionLarge, event.Date, event.Organizer, event.Place, event.State)

//     if err != nil {
//         log.Fatal(err)
//     }
// }

// func UpdateEvent(event models.Event) error {
//     db := GetDB()

//     // Recuperar el evento actual
//     var existingEvent models.Event
//     err := db.QueryRow("SELECT id, title, description_short, description_large, date, organizer, place, state FROM eventos WHERE id=$1", event.ID).Scan(
//         &existingEvent.ID, &existingEvent.Title, &existingEvent.DescriptionShort, &existingEvent.DescriptionLarge, &existingEvent.Date, &existingEvent.Organizer, &existingEvent.Place, &existingEvent.State)

//     if err == sql.ErrNoRows {
//         return fmt.Errorf("El evento con ID %d no existe", event.ID)
//     } else if err != nil {
//         log.Fatal(err)
//         return err
//     }

//     // Actualizar los campos solo si se enviaron en el JSON
//     if event.Title != "" {
//         existingEvent.Title = event.Title
//     }
//     if event.DescriptionShort != "" {
//         existingEvent.DescriptionShort = event.DescriptionShort
//     }
//     if event.DescriptionLarge != "" {
//         existingEvent.DescriptionLarge = event.DescriptionLarge
//     }
//     if !event.Date.IsZero() { 
//         existingEvent.Date = event.Date
//     }
//     if event.Organizer != "" {
//         existingEvent.Organizer = event.Organizer
//     }
//     if event.Place != "" {
//         existingEvent.Place = event.Place
//     }
//     if event.State != "" {
//         existingEvent.State = event.State
//     }

//     // Ejecutar la consulta de actualización
//     _, err = db.Exec(`UPDATE eventos 
//                       SET title=$1, description_short=$2, description_large=$3, date=$4, organizer=$5, place=$6, state=$7 
//                       WHERE id=$8`,
//         existingEvent.Title, existingEvent.DescriptionShort, existingEvent.DescriptionLarge, existingEvent.Date, existingEvent.Organizer, existingEvent.Place, existingEvent.State, event.ID)

//     if err != nil {
//         log.Println("Error al actualizar el evento:", err)
//         return err
//     }

//     return nil
// }


// func SuscribirUsuarioAEvento(eventoID int, usuarioID int) error {
//     db := GetDB()

//     var event models.Event
//     err := db.QueryRow("SELECT id, title, description_short, description_large, date, organizer, place, state FROM eventos WHERE id=$1 AND state='publicado'", eventoID).Scan(
//         &event.ID, &event.Title, &event.DescriptionShort, &event.DescriptionLarge, &event.Date, &event.Organizer, &event.Place, &event.State)

//     if err == sql.ErrNoRows {
//         return fmt.Errorf("no se encontró el evento o no está publicado")
//     } else if err != nil {
//         return fmt.Errorf("error al obtener el evento: %w", err)
//     }


//     today := time.Now().Truncate(24 * time.Hour)
//     if event.Date.Before(today) {
//         return fmt.Errorf("la fecha del evento ya pasó, no es posible suscribirse")
//     }
//     _, err = db.Exec("INSERT INTO suscripciones (eventoID, usuario_id) VALUES ($1, $2)", eventoID, usuarioID)
//     if err != nil {
//         return fmt.Errorf("error al suscribir usuario al evento: %w", err)
//     }

//     return nil
// }

