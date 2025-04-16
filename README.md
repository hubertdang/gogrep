# gogrep
My own implementation of the `grep` command, written in Go.

## 🔍 Usage

### Search a specific file

```bash
hubert@hubertlaptop:~/gogrep$ gogrep fakefile test
fakefile: there was a test
fakefile: with a test

2 occurences
```

### Search recursively in the current working directory

```bash
hubert@hubertlaptop:~/gogrep$ gogrep -r test
README.md: go install github.com/hubertdang/gogrep@latest
anotherfakefile: testing this
anotherfakefile: testmynuts
fakefile: there was a test
fakefile: with a test
testdir/testfile: test

6 occurences
```

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
