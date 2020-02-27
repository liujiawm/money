# money #
钱币的计算，不同币种的处理和显示

项目开发中...

# 安装 INSTALL #

``` $ go get github.com/liujiawm/money ```

# 例子 Examples #

```

func main() {

	var a,b,c Munit = 45896210,456789123,78964325123

	ma:=New(a,"CNY")
	mb:=New(b,"CNY")
	mc:=New(c,"BTC")

	// Munit值相加
	ab,err := a.Add(b)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(ab)
	}

	// Money相加
	mamb,err := ma.MoneyAdd(mb)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Printf("%d(%s)加%d(%s)等于%d(%s)\n",int64(ma.GetAmount()),ma.GetCurrency().Code,int64(mb.GetAmount()),mb.GetCurrency().Code,int64(mamb.GetAmount()),mamb.GetCurrency().Code)
	}

	// 转标准单位,BTC
	fmt.Println(mc.AsMajorUnits(true))

	// 转标准单位后千分分隔,CNY
	fmt.Println(ma.ThousandsSeparator(false))

}



```



# 引用库 Require #
``` github.com/shopspring/decimal ```