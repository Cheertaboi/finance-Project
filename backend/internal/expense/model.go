package expense

type Expense struct {
	ID          string `json:"id"`
	Amount      int64  `json:"amount"` // stored in paise
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
}
