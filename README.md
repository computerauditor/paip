# paip
Print all IPs subnet from 1 to 255

## paip
A simpe go tool that prints the IP addresses in the subnet from 1 to 255

## Usage:

```
paip 127.0.0.1
```
![1](https://github.com/computerauditor/paip/assets/117805200/5107a431-c8f0-4b29-87cd-4cde6f825162)

## Or can pipe using

```
echo '127.0.0.1' | paip  
```
![2](https://github.com/computerauditor/paip/assets/117805200/94193d05-fa4d-4998-af0c-f7563a4e7afa)


## Installation:

Simply use go build to build from binnary and set it up in your $GOPATH 

for example:

```
go build .\paip.go
```
Copy/Move to your defalt go path

```
mv paip.go \go\bin
```
