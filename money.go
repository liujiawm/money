package money

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// 钱的单位，指最小的整数值（SubUnit）
// 按currency值定义，如：人民币的最小单位是“分”，这个数字就表示“分”
// 因为currency内不包函分以下的单位，厘、毫之类的
// 同理比特币的最小单位是Satoshi
type Munit int64


// type Amount struct {
// 	val Munit
// }

type Money struct {
	amount Munit
	currency *Currency
}

const (
	minMunit Munit = math.MinInt64
	maxMunit Munit = math.MaxInt64
)




// const (
// 	Milli  Munit = 1 // 钱的最小单位 毫
// 	Li           = 10 * Milli // 厘
// 	Fen          = 10 * Li    // 分 Cent
// 	Jiao         = 10 * Fen   // 角 Dime
// 	Yuan         = 10 * Jiao  // 元 Dollar
// )

// Munit是否超出maxMunit与minMunit的范围
func validMunit(v Munit) error{
	if v > maxMunit || v < minMunit {
		return errors.New("溢出资金范围")
	}
	return nil
}
// Decimal是否超出maxMunit与minMunit的范围
func decimalValidMunit(v decimal.Decimal) error {
	if v.GreaterThan(maxMunit.munitToDecimal())==true || v.LessThan(minMunit.munitToDecimal()) {
		return errors.New("溢出资金范围")
	}
	return nil
}

func (m Munit)munitToDecimal() decimal.Decimal {
	return decimal.New(int64(m),0)
}


// 加，如果结果超过范围限制err != nil
func (m Munit)Add(m2 Munit) (result Munit, err error){
	md := m.munitToDecimal().Add(m2.munitToDecimal())
	if err = decimalValidMunit(md); err != nil {
		return 0,err
	}
	result = Munit(md.IntPart())

	return result,nil
}
// 减，如果结果超过范围限制err != nil
func (m Munit)Sub(m2 Munit) (result Munit, err error){
	md := m.munitToDecimal().Sub(m2.munitToDecimal())
	if err = decimalValidMunit(md); err != nil {
		return 0,err
	}
	result = Munit(md.IntPart())

	return result,nil
}

// 除，m/m2 因为是最小单位，所以结果四舍五入返回整数
func (m Munit)Div(m2 Munit) (result Munit, err error){
	md := m.munitToDecimal().DivRound(m2.munitToDecimal(),0)
	if err = decimalValidMunit(md); err != nil {
		return 0,err
	}
	result = Munit(md.IntPart())

	return result,nil
}

func New(v Munit, code string) *Money {
	return &Money{
		amount: v,
		currency: newCurrency(code).get(),
	}
}
func (m *Money)GetAmount()Munit{
	return m.amount
}
func (m *Money)GetCurrency()*Currency{
	return m.currency
}

func (m *Money)MoneyAdd(m2 *Money) (*Money,error) {
	if m.currency.Code != m2.currency.Code {
		return nil,errors.New("币种不同不能相加")
	}
	mm,err := m.amount.Add(m2.amount)
	if err != nil {
		return nil,err
	}
	return &Money{
		amount:mm,
		currency:m.currency,
	},nil

}

// 转小数
// per int 转小数的除数值
// pr int32 保留的小数位数
// asbank bool 是否用银行的四舍五入方式取数，详见decimal.RoundBank
func (m *Money)round(per int, pr int32, asbank bool)string{
	if m.amount == 0 {
		return "0"
	}

	if pr == 0 ||  per == 0 {
		return strconv.FormatInt(int64(m.amount),10)
	}

	d := m.amount.munitToDecimal().DivRound(decimal.NewFromInt(int64(per)),pr+1)
	ds := ""
	if asbank {
		ds = d.StringFixedBank(pr)
	} else {
		ds = d.StringFixed(pr)
	}
	if m.currency.DecimalMark != "." {
		return strings.Replace(ds,".",m.currency.DecimalMark,1)
	}
	return ds
}


// 转小数,用银行的四舍五入方式取数
// per int 转小数的除数值
// pr int32 保留的小数位数
// 详见decimal.RoundBank
func (m *Money)RoundBank(per int ,pr int32)string{
	return m.round(per,pr,true)
}

// 转小数，四舍五入
// per int 转小数的除数值
// pr int32 保留的小数位数
func (m *Money)Round(per int ,pr int32)string{
	return m.round(per,pr,false)
}


// 最小单位转标准单位(小数字符串)
// asbank是否用银行的四舍五入方式取数，详见decimal.RoundBank
func (m *Money)AsMajorUnits(asbank bool) string{

	pr := m.currency.precision()

	if asbank {
		return m.RoundBank(m.currency.SubUnitToUnit, pr)
	}

	return m.Round(m.currency.SubUnitToUnit, pr)
}

// 千分分隔
// asbank是否用银行的四舍五入方式取数，详见decimal.RoundBank
func (m *Money)ThousandsSeparator(asbank bool)string{
	ds := m.AsMajorUnits(asbank)
	if l := strings.LastIndex(ds,m.currency.DecimalMark); l > 0  && len(m.currency.ThousandsSeparator) > 0 {
		dsIntPart := ds[:l]
		dsDecimalPart := ds[l:]

		for i := len(dsIntPart)-3; i>0;i-=3 {
			dsIntPart = dsIntPart[:i] + m.currency.ThousandsSeparator + dsIntPart[i:]
		}
		return dsIntPart + dsDecimalPart
	}
	return ds
}
