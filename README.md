# maps
cartography

Map should become a cartography infrastructure for [go](https://www.golang.org).

# Status
- [x] `coordinates.go`: Spherical coordinates transformations
- [x] `tile.go`: Tile definitions and tile server interfaces
- [ ] path overlays
- [ ] file imports: gpx, fit, ...
- [ ] `cmd/map`: application of a slippery map
	- [ ] accepting file drops for path overlay
	- [ ] track summary
	- [ ] multi-stage support
	- [ ] altitude graphs
	- [ ] overlay all travelled points
	- [ ] answer the question: Have I been here before? When?
	- as a:
		- [ ] desktop application with shiny frontend
		- [ ] web server application
		- [ ] mobile application
- [ ] .osm.pbf -> extract -> render -> tile server pipeline
- A long way to go

# BUGS
- Go has still too many keywords. The package cannot be called `map`.
