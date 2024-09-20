package models

import (
    "time" 
   )


type Event struct {
    ID int`json:"id"`
    Title string`json:"title"`
    DescriptionShort string `json:"descriptionShort"` 
    DescriptionLarge string `json:"descriptionLarge"` 
    Date time.Time `json:"date"`
    Organizer string `json:"organizer"`
    Place string `json:"place"`
    State string `json:"state"` 
}

