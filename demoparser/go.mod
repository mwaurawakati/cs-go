module github.com/mwaurawakati/csgo/demoparser

go 1.19

replace github.com/mwaurawakati/csgo/playerdata => .././playerdata

require (
	github.com/David-Durst/head-position-model v0.0.0-20220824095659-398a29279792
	github.com/golang/geo v0.0.0-20210211234256-740aa86cb551
	github.com/markus-wa/demoinfocs-golang v1.11.0
	github.com/markus-wa/demoinfocs-golang/v2 v2.13.3
	github.com/mwaurawakati/csgo/playerdata v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/markus-wa/go-unassert v0.1.2 // indirect
	github.com/markus-wa/gobitread v0.2.3 // indirect
	github.com/markus-wa/godispatch v1.4.1 // indirect
	github.com/markus-wa/ice-cipher-go v0.0.0-20220126215401-a6adadccc817 // indirect
	github.com/markus-wa/quickhull-go/v2 v2.1.0 // indirect
)
