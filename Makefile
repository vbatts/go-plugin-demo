
all: code plugins/plugin.so

code: main.go
	go build -o $@ $^

plugins/plugin.so: plugin.go
	go build -buildmode=plugin -o $@ $^

clean:
	rm -rf *~ code plugins/plugin

