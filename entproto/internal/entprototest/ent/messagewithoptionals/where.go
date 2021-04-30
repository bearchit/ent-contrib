// Code generated by entc, DO NOT EDIT.

package messagewithoptionals

import (
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StrOptional applies equality check predicate on the "str_optional" field. It's identical to StrOptionalEQ.
func StrOptional(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrOptional), v))
	})
}

// IntOptional applies equality check predicate on the "int_optional" field. It's identical to IntOptionalEQ.
func IntOptional(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntOptional), v))
	})
}

// UintOptional applies equality check predicate on the "uint_optional" field. It's identical to UintOptionalEQ.
func UintOptional(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUintOptional), v))
	})
}

// FloatOptional applies equality check predicate on the "float_optional" field. It's identical to FloatOptionalEQ.
func FloatOptional(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFloatOptional), v))
	})
}

// BoolOptional applies equality check predicate on the "bool_optional" field. It's identical to BoolOptionalEQ.
func BoolOptional(v bool) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBoolOptional), v))
	})
}

// BytesOptional applies equality check predicate on the "bytes_optional" field. It's identical to BytesOptionalEQ.
func BytesOptional(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBytesOptional), v))
	})
}

// UUIDOptional applies equality check predicate on the "uuid_optional" field. It's identical to UUIDOptionalEQ.
func UUIDOptional(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUIDOptional), v))
	})
}

// TimeOptional applies equality check predicate on the "time_optional" field. It's identical to TimeOptionalEQ.
func TimeOptional(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimeOptional), v))
	})
}

// StrOptionalEQ applies the EQ predicate on the "str_optional" field.
func StrOptionalEQ(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrOptional), v))
	})
}

// StrOptionalNEQ applies the NEQ predicate on the "str_optional" field.
func StrOptionalNEQ(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStrOptional), v))
	})
}

// StrOptionalIn applies the In predicate on the "str_optional" field.
func StrOptionalIn(vs ...string) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStrOptional), v...))
	})
}

// StrOptionalNotIn applies the NotIn predicate on the "str_optional" field.
func StrOptionalNotIn(vs ...string) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStrOptional), v...))
	})
}

// StrOptionalGT applies the GT predicate on the "str_optional" field.
func StrOptionalGT(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStrOptional), v))
	})
}

// StrOptionalGTE applies the GTE predicate on the "str_optional" field.
func StrOptionalGTE(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStrOptional), v))
	})
}

// StrOptionalLT applies the LT predicate on the "str_optional" field.
func StrOptionalLT(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStrOptional), v))
	})
}

// StrOptionalLTE applies the LTE predicate on the "str_optional" field.
func StrOptionalLTE(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStrOptional), v))
	})
}

// StrOptionalContains applies the Contains predicate on the "str_optional" field.
func StrOptionalContains(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStrOptional), v))
	})
}

// StrOptionalHasPrefix applies the HasPrefix predicate on the "str_optional" field.
func StrOptionalHasPrefix(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStrOptional), v))
	})
}

// StrOptionalHasSuffix applies the HasSuffix predicate on the "str_optional" field.
func StrOptionalHasSuffix(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStrOptional), v))
	})
}

// StrOptionalIsNil applies the IsNil predicate on the "str_optional" field.
func StrOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStrOptional)))
	})
}

// StrOptionalNotNil applies the NotNil predicate on the "str_optional" field.
func StrOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStrOptional)))
	})
}

// StrOptionalEqualFold applies the EqualFold predicate on the "str_optional" field.
func StrOptionalEqualFold(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStrOptional), v))
	})
}

// StrOptionalContainsFold applies the ContainsFold predicate on the "str_optional" field.
func StrOptionalContainsFold(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStrOptional), v))
	})
}

// IntOptionalEQ applies the EQ predicate on the "int_optional" field.
func IntOptionalEQ(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntOptional), v))
	})
}

// IntOptionalNEQ applies the NEQ predicate on the "int_optional" field.
func IntOptionalNEQ(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntOptional), v))
	})
}

// IntOptionalIn applies the In predicate on the "int_optional" field.
func IntOptionalIn(vs ...int8) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIntOptional), v...))
	})
}

// IntOptionalNotIn applies the NotIn predicate on the "int_optional" field.
func IntOptionalNotIn(vs ...int8) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIntOptional), v...))
	})
}

// IntOptionalGT applies the GT predicate on the "int_optional" field.
func IntOptionalGT(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntOptional), v))
	})
}

// IntOptionalGTE applies the GTE predicate on the "int_optional" field.
func IntOptionalGTE(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntOptional), v))
	})
}

// IntOptionalLT applies the LT predicate on the "int_optional" field.
func IntOptionalLT(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntOptional), v))
	})
}

// IntOptionalLTE applies the LTE predicate on the "int_optional" field.
func IntOptionalLTE(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntOptional), v))
	})
}

// IntOptionalIsNil applies the IsNil predicate on the "int_optional" field.
func IntOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIntOptional)))
	})
}

// IntOptionalNotNil applies the NotNil predicate on the "int_optional" field.
func IntOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIntOptional)))
	})
}

// UintOptionalEQ applies the EQ predicate on the "uint_optional" field.
func UintOptionalEQ(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUintOptional), v))
	})
}

// UintOptionalNEQ applies the NEQ predicate on the "uint_optional" field.
func UintOptionalNEQ(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUintOptional), v))
	})
}

// UintOptionalIn applies the In predicate on the "uint_optional" field.
func UintOptionalIn(vs ...uint8) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUintOptional), v...))
	})
}

// UintOptionalNotIn applies the NotIn predicate on the "uint_optional" field.
func UintOptionalNotIn(vs ...uint8) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUintOptional), v...))
	})
}

// UintOptionalGT applies the GT predicate on the "uint_optional" field.
func UintOptionalGT(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUintOptional), v))
	})
}

// UintOptionalGTE applies the GTE predicate on the "uint_optional" field.
func UintOptionalGTE(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUintOptional), v))
	})
}

// UintOptionalLT applies the LT predicate on the "uint_optional" field.
func UintOptionalLT(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUintOptional), v))
	})
}

// UintOptionalLTE applies the LTE predicate on the "uint_optional" field.
func UintOptionalLTE(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUintOptional), v))
	})
}

// UintOptionalIsNil applies the IsNil predicate on the "uint_optional" field.
func UintOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUintOptional)))
	})
}

// UintOptionalNotNil applies the NotNil predicate on the "uint_optional" field.
func UintOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUintOptional)))
	})
}

// FloatOptionalEQ applies the EQ predicate on the "float_optional" field.
func FloatOptionalEQ(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFloatOptional), v))
	})
}

// FloatOptionalNEQ applies the NEQ predicate on the "float_optional" field.
func FloatOptionalNEQ(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFloatOptional), v))
	})
}

// FloatOptionalIn applies the In predicate on the "float_optional" field.
func FloatOptionalIn(vs ...float32) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFloatOptional), v...))
	})
}

// FloatOptionalNotIn applies the NotIn predicate on the "float_optional" field.
func FloatOptionalNotIn(vs ...float32) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFloatOptional), v...))
	})
}

// FloatOptionalGT applies the GT predicate on the "float_optional" field.
func FloatOptionalGT(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFloatOptional), v))
	})
}

// FloatOptionalGTE applies the GTE predicate on the "float_optional" field.
func FloatOptionalGTE(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFloatOptional), v))
	})
}

// FloatOptionalLT applies the LT predicate on the "float_optional" field.
func FloatOptionalLT(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFloatOptional), v))
	})
}

// FloatOptionalLTE applies the LTE predicate on the "float_optional" field.
func FloatOptionalLTE(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFloatOptional), v))
	})
}

// FloatOptionalIsNil applies the IsNil predicate on the "float_optional" field.
func FloatOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFloatOptional)))
	})
}

// FloatOptionalNotNil applies the NotNil predicate on the "float_optional" field.
func FloatOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFloatOptional)))
	})
}

// BoolOptionalEQ applies the EQ predicate on the "bool_optional" field.
func BoolOptionalEQ(v bool) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBoolOptional), v))
	})
}

// BoolOptionalNEQ applies the NEQ predicate on the "bool_optional" field.
func BoolOptionalNEQ(v bool) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBoolOptional), v))
	})
}

// BoolOptionalIsNil applies the IsNil predicate on the "bool_optional" field.
func BoolOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBoolOptional)))
	})
}

// BoolOptionalNotNil applies the NotNil predicate on the "bool_optional" field.
func BoolOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBoolOptional)))
	})
}

// BytesOptionalEQ applies the EQ predicate on the "bytes_optional" field.
func BytesOptionalEQ(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBytesOptional), v))
	})
}

// BytesOptionalNEQ applies the NEQ predicate on the "bytes_optional" field.
func BytesOptionalNEQ(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBytesOptional), v))
	})
}

// BytesOptionalIn applies the In predicate on the "bytes_optional" field.
func BytesOptionalIn(vs ...[]byte) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBytesOptional), v...))
	})
}

// BytesOptionalNotIn applies the NotIn predicate on the "bytes_optional" field.
func BytesOptionalNotIn(vs ...[]byte) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBytesOptional), v...))
	})
}

// BytesOptionalGT applies the GT predicate on the "bytes_optional" field.
func BytesOptionalGT(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBytesOptional), v))
	})
}

// BytesOptionalGTE applies the GTE predicate on the "bytes_optional" field.
func BytesOptionalGTE(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBytesOptional), v))
	})
}

// BytesOptionalLT applies the LT predicate on the "bytes_optional" field.
func BytesOptionalLT(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBytesOptional), v))
	})
}

// BytesOptionalLTE applies the LTE predicate on the "bytes_optional" field.
func BytesOptionalLTE(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBytesOptional), v))
	})
}

// BytesOptionalIsNil applies the IsNil predicate on the "bytes_optional" field.
func BytesOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBytesOptional)))
	})
}

// BytesOptionalNotNil applies the NotNil predicate on the "bytes_optional" field.
func BytesOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBytesOptional)))
	})
}

// UUIDOptionalEQ applies the EQ predicate on the "uuid_optional" field.
func UUIDOptionalEQ(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUIDOptional), v))
	})
}

// UUIDOptionalNEQ applies the NEQ predicate on the "uuid_optional" field.
func UUIDOptionalNEQ(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUIDOptional), v))
	})
}

// UUIDOptionalIn applies the In predicate on the "uuid_optional" field.
func UUIDOptionalIn(vs ...uuid.UUID) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUIDOptional), v...))
	})
}

// UUIDOptionalNotIn applies the NotIn predicate on the "uuid_optional" field.
func UUIDOptionalNotIn(vs ...uuid.UUID) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUIDOptional), v...))
	})
}

// UUIDOptionalGT applies the GT predicate on the "uuid_optional" field.
func UUIDOptionalGT(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUIDOptional), v))
	})
}

// UUIDOptionalGTE applies the GTE predicate on the "uuid_optional" field.
func UUIDOptionalGTE(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUIDOptional), v))
	})
}

// UUIDOptionalLT applies the LT predicate on the "uuid_optional" field.
func UUIDOptionalLT(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUIDOptional), v))
	})
}

// UUIDOptionalLTE applies the LTE predicate on the "uuid_optional" field.
func UUIDOptionalLTE(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUIDOptional), v))
	})
}

// UUIDOptionalIsNil applies the IsNil predicate on the "uuid_optional" field.
func UUIDOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUUIDOptional)))
	})
}

// UUIDOptionalNotNil applies the NotNil predicate on the "uuid_optional" field.
func UUIDOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUUIDOptional)))
	})
}

// TimeOptionalEQ applies the EQ predicate on the "time_optional" field.
func TimeOptionalEQ(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimeOptional), v))
	})
}

// TimeOptionalNEQ applies the NEQ predicate on the "time_optional" field.
func TimeOptionalNEQ(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimeOptional), v))
	})
}

// TimeOptionalIn applies the In predicate on the "time_optional" field.
func TimeOptionalIn(vs ...time.Time) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimeOptional), v...))
	})
}

// TimeOptionalNotIn applies the NotIn predicate on the "time_optional" field.
func TimeOptionalNotIn(vs ...time.Time) predicate.MessageWithOptionals {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimeOptional), v...))
	})
}

// TimeOptionalGT applies the GT predicate on the "time_optional" field.
func TimeOptionalGT(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimeOptional), v))
	})
}

// TimeOptionalGTE applies the GTE predicate on the "time_optional" field.
func TimeOptionalGTE(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimeOptional), v))
	})
}

// TimeOptionalLT applies the LT predicate on the "time_optional" field.
func TimeOptionalLT(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimeOptional), v))
	})
}

// TimeOptionalLTE applies the LTE predicate on the "time_optional" field.
func TimeOptionalLTE(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimeOptional), v))
	})
}

// TimeOptionalIsNil applies the IsNil predicate on the "time_optional" field.
func TimeOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTimeOptional)))
	})
}

// TimeOptionalNotNil applies the NotNil predicate on the "time_optional" field.
func TimeOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTimeOptional)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MessageWithOptionals) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MessageWithOptionals) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MessageWithOptionals) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(func(s *sql.Selector) {
		p(s.Not())
	})
}
