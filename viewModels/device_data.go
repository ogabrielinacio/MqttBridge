package viewModels

import "mqttbridge/enums"

type DeviceData struct {
	SerialNumber string        `json:serialNumber`
	SoilMoisture int           `json:soilMoisture`
	Status       enums.EStatus `json:status`
}
