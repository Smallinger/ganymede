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
	"github.com/zibbp/ganymede/ent/playlist"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/vod"
)

// PlaylistUpdate is the builder for updating Playlist entities.
type PlaylistUpdate struct {
	config
	hooks    []Hook
	mutation *PlaylistMutation
}

// Where appends a list predicates to the PlaylistUpdate builder.
func (pu *PlaylistUpdate) Where(ps ...predicate.Playlist) *PlaylistUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlaylistUpdate) SetName(s string) *PlaylistUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PlaylistUpdate) SetDescription(s string) *PlaylistUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PlaylistUpdate) SetNillableDescription(s *string) *PlaylistUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PlaylistUpdate) ClearDescription() *PlaylistUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (pu *PlaylistUpdate) SetThumbnailPath(s string) *PlaylistUpdate {
	pu.mutation.SetThumbnailPath(s)
	return pu
}

// SetNillableThumbnailPath sets the "thumbnail_path" field if the given value is not nil.
func (pu *PlaylistUpdate) SetNillableThumbnailPath(s *string) *PlaylistUpdate {
	if s != nil {
		pu.SetThumbnailPath(*s)
	}
	return pu
}

// ClearThumbnailPath clears the value of the "thumbnail_path" field.
func (pu *PlaylistUpdate) ClearThumbnailPath() *PlaylistUpdate {
	pu.mutation.ClearThumbnailPath()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PlaylistUpdate) SetUpdatedAt(t time.Time) *PlaylistUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (pu *PlaylistUpdate) AddVodIDs(ids ...uuid.UUID) *PlaylistUpdate {
	pu.mutation.AddVodIDs(ids...)
	return pu
}

// AddVods adds the "vods" edges to the Vod entity.
func (pu *PlaylistUpdate) AddVods(v ...*Vod) *PlaylistUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pu.AddVodIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pu *PlaylistUpdate) Mutation() *PlaylistMutation {
	return pu.mutation
}

// ClearVods clears all "vods" edges to the Vod entity.
func (pu *PlaylistUpdate) ClearVods() *PlaylistUpdate {
	pu.mutation.ClearVods()
	return pu
}

// RemoveVodIDs removes the "vods" edge to Vod entities by IDs.
func (pu *PlaylistUpdate) RemoveVodIDs(ids ...uuid.UUID) *PlaylistUpdate {
	pu.mutation.RemoveVodIDs(ids...)
	return pu
}

// RemoveVods removes "vods" edges to Vod entities.
func (pu *PlaylistUpdate) RemoveVods(v ...*Vod) *PlaylistUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pu.RemoveVodIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlaylistUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlaylistUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlaylistUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlaylistUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PlaylistUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := playlist.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PlaylistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlist.Table,
			Columns: playlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: playlist.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldName,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldDescription,
		})
	}
	if pu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playlist.FieldDescription,
		})
	}
	if value, ok := pu.mutation.ThumbnailPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldThumbnailPath,
		})
	}
	if pu.mutation.ThumbnailPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playlist.FieldThumbnailPath,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playlist.FieldUpdatedAt,
		})
	}
	if pu.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedVodsIDs(); len(nodes) > 0 && !pu.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlaylistUpdateOne is the builder for updating a single Playlist entity.
type PlaylistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlaylistMutation
}

// SetName sets the "name" field.
func (puo *PlaylistUpdateOne) SetName(s string) *PlaylistUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PlaylistUpdateOne) SetDescription(s string) *PlaylistUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PlaylistUpdateOne) SetNillableDescription(s *string) *PlaylistUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PlaylistUpdateOne) ClearDescription() *PlaylistUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (puo *PlaylistUpdateOne) SetThumbnailPath(s string) *PlaylistUpdateOne {
	puo.mutation.SetThumbnailPath(s)
	return puo
}

// SetNillableThumbnailPath sets the "thumbnail_path" field if the given value is not nil.
func (puo *PlaylistUpdateOne) SetNillableThumbnailPath(s *string) *PlaylistUpdateOne {
	if s != nil {
		puo.SetThumbnailPath(*s)
	}
	return puo
}

// ClearThumbnailPath clears the value of the "thumbnail_path" field.
func (puo *PlaylistUpdateOne) ClearThumbnailPath() *PlaylistUpdateOne {
	puo.mutation.ClearThumbnailPath()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PlaylistUpdateOne) SetUpdatedAt(t time.Time) *PlaylistUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (puo *PlaylistUpdateOne) AddVodIDs(ids ...uuid.UUID) *PlaylistUpdateOne {
	puo.mutation.AddVodIDs(ids...)
	return puo
}

// AddVods adds the "vods" edges to the Vod entity.
func (puo *PlaylistUpdateOne) AddVods(v ...*Vod) *PlaylistUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return puo.AddVodIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (puo *PlaylistUpdateOne) Mutation() *PlaylistMutation {
	return puo.mutation
}

// ClearVods clears all "vods" edges to the Vod entity.
func (puo *PlaylistUpdateOne) ClearVods() *PlaylistUpdateOne {
	puo.mutation.ClearVods()
	return puo
}

// RemoveVodIDs removes the "vods" edge to Vod entities by IDs.
func (puo *PlaylistUpdateOne) RemoveVodIDs(ids ...uuid.UUID) *PlaylistUpdateOne {
	puo.mutation.RemoveVodIDs(ids...)
	return puo
}

// RemoveVods removes "vods" edges to Vod entities.
func (puo *PlaylistUpdateOne) RemoveVods(v ...*Vod) *PlaylistUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return puo.RemoveVodIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlaylistUpdateOne) Select(field string, fields ...string) *PlaylistUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Playlist entity.
func (puo *PlaylistUpdateOne) Save(ctx context.Context) (*Playlist, error) {
	var (
		err  error
		node *Playlist
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Playlist)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlaylistMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlaylistUpdateOne) SaveX(ctx context.Context) *Playlist {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlaylistUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlaylistUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PlaylistUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := playlist.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PlaylistUpdateOne) sqlSave(ctx context.Context) (_node *Playlist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlist.Table,
			Columns: playlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: playlist.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Playlist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playlist.FieldID)
		for _, f := range fields {
			if !playlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != playlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldName,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldDescription,
		})
	}
	if puo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playlist.FieldDescription,
		})
	}
	if value, ok := puo.mutation.ThumbnailPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldThumbnailPath,
		})
	}
	if puo.mutation.ThumbnailPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playlist.FieldThumbnailPath,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playlist.FieldUpdatedAt,
		})
	}
	if puo.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedVodsIDs(); len(nodes) > 0 && !puo.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Playlist{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}