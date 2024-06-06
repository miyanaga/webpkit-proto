deps = libwebp/build/libwebp.a webpinfo/webpinfo.c cwebp/cwebp.c dwebp/dwebp.c gif2webp/gif2webp.c

webpkit: $(deps)
	go build -o webpkit .

libwebp/build:
	mkdir libwebp/build

libwebp/build/libwebp.a: libwebp/build
	cd libwebp/build && cmake cmake -WEBP_LINK_STATIC=ON -WEBP_LINK_STATIC_DEFAULT=ON ../ && make

webpinfo/webpinfo.c:
	cp libwebp/examples/webpinfo.c webpinfo/webpinfo.c
	sed -i 's/int main/__attribute__((weak))\nint internal_webpinfo/g' webpinfo/webpinfo.c

cwebp/cwebp.c:
	cp libwebp/examples/cwebp.c cwebp/cwebp.c
	sed -i 's/int main/__attribute__((weak))\nint internal_cwebp/g' cwebp/cwebp.c

dwebp/dwebp.c:
	cp libwebp/examples/dwebp.c dwebp/dwebp.c
	sed -i 's/int main/__attribute__((weak))\nint internal_dwebp/g' dwebp/dwebp.c

gif2webp/gif2webp.c:
	cp libwebp/examples/gif2webp.c gif2webp/gif2webp.c
	sed -i 's/int main/__attribute__((weak))\nint internal_gif2webp/g' gif2webp/gif2webp.c

.PHONY: test
test: $(deps)
	LOG_LEVEL=fatal go test ./app ./beside ./converter ./cwebp ./dwebp ./gif2webp ./imagetype ./l10n ./logger ./mirror ./webpinfo

.PHONY: debug
debug: $(deps)
	go run ./...