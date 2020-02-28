package money

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"sync"

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
	amount   Munit
	currency *Currency
	error    error

	mu sync.RWMutex
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

func (m Munit) String() string {
	return strconv.FormatInt(int64(m), 10)
}

// Decimal是否超出maxMunit与minMunit的范围
func decimalValidMunit(v decimal.Decimal) error {
	if v.GreaterThan(maxMunit.munitToDecimal()) == true || v.LessThan(minMunit.munitToDecimal()) {
		return errors.New("溢出资金范围")
	}
	return nil
}

func (m Munit) munitToDecimal() decimal.Decimal {
	return decimal.New(int64(m), 0)
}

//     -1 if m <  m2
//      0 if m == m2
//     +1 if m >  m2
//
func (m Munit) Cmp(m2 Munit) int {
	return m.munitToDecimal().Cmp(m2.munitToDecimal())
}

// m=m2
func (m Munit) Equal(m2 Munit) bool {
	return m.Cmp(m2) == 0
}

// m > m2
func (m Munit) GreaterThan(m2 Munit) bool {
	return m.Cmp(m2) == 1
}

// m >= m2
func (m Munit) GreaterThanOrEqual(m2 Munit) bool {
	cmp := m.Cmp(m2)
	return cmp == 1 || cmp == 0
}

// m < m2
func (m Munit) LessThan(m2 Munit) bool {
	return m.Cmp(m2) == -1
}

// m <= m2
func (m Munit) LessThanOrEqual(m2 Munit) bool {
	cmp := m.Cmp(m2)
	return cmp == -1 || cmp == 0
}

// 是否为0
func (m Munit) IsZero() bool {
	return m.munitToDecimal().IsZero()
}

// 绝对值
func (m Munit) Abs() (result Munit, err error) {
	if md := m.munitToDecimal(); md.IsZero() == true {
		return 0, nil
	} else {
		md = md.Abs()
		if err = decimalValidMunit(md); err != nil {
			return 0, err
		}
		result = Munit(md.IntPart())

		return result, nil
	}
}

// 负 returns -m
func (m Munit) Neg() (result Munit, err error) {
	if md := m.munitToDecimal(); md.IsZero() == true {
		return 0, nil
	} else {
		md = md.Neg()
		if err = decimalValidMunit(md); err != nil {
			return 0, err
		}
		result = Munit(md.IntPart())

		return result, nil
	}
}

// 加，如果结果超过范围限制err != nil
func (m Munit) Add(m2 Munit) (result Munit, err error) {
	md := m.munitToDecimal().Add(m2.munitToDecimal())
	if err = decimalValidMunit(md); err != nil {
		return 0, err
	}
	result = Munit(md.IntPart())

	return result, nil
}

// 减，如果结果超过范围限制err != nil
func (m Munit) Sub(m2 Munit) (result Munit, err error) {
	md := m.munitToDecimal().Sub(m2.munitToDecimal())
	if err = decimalValidMunit(md); err != nil {
		return 0, err
	}
	result = Munit(md.IntPart())

	return result, nil
}

// 乘，m * v 因为是最小单位，所以结果四舍五入返回整数
func (m Munit) Mul(v int) (result Munit, err error) {
	md := m.munitToDecimal().Mul(decimal.NewFromInt(int64(v)))
	if err = decimalValidMunit(md); err != nil {
		return 0, err
	}
	result = Munit(md.IntPart())

	return result, nil
}

// 除，m/v 因为是最小单位，所以结果四舍五入返回整数
func (m Munit) Div(v int) (result Munit, err error) {
	md := m.munitToDecimal().DivRound(decimal.NewFromInt(int64(v)), 0)
	if err = decimalValidMunit(md); err != nil {
		return 0, err
	}
	result = Munit(md.IntPart())

	return result, nil
}

func New(v Munit, code string) *Money {
	m := new(Money)
	currency := newCurrency(code).get()
	m.set(v, currency, nil)
	return m
}

func (m *Money) set(v Munit, currency *Currency, e error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.amount = v
	m.currency = currency
	m.error = e
}

func (m *Money) setAmount(v Munit) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.amount = v
}

func (m *Money) setError(e error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.error = e
}

func (m *Money) GetAmount() Munit {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.amount
}
func (m *Money) GetCurrency() *Currency {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.currency
}

func (m *Money) GetError() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.error
}

func (m *Money) MoneyAdd(m2 *Money) error {
	if m.currency.Code != m2.currency.Code {
		return errors.New("币种不同不能相加")
	}
	mm, err := m.amount.Add(m2.amount)
	if err != nil {
		return err
	}

	m.setAmount(mm)
	return nil
}

func MoneyAdd(m, m2 *Money) (*Money, error) {
	if m.currency.Code != m2.currency.Code {
		return nil, errors.New("币种不同不能计算")
	}
	mm, err := m.amount.Add(m2.amount)
	if err != nil {
		return nil, err
	}
	return &Money{
		amount:   mm,
		currency: m.currency,
		error:    nil,
	}, nil
	// nm := new(Money)
	// nm.set(mm,m.currency,nil)
	// return nm, nil
}

func (m *Money) MoneySub(m2 *Money) error {
	if m.currency.Code != m2.currency.Code {
		return errors.New("币种不同不能计算")
	}
	mm, err := m.amount.Sub(m2.amount)
	if err != nil {
		return err
	}

	m.setAmount(mm)
	return nil
}

func MoneySub(m, m2 *Money) (*Money, error) {
	if m.currency.Code != m2.currency.Code {
		return nil, errors.New("币种不同不能相加")
	}
	mm, err := m.amount.Sub(m2.amount)
	if err != nil {
		return nil, err
	}
	nm := new(Money)
	nm.set(mm, m.currency, nil)
	return nm, nil
}

// 转小数
// per int 转小数的除数值
// pr int32 保留的小数位数
// asbank bool 是否用银行的四舍五入方式取数，详见decimal.RoundBank
func (m *Money) round(per int, pr int32, asbank bool) string {
	if m.amount == 0 {
		return "0"
	}

	if pr == 0 || per == 0 {
		return strconv.FormatInt(int64(m.amount), 10)
	}

	d := m.amount.munitToDecimal().DivRound(decimal.NewFromInt(int64(per)), pr+1)
	ds := ""
	if asbank {
		ds = d.StringFixedBank(pr)
	} else {
		ds = d.StringFixed(pr)
	}
	if m.currency.DecimalMark != "." {
		return strings.Replace(ds, ".", m.currency.DecimalMark, 1)
	}
	return ds
}

// 转小数,用银行的四舍五入方式取数
// per int 转小数的除数值
// pr int32 保留的小数位数
// 详见decimal.RoundBank
func (m *Money) RoundBank(per int, pr int32) string {
	return m.round(per, pr, true)
}

// 转小数，四舍五入
// per int 转小数的除数值
// pr int32 保留的小数位数
func (m *Money) Round(per int, pr int32) string {
	return m.round(per, pr, false)
}

// 最小单位转标准单位(小数字符串)
// asbank是否用银行的四舍五入方式取数，详见decimal.RoundBank
func (m *Money) AsMajorUnits(asbank bool) string {

	pr := m.currency.precision()

	if asbank {
		return m.RoundBank(m.currency.SubUnitToUnit, pr)
	}

	return m.Round(m.currency.SubUnitToUnit, pr)
}

// 千分分隔
// asbank是否用银行的四舍五入方式取数，详见decimal.RoundBank
func (m *Money) ThousandsSeparator(asbank bool) string {
	ds := m.AsMajorUnits(asbank)
	if l := strings.LastIndex(ds, m.currency.DecimalMark); l > 0 && len(m.currency.ThousandsSeparator) > 0 {
		dsIntPart := ds[:l]
		dsDecimalPart := ds[l:]

		for i := len(dsIntPart) - 3; i > 0; i -= 3 {
			dsIntPart = dsIntPart[:i] + m.currency.ThousandsSeparator + dsIntPart[i:]
		}
		return dsIntPart + dsDecimalPart
	}
	return ds
}
