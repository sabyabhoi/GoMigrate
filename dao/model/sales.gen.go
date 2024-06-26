// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSale = "Sales"

// Sale mapped from table <Sales>
type Sale struct {
	TransactionID int32     `gorm:"column:TransactionID;primaryKey" json:"TransactionID"`
	ProductID     int32     `gorm:"column:ProductID" json:"ProductID"`
	Quantity      int32     `gorm:"column:Quantity" json:"Quantity"`
	SaleDate      time.Time `gorm:"column:SaleDate" json:"SaleDate"`
	TotalPrice    float64   `gorm:"column:TotalPrice" json:"TotalPrice"`
}

// TableName Sale's table name
func (*Sale) TableName() string {
	return TableNameSale
}
