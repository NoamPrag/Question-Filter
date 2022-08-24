# remvoe old build
rm -rf app/

# make folders for new build
mkdir app app/windows
mkdir -p app/macos/question-filter.app

# build for mac
go build -o app/macos/question-filter.app .

# build for windows
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o app/windows/question-filter.exe .