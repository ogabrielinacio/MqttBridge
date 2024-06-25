package utils

import "mqttbridge/enums"

func IsValidStatus(status enums.EStatus) bool {
	switch status {
	case enums.Online, enums.Offline, enums.SensorError:
		return true
	default:
		return false
	}
}
