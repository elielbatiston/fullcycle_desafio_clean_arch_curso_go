package graph

import "github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
}
