package domain

// Product represents product
type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
}
