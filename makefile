all: mac win lin

mac: 
		GOOS=darwin 	GOARCH=386 go build -o out/cyk-mac-bin

win: 
		GOOS=windows 	GOARCH=386 go build -o out/cyk-windows-bin

lin: 
		GOOS=linux 		GOARCH=386 go build -o out/cyk-linux-bin 

test:
	go test

