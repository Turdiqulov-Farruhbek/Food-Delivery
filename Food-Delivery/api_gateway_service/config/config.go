package config

import (
    "os"
)

var (
    CourierDeliveryServiceURL    = getEnv("Courier_Delivery_Service", "http://localhost:2020")
    ProductOrderingServiceURL = getEnv("Product_Ordering_service", "http://localhost:2030")
    NotificationsAlertsServiceURL = getEnv("Notifications_Alerts_service", "http://localhost:4040")
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
