package uast

import (
	"sync"
)

// Публичные конструкторы
func NewBuilder(currentDialect *SupportDialect) (*builder, error) {
	if currentDialect == nil {
		return nil, ErrInvalidDialect
	}
	return &builder{
		config: currentDialect.config,
		pool: &sync.Pool{
			New: func() any {
				return newContext()
			},
		},
		strateger: currentDialect.strateger,
	}, nil
}

// Публичные методы
func (builder *builder) Build(statement statement) (string, []any, error) {
	if statement == nil {
		return "", nil, ErrInvalidStatement
	}
	contexter := builder.pool.Get().(*contexter)
	defer func() {
		contexter.resetAll()
		builder.pool.Put(contexter)
	}()
	baseRenderer := newRenderer(builder.config, contexter, builder.strateger)
	baseTransformer := newTransformer(builder.config, contexter, builder.strateger)
	baseValidator := newValidator(builder.config, contexter, builder.strateger)
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
func (builder *builder) Close() {
	builder.pool = nil
}
func (builder *builder) Delete(from SourceBase) *stmtDelete {
	return &stmtDelete{
		command: uastManagementDelete,
		from:    from,
	}
}
func (builder *builder) Insert(columns ...markColumnable) *stmtInsert {
	return &stmtInsert{
		command: uastManagementInsert,
		column:  columns,
	}
}
func (builder *builder) Select(fields ...markFieldable) *stmtSelect {
	return &stmtSelect{
		command: uastManagementSelect,
		field:   fields,
	}
}
func (builder *builder) Update(onto SourceBase) *stmtUpdate {
	return &stmtUpdate{
		command: uastManagementUpdate,
		onto:    onto,
	}
}

// Приватные интерфейсы
type statement interface {
	render(baseRenderer *baseRenderer) error
	transform(baseTransformer *baseTransformer) error
	validate(baseValidator *baseValidator) error
}

// Приватные структуры
type builder struct {
	config    *config
	pool      *sync.Pool
	strateger strateger
}
