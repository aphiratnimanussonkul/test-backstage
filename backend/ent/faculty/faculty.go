// Code generated by entc, DO NOT EDIT.

package faculty

const (
	// Label holds the string label denoting the faculty type in the database.
	Label = "faculty"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFaculty holds the string denoting the faculty field in the database.
	FieldFaculty = "faculty"

	// EdgeFacuCour holds the string denoting the facu_cour edge name in mutations.
	EdgeFacuCour = "facu_cour"
	// EdgeFacuProf holds the string denoting the facu_prof edge name in mutations.
	EdgeFacuProf = "facu_prof"
	// EdgeFacuInst holds the string denoting the facu_inst edge name in mutations.
	EdgeFacuInst = "facu_inst"

	// Table holds the table name of the faculty in the database.
	Table = "faculties"
	// FacuCourTable is the table the holds the facu_cour relation/edge.
	FacuCourTable = "courses"
	// FacuCourInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	FacuCourInverseTable = "courses"
	// FacuCourColumn is the table column denoting the facu_cour relation/edge.
	FacuCourColumn = "faculty_facu_cour"
	// FacuProfTable is the table the holds the facu_prof relation/edge.
	FacuProfTable = "professors"
	// FacuProfInverseTable is the table name for the Professor entity.
	// It exists in this package in order to avoid circular dependency with the "professor" package.
	FacuProfInverseTable = "professors"
	// FacuProfColumn is the table column denoting the facu_prof relation/edge.
	FacuProfColumn = "faculty_facu_prof"
	// FacuInstTable is the table the holds the facu_inst relation/edge.
	FacuInstTable = "institutions"
	// FacuInstInverseTable is the table name for the Institution entity.
	// It exists in this package in order to avoid circular dependency with the "institution" package.
	FacuInstInverseTable = "institutions"
	// FacuInstColumn is the table column denoting the facu_inst relation/edge.
	FacuInstColumn = "faculty_facu_inst"
)

// Columns holds all SQL columns for faculty fields.
var Columns = []string{
	FieldID,
	FieldFaculty,
}

var (
	// FacultyValidator is a validator for the "faculty" field. It is called by the builders before save.
	FacultyValidator func(string) error
)
