# smug

Blame when a Go program crashes

* compact output
* simpler than https://github.com/maruel/panicparse
* print the offending line and git blame

# Example

without `smug`:

```
$ go run test/crash.go 
panic: attack!

goroutine 1 [running]:
main.bar(0x1, 0x2, 0x3)
	/home/bob/src/github.com/poofyleek/smug/test/crash.go:12 +0x64
main.foo(0x1, 0x2, 0x3)
	/home/bob/src/github.com/poofyleek/smug/test/crash.go:8 +0x3c
main.main()
	/home/bob/src/github.com/poofyleek/smug/test/crash.go:4 +0x39

goroutine 2 [runnable]:
runtime.forcegchelper()
	/usr/local/go/src/runtime/proc.go:90
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2232 +0x1

goroutine 3 [runnable]:
runtime.bgsweep()
	/usr/local/go/src/runtime/mgc0.go:82
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2232 +0x1
exit status 2
```

with `smug`:

```
$ go run test/crash.go |& go run smug.go
panic: attack!
goroutine 1 [running]:
 main.bar(0x1, 0x2, 0x3)        test/crash.go:12 +0x64
 main.foo(0x1, 0x2, 0x3)        /home/bob/src/github.com/poofyleek/smug/test/crash.go:8 +0x3c
 main.main()                    /home/bob/src/github.com/poofyleek/smug/test/crash.go:4 +0x39
goroutine 2 [runnable]:
 runtime.forcegchelper()        /usr/local/go/src/runtime/proc.go:90
 runtime.goexit()               /usr/local/go/src/runtime/asm_amd64.s:2232 +0x1
goroutine 3 [runnable]:
 runtime.bgsweep()              /usr/local/go/src/runtime/mgc0.go:82
 runtime.goexit()               /usr/local/go/src/runtime/asm_amd64.s:2232 +0x1
exit status 2
BLAME ba436098 (poofy 2015-12-09 17:08:37 -0800 12) 	panic("attack!")

```
