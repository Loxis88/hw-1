package storage

import (
	"errors"

	"hw-1/models"
)

// да это логично было бы поместить этот контракт в пакет сервиса т.к. он будет описывать ожидаемый сервисом сторадж но я пока оставлю его здесь

// общий интерфейс для реализации каждого хранилища который будет ожидать сервис
type OrderStorage interface {
	AddOrder(order models.Order) error
	UpdateOrder(order models.Order) error
	DeleteOrder(id uint) error
	GetOrders() []models.Order
	FindOrder(id uint) (*models.Order, error)
}

// New - функция которая достает нужную реализацию хранилища из мапы registry
func New(storageType string, config string) (OrderStorage, error) {
	factory, exists := registry[storageType]
	if !exists {
		return nil, ErrUnknownStorageType
	}
	return factory(config)
}

// StorageFactory - это сигнатура для функции, которая создаёт хранилище
type StorageFactory func(config string) (OrderStorage, error)

// registry - это мапа, где хранятся все доступные типы хранилищ
// идея в том что эта мапа хранит все определенные функции инициализации хранилищ которые будут определяться непосредственно в пакетах хранилищ
// что позволит расширять код без его изменения
var registry = make(map[string]StorageFactory)

// RegisterStorage - добавляет новый тип хранилища в мапу
func RegisterStorage(name string, factory StorageFactory) {
	registry[name] = factory
}

// определения общих ошибок для любой реализации хранилища
var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrUnknownStorageType = errors.New("unknown storage type")
)
