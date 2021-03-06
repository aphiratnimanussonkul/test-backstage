// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/country"
)

// Country is the model entity for the Country schema.
type Country struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges CountryEdges `json:"edges"`
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// CounProv holds the value of the coun_prov edge.
	CounProv []*Province
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CounProvOrErr returns the CounProv value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) CounProvOrErr() ([]*Province, error) {
	if e.loadedTypes[0] {
		return e.CounProv, nil
	}
	return nil, &NotLoadedError{edge: "coun_prov"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // country
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(values ...interface{}) error {
	if m, n := len(values), len(country.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field country", values[0])
	} else if value.Valid {
		c.Country = value.String
	}
	return nil
}

// QueryCounProv queries the coun_prov edge of the Country.
func (c *Country) QueryCounProv() *ProvinceQuery {
	return (&CountryClient{config: c.config}).QueryCounProv(c)
}

// Update returns a builder for updating this Country.
// Note that, you need to call Country.Unwrap() before calling this method, if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return (&CountryClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", country=")
	builder.WriteString(c.Country)
	builder.WriteByte(')')
	return builder.String()
}

// Countries is a parsable slice of Country.
type Countries []*Country

func (c Countries) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
