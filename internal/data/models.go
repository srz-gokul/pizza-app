package data

import "time"

const (
	CookingStage1 = "start"
	CookingStage2 = "dough-prep"
	CookingStage3 = "oven-bake"
	CookingStage4 = "topping-art"
	CookingStage5 = "done"
)

// OrderData holds the order related details.
type OrderData struct {
	ID           int        `json:"id,omitempty"`
	PizzaID      int        `json:"pizza_id,omitempty"`
	PizzaSize    string     `json:"pizza_size,omitempty"`
	CookingStage string     `json:"cooking_stage,omitempty"`
	UserID       int        `json:"user_id,omitempty"`
	StartTime    *time.Time `json:"start_time,omitempty"`
	EndTime      *time.Time `json:"end_time,omitempty"`
	IsActive     bool       `json:"is_active,omitempty"`
}

// newPizza holds the pizza related details.
type NewPizza struct {
	Pizza_id      int
	Pizza_name    string
	Pizza_type    string
	Pizza_size    string
	user_id       int
	starting_date *time.Time
	ending_date   *time.Time
}
