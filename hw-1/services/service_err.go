package services

import (
	"errors"
)

var (
	ErrOrderAlreadyExists          = errors.New("order already exists")
	ErrInvalidStorageDate          = errors.New("invalid storage date: storage date cannot be in the past")
	ErrOrderCannotBeReturned       = errors.New("the order cannot be returned: the order is either new or more than two days have passed since the date of issue")
	ErrOrderCannotBeDelivered      = errors.New("order cannot be delivered: order is not in new status")
	ErrStoragePeriodExpired        = errors.New("storage period expired: order cannot be returned")
	ErrStoragePeriodNotExpired     = errors.New("storage period not expired yet: order cannot be returned")
	ErrReturnPeriodExpired         = errors.New("return period expired: order cannot be returned")
	ErrOrderNotBelongToCustomer    = errors.New("order does not belong to customer: invalid customer ID")
	ErrOrderCannotBeReturnToCurier = errors.New("the order cannot be returned: order already delivered")
)
