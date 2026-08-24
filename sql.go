// UAST (Universal Abstract Syntax Tree)
// Copyright (C) 2026 Mikhail Dadaev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package uast

import (
	"database/sql"
	"sync"
)

// Публичные конструкторы
func NewSQL(options ...SQLOption) *SQLBuilder {
	sqlBuilder := &SQLBuilder{
		config: DialectDefault.config,
		mutate: false,
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
func WithMutate(mutate bool) SQLOption {
	return func(sqlBuilder *SQLBuilder) {
		sqlBuilder.mutate = mutate
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
	if !sqlBuilder.mutate {
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
	if sqlBuilder.mutate {
		return
	}
	sqlBuilder.config = dialect.config
	sqlBuilder.processor = dialect.processor
	sqlBuilder.strateger = dialect.strateger
}
func (sqlBuilder *SQLBuilder) SetMutate(mutate bool) {
	sqlBuilder.mutate = mutate
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
	mutate    bool
	pool      *sync.Pool
	processor processor
	strateger strateger
}
type SQLOption func(*SQLBuilder)
