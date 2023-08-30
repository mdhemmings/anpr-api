# Notes

## Cross-compile to Windows from Linux (using cgo)

```bash
apt-get install mingw-w64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o output.exe ./code/root/
```
