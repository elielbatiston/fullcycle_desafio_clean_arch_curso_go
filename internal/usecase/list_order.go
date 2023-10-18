package usecase

import (
	"github.com/elielbatiston/fullcycle_desafio_clean_arch_curso_go/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	listOrders := []ListOrdersOutputDTO{}

	for _, order := range orders {
		listOrders = append(listOrders, ListOrdersOutputDTO{ID: order.ID, Price: order.Price, Tax: order.Tax, FinalPrice: order.FinalPrice})
	}

	return listOrders, nil
}
