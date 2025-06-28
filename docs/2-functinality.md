## Full Test log for functionality

```bash
go test -v
```

```text
=== RUN   TestMain
=== RUN   TestMain/Check_main_is_called
=== RUN   TestMain/Check_main_1+2=3
=== RUN   TestMain/Check_main_1-2=-1
=== RUN   TestMain/Check_main_2*2=4
=== RUN   TestMain/Check_main_4/2=2
--- PASS: TestMain (0.00s)
    --- PASS: TestMain/Check_main_is_called (0.00s)
    --- PASS: TestMain/Check_main_1+2=3 (0.00s)
    --- PASS: TestMain/Check_main_1-2=-1 (0.00s)
    --- PASS: TestMain/Check_main_2*2=4 (0.00s)
    --- PASS: TestMain/Check_main_4/2=2 (0.00s)
=== RUN   TestFlags
=== RUN   TestFlags/Check_for_no_--oper_flag
=== RUN   TestFlags/Check_for_invalid_--oper_flag
--- PASS: TestFlags (0.00s)
    --- PASS: TestFlags/Check_for_no_--oper_flag (0.00s)
    --- PASS: TestFlags/Check_for_invalid_--oper_flag (0.00s)
=== RUN   TestAdd
=== RUN   TestAdd/Add:_1_+_1
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/Add:_1_+_1 (0.00s)
=== RUN   TestSub
=== RUN   TestSub/Sub:_1_-_1
--- PASS: TestSub (0.00s)
    --- PASS: TestSub/Sub:_1_-_1 (0.00s)
=== RUN   TestMult
=== RUN   TestMult/Mult:_1_*_1
--- PASS: TestMult (0.00s)
    --- PASS: TestMult/Mult:_1_*_1 (0.00s)
=== RUN   TestDiv
=== RUN   TestDiv/Div:_1_/_1
--- PASS: TestDiv (0.00s)
    --- PASS: TestDiv/Div:_1_/_1 (0.00s)
ok      github.com/rstoltzm-profile/go-calculator       0.002s
```