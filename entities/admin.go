package entities

import "time"

type Admin struct {
	ID        int       `json:"id"`                     
	Email     string    `json:"email"`                  
	Password  string    `json:"password"`               
	Role      string    `json:"role"`                   
	CreatedAt time.Time `json:"created_at"`            
	UpdatedAt time.Time `json:"updated_at"`             
}