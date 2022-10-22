package models

import "time"

type Tabler interface {
	TableName() string
}

type TodoModel struct {
	TodoId    uint      `gorm:"primaryKey" json:"order_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (TodoModel) TableName() string {
	return "public.Todo"
}
