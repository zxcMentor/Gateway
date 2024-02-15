package models

type Order struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Items    []Item `json:"items"`
}

type Item struct {
	ProductName string `json:"productId"`
	Quantity    int    `json:"quantity"`
}
