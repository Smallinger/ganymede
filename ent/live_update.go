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
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/predicate"
)

// LiveUpdate is the builder for updating Live entities.
type LiveUpdate struct {
	config
	hooks    []Hook
	mutation *LiveMutation
}

// Where appends a list predicates to the LiveUpdate builder.
func (lu *LiveUpdate) Where(ps ...predicate.Live) *LiveUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetWatchLive sets the "watch_live" field.
func (lu *LiveUpdate) SetWatchLive(b bool) *LiveUpdate {
	lu.mutation.SetWatchLive(b)
	return lu
}

// SetNillableWatchLive sets the "watch_live" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableWatchLive(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetWatchLive(*b)
	}
	return lu
}

// SetWatchVod sets the "watch_vod" field.
func (lu *LiveUpdate) SetWatchVod(b bool) *LiveUpdate {
	lu.mutation.SetWatchVod(b)
	return lu
}

// SetNillableWatchVod sets the "watch_vod" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableWatchVod(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetWatchVod(*b)
	}
	return lu
}

// SetDownloadArchives sets the "download_archives" field.
func (lu *LiveUpdate) SetDownloadArchives(b bool) *LiveUpdate {
	lu.mutation.SetDownloadArchives(b)
	return lu
}

// SetNillableDownloadArchives sets the "download_archives" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableDownloadArchives(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetDownloadArchives(*b)
	}
	return lu
}

// SetDownloadHighlights sets the "download_highlights" field.
func (lu *LiveUpdate) SetDownloadHighlights(b bool) *LiveUpdate {
	lu.mutation.SetDownloadHighlights(b)
	return lu
}

// SetNillableDownloadHighlights sets the "download_highlights" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableDownloadHighlights(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetDownloadHighlights(*b)
	}
	return lu
}

// SetDownloadUploads sets the "download_uploads" field.
func (lu *LiveUpdate) SetDownloadUploads(b bool) *LiveUpdate {
	lu.mutation.SetDownloadUploads(b)
	return lu
}

// SetNillableDownloadUploads sets the "download_uploads" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableDownloadUploads(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetDownloadUploads(*b)
	}
	return lu
}

// SetIsLive sets the "is_live" field.
func (lu *LiveUpdate) SetIsLive(b bool) *LiveUpdate {
	lu.mutation.SetIsLive(b)
	return lu
}

// SetNillableIsLive sets the "is_live" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableIsLive(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetIsLive(*b)
	}
	return lu
}

// SetArchiveChat sets the "archive_chat" field.
func (lu *LiveUpdate) SetArchiveChat(b bool) *LiveUpdate {
	lu.mutation.SetArchiveChat(b)
	return lu
}

// SetNillableArchiveChat sets the "archive_chat" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableArchiveChat(b *bool) *LiveUpdate {
	if b != nil {
		lu.SetArchiveChat(*b)
	}
	return lu
}

// SetResolution sets the "resolution" field.
func (lu *LiveUpdate) SetResolution(s string) *LiveUpdate {
	lu.mutation.SetResolution(s)
	return lu
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableResolution(s *string) *LiveUpdate {
	if s != nil {
		lu.SetResolution(*s)
	}
	return lu
}

// ClearResolution clears the value of the "resolution" field.
func (lu *LiveUpdate) ClearResolution() *LiveUpdate {
	lu.mutation.ClearResolution()
	return lu
}

// SetLastLive sets the "last_live" field.
func (lu *LiveUpdate) SetLastLive(t time.Time) *LiveUpdate {
	lu.mutation.SetLastLive(t)
	return lu
}

// SetNillableLastLive sets the "last_live" field if the given value is not nil.
func (lu *LiveUpdate) SetNillableLastLive(t *time.Time) *LiveUpdate {
	if t != nil {
		lu.SetLastLive(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LiveUpdate) SetUpdatedAt(t time.Time) *LiveUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetChannelID sets the "channel" edge to the Channel entity by ID.
func (lu *LiveUpdate) SetChannelID(id uuid.UUID) *LiveUpdate {
	lu.mutation.SetChannelID(id)
	return lu
}

// SetChannel sets the "channel" edge to the Channel entity.
func (lu *LiveUpdate) SetChannel(c *Channel) *LiveUpdate {
	return lu.SetChannelID(c.ID)
}

// Mutation returns the LiveMutation object of the builder.
func (lu *LiveUpdate) Mutation() *LiveMutation {
	return lu.mutation
}

// ClearChannel clears the "channel" edge to the Channel entity.
func (lu *LiveUpdate) ClearChannel() *LiveUpdate {
	lu.mutation.ClearChannel()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LiveUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lu.defaults()
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LiveMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LiveUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LiveUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LiveUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LiveUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := live.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LiveUpdate) check() error {
	if _, ok := lu.mutation.ChannelID(); lu.mutation.ChannelCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Live.channel"`)
	}
	return nil
}

func (lu *LiveUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   live.Table,
			Columns: live.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: live.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.WatchLive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldWatchLive,
		})
	}
	if value, ok := lu.mutation.WatchVod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldWatchVod,
		})
	}
	if value, ok := lu.mutation.DownloadArchives(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldDownloadArchives,
		})
	}
	if value, ok := lu.mutation.DownloadHighlights(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldDownloadHighlights,
		})
	}
	if value, ok := lu.mutation.DownloadUploads(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldDownloadUploads,
		})
	}
	if value, ok := lu.mutation.IsLive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldIsLive,
		})
	}
	if value, ok := lu.mutation.ArchiveChat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldArchiveChat,
		})
	}
	if value, ok := lu.mutation.Resolution(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: live.FieldResolution,
		})
	}
	if lu.mutation.ResolutionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: live.FieldResolution,
		})
	}
	if value, ok := lu.mutation.LastLive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: live.FieldLastLive,
		})
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: live.FieldUpdatedAt,
		})
	}
	if lu.mutation.ChannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   live.ChannelTable,
			Columns: []string{live.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: channel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   live.ChannelTable,
			Columns: []string{live.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{live.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LiveUpdateOne is the builder for updating a single Live entity.
type LiveUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LiveMutation
}

// SetWatchLive sets the "watch_live" field.
func (luo *LiveUpdateOne) SetWatchLive(b bool) *LiveUpdateOne {
	luo.mutation.SetWatchLive(b)
	return luo
}

// SetNillableWatchLive sets the "watch_live" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableWatchLive(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetWatchLive(*b)
	}
	return luo
}

// SetWatchVod sets the "watch_vod" field.
func (luo *LiveUpdateOne) SetWatchVod(b bool) *LiveUpdateOne {
	luo.mutation.SetWatchVod(b)
	return luo
}

// SetNillableWatchVod sets the "watch_vod" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableWatchVod(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetWatchVod(*b)
	}
	return luo
}

// SetDownloadArchives sets the "download_archives" field.
func (luo *LiveUpdateOne) SetDownloadArchives(b bool) *LiveUpdateOne {
	luo.mutation.SetDownloadArchives(b)
	return luo
}

// SetNillableDownloadArchives sets the "download_archives" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableDownloadArchives(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetDownloadArchives(*b)
	}
	return luo
}

// SetDownloadHighlights sets the "download_highlights" field.
func (luo *LiveUpdateOne) SetDownloadHighlights(b bool) *LiveUpdateOne {
	luo.mutation.SetDownloadHighlights(b)
	return luo
}

// SetNillableDownloadHighlights sets the "download_highlights" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableDownloadHighlights(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetDownloadHighlights(*b)
	}
	return luo
}

// SetDownloadUploads sets the "download_uploads" field.
func (luo *LiveUpdateOne) SetDownloadUploads(b bool) *LiveUpdateOne {
	luo.mutation.SetDownloadUploads(b)
	return luo
}

// SetNillableDownloadUploads sets the "download_uploads" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableDownloadUploads(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetDownloadUploads(*b)
	}
	return luo
}

// SetIsLive sets the "is_live" field.
func (luo *LiveUpdateOne) SetIsLive(b bool) *LiveUpdateOne {
	luo.mutation.SetIsLive(b)
	return luo
}

// SetNillableIsLive sets the "is_live" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableIsLive(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetIsLive(*b)
	}
	return luo
}

// SetArchiveChat sets the "archive_chat" field.
func (luo *LiveUpdateOne) SetArchiveChat(b bool) *LiveUpdateOne {
	luo.mutation.SetArchiveChat(b)
	return luo
}

// SetNillableArchiveChat sets the "archive_chat" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableArchiveChat(b *bool) *LiveUpdateOne {
	if b != nil {
		luo.SetArchiveChat(*b)
	}
	return luo
}

// SetResolution sets the "resolution" field.
func (luo *LiveUpdateOne) SetResolution(s string) *LiveUpdateOne {
	luo.mutation.SetResolution(s)
	return luo
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableResolution(s *string) *LiveUpdateOne {
	if s != nil {
		luo.SetResolution(*s)
	}
	return luo
}

// ClearResolution clears the value of the "resolution" field.
func (luo *LiveUpdateOne) ClearResolution() *LiveUpdateOne {
	luo.mutation.ClearResolution()
	return luo
}

// SetLastLive sets the "last_live" field.
func (luo *LiveUpdateOne) SetLastLive(t time.Time) *LiveUpdateOne {
	luo.mutation.SetLastLive(t)
	return luo
}

// SetNillableLastLive sets the "last_live" field if the given value is not nil.
func (luo *LiveUpdateOne) SetNillableLastLive(t *time.Time) *LiveUpdateOne {
	if t != nil {
		luo.SetLastLive(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LiveUpdateOne) SetUpdatedAt(t time.Time) *LiveUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetChannelID sets the "channel" edge to the Channel entity by ID.
func (luo *LiveUpdateOne) SetChannelID(id uuid.UUID) *LiveUpdateOne {
	luo.mutation.SetChannelID(id)
	return luo
}

// SetChannel sets the "channel" edge to the Channel entity.
func (luo *LiveUpdateOne) SetChannel(c *Channel) *LiveUpdateOne {
	return luo.SetChannelID(c.ID)
}

// Mutation returns the LiveMutation object of the builder.
func (luo *LiveUpdateOne) Mutation() *LiveMutation {
	return luo.mutation
}

// ClearChannel clears the "channel" edge to the Channel entity.
func (luo *LiveUpdateOne) ClearChannel() *LiveUpdateOne {
	luo.mutation.ClearChannel()
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LiveUpdateOne) Select(field string, fields ...string) *LiveUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Live entity.
func (luo *LiveUpdateOne) Save(ctx context.Context) (*Live, error) {
	var (
		err  error
		node *Live
	)
	luo.defaults()
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LiveMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Live)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LiveMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LiveUpdateOne) SaveX(ctx context.Context) *Live {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LiveUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LiveUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LiveUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := live.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LiveUpdateOne) check() error {
	if _, ok := luo.mutation.ChannelID(); luo.mutation.ChannelCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Live.channel"`)
	}
	return nil
}

func (luo *LiveUpdateOne) sqlSave(ctx context.Context) (_node *Live, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   live.Table,
			Columns: live.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: live.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Live.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, live.FieldID)
		for _, f := range fields {
			if !live.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != live.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.WatchLive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldWatchLive,
		})
	}
	if value, ok := luo.mutation.WatchVod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldWatchVod,
		})
	}
	if value, ok := luo.mutation.DownloadArchives(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldDownloadArchives,
		})
	}
	if value, ok := luo.mutation.DownloadHighlights(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldDownloadHighlights,
		})
	}
	if value, ok := luo.mutation.DownloadUploads(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldDownloadUploads,
		})
	}
	if value, ok := luo.mutation.IsLive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldIsLive,
		})
	}
	if value, ok := luo.mutation.ArchiveChat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: live.FieldArchiveChat,
		})
	}
	if value, ok := luo.mutation.Resolution(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: live.FieldResolution,
		})
	}
	if luo.mutation.ResolutionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: live.FieldResolution,
		})
	}
	if value, ok := luo.mutation.LastLive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: live.FieldLastLive,
		})
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: live.FieldUpdatedAt,
		})
	}
	if luo.mutation.ChannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   live.ChannelTable,
			Columns: []string{live.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: channel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   live.ChannelTable,
			Columns: []string{live.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Live{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{live.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
