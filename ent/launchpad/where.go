// Code generated by ent, DO NOT EDIT.

package launchpad

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/carbonable/carbonable-launchpad-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldLTE(FieldID, id))
}

// IsReady applies equality check predicate on the "is_ready" field. It's identical to IsReadyEQ.
func IsReady(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldIsReady, v))
}

// WhitelistedSaleOpen applies equality check predicate on the "whitelisted_sale_open" field. It's identical to WhitelistedSaleOpenEQ.
func WhitelistedSaleOpen(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldWhitelistedSaleOpen, v))
}

// PublicSaleOpen applies equality check predicate on the "public_sale_open" field. It's identical to PublicSaleOpenEQ.
func PublicSaleOpen(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldPublicSaleOpen, v))
}

// IsSoldOut applies equality check predicate on the "is_sold_out" field. It's identical to IsSoldOutEQ.
func IsSoldOut(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldIsSoldOut, v))
}

// IsCanceled applies equality check predicate on the "is_canceled" field. It's identical to IsCanceledEQ.
func IsCanceled(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldIsCanceled, v))
}

// IsReadyEQ applies the EQ predicate on the "is_ready" field.
func IsReadyEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldIsReady, v))
}

// IsReadyNEQ applies the NEQ predicate on the "is_ready" field.
func IsReadyNEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNEQ(FieldIsReady, v))
}

// WhitelistedSaleOpenEQ applies the EQ predicate on the "whitelisted_sale_open" field.
func WhitelistedSaleOpenEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldWhitelistedSaleOpen, v))
}

// WhitelistedSaleOpenNEQ applies the NEQ predicate on the "whitelisted_sale_open" field.
func WhitelistedSaleOpenNEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNEQ(FieldWhitelistedSaleOpen, v))
}

// PublicSaleOpenEQ applies the EQ predicate on the "public_sale_open" field.
func PublicSaleOpenEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldPublicSaleOpen, v))
}

// PublicSaleOpenNEQ applies the NEQ predicate on the "public_sale_open" field.
func PublicSaleOpenNEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNEQ(FieldPublicSaleOpen, v))
}

// IsSoldOutEQ applies the EQ predicate on the "is_sold_out" field.
func IsSoldOutEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldIsSoldOut, v))
}

// IsSoldOutNEQ applies the NEQ predicate on the "is_sold_out" field.
func IsSoldOutNEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNEQ(FieldIsSoldOut, v))
}

// IsCanceledEQ applies the EQ predicate on the "is_canceled" field.
func IsCanceledEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldEQ(FieldIsCanceled, v))
}

// IsCanceledNEQ applies the NEQ predicate on the "is_canceled" field.
func IsCanceledNEQ(v bool) predicate.Launchpad {
	return predicate.Launchpad(sql.FieldNEQ(FieldIsCanceled, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Launchpad {
	return predicate.Launchpad(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Launchpad {
	return predicate.Launchpad(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Launchpad) predicate.Launchpad {
	return predicate.Launchpad(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Launchpad) predicate.Launchpad {
	return predicate.Launchpad(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Launchpad) predicate.Launchpad {
	return predicate.Launchpad(sql.NotPredicates(p))
}
