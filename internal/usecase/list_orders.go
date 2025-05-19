package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrdersOutputDTO []OrderOutputDTO

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {

	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err

	}
	var dto ListOrdersOutputDTO
	for _, order := range orders {
		dto = append(dto, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}
	return dto, nil
}
