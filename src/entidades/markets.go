package entidades

type Markets struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Phone       int64  `json:"phone"`
	Mobile      int64  `json:"mobile"`
	Information string `json:"information"`
	Closed      int64  `json:"closed"`
	Image       string `json:"image"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
