# remvoe old build
rm -rf out/

# make folders for new build
mkdir out out/windows
mkdir -p out/macos/question-filter.app

# build for mac
go build -o out/macos/question-filter.app .

# build for windows
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o out/windows/question-filter.exe .