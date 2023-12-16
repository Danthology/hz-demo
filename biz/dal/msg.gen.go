// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"hz/demo/biz/model/mysql"
)

func newMsg(db *gorm.DB, opts ...gen.DOOption) msg {
	_msg := msg{}

	_msg.msgDo.UseDB(db, opts...)
	_msg.msgDo.UseModel(&mysql.Msg{})

	tableName := _msg.msgDo.TableName()
	_msg.ALL = field.NewAsterisk(tableName)
	_msg.ID = field.NewInt32(tableName, "Id")
	_msg.Num = field.NewInt32(tableName, "num")
	_msg.Code = field.NewString(tableName, "code")

	_msg.fillFieldMap()

	return _msg
}

type msg struct {
	msgDo

	ALL  field.Asterisk
	ID   field.Int32
	Num  field.Int32
	Code field.String

	fieldMap map[string]field.Expr
}

func (m msg) Table(newTableName string) *msg {
	m.msgDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msg) As(alias string) *msg {
	m.msgDo.DO = *(m.msgDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msg) updateTableName(table string) *msg {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "Id")
	m.Num = field.NewInt32(table, "num")
	m.Code = field.NewString(table, "code")

	m.fillFieldMap()

	return m
}

func (m *msg) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msg) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 3)
	m.fieldMap["Id"] = m.ID
	m.fieldMap["num"] = m.Num
	m.fieldMap["code"] = m.Code
}

func (m msg) clone(db *gorm.DB) msg {
	m.msgDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msg) replaceDB(db *gorm.DB) msg {
	m.msgDo.ReplaceDB(db)
	return m
}

type msgDo struct{ gen.DO }

type IMsgDo interface {
	gen.SubQuery
	Debug() IMsgDo
	WithContext(ctx context.Context) IMsgDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMsgDo
	WriteDB() IMsgDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMsgDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMsgDo
	Not(conds ...gen.Condition) IMsgDo
	Or(conds ...gen.Condition) IMsgDo
	Select(conds ...field.Expr) IMsgDo
	Where(conds ...gen.Condition) IMsgDo
	Order(conds ...field.Expr) IMsgDo
	Distinct(cols ...field.Expr) IMsgDo
	Omit(cols ...field.Expr) IMsgDo
	Join(table schema.Tabler, on ...field.Expr) IMsgDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMsgDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMsgDo
	Group(cols ...field.Expr) IMsgDo
	Having(conds ...gen.Condition) IMsgDo
	Limit(limit int) IMsgDo
	Offset(offset int) IMsgDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMsgDo
	Unscoped() IMsgDo
	Create(values ...*mysql.Msg) error
	CreateInBatches(values []*mysql.Msg, batchSize int) error
	Save(values ...*mysql.Msg) error
	First() (*mysql.Msg, error)
	Take() (*mysql.Msg, error)
	Last() (*mysql.Msg, error)
	Find() ([]*mysql.Msg, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*mysql.Msg, err error)
	FindInBatches(result *[]*mysql.Msg, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*mysql.Msg) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMsgDo
	Assign(attrs ...field.AssignExpr) IMsgDo
	Joins(fields ...field.RelationField) IMsgDo
	Preload(fields ...field.RelationField) IMsgDo
	FirstOrInit() (*mysql.Msg, error)
	FirstOrCreate() (*mysql.Msg, error)
	FindByPage(offset int, limit int) (result []*mysql.Msg, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMsgDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m msgDo) Debug() IMsgDo {
	return m.withDO(m.DO.Debug())
}

func (m msgDo) WithContext(ctx context.Context) IMsgDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgDo) ReadDB() IMsgDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgDo) WriteDB() IMsgDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgDo) Session(config *gorm.Session) IMsgDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgDo) Clauses(conds ...clause.Expression) IMsgDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgDo) Returning(value interface{}, columns ...string) IMsgDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgDo) Not(conds ...gen.Condition) IMsgDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgDo) Or(conds ...gen.Condition) IMsgDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgDo) Select(conds ...field.Expr) IMsgDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgDo) Where(conds ...gen.Condition) IMsgDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgDo) Order(conds ...field.Expr) IMsgDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgDo) Distinct(cols ...field.Expr) IMsgDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgDo) Omit(cols ...field.Expr) IMsgDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgDo) Join(table schema.Tabler, on ...field.Expr) IMsgDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMsgDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgDo) RightJoin(table schema.Tabler, on ...field.Expr) IMsgDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgDo) Group(cols ...field.Expr) IMsgDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgDo) Having(conds ...gen.Condition) IMsgDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgDo) Limit(limit int) IMsgDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgDo) Offset(offset int) IMsgDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMsgDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgDo) Unscoped() IMsgDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgDo) Create(values ...*mysql.Msg) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgDo) CreateInBatches(values []*mysql.Msg, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgDo) Save(values ...*mysql.Msg) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgDo) First() (*mysql.Msg, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*mysql.Msg), nil
	}
}

func (m msgDo) Take() (*mysql.Msg, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*mysql.Msg), nil
	}
}

func (m msgDo) Last() (*mysql.Msg, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*mysql.Msg), nil
	}
}

func (m msgDo) Find() ([]*mysql.Msg, error) {
	result, err := m.DO.Find()
	return result.([]*mysql.Msg), err
}

func (m msgDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*mysql.Msg, err error) {
	buf := make([]*mysql.Msg, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgDo) FindInBatches(result *[]*mysql.Msg, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgDo) Attrs(attrs ...field.AssignExpr) IMsgDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgDo) Assign(attrs ...field.AssignExpr) IMsgDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgDo) Joins(fields ...field.RelationField) IMsgDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgDo) Preload(fields ...field.RelationField) IMsgDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgDo) FirstOrInit() (*mysql.Msg, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*mysql.Msg), nil
	}
}

func (m msgDo) FirstOrCreate() (*mysql.Msg, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*mysql.Msg), nil
	}
}

func (m msgDo) FindByPage(offset int, limit int) (result []*mysql.Msg, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m msgDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgDo) Delete(models ...*mysql.Msg) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgDo) withDO(do gen.Dao) *msgDo {
	m.DO = *do.(*gen.DO)
	return m
}
