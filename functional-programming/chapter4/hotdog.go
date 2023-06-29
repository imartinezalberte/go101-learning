package chapter4

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const HotDogPrice = 4

var OperationNotAllowedErr = errors.New("operation not allowed")

type (
	DNI string

	CreditCard[T uuid.UUID|DNI] struct {
		ID     uuid.UUID
		Credit int
		DNI    T
	}

	Transaction[T uuid.UUID|DNI, K uuid.UUID|DNI] struct {
		Sender   CreditCard[T]
		Receiver CreditCard[K]
		Reason   string
	}

	Person struct {
		Name     string
		Surname  string
		BirthDay time.Time
		DNI      DNI
		Cards    map[uuid.UUID]CreditCard[DNI]
	}

	HotDog struct{
		ID uuid.UUID
		Cards map[uuid.UUID]CreditCard[uuid.UUID]
	}

	PaymentFn[T uuid.UUID|DNI] func(CreditCard[T], uint) (CreditCard[T], error)

	Identificator interface {
		ID() string
	}

	Pricer interface {
		Price() uint
	}
)

func (HotDog) Price() uint {
	return HotDogPrice
}

func (c CreditCard[T]) Add(amount uint) CreditCard[T] {
	c.Credit += int(amount)
	return c 
}

func (c CreditCard[T]) Substract(amount uint) (CreditCard[T], error) {
	{
		amount := int(amount)
		if amount > c.Credit {
			return c, OperationNotAllowedErr
		}
		c.Credit -= amount // side-effect because we are changing the state of a part of the system
		return c, nil
	}
}

func (c CreditCard[T]) Order(order Pricer) (CreditCard[T], error) {
	return c.Substract(order.Price())
}

func Order[T DNI|uuid.UUID](c CreditCard[T], order uint) (CreditCard[T], error) {
	return c.Substract(order)
}
