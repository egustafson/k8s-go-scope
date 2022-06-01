# Makefile

all:
	go build

package:
	docker build --tag k8s-go-scope .

clean:
	go clean
