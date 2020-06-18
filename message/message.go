package message

import (
	"fmt"
	"strconv"
	"strings"
)

var startRegion = "START REGION "
var endRegion="END REGION "
var startSendingTable="START SENDING TABLE "
var endSendingTable="END SENDING TABLE "

func MakeStartRegion(region int) string {
	return fmt.Sprintf(startRegion+"%d", region)
}
func ReadStartRegion(recv string) int {
	if !strings.Contains(recv, startRegion) {
		return -1
	}
	ls := strings.Split(recv, " ")
	if len(ls) != 3 {
		return -1
	}
	res, err := strconv.ParseInt(ls[2], 10, 32)
	if err != nil {
		return -1
	}
	return int(res)
}
func MakeEndRegion(region int) string {
	return fmt.Sprintf(endRegion+"%d", region)
}
func ReadEndRegion(recv string) int {
	if !strings.Contains(recv, endRegion) {
		return -1
	}
	ls := strings.Split(recv, " ")
	if len(ls) != 3 {
		return -1
	}
	res, err := strconv.ParseInt(ls[2], 10, 32)
	if err != nil {
		return -1
	}
	return int(res)
}
func MakeStartSendingTable(table string) string {
	return fmt.Sprintf(startSendingTable+"%s", table)
}
func ReadStartSendingTable(recv string) (bool,string) {
	if !strings.Contains(recv, startSendingTable) {
		return false,""
	}
	ls := strings.Split(recv, " ")
	if len(ls) != 4 {
		return false,""
	}
	return true,ls[3]
}
func MakeEndSendingTable(table string) string {
	return fmt.Sprintf(endSendingTable+"%s", table)
}
func ReadEndSendingTable(recv string) (bool,string) {
	if !strings.Contains(recv, endSendingTable) {
		return false,""
	}
	ls := strings.Split(recv, " ")
	if len(ls) != 4 {
		return false,""
	}
	return true,ls[3]
}

