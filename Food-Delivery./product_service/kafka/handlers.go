package kafka

import (
	"context"
	// "encoding/json"
	"log"

	pb "delivery/product_service/genproto"
	"delivery/product_service/service"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/encoding/protojson"
)

func OrderCreateHandler(orderService *service.OrderService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var order pb.OrderCreateReq

		if err := protojson.Unmarshal([]byte(message.Value), &order); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)

			return
		}

		res, err := orderService.CreateOrder(context.Background(), &order)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Printf("Created order: %+v", res)
	}
}

func OrderUpdateHandler(orderService *service.OrderService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var order pb.OrderUpdate
		if err := protojson.Unmarshal([]byte(message.Value), &order); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)

			return
		}

		res, err := orderService.UpdateOrder(context.Background(), &order)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Printf("Updated Ordere: %+v", res)
	}
}

func ProductCreateHandler(prodService *service.ProductService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var prod pb.ProductCreateReq
		if err := protojson.Unmarshal([]byte(message.Value), &prod); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)

			return
		}

		res, err := prodService.CreateProduct(context.Background(), &prod)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Created product: %+v", res)
	}
}
func ProductUpdateHandler(prodService *service.ProductService) func(message kafka.Message) {
	return func(message kafka.Message) {
		//unmarshal the message
		var prod pb.ProductCreateReq
		if err := protojson.Unmarshal([]byte(message.Value), &prod); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := prodService.UpdateProduct(context.Background(), &prod)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Updated product: %+v", res)
	}
}

func ItemCreateHandler(itemService *service.OrderItemService) func(message kafka.Message) {
	return func(message kafka.Message) {
		log.Println("kafka- handlers- intemcreate handler!!!!!!!!!!!!!!!!!!1")

		//unmarshal the message
		var item pb.ItemCreateReq
		log.Println("kafka- handlers- intemcreate handler initialized model")
		if err := protojson.Unmarshal([]byte(message.Value), &item); err != nil {
			log.Println("kafka- handlers- intemcreate handler inside the errro ")

			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)

			return
		}
		log.Println("kafka- handlers- intemcreate handler after the error ")

		res, err := itemService.AddItem(context.Background(), &item)
		log.Println("kafka- handlers- intemcreate handler called the function errorr ", err)

		if err != nil {
			log.Println(err)
			log.Fatal(err)
			return
		}
		log.Printf("Created item: %+v", res)
	}
}
func ItemUpdateHandler(itemService *service.OrderItemService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var item pb.ItemCreateReq
		if err := protojson.Unmarshal(message.Value, &item); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := itemService.UpdateItem(context.Background(), &item)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Updated item: %+v", res)
	}
}

func KitchenCreateHandler(kitchenService *service.KitchenService) func(message kafka.Message) {
	return func(message kafka.Message) {
		//unmarshal the message
		var kitchen pb.KitchenCreateReq
		if err := protojson.Unmarshal(message.Value, &kitchen); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			log.Fatal(err)
			return
		}

		res, err := kitchenService.CreateKitchen(context.Background(), &kitchen)
		if err != nil {
			log.Panicln(err)
			return
		}

		log.Printf("Created kitchen: %+v", res)
	}
}

func KitchenUpdateHandler(kitchenService *service.KitchenService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var kitchen pb.KitchenCreateReq
		if err := protojson.Unmarshal(message.Value, &kitchen); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := kitchenService.UpdateKitchen(context.Background(), &kitchen)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Updated kitchen: %+v", res)
	}
}

func LocationCreateHandler(locationService *service.LocationService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var location pb.LocationCreateReq
		if err := protojson.Unmarshal(message.Value, &location); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := locationService.CreateLocation(context.Background(), &location)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Created location: %+v", res)
	}
}
func LocationUpdateHandler(locationService *service.LocationService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var location pb.LocationCreateReq
		if err := protojson.Unmarshal(message.Value, &location); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := locationService.UpdateLocation(context.Background(), &location)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Updated location: %+v", res)
	}
}
func OrderAcceptHandler(courierService *service.CourierService) func(message kafka.Message) {
	return func(message kafka.Message) {
		var order_accept pb.OrderAcept
		if err := protojson.Unmarshal(message.Value, &order_accept); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := courierService.AcceptOrder(context.Background(), &order_accept)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Created location: %+v", res)
	}
}
