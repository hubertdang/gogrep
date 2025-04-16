# gogrep
My own implementation of the `grep` command, written in Go.

## 🔧 Installation

### Requires Go 1.16 or newer

To install `gogrep`:

```bash
go install github.com/hubertdang/gogrep@latest
```

Make sure your Go `bin` directory is in your `PATH`. Add this line at the end of your `~/.bashrc`:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
