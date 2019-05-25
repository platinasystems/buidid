This package gets a GO program's [BuildId].
Usage,

```golang
	s, err := buildid.New("/proc/self/exe")
	if err != nil {
		fmt.Frintln(os.Stderr, err)
	} else {
		fmt.Println(s)
	}
```

Which is equivalent to,

```console
$ go tool buildid PROGRAM
```

---

*&copy; 2018-2019 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE].*

[LICENSE]: LICENSE
[BuildId]: https://godoc.org/cmd/buildid
