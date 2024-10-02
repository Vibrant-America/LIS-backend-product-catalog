// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"productCatalog/ent/merchandise"
	"productCatalog/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MerchandiseUpdate is the builder for updating Merchandise entities.
type MerchandiseUpdate struct {
	config
	hooks    []Hook
	mutation *MerchandiseMutation
}

// Where appends a list predicates to the MerchandiseUpdate builder.
func (mu *MerchandiseUpdate) Where(ps ...predicate.Merchandise) *MerchandiseUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetType sets the "type" field.
func (mu *MerchandiseUpdate) SetType(m merchandise.Type) *MerchandiseUpdate {
	mu.mutation.SetType(m)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableType(m *merchandise.Type) *MerchandiseUpdate {
	if m != nil {
		mu.SetType(*m)
	}
	return mu
}

// SetTypeID sets the "type_id" field.
func (mu *MerchandiseUpdate) SetTypeID(s string) *MerchandiseUpdate {
	mu.mutation.SetTypeID(s)
	return mu
}

// SetNillableTypeID sets the "type_id" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableTypeID(s *string) *MerchandiseUpdate {
	if s != nil {
		mu.SetTypeID(*s)
	}
	return mu
}

// ClearTypeID clears the value of the "type_id" field.
func (mu *MerchandiseUpdate) ClearTypeID() *MerchandiseUpdate {
	mu.mutation.ClearTypeID()
	return mu
}

// SetStock sets the "stock" field.
func (mu *MerchandiseUpdate) SetStock(i int) *MerchandiseUpdate {
	mu.mutation.ResetStock()
	mu.mutation.SetStock(i)
	return mu
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableStock(i *int) *MerchandiseUpdate {
	if i != nil {
		mu.SetStock(*i)
	}
	return mu
}

// AddStock adds i to the "stock" field.
func (mu *MerchandiseUpdate) AddStock(i int) *MerchandiseUpdate {
	mu.mutation.AddStock(i)
	return mu
}

// SetPendingPeriod sets the "pending_period" field.
func (mu *MerchandiseUpdate) SetPendingPeriod(i int) *MerchandiseUpdate {
	mu.mutation.ResetPendingPeriod()
	mu.mutation.SetPendingPeriod(i)
	return mu
}

// SetNillablePendingPeriod sets the "pending_period" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillablePendingPeriod(i *int) *MerchandiseUpdate {
	if i != nil {
		mu.SetPendingPeriod(*i)
	}
	return mu
}

// AddPendingPeriod adds i to the "pending_period" field.
func (mu *MerchandiseUpdate) AddPendingPeriod(i int) *MerchandiseUpdate {
	mu.mutation.AddPendingPeriod(i)
	return mu
}

// SetName sets the "name" field.
func (mu *MerchandiseUpdate) SetName(s string) *MerchandiseUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableName(s *string) *MerchandiseUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// ClearName clears the value of the "name" field.
func (mu *MerchandiseUpdate) ClearName() *MerchandiseUpdate {
	mu.mutation.ClearName()
	return mu
}

// SetDisplayName sets the "display_name" field.
func (mu *MerchandiseUpdate) SetDisplayName(s string) *MerchandiseUpdate {
	mu.mutation.SetDisplayName(s)
	return mu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableDisplayName(s *string) *MerchandiseUpdate {
	if s != nil {
		mu.SetDisplayName(*s)
	}
	return mu
}

// ClearDisplayName clears the value of the "display_name" field.
func (mu *MerchandiseUpdate) ClearDisplayName() *MerchandiseUpdate {
	mu.mutation.ClearDisplayName()
	return mu
}

// SetPictureURL sets the "picture_url" field.
func (mu *MerchandiseUpdate) SetPictureURL(s string) *MerchandiseUpdate {
	mu.mutation.SetPictureURL(s)
	return mu
}

// SetNillablePictureURL sets the "picture_url" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillablePictureURL(s *string) *MerchandiseUpdate {
	if s != nil {
		mu.SetPictureURL(*s)
	}
	return mu
}

// ClearPictureURL clears the value of the "picture_url" field.
func (mu *MerchandiseUpdate) ClearPictureURL() *MerchandiseUpdate {
	mu.mutation.ClearPictureURL()
	return mu
}

// SetDescription sets the "description" field.
func (mu *MerchandiseUpdate) SetDescription(s string) *MerchandiseUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableDescription(s *string) *MerchandiseUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// ClearDescription clears the value of the "description" field.
func (mu *MerchandiseUpdate) ClearDescription() *MerchandiseUpdate {
	mu.mutation.ClearDescription()
	return mu
}

// SetPrice sets the "price" field.
func (mu *MerchandiseUpdate) SetPrice(f float64) *MerchandiseUpdate {
	mu.mutation.ResetPrice()
	mu.mutation.SetPrice(f)
	return mu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillablePrice(f *float64) *MerchandiseUpdate {
	if f != nil {
		mu.SetPrice(*f)
	}
	return mu
}

// AddPrice adds f to the "price" field.
func (mu *MerchandiseUpdate) AddPrice(f float64) *MerchandiseUpdate {
	mu.mutation.AddPrice(f)
	return mu
}

// ClearPrice clears the value of the "price" field.
func (mu *MerchandiseUpdate) ClearPrice() *MerchandiseUpdate {
	mu.mutation.ClearPrice()
	return mu
}

// SetRefundableRatio sets the "refundable_ratio" field.
func (mu *MerchandiseUpdate) SetRefundableRatio(f float64) *MerchandiseUpdate {
	mu.mutation.ResetRefundableRatio()
	mu.mutation.SetRefundableRatio(f)
	return mu
}

// SetNillableRefundableRatio sets the "refundable_ratio" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableRefundableRatio(f *float64) *MerchandiseUpdate {
	if f != nil {
		mu.SetRefundableRatio(*f)
	}
	return mu
}

// AddRefundableRatio adds f to the "refundable_ratio" field.
func (mu *MerchandiseUpdate) AddRefundableRatio(f float64) *MerchandiseUpdate {
	mu.mutation.AddRefundableRatio(f)
	return mu
}

// ClearRefundableRatio clears the value of the "refundable_ratio" field.
func (mu *MerchandiseUpdate) ClearRefundableRatio() *MerchandiseUpdate {
	mu.mutation.ClearRefundableRatio()
	return mu
}

// SetDeductionRatio sets the "deduction_ratio" field.
func (mu *MerchandiseUpdate) SetDeductionRatio(f float64) *MerchandiseUpdate {
	mu.mutation.ResetDeductionRatio()
	mu.mutation.SetDeductionRatio(f)
	return mu
}

// SetNillableDeductionRatio sets the "deduction_ratio" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableDeductionRatio(f *float64) *MerchandiseUpdate {
	if f != nil {
		mu.SetDeductionRatio(*f)
	}
	return mu
}

// AddDeductionRatio adds f to the "deduction_ratio" field.
func (mu *MerchandiseUpdate) AddDeductionRatio(f float64) *MerchandiseUpdate {
	mu.mutation.AddDeductionRatio(f)
	return mu
}

// ClearDeductionRatio clears the value of the "deduction_ratio" field.
func (mu *MerchandiseUpdate) ClearDeductionRatio() *MerchandiseUpdate {
	mu.mutation.ClearDeductionRatio()
	return mu
}

// SetDeductionReason sets the "deduction_reason" field.
func (mu *MerchandiseUpdate) SetDeductionReason(s string) *MerchandiseUpdate {
	mu.mutation.SetDeductionReason(s)
	return mu
}

// SetNillableDeductionReason sets the "deduction_reason" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableDeductionReason(s *string) *MerchandiseUpdate {
	if s != nil {
		mu.SetDeductionReason(*s)
	}
	return mu
}

// ClearDeductionReason clears the value of the "deduction_reason" field.
func (mu *MerchandiseUpdate) ClearDeductionReason() *MerchandiseUpdate {
	mu.mutation.ClearDeductionReason()
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MerchandiseUpdate) SetCreatedAt(t time.Time) *MerchandiseUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableCreatedAt(t *time.Time) *MerchandiseUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MerchandiseUpdate) SetUpdatedAt(t time.Time) *MerchandiseUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MerchandiseUpdate) SetDeletedAt(t time.Time) *MerchandiseUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MerchandiseUpdate) SetNillableDeletedAt(t *time.Time) *MerchandiseUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MerchandiseUpdate) ClearDeletedAt() *MerchandiseUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// Mutation returns the MerchandiseMutation object of the builder.
func (mu *MerchandiseUpdate) Mutation() *MerchandiseMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MerchandiseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchandiseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MerchandiseUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MerchandiseUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MerchandiseUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MerchandiseUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := merchandise.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MerchandiseUpdate) check() error {
	if v, ok := mu.mutation.GetType(); ok {
		if err := merchandise.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Merchandise.type": %w`, err)}
		}
	}
	return nil
}

func (mu *MerchandiseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchandise.Table,
			Columns: merchandise.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchandise.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: merchandise.FieldType,
		})
	}
	if value, ok := mu.mutation.TypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldTypeID,
		})
	}
	if mu.mutation.TypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldTypeID,
		})
	}
	if value, ok := mu.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldStock,
		})
	}
	if value, ok := mu.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldStock,
		})
	}
	if value, ok := mu.mutation.PendingPeriod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldPendingPeriod,
		})
	}
	if value, ok := mu.mutation.AddedPendingPeriod(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldPendingPeriod,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldName,
		})
	}
	if mu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldName,
		})
	}
	if value, ok := mu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldDisplayName,
		})
	}
	if mu.mutation.DisplayNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldDisplayName,
		})
	}
	if value, ok := mu.mutation.PictureURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldPictureURL,
		})
	}
	if mu.mutation.PictureURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldPictureURL,
		})
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldDescription,
		})
	}
	if mu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldDescription,
		})
	}
	if value, ok := mu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldPrice,
		})
	}
	if value, ok := mu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldPrice,
		})
	}
	if mu.mutation.PriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: merchandise.FieldPrice,
		})
	}
	if value, ok := mu.mutation.RefundableRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldRefundableRatio,
		})
	}
	if value, ok := mu.mutation.AddedRefundableRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldRefundableRatio,
		})
	}
	if mu.mutation.RefundableRatioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: merchandise.FieldRefundableRatio,
		})
	}
	if value, ok := mu.mutation.DeductionRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldDeductionRatio,
		})
	}
	if value, ok := mu.mutation.AddedDeductionRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldDeductionRatio,
		})
	}
	if mu.mutation.DeductionRatioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: merchandise.FieldDeductionRatio,
		})
	}
	if value, ok := mu.mutation.DeductionReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldDeductionReason,
		})
	}
	if mu.mutation.DeductionReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldDeductionReason,
		})
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchandise.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchandise.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchandise.FieldDeletedAt,
		})
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: merchandise.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchandise.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MerchandiseUpdateOne is the builder for updating a single Merchandise entity.
type MerchandiseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchandiseMutation
}

// SetType sets the "type" field.
func (muo *MerchandiseUpdateOne) SetType(m merchandise.Type) *MerchandiseUpdateOne {
	muo.mutation.SetType(m)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableType(m *merchandise.Type) *MerchandiseUpdateOne {
	if m != nil {
		muo.SetType(*m)
	}
	return muo
}

// SetTypeID sets the "type_id" field.
func (muo *MerchandiseUpdateOne) SetTypeID(s string) *MerchandiseUpdateOne {
	muo.mutation.SetTypeID(s)
	return muo
}

// SetNillableTypeID sets the "type_id" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableTypeID(s *string) *MerchandiseUpdateOne {
	if s != nil {
		muo.SetTypeID(*s)
	}
	return muo
}

// ClearTypeID clears the value of the "type_id" field.
func (muo *MerchandiseUpdateOne) ClearTypeID() *MerchandiseUpdateOne {
	muo.mutation.ClearTypeID()
	return muo
}

// SetStock sets the "stock" field.
func (muo *MerchandiseUpdateOne) SetStock(i int) *MerchandiseUpdateOne {
	muo.mutation.ResetStock()
	muo.mutation.SetStock(i)
	return muo
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableStock(i *int) *MerchandiseUpdateOne {
	if i != nil {
		muo.SetStock(*i)
	}
	return muo
}

// AddStock adds i to the "stock" field.
func (muo *MerchandiseUpdateOne) AddStock(i int) *MerchandiseUpdateOne {
	muo.mutation.AddStock(i)
	return muo
}

// SetPendingPeriod sets the "pending_period" field.
func (muo *MerchandiseUpdateOne) SetPendingPeriod(i int) *MerchandiseUpdateOne {
	muo.mutation.ResetPendingPeriod()
	muo.mutation.SetPendingPeriod(i)
	return muo
}

// SetNillablePendingPeriod sets the "pending_period" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillablePendingPeriod(i *int) *MerchandiseUpdateOne {
	if i != nil {
		muo.SetPendingPeriod(*i)
	}
	return muo
}

// AddPendingPeriod adds i to the "pending_period" field.
func (muo *MerchandiseUpdateOne) AddPendingPeriod(i int) *MerchandiseUpdateOne {
	muo.mutation.AddPendingPeriod(i)
	return muo
}

// SetName sets the "name" field.
func (muo *MerchandiseUpdateOne) SetName(s string) *MerchandiseUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableName(s *string) *MerchandiseUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// ClearName clears the value of the "name" field.
func (muo *MerchandiseUpdateOne) ClearName() *MerchandiseUpdateOne {
	muo.mutation.ClearName()
	return muo
}

// SetDisplayName sets the "display_name" field.
func (muo *MerchandiseUpdateOne) SetDisplayName(s string) *MerchandiseUpdateOne {
	muo.mutation.SetDisplayName(s)
	return muo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableDisplayName(s *string) *MerchandiseUpdateOne {
	if s != nil {
		muo.SetDisplayName(*s)
	}
	return muo
}

// ClearDisplayName clears the value of the "display_name" field.
func (muo *MerchandiseUpdateOne) ClearDisplayName() *MerchandiseUpdateOne {
	muo.mutation.ClearDisplayName()
	return muo
}

// SetPictureURL sets the "picture_url" field.
func (muo *MerchandiseUpdateOne) SetPictureURL(s string) *MerchandiseUpdateOne {
	muo.mutation.SetPictureURL(s)
	return muo
}

// SetNillablePictureURL sets the "picture_url" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillablePictureURL(s *string) *MerchandiseUpdateOne {
	if s != nil {
		muo.SetPictureURL(*s)
	}
	return muo
}

// ClearPictureURL clears the value of the "picture_url" field.
func (muo *MerchandiseUpdateOne) ClearPictureURL() *MerchandiseUpdateOne {
	muo.mutation.ClearPictureURL()
	return muo
}

// SetDescription sets the "description" field.
func (muo *MerchandiseUpdateOne) SetDescription(s string) *MerchandiseUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableDescription(s *string) *MerchandiseUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// ClearDescription clears the value of the "description" field.
func (muo *MerchandiseUpdateOne) ClearDescription() *MerchandiseUpdateOne {
	muo.mutation.ClearDescription()
	return muo
}

// SetPrice sets the "price" field.
func (muo *MerchandiseUpdateOne) SetPrice(f float64) *MerchandiseUpdateOne {
	muo.mutation.ResetPrice()
	muo.mutation.SetPrice(f)
	return muo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillablePrice(f *float64) *MerchandiseUpdateOne {
	if f != nil {
		muo.SetPrice(*f)
	}
	return muo
}

// AddPrice adds f to the "price" field.
func (muo *MerchandiseUpdateOne) AddPrice(f float64) *MerchandiseUpdateOne {
	muo.mutation.AddPrice(f)
	return muo
}

// ClearPrice clears the value of the "price" field.
func (muo *MerchandiseUpdateOne) ClearPrice() *MerchandiseUpdateOne {
	muo.mutation.ClearPrice()
	return muo
}

// SetRefundableRatio sets the "refundable_ratio" field.
func (muo *MerchandiseUpdateOne) SetRefundableRatio(f float64) *MerchandiseUpdateOne {
	muo.mutation.ResetRefundableRatio()
	muo.mutation.SetRefundableRatio(f)
	return muo
}

// SetNillableRefundableRatio sets the "refundable_ratio" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableRefundableRatio(f *float64) *MerchandiseUpdateOne {
	if f != nil {
		muo.SetRefundableRatio(*f)
	}
	return muo
}

// AddRefundableRatio adds f to the "refundable_ratio" field.
func (muo *MerchandiseUpdateOne) AddRefundableRatio(f float64) *MerchandiseUpdateOne {
	muo.mutation.AddRefundableRatio(f)
	return muo
}

// ClearRefundableRatio clears the value of the "refundable_ratio" field.
func (muo *MerchandiseUpdateOne) ClearRefundableRatio() *MerchandiseUpdateOne {
	muo.mutation.ClearRefundableRatio()
	return muo
}

// SetDeductionRatio sets the "deduction_ratio" field.
func (muo *MerchandiseUpdateOne) SetDeductionRatio(f float64) *MerchandiseUpdateOne {
	muo.mutation.ResetDeductionRatio()
	muo.mutation.SetDeductionRatio(f)
	return muo
}

// SetNillableDeductionRatio sets the "deduction_ratio" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableDeductionRatio(f *float64) *MerchandiseUpdateOne {
	if f != nil {
		muo.SetDeductionRatio(*f)
	}
	return muo
}

// AddDeductionRatio adds f to the "deduction_ratio" field.
func (muo *MerchandiseUpdateOne) AddDeductionRatio(f float64) *MerchandiseUpdateOne {
	muo.mutation.AddDeductionRatio(f)
	return muo
}

// ClearDeductionRatio clears the value of the "deduction_ratio" field.
func (muo *MerchandiseUpdateOne) ClearDeductionRatio() *MerchandiseUpdateOne {
	muo.mutation.ClearDeductionRatio()
	return muo
}

// SetDeductionReason sets the "deduction_reason" field.
func (muo *MerchandiseUpdateOne) SetDeductionReason(s string) *MerchandiseUpdateOne {
	muo.mutation.SetDeductionReason(s)
	return muo
}

// SetNillableDeductionReason sets the "deduction_reason" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableDeductionReason(s *string) *MerchandiseUpdateOne {
	if s != nil {
		muo.SetDeductionReason(*s)
	}
	return muo
}

// ClearDeductionReason clears the value of the "deduction_reason" field.
func (muo *MerchandiseUpdateOne) ClearDeductionReason() *MerchandiseUpdateOne {
	muo.mutation.ClearDeductionReason()
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MerchandiseUpdateOne) SetCreatedAt(t time.Time) *MerchandiseUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableCreatedAt(t *time.Time) *MerchandiseUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MerchandiseUpdateOne) SetUpdatedAt(t time.Time) *MerchandiseUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MerchandiseUpdateOne) SetDeletedAt(t time.Time) *MerchandiseUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MerchandiseUpdateOne) SetNillableDeletedAt(t *time.Time) *MerchandiseUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MerchandiseUpdateOne) ClearDeletedAt() *MerchandiseUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// Mutation returns the MerchandiseMutation object of the builder.
func (muo *MerchandiseUpdateOne) Mutation() *MerchandiseMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MerchandiseUpdateOne) Select(field string, fields ...string) *MerchandiseUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Merchandise entity.
func (muo *MerchandiseUpdateOne) Save(ctx context.Context) (*Merchandise, error) {
	var (
		err  error
		node *Merchandise
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchandiseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MerchandiseUpdateOne) SaveX(ctx context.Context) *Merchandise {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MerchandiseUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MerchandiseUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MerchandiseUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := merchandise.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MerchandiseUpdateOne) check() error {
	if v, ok := muo.mutation.GetType(); ok {
		if err := merchandise.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Merchandise.type": %w`, err)}
		}
	}
	return nil
}

func (muo *MerchandiseUpdateOne) sqlSave(ctx context.Context) (_node *Merchandise, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchandise.Table,
			Columns: merchandise.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchandise.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Merchandise.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchandise.FieldID)
		for _, f := range fields {
			if !merchandise.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchandise.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: merchandise.FieldType,
		})
	}
	if value, ok := muo.mutation.TypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldTypeID,
		})
	}
	if muo.mutation.TypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldTypeID,
		})
	}
	if value, ok := muo.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldStock,
		})
	}
	if value, ok := muo.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldStock,
		})
	}
	if value, ok := muo.mutation.PendingPeriod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldPendingPeriod,
		})
	}
	if value, ok := muo.mutation.AddedPendingPeriod(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchandise.FieldPendingPeriod,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldName,
		})
	}
	if muo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldName,
		})
	}
	if value, ok := muo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldDisplayName,
		})
	}
	if muo.mutation.DisplayNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldDisplayName,
		})
	}
	if value, ok := muo.mutation.PictureURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldPictureURL,
		})
	}
	if muo.mutation.PictureURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldPictureURL,
		})
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldDescription,
		})
	}
	if muo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldDescription,
		})
	}
	if value, ok := muo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldPrice,
		})
	}
	if value, ok := muo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldPrice,
		})
	}
	if muo.mutation.PriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: merchandise.FieldPrice,
		})
	}
	if value, ok := muo.mutation.RefundableRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldRefundableRatio,
		})
	}
	if value, ok := muo.mutation.AddedRefundableRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldRefundableRatio,
		})
	}
	if muo.mutation.RefundableRatioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: merchandise.FieldRefundableRatio,
		})
	}
	if value, ok := muo.mutation.DeductionRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldDeductionRatio,
		})
	}
	if value, ok := muo.mutation.AddedDeductionRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchandise.FieldDeductionRatio,
		})
	}
	if muo.mutation.DeductionRatioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: merchandise.FieldDeductionRatio,
		})
	}
	if value, ok := muo.mutation.DeductionReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchandise.FieldDeductionReason,
		})
	}
	if muo.mutation.DeductionReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: merchandise.FieldDeductionReason,
		})
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchandise.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchandise.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchandise.FieldDeletedAt,
		})
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: merchandise.FieldDeletedAt,
		})
	}
	_node = &Merchandise{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchandise.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
