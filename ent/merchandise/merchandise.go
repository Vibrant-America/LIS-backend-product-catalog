// Code generated by entc, DO NOT EDIT.

package merchandise

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the merchandise type in the database.
	Label = "merchandise"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTypeID holds the string denoting the type_id field in the database.
	FieldTypeID = "type_id"
	// FieldStock holds the string denoting the stock field in the database.
	FieldStock = "stock"
	// FieldPendingPeriod holds the string denoting the pending_period field in the database.
	FieldPendingPeriod = "pending_period"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldPictureURL holds the string denoting the picture_url field in the database.
	FieldPictureURL = "picture_url"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldRefundableRatio holds the string denoting the refundable_ratio field in the database.
	FieldRefundableRatio = "refundable_ratio"
	// FieldDeductionRatio holds the string denoting the deduction_ratio field in the database.
	FieldDeductionRatio = "deduction_ratio"
	// FieldDeductionReason holds the string denoting the deduction_reason field in the database.
	FieldDeductionReason = "deduction_reason"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the merchandise in the database.
	Table = "merchandises"
)

// Columns holds all SQL columns for merchandise fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldTypeID,
	FieldStock,
	FieldPendingPeriod,
	FieldName,
	FieldDisplayName,
	FieldPictureURL,
	FieldDescription,
	FieldPrice,
	FieldRefundableRatio,
	FieldDeductionRatio,
	FieldDeductionReason,
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
	// DefaultStock holds the default value on creation for the "stock" field.
	DefaultStock int
	// DefaultPendingPeriod holds the default value on creation for the "pending_period" field.
	DefaultPendingPeriod int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypeItem is the default value of the Type enum.
const DefaultType = TypeItem

// Type values.
const (
	TypeItem    Type = "item"
	TypeService Type = "service"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeItem, TypeService:
		return nil
	default:
		return fmt.Errorf("merchandise: invalid enum value for type field: %q", _type)
	}
}
