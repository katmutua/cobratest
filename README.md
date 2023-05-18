# Cobratest


Quick 1, 2 example on how to use Cobra

### create the project folder
```bash 
mkdir cobratest && cd cobratest
```
### initialize the project
```bash 
go mod init https://github.com/katmutua/cobratest
```

### create the repository structure
```bash 
mkdir -p cmd/cobratest
mkdir -p pkg/cobratest

touch cmd/cobratest/root.go
touch pkg/cobratest/cobratest.go
touch main.go
```
### add Cobra as a dependency
```bash
go get -u github.com/spf13/cobra@latest
```
### open project in VSCode
```bash 
code .
```
