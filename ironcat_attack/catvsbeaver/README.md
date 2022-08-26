# Cat vs Beaver (Linux)

Compile instructions (From Windows Powershell):

```
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build ../src/catvsbeaver.go
```

This will generate an executable elf file.

Transfer to a linux system and execute with the following commands:

`sudo chmod +x ./catvsbeaver`
`sudo ./catvsbeaver &`
