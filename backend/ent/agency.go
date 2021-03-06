// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/agency"
)

// Agency is the model entity for the Agency schema.
type Agency struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AGENCY holds the value of the "AGENCY" field.
	AGENCY string `json:"AGENCY,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AgencyQuery when eager-loading is set.
	Edges AgencyEdges `json:"edges"`
}

// AgencyEdges holds the relations/edges for other nodes in the graph.
type AgencyEdges struct {
	// AgenActi holds the value of the agen_acti edge.
	AgenActi []*Activity
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AgenActiOrErr returns the AgenActi value or an error if the edge
// was not loaded in eager-loading.
func (e AgencyEdges) AgenActiOrErr() ([]*Activity, error) {
	if e.loadedTypes[0] {
		return e.AgenActi, nil
	}
	return nil, &NotLoadedError{edge: "agen_acti"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Agency) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // AGENCY
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Agency fields.
func (a *Agency) assignValues(values ...interface{}) error {
	if m, n := len(values), len(agency.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field AGENCY", values[0])
	} else if value.Valid {
		a.AGENCY = value.String
	}
	return nil
}

// QueryAgenActi queries the agen_acti edge of the Agency.
func (a *Agency) QueryAgenActi() *ActivityQuery {
	return (&AgencyClient{config: a.config}).QueryAgenActi(a)
}

// Update returns a builder for updating this Agency.
// Note that, you need to call Agency.Unwrap() before calling this method, if this Agency
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Agency) Update() *AgencyUpdateOne {
	return (&AgencyClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Agency) Unwrap() *Agency {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Agency is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Agency) String() string {
	var builder strings.Builder
	builder.WriteString("Agency(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", AGENCY=")
	builder.WriteString(a.AGENCY)
	builder.WriteByte(')')
	return builder.String()
}

// Agencies is a parsable slice of Agency.
type Agencies []*Agency

func (a Agencies) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
