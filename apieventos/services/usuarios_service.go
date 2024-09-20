package services

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
    
    "apieventos/models"
)

var admins = []models.Admin{
    {ID: 1, Adminname: "admin", Password: hashPassword("adminpass"), Role: "admin"},
}

func hashPassword(password string) string {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashed)
}

func ValidateAdmin(adminname, password string) (*models.Admin, bool) {
    for _, admin := range admins {
        if admin.Adminname == adminname {
            if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err == nil {
                return &admin, true
            }
        }
    }
    return nil, false
}

var usuarios = []models.User{
    {
        ID:  1,
        Email: "usuario1@example.com",
    },
    {
        ID:  2,
        Email: "usuario2@example.com",
    },
}

var nextUserID = 3

func GetUsuarioIDByEmail(email string) (int, error) {
    for _, usuario := range usuarios {
        if usuario.Email == email {
            return usuario.ID, nil
        }
    }
    return 0, fmt.Errorf("Usuario no encontrado")
}

func SetUser(user models.User) error {
    for _, u := range usuarios {
        if u.Email == user.Email {
            return fmt.Errorf("El usuario con el email %s ya existe", user.Email)
        }
    }

    user.ID = nextUserID
    usuarios = append(usuarios, user) 
    nextUserID++

    fmt.Printf("Usuario creado: ID=%d, Email=%s\n", user.ID, user.Email)
    return nil
}


// func GetUsuarioIDByEmail(email string) (int, error) {
//     db := GetDB()

//     var usuarioID int
//     err := db.QueryRow("SELECT id FROM usuarios WHERE email = $1", email).Scan(&usuarioID)
//     if err == sql.ErrNoRows {
//         return 0, fmt.Errorf("error al insertar usuario: %w", err)
//     } else if err != nil {
//         return 0, fmt.Errorf("Error al buscar usuario por email: %w", err)
//     }

//     return usuarioID, nil
// }

// func SetUser(user models.User) error {
//     db := GetDB()

//     _, err := db.Exec("INSERT INTO usuarios (email) VALUES ($1)", user.Email)
//     if err != nil {
//         return fmt.Errorf("error al insertar usuario: %w", err)
//     }

//     return nil
// }

