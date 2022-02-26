package route

import (
	"project_we/delivery"
)

func Init(delivery *delivery.Deliveries) {
	Common(delivery)
	Location(delivery)
	User(delivery)
}
