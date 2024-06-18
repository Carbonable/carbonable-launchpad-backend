// Code generated by ent, DO NOT EDIT.

package project

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/carbonable/carbonable-launchpad-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldAddress, v))
}

// Slot applies equality check predicate on the "slot" field. It's identical to SlotEQ.
func Slot(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldSlot, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldSlug, v))
}

// ValueDecimal applies equality check predicate on the "value_decimal" field. It's identical to ValueDecimalEQ.
func ValueDecimal(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldValueDecimal, v))
}

// ForecastedApr applies equality check predicate on the "forecasted_apr" field. It's identical to ForecastedAprEQ.
func ForecastedApr(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldForecastedApr, v))
}

// TotalValue applies equality check predicate on the "total_value" field. It's identical to TotalValueEQ.
func TotalValue(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldTotalValue, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldAddress, v))
}

// SlotEQ applies the EQ predicate on the "slot" field.
func SlotEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldSlot, v))
}

// SlotNEQ applies the NEQ predicate on the "slot" field.
func SlotNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldSlot, v))
}

// SlotIn applies the In predicate on the "slot" field.
func SlotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldSlot, vs...))
}

// SlotNotIn applies the NotIn predicate on the "slot" field.
func SlotNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldSlot, vs...))
}

// SlotGT applies the GT predicate on the "slot" field.
func SlotGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldSlot, v))
}

// SlotGTE applies the GTE predicate on the "slot" field.
func SlotGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldSlot, v))
}

// SlotLT applies the LT predicate on the "slot" field.
func SlotLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldSlot, v))
}

// SlotLTE applies the LTE predicate on the "slot" field.
func SlotLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldSlot, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldSlug, v))
}

// ValueDecimalEQ applies the EQ predicate on the "value_decimal" field.
func ValueDecimalEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldValueDecimal, v))
}

// ValueDecimalNEQ applies the NEQ predicate on the "value_decimal" field.
func ValueDecimalNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldValueDecimal, v))
}

// ValueDecimalIn applies the In predicate on the "value_decimal" field.
func ValueDecimalIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldValueDecimal, vs...))
}

// ValueDecimalNotIn applies the NotIn predicate on the "value_decimal" field.
func ValueDecimalNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldValueDecimal, vs...))
}

// ValueDecimalGT applies the GT predicate on the "value_decimal" field.
func ValueDecimalGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldValueDecimal, v))
}

// ValueDecimalGTE applies the GTE predicate on the "value_decimal" field.
func ValueDecimalGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldValueDecimal, v))
}

// ValueDecimalLT applies the LT predicate on the "value_decimal" field.
func ValueDecimalLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldValueDecimal, v))
}

// ValueDecimalLTE applies the LTE predicate on the "value_decimal" field.
func ValueDecimalLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldValueDecimal, v))
}

// ForecastedAprEQ applies the EQ predicate on the "forecasted_apr" field.
func ForecastedAprEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldForecastedApr, v))
}

// ForecastedAprNEQ applies the NEQ predicate on the "forecasted_apr" field.
func ForecastedAprNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldForecastedApr, v))
}

// ForecastedAprIn applies the In predicate on the "forecasted_apr" field.
func ForecastedAprIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldForecastedApr, vs...))
}

// ForecastedAprNotIn applies the NotIn predicate on the "forecasted_apr" field.
func ForecastedAprNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldForecastedApr, vs...))
}

// ForecastedAprGT applies the GT predicate on the "forecasted_apr" field.
func ForecastedAprGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldForecastedApr, v))
}

// ForecastedAprGTE applies the GTE predicate on the "forecasted_apr" field.
func ForecastedAprGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldForecastedApr, v))
}

// ForecastedAprLT applies the LT predicate on the "forecasted_apr" field.
func ForecastedAprLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldForecastedApr, v))
}

// ForecastedAprLTE applies the LTE predicate on the "forecasted_apr" field.
func ForecastedAprLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldForecastedApr, v))
}

// ForecastedAprContains applies the Contains predicate on the "forecasted_apr" field.
func ForecastedAprContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldForecastedApr, v))
}

// ForecastedAprHasPrefix applies the HasPrefix predicate on the "forecasted_apr" field.
func ForecastedAprHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldForecastedApr, v))
}

// ForecastedAprHasSuffix applies the HasSuffix predicate on the "forecasted_apr" field.
func ForecastedAprHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldForecastedApr, v))
}

// ForecastedAprIsNil applies the IsNil predicate on the "forecasted_apr" field.
func ForecastedAprIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldForecastedApr))
}

// ForecastedAprNotNil applies the NotNil predicate on the "forecasted_apr" field.
func ForecastedAprNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldForecastedApr))
}

// ForecastedAprEqualFold applies the EqualFold predicate on the "forecasted_apr" field.
func ForecastedAprEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldForecastedApr, v))
}

// ForecastedAprContainsFold applies the ContainsFold predicate on the "forecasted_apr" field.
func ForecastedAprContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldForecastedApr, v))
}

// TotalValueEQ applies the EQ predicate on the "total_value" field.
func TotalValueEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldTotalValue, v))
}

// TotalValueNEQ applies the NEQ predicate on the "total_value" field.
func TotalValueNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldTotalValue, v))
}

// TotalValueIn applies the In predicate on the "total_value" field.
func TotalValueIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldTotalValue, vs...))
}

// TotalValueNotIn applies the NotIn predicate on the "total_value" field.
func TotalValueNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldTotalValue, vs...))
}

// TotalValueGT applies the GT predicate on the "total_value" field.
func TotalValueGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldTotalValue, v))
}

// TotalValueGTE applies the GTE predicate on the "total_value" field.
func TotalValueGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldTotalValue, v))
}

// TotalValueLT applies the LT predicate on the "total_value" field.
func TotalValueLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldTotalValue, v))
}

// TotalValueLTE applies the LTE predicate on the "total_value" field.
func TotalValueLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldTotalValue, v))
}

// TotalValueContains applies the Contains predicate on the "total_value" field.
func TotalValueContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldTotalValue, v))
}

// TotalValueHasPrefix applies the HasPrefix predicate on the "total_value" field.
func TotalValueHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldTotalValue, v))
}

// TotalValueHasSuffix applies the HasSuffix predicate on the "total_value" field.
func TotalValueHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldTotalValue, v))
}

// TotalValueEqualFold applies the EqualFold predicate on the "total_value" field.
func TotalValueEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldTotalValue, v))
}

// TotalValueContainsFold applies the ContainsFold predicate on the "total_value" field.
func TotalValueContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldTotalValue, v))
}

// HasMint applies the HasEdge predicate on the "mint" edge.
func HasMint() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MintTable, MintColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMintWith applies the HasEdge predicate on the "mint" edge with a given conditions (other predicates).
func HasMintWith(preds ...predicate.Mint) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newMintStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLaunchpad applies the HasEdge predicate on the "launchpad" edge.
func HasLaunchpad() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, LaunchpadTable, LaunchpadColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLaunchpadWith applies the HasEdge predicate on the "launchpad" edge with a given conditions (other predicates).
func HasLaunchpadWith(preds ...predicate.Launchpad) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newLaunchpadStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(sql.NotPredicates(p))
}
