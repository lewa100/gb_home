package main

import (
	"github.com/gb_home/hw8/chat"
	"github.com/gb_home/hw8/scanner"
)
func main() {
	scanner.SingleScanner()
	scanner.MultiScanner()
	chat.Server()
}
