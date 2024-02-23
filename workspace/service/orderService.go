package service

import (
	"context"

	"github.com/AdiLambe/TestGoLangCodes/workspace/domain"
	"github.com/AdiLambe/TestGoLangCodes/workspace/dto"
	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
)

type OrderService interface {
	GetRepositry() domain.OrderRepository
	CreateOrder(req *dto.NewOrderRequest) (*dto.NewOrderResponse, *errs.AppError)
	GetOrder(id string) (*dto.OrderResponse, *errs.AppError)
	SaveOrder(order domain.Order) (*dto.OrderResponse, *errs.AppError)
	GetOrdersList(ctx context.Context, status string) ([]dto.OrderResponse, *errs.AppError)
}

// service implementation
type DefaultOrderService struct {
	repo domain.OrderRepository
}

func (s DefaultOrderService) GetRepositry() domain.OrderRepository {
	return s.repo
}

func (s DefaultOrderService) CreateOrder(req *dto.NewOrderRequest) (*dto.NewOrderResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Order{
		Id:          req.Id,
		Name:        req.Name,
		Number:      req.Number,
		Description: req.Description,
		Status:      req.Status,
	}

	newOrder, err := s.repo.SaveOrder(a)
	if err != nil {
		return nil, err
	}
	response := newOrder.ToNewOrderResponseDto()
	return &response, nil
}

func (s DefaultOrderService) SaveOrder(order domain.Order) (*dto.OrderResponse, *errs.AppError) {
	c, err := s.repo.SaveOrder(order)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultOrderService) GetOrdersList(ctx context.Context, status string) ([]dto.OrderResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	orders, err := s.repo.FindAll(ctx, status)
	if err != nil {
		return nil, err
	}

	response := make([]dto.OrderResponse, 0)
	for _, c := range orders {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s DefaultOrderService) GetOrder(id string) (*dto.OrderResponse, *errs.AppError) {

	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewOrderService(repository domain.OrderRepository) OrderService {
	return DefaultOrderService{repository}
}
