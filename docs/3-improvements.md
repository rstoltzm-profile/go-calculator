## test coverage
```text
go test -cover
PASS
coverage: 90.9% of statements
ok      github.com/rstoltzm-profile/go-calculator       0.005s
```

## test output
```text
go test -v
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
=== RUN   TestFlags/Check_for_invalid_inputs_to_num1
=== RUN   TestFlags/Check_for_invalid_inputs_to_num2
--- PASS: TestFlags (0.00s)
    --- PASS: TestFlags/Check_for_no_--oper_flag (0.00s)
    --- PASS: TestFlags/Check_for_invalid_--oper_flag (0.00s)
    --- PASS: TestFlags/Check_for_invalid_inputs_to_num1 (0.00s)
    --- PASS: TestFlags/Check_for_invalid_inputs_to_num2 (0.00s)
=== RUN   TestAdd
=== RUN   TestAdd/Add:_1_+_1
=== RUN   TestAdd/Add:_1_+_-5
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/Add:_1_+_1 (0.00s)
    --- PASS: TestAdd/Add:_1_+_-5 (0.00s)
=== RUN   TestSub
=== RUN   TestSub/Sub:_1_-_1
=== RUN   TestSub/Sub:_1_-_-2
--- PASS: TestSub (0.00s)
    --- PASS: TestSub/Sub:_1_-_1 (0.00s)
    --- PASS: TestSub/Sub:_1_-_-2 (0.00s)
=== RUN   TestMult
=== RUN   TestMult/Mult:_1_*_1
=== RUN   TestMult/Mult:_0_*_1
--- PASS: TestMult (0.00s)
    --- PASS: TestMult/Mult:_1_*_1 (0.00s)
    --- PASS: TestMult/Mult:_0_*_1 (0.00s)
=== RUN   TestDiv
=== RUN   TestDiv/Div:_1_/_1
=== RUN   TestDiv/Div:_1_/_0
=== RUN   TestDiv/Div:_0_/_1
--- PASS: TestDiv (0.00s)
    --- PASS: TestDiv/Div:_1_/_1 (0.00s)
    --- PASS: TestDiv/Div:_1_/_0 (0.00s)
    --- PASS: TestDiv/Div:_0_/_1 (0.00s)
PASS
ok      github.com/rstoltzm-profile/go-calculator       0.005s
```

## cli testing
```text
make test-bin
CG0_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-calculator-linux main.go
./test-go-calculator-cli.sh
basic cli testing

add
GO Calculator App
Oper:    add
Inputs:  1 3
Output: 4

sub
GO Calculator App
Oper:    sub
Inputs:  5 3
Output: 2

mult
GO Calculator App
Oper:    mult
Inputs:  2 3
Output: 6

div
GO Calculator App
Oper:    div
Inputs:  4 2
Output: 2

done
```