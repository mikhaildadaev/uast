package uast

import (
	"sync"
)

// Публичные конструкторы
func NewSQL(currentDialect *SupportDialect) (*sql, error) {
	if currentDialect == nil {
		return nil, ErrInvalidDialect
	}
	return &sql{
		config: currentDialect.config,
		pool: &sync.Pool{
			New: func() any {
				return newContext()
			},
		},
		processor: currentDialect.processor,
		strateger: currentDialect.strateger,
	}, nil
}

// Публичные методы
func (sql *sql) Build(statement statement) (string, []any, error) {
	if statement == nil {
		return "", nil, ErrInvalidStatement
	}
	contexter := sql.pool.Get().(*contexter)
	defer func() {
		contexter.resetAll()
		sql.pool.Put(contexter)
	}()
	baseRenderer := sql.processor.createRenderer(sql.config, contexter, sql.strateger)
	baseTransformer := sql.processor.createTransformer(sql.config, contexter, sql.strateger)
	baseValidator := sql.processor.createValidator(sql.config, contexter, sql.strateger)
	if err := statement.validate(baseValidator); err != nil {
		return "", nil, err
	}
	if err := statement.transform(baseTransformer); err != nil {
		return "", nil, err
	}
	if err := statement.render(baseRenderer); err != nil {
		return "", nil, err
	}
	return contexter.bufferQuery.String(), contexter.bufferValue, nil
}
func (sql *sql) Close() {
	sql.pool = nil
}
func (sql *sql) SetDialect(dialect *SupportDialect) {
	sql.config = dialect.config
	sql.processor = dialect.processor
	sql.strateger = dialect.strateger
}

// Приватные интерфейсы
type statement interface {
	render(baseRenderer *baseRenderer) error
	transform(baseTransformer *baseTransformer) error
	validate(baseValidator *baseValidator) error
}

// Приватные структуры
type sql struct {
	config    *config
	pool      *sync.Pool
	processor processor
	strateger strateger
}
