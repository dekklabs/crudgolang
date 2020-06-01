package entidades

//Producto entidad de producto
type Producto struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Discount_price float64 `json:"discount_price,omitempty"`
	Description    string  `json:"description,omitempty"`
	Capacity       float64 `json:"capacity,omitempty"`
	Deliverable    int64   `json:"deliverable,omitempty"`
	Image          string  `json:"image,omitempty"`
	Category_id    int64   `json:"category_id,omitempty"`
	Market_id      int64   `json:"market_id,omitempty"`
	Created_at     string  `json:"created_at,omitempty"`
	Updated_at     string  `json:"updated_at,omitempty"`
}
