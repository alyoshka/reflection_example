package modules

func (dialog *Dialog) Echo() {
    dialog.response = dialog.request
}
