package entidades

//Producto entidad de producto
type Producto struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Discount_price float64 `json:"discount_price"`
	Description    string  `json:"description"`
	Capacity       float64 `json:"capacity"`
	Deliverable    int64   `json:"deliverable"`
	Image          string  `json:"image"`
	Category_id    int64   `json:"category_id"`
	Market_id      int64   `json:"market_id"`
	Created_at     string  `json:"created_at"`
	Updated_at     string  `json:"updated_at"`
}
