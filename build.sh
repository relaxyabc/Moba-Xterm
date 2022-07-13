#!/bin/bash
# remove origin files
echo "======= start remove go mod and vendor ======="
rm -rf go.mod go.sum vendor/
echo "======= remove go mod and vendor end   ======="
# rebuild mod and init vendor
echo "======= start rebuild go mod and init vendor ======="
go mod init mobaxterm && go mod tidy && go mod vendor
echo "======= rebuild go mod and init vendor end   ======="
# install fyne-cross
echo "======= start install fyne-cross ======="
go install github.com/fyne-io/fyne-cross@latest
echo "======= fyne-cross installed     ======="
# build windows executable file

echo "======= start build windows executable file by fyne-cross ======="
sudo ~/go/bin/fyne-cross windows -arch=* -env "-mod=vendor"
# build linux executable file
echo "======= start build linux executable file by fyne-cross ======="
sudo ~/go/bin/fyne-cross linux -arch=* -tags "GUI" -env "-mod=vendor"
echo "======= executable file build success ======="
