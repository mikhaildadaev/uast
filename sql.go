package uast

import (
	"database/sql"
	"sync"
)

// Публичные конструкторы
func NewSQL(options ...SQLOption) *SQLBuilder {
	sqlBuilder := &SQLBuilder{
		config:  DialectDefault.config,
		mutable: false,
		pool: &sync.Pool{
			New: func() any {
				return newContext()
			},
		},
		processor: DialectDefault.processor,
		strateger: DialectDefault.strateger,
	}
	for _, opt := range options {
		opt(sqlBuilder)
	}
	return sqlBuilder
}

// Публичные функции
func WithDialect(dialect *SupportDialect) SQLOption {
	return func(sqlBuilder *SQLBuilder) {
		if dialect != nil {
			sqlBuilder.config = dialect.config
			sqlBuilder.processor = dialect.processor
			sqlBuilder.strateger = dialect.strateger
		}
	}
}
func WithMutable() SQLOption {
	return func(sqlBuilder *SQLBuilder) {
		sqlBuilder.mutable = true
	}
}

// Публичные методы
func (sqlBuilder *SQLBuilder) Build(statement statement) (string, []any, error) {
	if statement == nil {
		return "", nil, ErrInvalidStatement
	}
	contexter := sqlBuilder.pool.Get().(*contexter)
	defer func() {
		contexter.resetAll()
		sqlBuilder.pool.Put(contexter)
	}()
	stmt := statement
	if !sqlBuilder.mutable {
		stmt = statement.clone()
	}
	baseRenderer := sqlBuilder.processor.createRenderer(sqlBuilder.config, contexter, sqlBuilder.strateger)
	baseTransformer := sqlBuilder.processor.createTransformer(sqlBuilder.config, contexter, sqlBuilder.strateger)
	baseValidator := sqlBuilder.processor.createValidator(sqlBuilder.config, contexter, sqlBuilder.strateger)
	if err := stmt.validate(baseValidator); err != nil {
		return "", nil, err
	}
	if err := stmt.transform(baseTransformer); err != nil {
		return "", nil, err
	}
	if err := stmt.render(baseRenderer); err != nil {
		return "", nil, err
	}
	return contexter.bufferQuery.String(), contexter.bufferValue, nil
}
func (sqlBuilder *SQLBuilder) Close() {
	sqlBuilder.pool = nil
}
func (sqlBuilder *SQLBuilder) Exec(stmt statement, db *sql.DB) (sql.Result, error) {
	query, args, err := sqlBuilder.Build(stmt)
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}
func (sqlBuilder *SQLBuilder) SetDialect(dialect *SupportDialect) {
	if sqlBuilder.mutable {
		return
	}
	sqlBuilder.config = dialect.config
	sqlBuilder.processor = dialect.processor
	sqlBuilder.strateger = dialect.strateger
}
func (sqlBuilder *SQLBuilder) SetMutable() {
	sqlBuilder.mutable = true
}
func (sqlBuilder *SQLBuilder) Query(stmt statement, db *sql.DB) (*sql.Rows, error) {
	query, args, err := sqlBuilder.Build(stmt)
	if err != nil {
		return nil, err
	}
	return db.Query(query, args...)
}
func (sqlBuilder *SQLBuilder) QueryRow(stmt statement, db *sql.DB) (*sql.Row, error) {
	query, args, err := sqlBuilder.Build(stmt)
	if err != nil {
		return nil, err
	}
	return db.QueryRow(query, args...), nil
}

// Приватные интерфейсы
type statement interface {
	clone() statement
	render(baseRenderer *baseRenderer) error
	transform(baseTransformer *baseTransformer) error
	validate(baseValidator *baseValidator) error
}

// Приватные структуры
type SQLBuilder struct {
	config    *config
	mutable   bool
	pool      *sync.Pool
	processor processor
	strateger strateger
}
type SQLOption func(*SQLBuilder)
