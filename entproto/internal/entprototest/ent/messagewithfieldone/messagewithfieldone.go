// Code generated by entc, DO NOT EDIT.

package messagewithfieldone

const (
	// Label holds the string label denoting the messagewithfieldone type in the database.
	Label = "message_with_field_one"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFieldOne holds the string denoting the field_one field in the database.
	FieldFieldOne = "field_one"
	// Table holds the table name of the messagewithfieldone in the database.
	Table = "message_with_field_ones"
)

// Columns holds all SQL columns for messagewithfieldone fields.
var Columns = []string{
	FieldID,
	FieldFieldOne,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
