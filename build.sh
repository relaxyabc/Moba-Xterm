#!/bin/bash

# install fyne-cross
go install github.com/fyne-io/fyne-cross

sudo ~/go/bin/fyne-cross windows -arch=*
