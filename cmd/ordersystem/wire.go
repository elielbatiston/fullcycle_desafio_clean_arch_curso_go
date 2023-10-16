//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/entity"
	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/event"
	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/infra/database"
	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/infra/web"
	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/usecase"
	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}