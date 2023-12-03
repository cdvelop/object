module github.com/cdvelop/object

go 1.20

require (
	github.com/cdvelop/input v0.0.58
	github.com/cdvelop/model v0.0.75
	github.com/cdvelop/unixid v0.0.24
)

require (
	github.com/cdvelop/strings v0.0.7
	github.com/cdvelop/timetools v0.0.24 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/input => ../input
