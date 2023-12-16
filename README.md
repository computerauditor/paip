# paip
Print all IPs from 1 to 255

## paip
A simple go tool that prints the IP addresses from 1 to 255

## Usage:

```
paip 127.0.0.1
```
![1](https://github.com/computerauditor/paip/assets/117805200/5107a431-c8f0-4b29-87cd-4cde6f825162)

## Or can pipe using

```
echo '127.0.0.1' | paip  
```
![2](https://github.com/computerauditor/paip/assets/117805200/97a6c014-66b6-46b6-b37b-2827b9a6b6de)

## Installation:

Simply use go build to build from binnary and set it up in your $GOPATH 

for example:

```
go build .\paip.go
```
Copy/Move to your defalt go path

```
mv paip.exe \go\bin
```
