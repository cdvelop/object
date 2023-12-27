module github.com/cdvelop/object

go 1.20

require (
	github.com/cdvelop/input v0.0.80
	github.com/cdvelop/model v0.0.107
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/structs v0.0.1
	github.com/cdvelop/unixid v0.0.47
)

require github.com/cdvelop/timetools v0.0.32 // indirect

replace github.com/cdvelop/model => ../model
