#!/bin/bash

# install fyne-cross
go install github.com/fyne-io/fyne-cross@latest

sudo ~/go/bin/fyne-cross windows -arch=*
