package viewModels

import "mqttbridge/enums"

type DeviceDataViewModel struct {
	SerialNumber string        `json:serialNumber`
	SoilMoisture int           `json:soilMoisture`
	Status       enums.EStatus `json:status`
}
