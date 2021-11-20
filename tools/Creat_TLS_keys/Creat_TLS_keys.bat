echo ##Set your GoRoot path#
set GOROOT=0
set /P GOROOT=

if "%GOROOT%" neq "0" (
    set GO111MODULE=off
    go run "%GOROOT%/src/crypto/tls/generate_cert.go" --host localhost
)