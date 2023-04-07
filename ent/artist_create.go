// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/artist"
	"github.com/scarlet0725/prism-api/ent/event"
)

// ArtistCreate is the builder for creating a Artist entity.
type ArtistCreate struct {
	config
	mutation *ArtistMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetArtistID sets the "artist_id" field.
func (ac *ArtistCreate) SetArtistID(s string) *ArtistCreate {
	ac.mutation.SetArtistID(s)
	return ac
}

// SetName sets the "name" field.
func (ac *ArtistCreate) SetName(s string) *ArtistCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetURL sets the "url" field.
func (ac *ArtistCreate) SetURL(s string) *ArtistCreate {
	ac.mutation.SetURL(s)
	return ac
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ac *ArtistCreate) SetNillableURL(s *string) *ArtistCreate {
	if s != nil {
		ac.SetURL(*s)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArtistCreate) SetCreatedAt(t time.Time) *ArtistCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArtistCreate) SetNillableCreatedAt(t *time.Time) *ArtistCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArtistCreate) SetUpdatedAt(t time.Time) *ArtistCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArtistCreate) SetNillableUpdatedAt(t *time.Time) *ArtistCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *ArtistCreate) SetDeletedAt(t time.Time) *ArtistCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *ArtistCreate) SetNillableDeletedAt(t *time.Time) *ArtistCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (ac *ArtistCreate) AddEventIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddEventIDs(ids...)
	return ac
}

// AddEvents adds the "events" edges to the Event entity.
func (ac *ArtistCreate) AddEvents(e ...*Event) *ArtistCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ac.AddEventIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (ac *ArtistCreate) Mutation() *ArtistMutation {
	return ac.mutation
}

// Save creates the Artist in the database.
func (ac *ArtistCreate) Save(ctx context.Context) (*Artist, error) {
	ac.defaults()
	return withHooks[*Artist, ArtistMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtistCreate) SaveX(ctx context.Context) *Artist {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtistCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtistCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArtistCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := artist.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := artist.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtistCreate) check() error {
	if _, ok := ac.mutation.ArtistID(); !ok {
		return &ValidationError{Name: "artist_id", err: errors.New(`ent: missing required field "Artist.artist_id"`)}
	}
	if v, ok := ac.mutation.ArtistID(); ok {
		if err := artist.ArtistIDValidator(v); err != nil {
			return &ValidationError{Name: "artist_id", err: fmt.Errorf(`ent: validator failed for field "Artist.artist_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Artist.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := artist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artist.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Artist.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Artist.updated_at"`)}
	}
	return nil
}

func (ac *ArtistCreate) sqlSave(ctx context.Context) (*Artist, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArtistCreate) createSpec() (*Artist, *sqlgraph.CreateSpec) {
	var (
		_node = &Artist{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(artist.Table, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.ArtistID(); ok {
		_spec.SetField(artist.FieldArtistID, field.TypeString, value)
		_node.ArtistID = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.URL(); ok {
		_spec.SetField(artist.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(artist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(artist.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(artist.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := ac.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.EventsTable,
			Columns: artist.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Artist.Create().
//		SetArtistID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtistUpsert) {
//			SetArtistID(v+v).
//		}).
//		Exec(ctx)
func (ac *ArtistCreate) OnConflict(opts ...sql.ConflictOption) *ArtistUpsertOne {
	ac.conflict = opts
	return &ArtistUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *ArtistCreate) OnConflictColumns(columns ...string) *ArtistUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ArtistUpsertOne{
		create: ac,
	}
}

type (
	// ArtistUpsertOne is the builder for "upsert"-ing
	//  one Artist node.
	ArtistUpsertOne struct {
		create *ArtistCreate
	}

	// ArtistUpsert is the "OnConflict" setter.
	ArtistUpsert struct {
		*sql.UpdateSet
	}
)

// SetArtistID sets the "artist_id" field.
func (u *ArtistUpsert) SetArtistID(v string) *ArtistUpsert {
	u.Set(artist.FieldArtistID, v)
	return u
}

// UpdateArtistID sets the "artist_id" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateArtistID() *ArtistUpsert {
	u.SetExcluded(artist.FieldArtistID)
	return u
}

// SetName sets the "name" field.
func (u *ArtistUpsert) SetName(v string) *ArtistUpsert {
	u.Set(artist.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateName() *ArtistUpsert {
	u.SetExcluded(artist.FieldName)
	return u
}

// SetURL sets the "url" field.
func (u *ArtistUpsert) SetURL(v string) *ArtistUpsert {
	u.Set(artist.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateURL() *ArtistUpsert {
	u.SetExcluded(artist.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *ArtistUpsert) ClearURL() *ArtistUpsert {
	u.SetNull(artist.FieldURL)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArtistUpsert) SetUpdatedAt(v time.Time) *ArtistUpsert {
	u.Set(artist.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateUpdatedAt() *ArtistUpsert {
	u.SetExcluded(artist.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ArtistUpsert) SetDeletedAt(v time.Time) *ArtistUpsert {
	u.Set(artist.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateDeletedAt() *ArtistUpsert {
	u.SetExcluded(artist.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ArtistUpsert) ClearDeletedAt() *ArtistUpsert {
	u.SetNull(artist.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ArtistUpsertOne) UpdateNewValues() *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(artist.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Artist.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ArtistUpsertOne) Ignore() *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtistUpsertOne) DoNothing() *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtistCreate.OnConflict
// documentation for more info.
func (u *ArtistUpsertOne) Update(set func(*ArtistUpsert)) *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtistUpsert{UpdateSet: update})
	}))
	return u
}

// SetArtistID sets the "artist_id" field.
func (u *ArtistUpsertOne) SetArtistID(v string) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetArtistID(v)
	})
}

// UpdateArtistID sets the "artist_id" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateArtistID() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateArtistID()
	})
}

// SetName sets the "name" field.
func (u *ArtistUpsertOne) SetName(v string) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateName() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateName()
	})
}

// SetURL sets the "url" field.
func (u *ArtistUpsertOne) SetURL(v string) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateURL() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *ArtistUpsertOne) ClearURL() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.ClearURL()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArtistUpsertOne) SetUpdatedAt(v time.Time) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateUpdatedAt() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ArtistUpsertOne) SetDeletedAt(v time.Time) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateDeletedAt() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ArtistUpsertOne) ClearDeletedAt() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ArtistUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArtistCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtistUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArtistUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArtistUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArtistCreateBulk is the builder for creating many Artist entities in bulk.
type ArtistCreateBulk struct {
	config
	builders []*ArtistCreate
	conflict []sql.ConflictOption
}

// Save creates the Artist entities in the database.
func (acb *ArtistCreateBulk) Save(ctx context.Context) ([]*Artist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artist, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtistCreateBulk) SaveX(ctx context.Context) []*Artist {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtistCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtistCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Artist.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtistUpsert) {
//			SetArtistID(v+v).
//		}).
//		Exec(ctx)
func (acb *ArtistCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArtistUpsertBulk {
	acb.conflict = opts
	return &ArtistUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *ArtistCreateBulk) OnConflictColumns(columns ...string) *ArtistUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ArtistUpsertBulk{
		create: acb,
	}
}

// ArtistUpsertBulk is the builder for "upsert"-ing
// a bulk of Artist nodes.
type ArtistUpsertBulk struct {
	create *ArtistCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ArtistUpsertBulk) UpdateNewValues() *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(artist.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ArtistUpsertBulk) Ignore() *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtistUpsertBulk) DoNothing() *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtistCreateBulk.OnConflict
// documentation for more info.
func (u *ArtistUpsertBulk) Update(set func(*ArtistUpsert)) *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtistUpsert{UpdateSet: update})
	}))
	return u
}

// SetArtistID sets the "artist_id" field.
func (u *ArtistUpsertBulk) SetArtistID(v string) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetArtistID(v)
	})
}

// UpdateArtistID sets the "artist_id" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateArtistID() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateArtistID()
	})
}

// SetName sets the "name" field.
func (u *ArtistUpsertBulk) SetName(v string) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateName() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateName()
	})
}

// SetURL sets the "url" field.
func (u *ArtistUpsertBulk) SetURL(v string) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateURL() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *ArtistUpsertBulk) ClearURL() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.ClearURL()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArtistUpsertBulk) SetUpdatedAt(v time.Time) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateUpdatedAt() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ArtistUpsertBulk) SetDeletedAt(v time.Time) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateDeletedAt() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ArtistUpsertBulk) ClearDeletedAt() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ArtistUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArtistCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArtistCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtistUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
