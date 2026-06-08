package convert

import "strconv"

func Flag(value int) bool {
	return value == 1
}

func Success(value int) bool {
	return Flag(value)
}

func InventoryKey(appID int, classID, instanceID string) string {
	return strconv.Itoa(appID) + "_" + classID + "_" + instanceID
}
