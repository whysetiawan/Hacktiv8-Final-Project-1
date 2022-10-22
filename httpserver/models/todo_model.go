package models

import "time"

type Tabler interface {
	TableName() string
}

type TodoModel struct {
	TodoId    uint      `gorm:"primaryKey" json:"order_id"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (TodoModel) TableName() string {
	return "public.Todo"
}

type TodoParams struct {
	ID int64 `uri:"id" binding:"required"`
}
