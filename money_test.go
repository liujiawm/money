package money

import (
	"fmt"
	"testing"

)

func TestAdd(t *testing.T){

	var a,b,c Munit = 45896210,456789123,789461244567123453

	ma:=New(a,"CNY")
	mb:=New(b,"CNY")
	mc:=New(c,"BTC")

	fmt.Println("//////////////////////////////////////////////////////////////")
	fmt.Println("Munit相加")
	ab,err := a.Add(b)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(ab)
		fmt.Printf("%d加%d等于%d\n",int64(a),int64(b),int64(ab))
	}

	fmt.Println("//////////////////////////////////////////////////////////////")
	fmt.Println("相同币种的Money相加")
	mamb,err := ma.MoneyAdd(mb)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Printf("%d(%s)加%d(%s)等于%d(%s)\n",int64(ma.GetAmount()),ma.GetCurrency().Code,int64(mb.GetAmount()),mb.GetCurrency().Code,int64(mamb.GetAmount()),mamb.GetCurrency().Code)
	}

	fmt.Println("//////////////////////////////////////////////////////////////")
	fmt.Println("不同币种的Money相加")
	mamc,err := ma.MoneyAdd(mc)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Printf("%d(%s)加%d(%s)等于%d(%s)\n",int64(ma.GetAmount()),ma.GetCurrency().Code,int64(mc.GetAmount()),mc.GetCurrency().Code,int64(mamc.GetAmount()),mamc.GetCurrency().Code)
	}


}

func TestAsMajorUnits(t *testing.T){
	var m Munit =8223372036854775855
	mm := New(m,"CNY")

	fmt.Println("//////////////////////////////////////////////////////////////")
	fmt.Println(m)
	fmt.Println("转标准单位")
	fmt.Println(mm.AsMajorUnits(true))
}

func TestThousandsSeparator(t *testing.T){
	var m Munit =8223372036854775855
	mm := New(m,"CNY")

	fmt.Println("//////////////////////////////////////////////////////////////")
	fmt.Println(m)
	fmt.Println("转标准单位后千分分隔")
	fmt.Println(mm.ThousandsSeparator(false))
}
