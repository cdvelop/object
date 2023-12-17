module github.com/cdvelop/object

go 1.20

require (
	github.com/cdvelop/input v0.0.75
	github.com/cdvelop/model v0.0.103
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/unixid v0.0.44
)

require github.com/cdvelop/timetools v0.0.32 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input
