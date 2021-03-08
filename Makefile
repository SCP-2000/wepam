module:
	go build -buildmode=c-shared -o ./build/wepam.so ./pkg/wepam

clean:
	rm -r ./build

.PHONY: module
