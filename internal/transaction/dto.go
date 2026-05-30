package transaction

import "time"

type CreateTransactionRequest struct {
	CategoryID      int64     `json:"categoryId" binding:"required"`
	Type            string    `json:"type" binding:"required,oneof=income expense"`
	Amount          float64   `json:"amount" binding:"required,gt=0"`
	Note            string    `json:"note"`
	TransactionDate time.Time `json:"transactionDate" binding:"required"`
}

type TransactionResponse struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"userId"`
	CategoryID      int64     `json:"categoryId"`
	Type            string    `json:"type"`
	Amount          float64   `json:"amount"`
	Note            string    `json:"note"`
	TransactionDate time.Time `json:"transactionDate"`
}

type UpdateTransactionRequest struct {
	CategoryID int64 `json:"categoryId" binding:"required"`

	Amount float64 `json:"amount" binding:"required,gt=0"`

	Type string `json:"type" binding:"required,oneof=income expense"`

	Note string `json:"note"`

	TransactionDate time.Time `json:"transactionDate" binding:"required"`
}

type ListTransactionResponse struct {
	ID int64 `json:"id"`

	CategoryID int64 `json:"categoryId"`

	CategoryName string `json:"categoryName"`

	Type string `json:"type"`

	Amount float64 `json:"amount"`

	Note string `json:"note"`

	TransactionDate time.Time `json:"transactionDate"`
}

type SummaryResponse struct {
	TotalIncome  float64 `json:"totalIncome"`
	TotalExpense float64 `json:"totalExpense"`
	Balance      float64 `json:"balance"`
}
