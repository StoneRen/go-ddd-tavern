package entity

// Customer 聚合
// 可变的为指针类型
// 不可变的为值类型
type Customer struct {
	person       *Person
	products     []*Item
	transactions []Transaction
}
