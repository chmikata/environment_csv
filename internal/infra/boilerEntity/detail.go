// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boilerEntity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Detail is an object representing the database table.
type Detail struct {
	ID      string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Detail1 null.String `boil:"detail1" json:"detail1,omitempty" toml:"detail1" yaml:"detail1,omitempty"`
	Detail2 null.String `boil:"detail2" json:"detail2,omitempty" toml:"detail2" yaml:"detail2,omitempty"`
	Detail3 null.String `boil:"detail3" json:"detail3,omitempty" toml:"detail3" yaml:"detail3,omitempty"`

	R *detailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L detailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DetailColumns = struct {
	ID      string
	Detail1 string
	Detail2 string
	Detail3 string
}{
	ID:      "id",
	Detail1: "detail1",
	Detail2: "detail2",
	Detail3: "detail3",
}

var DetailTableColumns = struct {
	ID      string
	Detail1 string
	Detail2 string
	Detail3 string
}{
	ID:      "detail.id",
	Detail1: "detail.detail1",
	Detail2: "detail.detail2",
	Detail3: "detail.detail3",
}

// Generated where

var DetailWhere = struct {
	ID      whereHelperstring
	Detail1 whereHelpernull_String
	Detail2 whereHelpernull_String
	Detail3 whereHelpernull_String
}{
	ID:      whereHelperstring{field: "\"detail\".\"id\""},
	Detail1: whereHelpernull_String{field: "\"detail\".\"detail1\""},
	Detail2: whereHelpernull_String{field: "\"detail\".\"detail2\""},
	Detail3: whereHelpernull_String{field: "\"detail\".\"detail3\""},
}

// DetailRels is where relationship names are stored.
var DetailRels = struct {
}{}

// detailR is where relationships are stored.
type detailR struct {
}

// NewStruct creates a new relationship struct
func (*detailR) NewStruct() *detailR {
	return &detailR{}
}

// detailL is where Load methods for each relationship are stored.
type detailL struct{}

var (
	detailAllColumns            = []string{"id", "detail1", "detail2", "detail3"}
	detailColumnsWithoutDefault = []string{"id"}
	detailColumnsWithDefault    = []string{"detail1", "detail2", "detail3"}
	detailPrimaryKeyColumns     = []string{"id"}
	detailGeneratedColumns      = []string{}
)

type (
	// DetailSlice is an alias for a slice of pointers to Detail.
	// This should almost always be used instead of []Detail.
	DetailSlice []*Detail
	// DetailHook is the signature for custom Detail hook methods
	DetailHook func(context.Context, boil.ContextExecutor, *Detail) error

	detailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	detailType                 = reflect.TypeOf(&Detail{})
	detailMapping              = queries.MakeStructMapping(detailType)
	detailPrimaryKeyMapping, _ = queries.BindMapping(detailType, detailMapping, detailPrimaryKeyColumns)
	detailInsertCacheMut       sync.RWMutex
	detailInsertCache          = make(map[string]insertCache)
	detailUpdateCacheMut       sync.RWMutex
	detailUpdateCache          = make(map[string]updateCache)
	detailUpsertCacheMut       sync.RWMutex
	detailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var detailAfterSelectHooks []DetailHook

var detailBeforeInsertHooks []DetailHook
var detailAfterInsertHooks []DetailHook

var detailBeforeUpdateHooks []DetailHook
var detailAfterUpdateHooks []DetailHook

var detailBeforeDeleteHooks []DetailHook
var detailAfterDeleteHooks []DetailHook

var detailBeforeUpsertHooks []DetailHook
var detailAfterUpsertHooks []DetailHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Detail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Detail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Detail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Detail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Detail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Detail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Detail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Detail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Detail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range detailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDetailHook registers your hook function for all future operations.
func AddDetailHook(hookPoint boil.HookPoint, detailHook DetailHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		detailAfterSelectHooks = append(detailAfterSelectHooks, detailHook)
	case boil.BeforeInsertHook:
		detailBeforeInsertHooks = append(detailBeforeInsertHooks, detailHook)
	case boil.AfterInsertHook:
		detailAfterInsertHooks = append(detailAfterInsertHooks, detailHook)
	case boil.BeforeUpdateHook:
		detailBeforeUpdateHooks = append(detailBeforeUpdateHooks, detailHook)
	case boil.AfterUpdateHook:
		detailAfterUpdateHooks = append(detailAfterUpdateHooks, detailHook)
	case boil.BeforeDeleteHook:
		detailBeforeDeleteHooks = append(detailBeforeDeleteHooks, detailHook)
	case boil.AfterDeleteHook:
		detailAfterDeleteHooks = append(detailAfterDeleteHooks, detailHook)
	case boil.BeforeUpsertHook:
		detailBeforeUpsertHooks = append(detailBeforeUpsertHooks, detailHook)
	case boil.AfterUpsertHook:
		detailAfterUpsertHooks = append(detailAfterUpsertHooks, detailHook)
	}
}

// One returns a single detail record from the query.
func (q detailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Detail, error) {
	o := &Detail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boilerEntity: failed to execute a one query for detail")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Detail records from the query.
func (q detailQuery) All(ctx context.Context, exec boil.ContextExecutor) (DetailSlice, error) {
	var o []*Detail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boilerEntity: failed to assign all query results to Detail slice")
	}

	if len(detailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Detail records in the query.
func (q detailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: failed to count detail rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q detailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boilerEntity: failed to check if detail exists")
	}

	return count > 0, nil
}

// Details retrieves all the records using an executor.
func Details(mods ...qm.QueryMod) detailQuery {
	mods = append(mods, qm.From("\"detail\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"detail\".*"})
	}

	return detailQuery{q}
}

// FindDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDetail(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Detail, error) {
	detailObj := &Detail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"detail\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, detailObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boilerEntity: unable to select from detail")
	}

	if err = detailObj.doAfterSelectHooks(ctx, exec); err != nil {
		return detailObj, err
	}

	return detailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Detail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boilerEntity: no detail provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(detailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	detailInsertCacheMut.RLock()
	cache, cached := detailInsertCache[key]
	detailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			detailAllColumns,
			detailColumnsWithDefault,
			detailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(detailType, detailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(detailType, detailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"detail\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"detail\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "boilerEntity: unable to insert into detail")
	}

	if !cached {
		detailInsertCacheMut.Lock()
		detailInsertCache[key] = cache
		detailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Detail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Detail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	detailUpdateCacheMut.RLock()
	cache, cached := detailUpdateCache[key]
	detailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			detailAllColumns,
			detailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boilerEntity: unable to update detail, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"detail\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, detailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(detailType, detailMapping, append(wl, detailPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to update detail row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: failed to get rows affected by update for detail")
	}

	if !cached {
		detailUpdateCacheMut.Lock()
		detailUpdateCache[key] = cache
		detailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q detailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to update all for detail")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to retrieve rows affected for detail")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boilerEntity: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), detailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"detail\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, detailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to update all in detail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to retrieve rows affected all in update all detail")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Detail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boilerEntity: no detail provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(detailColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	detailUpsertCacheMut.RLock()
	cache, cached := detailUpsertCache[key]
	detailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			detailAllColumns,
			detailColumnsWithDefault,
			detailColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			detailAllColumns,
			detailPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boilerEntity: unable to upsert detail, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(detailPrimaryKeyColumns))
			copy(conflict, detailPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"detail\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(detailType, detailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(detailType, detailMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "boilerEntity: unable to upsert detail")
	}

	if !cached {
		detailUpsertCacheMut.Lock()
		detailUpsertCache[key] = cache
		detailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Detail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Detail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boilerEntity: no Detail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), detailPrimaryKeyMapping)
	sql := "DELETE FROM \"detail\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to delete from detail")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: failed to get rows affected by delete for detail")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q detailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boilerEntity: no detailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to delete all from detail")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: failed to get rows affected by deleteall for detail")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(detailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), detailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"detail\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, detailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: unable to delete all from detail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boilerEntity: failed to get rows affected by deleteall for detail")
	}

	if len(detailAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Detail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDetail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), detailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"detail\".* FROM \"detail\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, detailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boilerEntity: unable to reload all in DetailSlice")
	}

	*o = slice

	return nil
}

// DetailExists checks if the Detail row exists.
func DetailExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"detail\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boilerEntity: unable to check if detail exists")
	}

	return exists, nil
}

// Exists checks if the Detail row exists.
func (o *Detail) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DetailExists(ctx, exec, o.ID)
}
