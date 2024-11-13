package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderRepository struct {
	saveFunc func(order Order) error
}

func (m *mockOrderRepository) Save(order Order) error {
	return m.saveFunc(order)
}

func TestCreateOrder(t *testing.T) {
	// Success case
	t.Run("success", func(t *testing.T) {
		repo := &mockOrderRepository{
			saveFunc: func(order Order) error {
				// Simulate successful save
				return nil
			},
		}
		service := NewOrderService(repo)

		err := service.CreateOrder(Order{Total: 100})
		assert.NoError(t, err)
	})

	t.Run("total less than 0", func(t *testing.T) {
		repo := &mockOrderRepository{
			saveFunc: func(order Order) error {
				// This won't be called due to validation
				return nil
			},
		}
		service := NewOrderService(repo)

		err := service.CreateOrder(Order{Total: -10})
		assert.Error(t, err)
		assert.Equal(t, "total must be positive", err.Error())
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &mockOrderRepository{
			saveFunc: func(order Order) error {
				return errors.New("database error")
			},
		}
		service := NewOrderService(repo)

		err := service.CreateOrder(Order{Total: 10})
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
