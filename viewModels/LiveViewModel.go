package viewModels

import "mqttbridge/enums"

type LiveViewModel struct {
	SerialNumber string        `json:serialNumber`
	Status       enums.EStatus `json:status`
}
