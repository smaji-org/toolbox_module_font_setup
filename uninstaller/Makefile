debug:
	go build -gcflags=all="-N -l" .
	GOOS=windows go build -gcflags=all="-N -l"

release:
	go build -ldflags=-w .
	GOOS=windows go build -ldflags "-s -w -H=windowsgui" .

clean:
	rm -f uninstaller uninstaller.exe
