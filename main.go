package main

import (
	_ "github.com/Abdulhossein/hiddify-app-extension/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
