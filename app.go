package main

type app struct {
	Name    string
	Version string
}

// App include application name and version
var App = app{
	Name:    "tcpeek_exporter",
	Version: "0.1.0",
}
