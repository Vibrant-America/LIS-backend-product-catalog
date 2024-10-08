// Code generated by entc, DO NOT EDIT.

package item

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTypeID holds the string denoting the type_id field in the database.
	FieldTypeID = "type_id"
	// FieldOrderTypeID holds the string denoting the order_type_id field in the database.
	FieldOrderTypeID = "order_type_id"
	// FieldIsOrderable holds the string denoting the is_orderable field in the database.
	FieldIsOrderable = "is_orderable"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldUniqueEmrCode holds the string denoting the unique_emr_code field in the database.
	FieldUniqueEmrCode = "unique_emr_code"
	// FieldTaxCode holds the string denoting the tax_code field in the database.
	FieldTaxCode = "tax_code"
	// FieldWeblink holds the string denoting the weblink field in the database.
	FieldWeblink = "weblink"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSampleTypeList holds the string denoting the sample_type_list field in the database.
	FieldSampleTypeList = "sample_type_list"
	// FieldSubpackagesList holds the string denoting the subpackages_list field in the database.
	FieldSubpackagesList = "subpackages_list"
	// FieldSubtestsList holds the string denoting the subtests_list field in the database.
	FieldSubtestsList = "subtests_list"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the item in the database.
	Table = "items"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldTypeID,
	FieldOrderTypeID,
	FieldIsOrderable,
	FieldName,
	FieldDisplayName,
	FieldUniqueEmrCode,
	FieldTaxCode,
	FieldWeblink,
	FieldDescription,
	FieldSampleTypeList,
	FieldSubpackagesList,
	FieldSubtestsList,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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

var (
	// DefaultIsOrderable holds the default value on creation for the "is_orderable" field.
	DefaultIsOrderable bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypeTEST is the default value of the Type enum.
const DefaultType = TypeTEST

// Type values.
const (
	TypeTEST  Type = "TEST"
	TypeGROUP Type = "GROUP"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeTEST, TypeGROUP:
		return nil
	default:
		return fmt.Errorf("item: invalid enum value for type field: %q", _type)
	}
}
