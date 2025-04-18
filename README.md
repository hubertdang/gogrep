# gogrep
My own implementation of the `grep` command, written in Go (yes, the picture is my original work).

![grepping_gopher](https://github.com/user-attachments/assets/ed3d41b0-e03b-4d49-8241-1ece9f5dc9f3)

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
anotherfakefile: testmygrep
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
