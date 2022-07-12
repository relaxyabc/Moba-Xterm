#!/bin/bash
# install fyne v2
go install fyne.io/fyne/v2/cmd/fyne@latest

# bundle static resource .
# command: fyne bundle image.png >> bundled.go

# bundle static resource
fyne bundle -package asset asset/logo.png > asset/bundled.go
# append static resource
fyne bundle -package asset asset/image.png >> asset/bundled.go
