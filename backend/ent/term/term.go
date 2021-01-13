// Code generated by entc, DO NOT EDIT.

package term

const (
	// Label holds the string label denoting the term type in the database.
	Label = "term"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSemester holds the string denoting the semester field in the database.
	FieldSemester = "semester"

	// EdgeTermResu holds the string denoting the term_resu edge name in mutations.
	EdgeTermResu = "term_resu"
	// EdgeTermActi holds the string denoting the term_acti edge name in mutations.
	EdgeTermActi = "term_acti"

	// Table holds the table name of the term in the database.
	Table = "terms"
	// TermResuTable is the table the holds the term_resu relation/edge.
	TermResuTable = "results"
	// TermResuInverseTable is the table name for the Results entity.
	// It exists in this package in order to avoid circular dependency with the "results" package.
	TermResuInverseTable = "results"
	// TermResuColumn is the table column denoting the term_resu relation/edge.
	TermResuColumn = "term_term_resu"
	// TermActiTable is the table the holds the term_acti relation/edge.
	TermActiTable = "activities"
	// TermActiInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	TermActiInverseTable = "activities"
	// TermActiColumn is the table column denoting the term_acti relation/edge.
	TermActiColumn = "term_term_acti"
)

// Columns holds all SQL columns for term fields.
var Columns = []string{
	FieldID,
	FieldSemester,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Term type.
var ForeignKeys = []string{
	"year_year_term",
}

var (
	// SemesterValidator is a validator for the "semester" field. It is called by the builders before save.
	SemesterValidator func(int) error
)
