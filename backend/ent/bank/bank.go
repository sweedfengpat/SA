// Code generated by entc, DO NOT EDIT.

package bank

const (
	// Label holds the string label denoting the bank type in the database.
	Label = "bank"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBank holds the string denoting the bank field in the database.
	FieldBank = "bank"

	// EdgeBill holds the string denoting the bill edge name in mutations.
	EdgeBill = "bill"

	// Table holds the table name of the bank in the database.
	Table = "banks"
	// BillTable is the table the holds the bill relation/edge.
	BillTable = "bills"
	// BillInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillInverseTable = "bills"
	// BillColumn is the table column denoting the bill relation/edge.
	BillColumn = "bank_id"
)

// Columns holds all SQL columns for bank fields.
var Columns = []string{
	FieldID,
	FieldBank,
}
