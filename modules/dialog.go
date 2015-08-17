package modules

import (
    "fmt"
    "reflect"
)

type Dialog struct {
    request  string
    response string
}

func NewDialog(request string) (Dialog) {
    return Dialog{request, ""}
}

func (dialog Dialog) Print() {
    fmt.Println("request: ", dialog.request)
    fmt.Println("response: ", dialog.response)
}

func (dialog *Dialog) Invoke(method string) {
    reflect.ValueOf(dialog).MethodByName(method).Call([]reflect.Value{})
}
