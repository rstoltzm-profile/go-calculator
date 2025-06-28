#!/usr/bin/bash
echo "basic cli testing"
echo ""
echo "add"
bin/go-calculator-linux --num1 1 --num2 3 --oper add
echo ""
echo "sub"
bin/go-calculator-linux --num1 5 --num2 3 --oper sub
echo ""
echo "mult"
bin/go-calculator-linux --num1 2 --num2 3 --oper mult
echo ""
echo "div"
bin/go-calculator-linux --num1 4 --num2 2 --oper div
echo ""
echo "done"