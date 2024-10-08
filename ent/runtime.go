// Code generated by entc, DO NOT EDIT.

package ent

import (
	"productCatalog/ent/item"
	"productCatalog/ent/merchandise"
	"productCatalog/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescIsOrderable is the schema descriptor for is_orderable field.
	itemDescIsOrderable := itemFields[3].Descriptor()
	// item.DefaultIsOrderable holds the default value on creation for the is_orderable field.
	item.DefaultIsOrderable = itemDescIsOrderable.Default.(bool)
	// itemDescCreatedAt is the schema descriptor for created_at field.
	itemDescCreatedAt := itemFields[13].Descriptor()
	// item.DefaultCreatedAt holds the default value on creation for the created_at field.
	item.DefaultCreatedAt = itemDescCreatedAt.Default.(func() time.Time)
	// itemDescUpdatedAt is the schema descriptor for updated_at field.
	itemDescUpdatedAt := itemFields[14].Descriptor()
	// item.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	item.DefaultUpdatedAt = itemDescUpdatedAt.Default.(func() time.Time)
	// item.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	item.UpdateDefaultUpdatedAt = itemDescUpdatedAt.UpdateDefault.(func() time.Time)
	merchandiseFields := schema.Merchandise{}.Fields()
	_ = merchandiseFields
	// merchandiseDescStock is the schema descriptor for stock field.
	merchandiseDescStock := merchandiseFields[2].Descriptor()
	// merchandise.DefaultStock holds the default value on creation for the stock field.
	merchandise.DefaultStock = merchandiseDescStock.Default.(int)
	// merchandiseDescPendingPeriod is the schema descriptor for pending_period field.
	merchandiseDescPendingPeriod := merchandiseFields[3].Descriptor()
	// merchandise.DefaultPendingPeriod holds the default value on creation for the pending_period field.
	merchandise.DefaultPendingPeriod = merchandiseDescPendingPeriod.Default.(int)
	// merchandiseDescCreatedAt is the schema descriptor for created_at field.
	merchandiseDescCreatedAt := merchandiseFields[12].Descriptor()
	// merchandise.DefaultCreatedAt holds the default value on creation for the created_at field.
	merchandise.DefaultCreatedAt = merchandiseDescCreatedAt.Default.(func() time.Time)
	// merchandiseDescUpdatedAt is the schema descriptor for updated_at field.
	merchandiseDescUpdatedAt := merchandiseFields[13].Descriptor()
	// merchandise.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	merchandise.DefaultUpdatedAt = merchandiseDescUpdatedAt.Default.(func() time.Time)
	// merchandise.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	merchandise.UpdateDefaultUpdatedAt = merchandiseDescUpdatedAt.UpdateDefault.(func() time.Time)
}
