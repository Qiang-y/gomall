package types

type OrderItem struct {
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float32
}

// Order 订单相关结构�?
type Order struct {
	OrderId     string
	CreatedDate string
	Cost        float32
	Items       []OrderItem
}
