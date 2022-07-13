#!/bin/bash
# remove origin files
rm -rf go.mod go.sum vendor/
# rebuild mod and init vendor
go mod init mobaxterm && go mod tidy && go mod vendor
# install fyne-cross
go get github.com/fyne-io/fyne-cross
# build windows executable file
sudo ~/go/bin/fyne-cross windows -arch=* -env "-mod=vendor"
# build linux executable file
sudo ~/go/bin/fyne-cross linux -arch=* -tags "GUI" -env "-mod=vendor"
