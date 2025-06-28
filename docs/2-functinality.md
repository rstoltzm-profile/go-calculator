## Tests for add, sub, mult, and div functions
```text
go test -v
=== RUN   TestAdd
=== RUN   TestAdd/Add:_1_+_1
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/Add:_1_+_1 (0.00s)
=== RUN   TestMult
=== RUN   TestMult/Mult:_1_*_1
--- PASS: TestMult (0.00s)
    --- PASS: TestMult/Mult:_1_*_1 (0.00s)
=== RUN   TestDiv
=== RUN   TestDiv/Div:_1_/_1
--- PASS: TestDiv (0.00s)
    --- PASS: TestDiv/Div:_1_/_1 (0.00s)
=== RUN   TestSub
=== RUN   TestSub/Sub:_1_-_1
--- PASS: TestSub (0.00s)
    --- PASS: TestSub/Sub:_1_-_1 (0.00s)
PASS
ok      github.com/rstoltzm-profile/go-calculator       0.004s
```