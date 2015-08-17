package main

import (
    "reflection_example/modules"
)

func main() {
    methodName := "Echo"
    dialog := modules.NewDialog("hello")
    dialog.Invoke(methodName)
    dialog.Print()
}
