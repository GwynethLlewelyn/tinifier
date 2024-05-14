# tinifier.

A simple GUI to run the Tinify API, courtesy of [TinyPNG](https://tinypng.com).

Requires a developer key (free to get) in order to connect to TinyPNG's web services.

GUI environment (multi-platform) developed by [Fyne](https://fyne.io).

Licensed under a [MIT License](https://gwyneth-llewelyn.mit-license.org/).

## Minimalist instructions:

Most of these instructions are required to use the GUI framework [Fyne](https://fyne.io) and dealing with cross-compilation & packaging.

1. Compile as usual, e.g. `go build .`
2. Optionally, create whatever packages you need for the different operating systems, and use the `fyne` command-line tool, which you can install with `go get fyne.io/fyne/v2/cmd/fyne` (on the same directory where you're working on this repo)

Because Fyne has some OpenGL dependencies, and these are different from system to system, you will need to have the adequate libraries for cross-compilation.

Here is an example: [compiling in macOS/Debian, target is Windows 64bits](https://stackoverflow.com/a/36916044/1035977):

1. `brew install mingw-w64` (macOS, Darwin); or
2. `apt install mingw-64` (Debian, Ubuntu)
3. `env GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc"`

