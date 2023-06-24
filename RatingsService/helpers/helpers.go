package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func LogMessageBody(data []byte) {
	var out bytes.Buffer
	json.Indent(&out, data, "", " ")
	out.WriteTo(os.Stdout)
	fmt.Println()
}
