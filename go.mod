module github.com/markxp/gomod

go 1.12

require (
	github.com/markxp/gomod/sub v0.0.0
	github.com/markxp/gomod/sub2 v0.0.0
)

replace (
	github.com/markxp/gomod/sub => ./sub
	github.com/markxp/gomod/sub2 => ./sub2
)
