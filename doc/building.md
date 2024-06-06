# ビルド手順

## Debian / Ubuntu

```bash
sudo apt install -y build-essential cmake libjpeg-dev libpng-dev libtiff-dev libgif-dev golang
git clone --recursive https://github.com/ideamans/webpkit.git
cd webpkit
make test && make
```

## FreeBSD

```bash
sudo pkg install -y git cmake jpeg-turbo png tiff giflib gsed gmake go
git clone --recursive https://github.com/ideamans/webpkit.git
cd webpkit
gmake test && gmake
```

## OpenBSD

```bash
sudo pkg_add git cmake jpeg png giflib tiff gsed gmake go
git clone --recursive https://github.com/ideamans/webpkit.git
cd webpkit
make test && make
```

## Windows (MSYS2 UCRT) 試行中

```bash
pacman -S zlib mingw-w64-ucrt-x86_64-libjpeg-turbo mingw-w64-ucrt-x86_64-libpng mingw-w64-ucrt-x86_64-giflib mingw-w64-ucrt-x86_64-libtiff mingw-w64-ucrt-x86_64-go
git clone --recursive https://github.com/ideamans/webpkit.git
cd webpkit
make test && make
```
