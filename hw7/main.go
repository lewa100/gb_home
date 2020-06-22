package main

import (
 "github.com/gb_home/hw7/pipeline"
 "github.com/gb_home/hw7/spinner"
 "github.com/gb_home/hw7/timeServer"
)

func main() {
 spinner.Spinner()
 pipeline.PipeLine()
 timeServer.TimeServer()
}