package kafka

import (
	"context"
	// "encoding/json"
	"log"

	pb "delivery/notification_service/genproto"
	"delivery/notification_service/service"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/encoding/protojson"
)

func NotificationCreateHandler(notifService *service.NotificationService) func(message kafka.Message) {
	return func(message kafka.Message) {

		//unmarshal the message
		var notif pb.NotificationCreate
		if err := protojson.Unmarshal([]byte(message.Value), &notif); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := notifService.CreateNotification(context.Background(), &notif)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Created notification: %+v", res)
	}
}
func NotifyAllHandler(notifService *service.NotificationService) func(message kafka.Message) {
	return func(message kafka.Message) {
		//unmarshal the message
		var req pb.NotificationMessage
		if err := protojson.Unmarshal([]byte(message.Value), &req); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := notifService.NotifyAll(context.Background(), &req)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Notified All: %+v", res)
	}
}
