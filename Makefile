static:
	cd vue_src; npm run build; cd ..;
gen:
	go generate
compile:
	go build  -ldflags="-s -w"
build:	static gen compile
run: build
	./anychart-vue-back



