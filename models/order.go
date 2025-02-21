package models

import (
	"fmt"
	"time"
)

type OrderStatus string

const (
	StatusNew       OrderStatus = "new"
	StatusExpired   OrderStatus = "expired"
	StatusReturned  OrderStatus = "returned"
	StatusDelivered OrderStatus = "delivered"
)

type Order struct {
	ID           uint        `json:"id"`
	CustomerID   uint        `json:"customer_id"`
	StorageUntil time.Time   `json:"storage_until"`
	Status       OrderStatus `json:"status"`
	DeliveredAt  time.Time   `json:"delivered_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (o Order) String() string {
	storageUntil := o.StorageUntil.Format(time.ANSIC)
	updatedAt := o.UpdatedAt.Format(time.ANSIC)
	deliveredAt := o.DeliveredAt.Format(time.ANSIC)

	// Если DeliveredAt нулевое, выводим "N/A"
	if o.DeliveredAt.IsZero() {
		deliveredAt = "N/A"
	}

	return fmt.Sprintf(
		"\nOrder ID: %d\nCustomer ID: %d\nStatus: %s\nStorage Until: %s\nUpdated At: %s\nDelivered At: %s\n",
		o.ID,
		o.CustomerID,
		o.Status,
		storageUntil,
		updatedAt,
		deliveredAt,
	)
}
