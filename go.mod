module github.com/cdvelop/object

go 1.20

require (
	github.com/cdvelop/input v0.0.41
	github.com/cdvelop/model v0.0.59
	github.com/cdvelop/unixid v0.0.7
)

require (
	github.com/cdvelop/strings v0.0.2
	github.com/cdvelop/timetools v0.0.7 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/input => ../input
