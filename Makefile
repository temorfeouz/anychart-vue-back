static:
	cd vue_src; npm run build; cd ..;
gen:
	go generate
compile:
	go build  -ldflags="-s -w"
compile-mips:
	GOOS=linux GOARCH=mipsle go build -ldflags="-s -w"
build:	static gen compile
build-mips: static gen compile-mips
run: build
	./anychart-vue-back
rf:
	cd vue_src; npm run dev


