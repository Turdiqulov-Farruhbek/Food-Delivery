package app

import (
	"errors"

	"delivery/product_service/config"
	"delivery/product_service/kafka"
	"delivery/product_service/service"
	"delivery/product_service/storage/postgres"
)

func Regitries(cfg config.Config, db *postgres.Storage) error {
	//services
	orderService := service.NewOrderService(db)
	productService := service.NewProductService(db)
	itemService := service.NewOrderItemService(db)
	kitchenService := service.NewKitchenService(db)
	locationService := service.NewLocationService(db)
	courierService := service.NewCourierService(db)

	//brokers
	brokers := []string{cfg.KafkaHost + cfg.KafkaPort}

	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "order-create", "order-create", kafka.OrderCreateHandler(orderService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'order-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "order-update", "order-update", kafka.OrderUpdateHandler(orderService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'order-update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "product-create", "product-create", kafka.ProductCreateHandler(productService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'product-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "product-update", "product-update", kafka.ProductUpdateHandler(productService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'product-update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "item-create", "item-create", kafka.ItemCreateHandler(itemService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'item-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "item-update", "item-update", kafka.ItemUpdateHandler(itemService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'item-update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "kitchen-create", "kitchen-create", kafka.KitchenCreateHandler(kitchenService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'kitchen-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "kitchen-update", "kitchen-update", kafka.KitchenUpdateHandler(kitchenService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'kitchen-update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "location-create", "location-create", kafka.LocationCreateHandler(locationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'location-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "location-update", "location-update", kafka.LocationUpdateHandler(locationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'location-update' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "order-accept", "order-accept", kafka.OrderAcceptHandler(courierService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'order-accept' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}

	return nil
}
