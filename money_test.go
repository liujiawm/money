package money

import (
	"fmt"
	"testing"
)

var (
	a Munit = 8425896210
	b Munit = 5456789123
	c Munit = 789461244567123453

	ma = New(a, "CNY")
	mb = New(b, "CNY")
	mc = New(c, "BTC")
)

func TestMunitCompare(t *testing.T) {
	fmt.Println("************************** 两值比较 **************************")
	cm := a.Cmp(b)
	if cm != 1 {
		t.Error("比较错误")
	}
	// ==
	if a.Equal(b) {
		t.Error("比较错误")
	}

	// >
	if b.GreaterThan(a) {
		t.Error("比较错误")
	}

	// >=
	if b.GreaterThanOrEqual(a) {
		t.Error("比较错误")
	}

	// <
	if a.LessThan(b) {
		t.Error("比较错误")
	}
	// <=
	if a.LessThanOrEqual(b) {
		t.Error("比较错误")
	}

}

func TestMunitCalculate(t *testing.T) {
	fmt.Println("************************** 两值计算 **************************")

	fmt.Println("Munit相加")
	ab, err := a.Add(b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v加%v等于%v\n", a, b, ab)
	}
	fmt.Println("Munit相减")
	ab, err = a.Sub(b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v减%v等于%v\n", a, b, ab)
	}

	fmt.Println("Munit乘5")
	mul, err := a.Mul(5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// mul == 42129481050
		fmt.Printf("%v乘%v等于%v\n", a, 5, mul)
	}

	fmt.Println("Munit除以5")
	div, err := a.Div(5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// div == 1685179242
		fmt.Printf("%v除以%v等于%v\n", a, 5, div)
	}
}

func TestMoneyCalculate(t *testing.T) {
	fmt.Println("************************** 钱计算 **************************")

	fmt.Println("相同币种的Money相加")
	mamb, err := MoneyAdd(ma, mb)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d(%s)加%d(%s)等于%d(%s)\n", int64(ma.GetAmount()), ma.GetCurrency().Code, int64(mb.GetAmount()), mb.GetCurrency().Code, int64(mamb.GetAmount()), mamb.GetCurrency().Code)
	}

	fmt.Println("不同币种的Money相加")
	if err := ma.MoneyAdd(mc); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error("Money计算错误")
	}

	//ch := make(chan bool)
	go func() {
		fmt.Println("相加后附给ma")
		if err := ma.MoneyAdd(mb); err != nil {
			fmt.Println(err.Error())
		}
		//ch <- true
	}()
	//<-ch

	fmt.Println("相同币种的Money相减")
	mamb, err = MoneySub(ma, mb)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d(%s)减%d(%s)等于%d(%s)\n", int64(ma.GetAmount()), ma.GetCurrency().Code, int64(mb.GetAmount()), mb.GetCurrency().Code, int64(mamb.GetAmount()), mamb.GetCurrency().Code)
	}

	mamb, err = MoneySub(mb, ma)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d(%s)减%d(%s)等于%d(%s)\n", int64(mb.GetAmount()), mb.GetCurrency().Code, int64(ma.GetAmount()), ma.GetCurrency().Code, int64(mamb.GetAmount()), mamb.GetCurrency().Code)
	}

}

func TestAsMajorUnits(t *testing.T) {
	fmt.Println("************************** 转标准单位 **************************")
	fmt.Println(a)
	fmt.Println(ma.AsMajorUnits(true))
}

func TestThousandsSeparator(t *testing.T) {
	fmt.Println("************************** 转标准单位后千分分隔 **************************")
	fmt.Println(c)
	fmt.Printf("%s %s \n", mc.GetCurrency().Symbol, mc.ThousandsSeparator(false))
}
