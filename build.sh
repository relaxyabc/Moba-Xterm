#!/bin/bash

# install fyne-cross
go install github.com/fyne-io/fyne-cross@latest
# build windows executable file
sudo ~/go/bin/fyne-cross windows -arch=*
