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
	"github.com/scarlet0725/prism-api/ent/googleoauthtoken"
	"github.com/scarlet0725/prism-api/ent/user"
)

// GoogleOauthTokenCreate is the builder for creating a GoogleOauthToken entity.
type GoogleOauthTokenCreate struct {
	config
	mutation *GoogleOauthTokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRefreshToken sets the "refresh_token" field.
func (gotc *GoogleOauthTokenCreate) SetRefreshToken(s string) *GoogleOauthTokenCreate {
	gotc.mutation.SetRefreshToken(s)
	return gotc
}

// SetAccessToken sets the "access_token" field.
func (gotc *GoogleOauthTokenCreate) SetAccessToken(s string) *GoogleOauthTokenCreate {
	gotc.mutation.SetAccessToken(s)
	return gotc
}

// SetExpiry sets the "expiry" field.
func (gotc *GoogleOauthTokenCreate) SetExpiry(t time.Time) *GoogleOauthTokenCreate {
	gotc.mutation.SetExpiry(t)
	return gotc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gotc *GoogleOauthTokenCreate) SetUserID(id int) *GoogleOauthTokenCreate {
	gotc.mutation.SetUserID(id)
	return gotc
}

// SetUser sets the "user" edge to the User entity.
func (gotc *GoogleOauthTokenCreate) SetUser(u *User) *GoogleOauthTokenCreate {
	return gotc.SetUserID(u.ID)
}

// Mutation returns the GoogleOauthTokenMutation object of the builder.
func (gotc *GoogleOauthTokenCreate) Mutation() *GoogleOauthTokenMutation {
	return gotc.mutation
}

// Save creates the GoogleOauthToken in the database.
func (gotc *GoogleOauthTokenCreate) Save(ctx context.Context) (*GoogleOauthToken, error) {
	return withHooks[*GoogleOauthToken, GoogleOauthTokenMutation](ctx, gotc.sqlSave, gotc.mutation, gotc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gotc *GoogleOauthTokenCreate) SaveX(ctx context.Context) *GoogleOauthToken {
	v, err := gotc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gotc *GoogleOauthTokenCreate) Exec(ctx context.Context) error {
	_, err := gotc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gotc *GoogleOauthTokenCreate) ExecX(ctx context.Context) {
	if err := gotc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gotc *GoogleOauthTokenCreate) check() error {
	if _, ok := gotc.mutation.RefreshToken(); !ok {
		return &ValidationError{Name: "refresh_token", err: errors.New(`ent: missing required field "GoogleOauthToken.refresh_token"`)}
	}
	if _, ok := gotc.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New(`ent: missing required field "GoogleOauthToken.access_token"`)}
	}
	if _, ok := gotc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`ent: missing required field "GoogleOauthToken.expiry"`)}
	}
	if _, ok := gotc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GoogleOauthToken.user"`)}
	}
	return nil
}

func (gotc *GoogleOauthTokenCreate) sqlSave(ctx context.Context) (*GoogleOauthToken, error) {
	if err := gotc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gotc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gotc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gotc.mutation.id = &_node.ID
	gotc.mutation.done = true
	return _node, nil
}

func (gotc *GoogleOauthTokenCreate) createSpec() (*GoogleOauthToken, *sqlgraph.CreateSpec) {
	var (
		_node = &GoogleOauthToken{config: gotc.config}
		_spec = sqlgraph.NewCreateSpec(googleoauthtoken.Table, sqlgraph.NewFieldSpec(googleoauthtoken.FieldID, field.TypeInt))
	)
	_spec.OnConflict = gotc.conflict
	if value, ok := gotc.mutation.RefreshToken(); ok {
		_spec.SetField(googleoauthtoken.FieldRefreshToken, field.TypeString, value)
		_node.RefreshToken = value
	}
	if value, ok := gotc.mutation.AccessToken(); ok {
		_spec.SetField(googleoauthtoken.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := gotc.mutation.Expiry(); ok {
		_spec.SetField(googleoauthtoken.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if nodes := gotc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   googleoauthtoken.UserTable,
			Columns: []string{googleoauthtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoogleOauthToken.Create().
//		SetRefreshToken(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoogleOauthTokenUpsert) {
//			SetRefreshToken(v+v).
//		}).
//		Exec(ctx)
func (gotc *GoogleOauthTokenCreate) OnConflict(opts ...sql.ConflictOption) *GoogleOauthTokenUpsertOne {
	gotc.conflict = opts
	return &GoogleOauthTokenUpsertOne{
		create: gotc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoogleOauthToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gotc *GoogleOauthTokenCreate) OnConflictColumns(columns ...string) *GoogleOauthTokenUpsertOne {
	gotc.conflict = append(gotc.conflict, sql.ConflictColumns(columns...))
	return &GoogleOauthTokenUpsertOne{
		create: gotc,
	}
}

type (
	// GoogleOauthTokenUpsertOne is the builder for "upsert"-ing
	//  one GoogleOauthToken node.
	GoogleOauthTokenUpsertOne struct {
		create *GoogleOauthTokenCreate
	}

	// GoogleOauthTokenUpsert is the "OnConflict" setter.
	GoogleOauthTokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetRefreshToken sets the "refresh_token" field.
func (u *GoogleOauthTokenUpsert) SetRefreshToken(v string) *GoogleOauthTokenUpsert {
	u.Set(googleoauthtoken.FieldRefreshToken, v)
	return u
}

// UpdateRefreshToken sets the "refresh_token" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsert) UpdateRefreshToken() *GoogleOauthTokenUpsert {
	u.SetExcluded(googleoauthtoken.FieldRefreshToken)
	return u
}

// SetAccessToken sets the "access_token" field.
func (u *GoogleOauthTokenUpsert) SetAccessToken(v string) *GoogleOauthTokenUpsert {
	u.Set(googleoauthtoken.FieldAccessToken, v)
	return u
}

// UpdateAccessToken sets the "access_token" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsert) UpdateAccessToken() *GoogleOauthTokenUpsert {
	u.SetExcluded(googleoauthtoken.FieldAccessToken)
	return u
}

// SetExpiry sets the "expiry" field.
func (u *GoogleOauthTokenUpsert) SetExpiry(v time.Time) *GoogleOauthTokenUpsert {
	u.Set(googleoauthtoken.FieldExpiry, v)
	return u
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsert) UpdateExpiry() *GoogleOauthTokenUpsert {
	u.SetExcluded(googleoauthtoken.FieldExpiry)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GoogleOauthToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GoogleOauthTokenUpsertOne) UpdateNewValues() *GoogleOauthTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoogleOauthToken.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GoogleOauthTokenUpsertOne) Ignore() *GoogleOauthTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoogleOauthTokenUpsertOne) DoNothing() *GoogleOauthTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoogleOauthTokenCreate.OnConflict
// documentation for more info.
func (u *GoogleOauthTokenUpsertOne) Update(set func(*GoogleOauthTokenUpsert)) *GoogleOauthTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoogleOauthTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetRefreshToken sets the "refresh_token" field.
func (u *GoogleOauthTokenUpsertOne) SetRefreshToken(v string) *GoogleOauthTokenUpsertOne {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.SetRefreshToken(v)
	})
}

// UpdateRefreshToken sets the "refresh_token" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsertOne) UpdateRefreshToken() *GoogleOauthTokenUpsertOne {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.UpdateRefreshToken()
	})
}

// SetAccessToken sets the "access_token" field.
func (u *GoogleOauthTokenUpsertOne) SetAccessToken(v string) *GoogleOauthTokenUpsertOne {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.SetAccessToken(v)
	})
}

// UpdateAccessToken sets the "access_token" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsertOne) UpdateAccessToken() *GoogleOauthTokenUpsertOne {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.UpdateAccessToken()
	})
}

// SetExpiry sets the "expiry" field.
func (u *GoogleOauthTokenUpsertOne) SetExpiry(v time.Time) *GoogleOauthTokenUpsertOne {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsertOne) UpdateExpiry() *GoogleOauthTokenUpsertOne {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.UpdateExpiry()
	})
}

// Exec executes the query.
func (u *GoogleOauthTokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoogleOauthTokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoogleOauthTokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GoogleOauthTokenUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GoogleOauthTokenUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GoogleOauthTokenCreateBulk is the builder for creating many GoogleOauthToken entities in bulk.
type GoogleOauthTokenCreateBulk struct {
	config
	builders []*GoogleOauthTokenCreate
	conflict []sql.ConflictOption
}

// Save creates the GoogleOauthToken entities in the database.
func (gotcb *GoogleOauthTokenCreateBulk) Save(ctx context.Context) ([]*GoogleOauthToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gotcb.builders))
	nodes := make([]*GoogleOauthToken, len(gotcb.builders))
	mutators := make([]Mutator, len(gotcb.builders))
	for i := range gotcb.builders {
		func(i int, root context.Context) {
			builder := gotcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoogleOauthTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, gotcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gotcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gotcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gotcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gotcb *GoogleOauthTokenCreateBulk) SaveX(ctx context.Context) []*GoogleOauthToken {
	v, err := gotcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gotcb *GoogleOauthTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := gotcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gotcb *GoogleOauthTokenCreateBulk) ExecX(ctx context.Context) {
	if err := gotcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoogleOauthToken.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoogleOauthTokenUpsert) {
//			SetRefreshToken(v+v).
//		}).
//		Exec(ctx)
func (gotcb *GoogleOauthTokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *GoogleOauthTokenUpsertBulk {
	gotcb.conflict = opts
	return &GoogleOauthTokenUpsertBulk{
		create: gotcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoogleOauthToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gotcb *GoogleOauthTokenCreateBulk) OnConflictColumns(columns ...string) *GoogleOauthTokenUpsertBulk {
	gotcb.conflict = append(gotcb.conflict, sql.ConflictColumns(columns...))
	return &GoogleOauthTokenUpsertBulk{
		create: gotcb,
	}
}

// GoogleOauthTokenUpsertBulk is the builder for "upsert"-ing
// a bulk of GoogleOauthToken nodes.
type GoogleOauthTokenUpsertBulk struct {
	create *GoogleOauthTokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GoogleOauthToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GoogleOauthTokenUpsertBulk) UpdateNewValues() *GoogleOauthTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoogleOauthToken.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GoogleOauthTokenUpsertBulk) Ignore() *GoogleOauthTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoogleOauthTokenUpsertBulk) DoNothing() *GoogleOauthTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoogleOauthTokenCreateBulk.OnConflict
// documentation for more info.
func (u *GoogleOauthTokenUpsertBulk) Update(set func(*GoogleOauthTokenUpsert)) *GoogleOauthTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoogleOauthTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetRefreshToken sets the "refresh_token" field.
func (u *GoogleOauthTokenUpsertBulk) SetRefreshToken(v string) *GoogleOauthTokenUpsertBulk {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.SetRefreshToken(v)
	})
}

// UpdateRefreshToken sets the "refresh_token" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsertBulk) UpdateRefreshToken() *GoogleOauthTokenUpsertBulk {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.UpdateRefreshToken()
	})
}

// SetAccessToken sets the "access_token" field.
func (u *GoogleOauthTokenUpsertBulk) SetAccessToken(v string) *GoogleOauthTokenUpsertBulk {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.SetAccessToken(v)
	})
}

// UpdateAccessToken sets the "access_token" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsertBulk) UpdateAccessToken() *GoogleOauthTokenUpsertBulk {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.UpdateAccessToken()
	})
}

// SetExpiry sets the "expiry" field.
func (u *GoogleOauthTokenUpsertBulk) SetExpiry(v time.Time) *GoogleOauthTokenUpsertBulk {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.SetExpiry(v)
	})
}

// UpdateExpiry sets the "expiry" field to the value that was provided on create.
func (u *GoogleOauthTokenUpsertBulk) UpdateExpiry() *GoogleOauthTokenUpsertBulk {
	return u.Update(func(s *GoogleOauthTokenUpsert) {
		s.UpdateExpiry()
	})
}

// Exec executes the query.
func (u *GoogleOauthTokenUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GoogleOauthTokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoogleOauthTokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoogleOauthTokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
