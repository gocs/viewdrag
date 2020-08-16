# viewdrag

[![PkgGoDev](https://pkg.go.dev/badge/github.com/gocs/viewdrag)](https://pkg.go.dev/github.com/gocs/viewdrag)
[![Go Report Card](https://goreportcard.com/badge/github.com/gocs/viewdrag)](https://goreportcard.com/report/github.com/gocs/viewdrag)

viewdrag will drag and drop the image in the viewport.

accepts ebiten.image only as sprite or;
accepts ebiten.image with vertices and indeces as mesh

### running the app

This is basic dragging the view implementation.

This is ebiten's dragging example (with only one "ebiten"): `go run ./cmd/1_basic/main.go`

This creates a mesh if you want to set the mesh upon initializing the app: `go run ./cmd/1_basic_with_mesher/main.go`


This sets the mesh every loop: `go run ./cmd/1_basic_with_mesh/main.go`

## license

`Apache License 2.0`
