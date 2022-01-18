package catalog

//easyjson:json
type Category struct {
	uuid        string
	description string
	name        string
}

//easyjson:json
type Product struct {
	uuid        string
	categoryId  string
	description string
	name        string
	price       float32
}

//easyjson:json
type Order struct {
	uuid         string
	customerData string
	deliveryData string
}

//easyjson:json
type OrderCustomerData struct {
	email string
	name  string
	phone string
}

//easyjson:json
type OrderDeliveryData struct {
	address string
	price   float32
}

//easyjson:json
type OrderItem struct {
	amount    string
	orderId   string
	productId string
}
