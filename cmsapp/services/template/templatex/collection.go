package templatex

import "github.com/goatcms/goatcore/varutil"

// Contains return true if a string array contains str
func Contains(arr []string, str string) bool {
	return varutil.IsArrContainStr(arr, str)
}
