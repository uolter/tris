## tris


   0 1 2<br/>
0 0X=<br/>
1 0=X<br/>
2 XX=<br/>
Command line tris game, developed with Go Language. .

After git clone just run go run main.go

### Build
If you want to build the code run:
```
>> go build
>> chmod +x tris
>> ./tris
```

### Running the usnit tests

```
>>  go test -v ./...

=== RUN TestAddOK
--- PASS: TestAddOK (0.00s)
=== RUN TestAddKO
--- PASS: TestAddKO (0.00s)
=== RUN TestAddKOVal
--- PASS: TestAddKOVal (0.00s)
=== RUN TestGetOk
--- PASS: TestGetOk (0.00s)
=== RUN TestOneRowWin
--- PASS: TestOneRowWin (0.00s)
=== RUN TestOneColsWin
--- PASS: TestOneColsWin (0.00s)
=== RUN TestCrossWin
--- PASS: TestCrossWin (0.00s)
=== RUN TestFreeCells
--- PASS: TestFreeCells (0.00s)
PASS
ok  	github.com/uolter/tris/rules	0.002s
>>
```



