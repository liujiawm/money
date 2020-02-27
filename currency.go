package money

import (
	"strconv"
	"strings"
)

type Currency struct {
	Code               string   // 字母代码
	IsoNumeric         int      // 数字代码
	Name               string   // 名称
	CnName             string   // 中文名称
	Symbol             string   // 标准符号
	SymbolFirst        bool     // 标准符号优先
	AlternateSymbols   []string // 备用符号
	ThousandsSeparator string   // 千位分隔符
	DecimalMark        string   // 小数分隔符
	SubUnit            string   // 辅币单位
	SubUnitToUnit      int      // 辅币进位制
	HTMLEntity         string   // HTML代码显示符号
}

var currencies = map[string]*Currency{
	"AED": {Code: "AED", IsoNumeric: 784, Name: "United Arab Emirates Dirham", CnName: "阿拉伯联合酋长国迪拉姆", Symbol: "د.إ", SymbolFirst: true, AlternateSymbols: []string{"DH", "Dhs"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fils", SubUnitToUnit: 100, HTMLEntity: ""},
	"AFN": {Code: "AFN", IsoNumeric: 971, Name: "Afghan Afghani", CnName: "阿富汗阿富汗尼", Symbol: "؋", SymbolFirst: false, AlternateSymbols: []string{"Af", "Afs"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Pul", SubUnitToUnit: 100, HTMLEntity: ""},
	"ALL": {Code: "ALL", IsoNumeric: 8, Name: "Albanian Lek", CnName: "阿尔巴尼亚勒克", Symbol: "L", SymbolFirst: false, AlternateSymbols: []string{"Lek"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Qintar", SubUnitToUnit: 100, HTMLEntity: ""},
	"AMD": {Code: "AMD", IsoNumeric: 51, Name: "Armenian Dram", CnName: "美尼亚德拉姆", Symbol: "դր.", SymbolFirst: false, AlternateSymbols: []string{"dram"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Luma", SubUnitToUnit: 100, HTMLEntity: ""},
	"ANG": {Code: "ANG", IsoNumeric: 532, Name: "Netherlands Antillean Gulden", CnName: "荷兰盾", Symbol: "ƒ", SymbolFirst: true, AlternateSymbols: []string{"NAƒ", "NAf", "f"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x0192;"},
	"AOA": {Code: "AOA", IsoNumeric: 973, Name: "Angolan Kwanza", CnName: "安哥拉宽扎", Symbol: "Kz", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cêntimo", SubUnitToUnit: 100, HTMLEntity: ""},
	"ARS": {Code: "ARS", IsoNumeric: 32, Name: "Argentine Peso", CnName: "阿根廷比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"$m/n", "m$n"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "&#x20B1;"},
	"AUD": {Code: "AUD", IsoNumeric: 36, Name: "Australian Dollar", CnName: "澳大利亚元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"A$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"AWG": {Code: "AWG", IsoNumeric: 533, Name: "Aruban Florin", CnName: "阿鲁巴或荷兰盾", Symbol: "ƒ", SymbolFirst: false, AlternateSymbols: []string{"Afl"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x0192;"},
	"AZN": {Code: "AZN", IsoNumeric: 944, Name: "Azerbaijani Manat", CnName: "阿塞拜疆新马纳特", Symbol: "₼", SymbolFirst: true, AlternateSymbols: []string{"m", "man"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Qəpik", SubUnitToUnit: 100, HTMLEntity: ""},
	"BAM": {Code: "BAM", IsoNumeric: 977, Name: "Bosnia and Herzegovina Convertible Mark", CnName: "波斯尼亚兑换马尔卡", Symbol: "КМ", SymbolFirst: true, AlternateSymbols: []string{"KM"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fening", SubUnitToUnit: 100, HTMLEntity: ""},
	"BBD": {Code: "BBD", IsoNumeric: 52, Name: "Barbadian Dollar", CnName: "巴巴多斯元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"Bds$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"BDT": {Code: "BDT", IsoNumeric: 50, Name: "Bangladeshi Taka", CnName: "孟加拉国塔卡", Symbol: "৳", SymbolFirst: true, AlternateSymbols: []string{"Tk"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Paisa", SubUnitToUnit: 100, HTMLEntity: ""},
	"BGN": {Code: "BGN", IsoNumeric: 975, Name: "Bulgarian Lev", CnName: "保加利亚利瓦", Symbol: "лв", SymbolFirst: false, AlternateSymbols: []string{"lev", "leva", "лев", "лева"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Stotinka", SubUnitToUnit: 100, HTMLEntity: ""},
	"BHD": {Code: "BHD", IsoNumeric: 48, Name: "Bahraini Dinar", CnName: "巴林第纳尔", Symbol: "ب.د", SymbolFirst: true, AlternateSymbols: []string{"BD"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fils", SubUnitToUnit: 1000, HTMLEntity: ""},
	"BIF": {Code: "BIF", IsoNumeric: 108, Name: "Burundian Franc", CnName: "布隆迪法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"FBu"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"BMD": {Code: "BMD", IsoNumeric: 60, Name: "Bermudian Dollar", CnName: "百慕大元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"BD$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"BND": {Code: "BND", IsoNumeric: 96, Name: "Brunei Dollar", CnName: "文莱元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"B$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Sen", SubUnitToUnit: 100, HTMLEntity: "$"},
	"BOB": {Code: "BOB", IsoNumeric: 68, Name: "Bolivian Boliviano", CnName: "玻利维亚币", Symbol: "Bs.", SymbolFirst: true, AlternateSymbols: []string{"Bs"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"BRL": {Code: "BRL", IsoNumeric: 986, Name: "Brazilian Real", CnName: "巴西雷亚尔", Symbol: "R$", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "R$"},
	"BSD": {Code: "BSD", IsoNumeric: 44, Name: "Bahamian Dollar", CnName: "巴哈马元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"B$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"BTC": {Code: "BTC", IsoNumeric: 0, Name: "Bitcoin", CnName: "比特币", Symbol: "B⃦", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Satoshi", SubUnitToUnit: 100000000, HTMLEntity: ""},
	"BTN": {Code: "BTN", IsoNumeric: 64, Name: "Bhutanese Ngultrum", CnName: "不丹努尔特鲁姆", Symbol: "Nu.", SymbolFirst: false, AlternateSymbols: []string{"Nu"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Chertrum", SubUnitToUnit: 100, HTMLEntity: ""},
	"BWP": {Code: "BWP", IsoNumeric: 72, Name: "Botswana Pula", CnName: "博茨瓦纳普拉", Symbol: "P", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Thebe", SubUnitToUnit: 100, HTMLEntity: ""},
	"BYR": {Code: "BYR", IsoNumeric: 974, Name: "Belarusian Ruble", CnName: "白俄罗斯卢布", Symbol: "Br", SymbolFirst: false, AlternateSymbols: []string{""}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Kapyeyka", SubUnitToUnit: 100, HTMLEntity: ""},
	"BZD": {Code: "BZD", IsoNumeric: 84, Name: "Belize Dollar", CnName: "伯利兹元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"BZ$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"CAD": {Code: "CAD", IsoNumeric: 124, Name: "Canadian Dollar", CnName: "加拿大元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"C$", "CAD$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"CDF": {Code: "CDF", IsoNumeric: 976, Name: "Congolese Franc", CnName: "刚果法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"FC"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"CHF": {Code: "CHF", IsoNumeric: 756, Name: "Swiss Franc", CnName: "瑞士法郎", Symbol: "Fr", SymbolFirst: true, AlternateSymbols: []string{"SFr", "CHF"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Rappen", SubUnitToUnit: 100, HTMLEntity: ""},
	"CLF": {Code: "CLF", IsoNumeric: 990, Name: "Unidad de Fomento", CnName: "", Symbol: "UF", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Peso", SubUnitToUnit: 1, HTMLEntity: "&#x20B1;"},
	"CLP": {Code: "CLP", IsoNumeric: 152, Name: "Chilean Peso", CnName: "智利比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Peso", SubUnitToUnit: 100, HTMLEntity: "&#36;"},
	"CNY": {Code: "CNY", IsoNumeric: 156, Name: "Chinese Renminbi Yuan", CnName: "人民币", Symbol: "¥", SymbolFirst: true, AlternateSymbols: []string{"元", "CN元", "CN¥"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fen", SubUnitToUnit: 100, HTMLEntity: "￥"},
	"COP": {Code: "COP", IsoNumeric: 170, Name: "Colombian Peso", CnName: "哥伦比亚比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"COL$"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "&#x20B1;"},
	"CRC": {Code: "CRC", IsoNumeric: 188, Name: "Costa Rican Colón", CnName: "哥斯达黎加科朗", Symbol: "₡", SymbolFirst: true, AlternateSymbols: []string{"¢"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Céntimo", SubUnitToUnit: 100, HTMLEntity: "&#x20A1;"},
	"CUC": {Code: "CUC", IsoNumeric: 931, Name: "Cuban Convertible Peso", CnName: "古巴可兑换比索", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"CUC$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"CUP": {Code: "CUP", IsoNumeric: 192, Name: "Cuban Peso", CnName: "古巴比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"$MN"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "&#x20B1;"},
	"CVE": {Code: "CVE", IsoNumeric: 132, Name: "Cape Verdean Escudo", CnName: "佛得角埃斯库多", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"Esc"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"CZK": {Code: "CZK", IsoNumeric: 203, Name: "Czech Koruna", CnName: "捷克共和国克朗", Symbol: "Kč", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Haléř", SubUnitToUnit: 100, HTMLEntity: ""},
	"DJF": {Code: "DJF", IsoNumeric: 262, Name: "Djiboutian Franc", CnName: "吉布提法郎", Symbol: "Fdj", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"DKK": {Code: "DKK", IsoNumeric: 208, Name: "Danish Krone", CnName: "丹麦克朗", Symbol: "kr", SymbolFirst: false, AlternateSymbols: []string{",-"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Øre", SubUnitToUnit: 100, HTMLEntity: ""},
	"DOP": {Code: "DOP", IsoNumeric: 214, Name: "Dominican Peso", CnName: "多米尼加比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"RD$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "&#x20B1;"},
	"DZD": {Code: "DZD", IsoNumeric: 12, Name: "Algerian Dinar", CnName: "阿尔及利亚第纳尔", Symbol: "د.ج", SymbolFirst: false, AlternateSymbols: []string{"DA"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"EEK": {Code: "EEK", IsoNumeric: 233, Name: "Estonian Kroon", CnName: "爱沙尼亚克伦尼", Symbol: "KR", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Sent", SubUnitToUnit: 100, HTMLEntity: ""},
	"EGP": {Code: "EGP", IsoNumeric: 818, Name: "Egyptian Pound", CnName: "埃及镑", Symbol: "ج.م", SymbolFirst: true, AlternateSymbols: []string{"LE", "E£", "L.E."}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Piastre", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"ERN": {Code: "ERN", IsoNumeric: 232, Name: "Eritrean Nakfa", CnName: "厄立特里亚纳克法", Symbol: "Nfk", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"ETB": {Code: "ETB", IsoNumeric: 230, Name: "Ethiopian Birr", CnName: "埃塞俄比亚比尔", Symbol: "Br", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Santim", SubUnitToUnit: 100, HTMLEntity: ""},
	"EUR": {Code: "EUR", IsoNumeric: 978, Name: "Euro", CnName: "欧元", Symbol: "€", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x20AC;"},
	"FJD": {Code: "FJD", IsoNumeric: 242, Name: "Fijian Dollar", CnName: "斐济元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"FJ$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"FKP": {Code: "FKP", IsoNumeric: 238, Name: "Falkland Pound", CnName: "福克兰岛磅", Symbol: "£", SymbolFirst: false, AlternateSymbols: []string{"FK£"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Penny", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"GBP": {Code: "GBP", IsoNumeric: 826, Name: "British Pound", CnName: "英镑", Symbol: "£", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Penny", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"GEL": {Code: "GEL", IsoNumeric: 981, Name: "Georgian Lari", CnName: "格鲁吉亚拉里", Symbol: "ლ", SymbolFirst: false, AlternateSymbols: []string{"lari"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Tetri", SubUnitToUnit: 100, HTMLEntity: ""},
	"GHS": {Code: "GHS", IsoNumeric: 936, Name: "Ghanaian Cedi", CnName: "加纳塞地", Symbol: "₵", SymbolFirst: true, AlternateSymbols: []string{"GH¢", "GH₵"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Pesewa", SubUnitToUnit: 100, HTMLEntity: "&#x20B5;"},
	"GIP": {Code: "GIP", IsoNumeric: 292, Name: "Gibraltar Pound", CnName: "直布罗陀镑", Symbol: "£", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Penny", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"GMD": {Code: "GMD", IsoNumeric: 270, Name: "Gambian Dalasi", CnName: "冈比亚法拉西", Symbol: "D", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Butut", SubUnitToUnit: 100, HTMLEntity: ""},
	"GNF": {Code: "GNF", IsoNumeric: 324, Name: "Guinean Franc", CnName: "几内亚法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"FG", "GFr"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"GTQ": {Code: "GTQ", IsoNumeric: 320, Name: "Guatemalan Quetzal", CnName: "危地马拉格查尔", Symbol: "Q", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"GYD": {Code: "GYD", IsoNumeric: 328, Name: "Guyanese Dollar", CnName: "圭亚那元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"G$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"HKD": {Code: "HKD", IsoNumeric: 344, Name: "Hong Kong Dollar", CnName: "港元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"HK$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"HNL": {Code: "HNL", IsoNumeric: 340, Name: "Honduran Lempira", CnName: "洪都拉斯伦皮拉", Symbol: "L", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"HRK": {Code: "HRK", IsoNumeric: 191, Name: "Croatian Kuna", CnName: "克罗地亚库纳", Symbol: "kn", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Lipa", SubUnitToUnit: 100, HTMLEntity: ""},
	"HTG": {Code: "HTG", IsoNumeric: 332, Name: "Haitian Gourde", CnName: "海地古德", Symbol: "G", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"HUF": {Code: "HUF", IsoNumeric: 348, Name: "Hungarian Forint", CnName: "匈牙利福林", Symbol: "Ft", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Fillér", SubUnitToUnit: 100, HTMLEntity: ""},
	"IDR": {Code: "IDR", IsoNumeric: 360, Name: "Indonesian Rupiah", CnName: "印尼卢比", Symbol: "Rp", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Sen", SubUnitToUnit: 100, HTMLEntity: ""},
	"ILS": {Code: "ILS", IsoNumeric: 376, Name: "Israeli New Sheqel", CnName: "以色列新谢克尔", Symbol: "₪", SymbolFirst: true, AlternateSymbols: []string{"ש״ח", "NIS"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Agora", SubUnitToUnit: 100, HTMLEntity: "&#x20AA;"},
	"INR": {Code: "INR", IsoNumeric: 356, Name: "Indian Rupee", CnName: "印度卢比", Symbol: "₹", SymbolFirst: true, AlternateSymbols: []string{"Rs", "৳", "૱", "௹", "रु", "₨"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Paisa", SubUnitToUnit: 100, HTMLEntity: "&#x20b9;"},
	"IQD": {Code: "IQD", IsoNumeric: 368, Name: "Iraqi Dinar", CnName: "伊拉克第纳尔", Symbol: "ع.د", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fils", SubUnitToUnit: 1000, HTMLEntity: ""},
	"IRR": {Code: "IRR", IsoNumeric: 364, Name: "Iranian Rial", CnName: "伊朗里亚尔", Symbol: "﷼", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Dinar", SubUnitToUnit: 100, HTMLEntity: "&#xFDFC;"},
	"ISK": {Code: "ISK", IsoNumeric: 352, Name: "Icelandic Króna", CnName: "冰岛克朗", Symbol: "kr", SymbolFirst: true, AlternateSymbols: []string{"Íkr"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Eyrir", SubUnitToUnit: 100, HTMLEntity: ""},
	"JEP": {Code: "JEP", IsoNumeric: 0, Name: "Jersey Pound", CnName: "泽西镑", Symbol: "£", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Penny", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"JMD": {Code: "JMD", IsoNumeric: 388, Name: "Jamaican Dollar", CnName: "牙买加元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"J$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"JOD": {Code: "JOD", IsoNumeric: 400, Name: "Jordanian Dinar", CnName: "约旦第纳尔", Symbol: "د.ا", SymbolFirst: true, AlternateSymbols: []string{"JD"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Piastre", SubUnitToUnit: 100, HTMLEntity: ""},
	"JPY": {Code: "JPY", IsoNumeric: 392, Name: "Japanese Yen", CnName: "日本日圆", Symbol: "¥", SymbolFirst: true, AlternateSymbols: []string{"円", "圓"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "", SubUnitToUnit: 1, HTMLEntity: "&#x00A5;"},
	"KES": {Code: "KES", IsoNumeric: 404, Name: "Kenyan Shilling", CnName: "肯尼亚先令", Symbol: "KSh", SymbolFirst: true, AlternateSymbols: []string{"Sh"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"KGS": {Code: "KGS", IsoNumeric: 417, Name: "Kyrgyzstani Som", CnName: "吉尔吉斯斯坦索姆", Symbol: "som", SymbolFirst: false, AlternateSymbols: []string{"сом"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Tyiyn", SubUnitToUnit: 100, HTMLEntity: ""},
	"KHR": {Code: "KHR", IsoNumeric: 116, Name: "Cambodian Riel", CnName: "柬埔寨瑞尔", Symbol: "៛", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Sen", SubUnitToUnit: 100, HTMLEntity: "&#x17DB;"},
	"KMF": {Code: "KMF", IsoNumeric: 174, Name: "Comorian Franc", CnName: "科摩罗法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"CF"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"KPW": {Code: "KPW", IsoNumeric: 408, Name: "North Korean Won", CnName: "朝鲜圆", Symbol: "₩", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Chŏn", SubUnitToUnit: 100, HTMLEntity: "&#x20A9;"},
	"KRW": {Code: "KRW", IsoNumeric: 410, Name: "South Korean Won", CnName: "韩国圆", Symbol: "₩", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "", SubUnitToUnit: 100, HTMLEntity: "&#x20A9;"},
	"KWD": {Code: "KWD", IsoNumeric: 414, Name: "Kuwaiti Dinar", CnName: "科威特第纳尔", Symbol: "د.ك", SymbolFirst: true, AlternateSymbols: []string{"K.D."}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fils", SubUnitToUnit: 1000, HTMLEntity: ""},
	"KYD": {Code: "KYD", IsoNumeric: 136, Name: "Cayman Islands Dollar", CnName: "开曼岛元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"CI$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"KZT": {Code: "KZT", IsoNumeric: 398, Name: "Kazakhstani Tenge", CnName: "哈萨克斯坦坚戈", Symbol: "〒", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Tiyn", SubUnitToUnit: 100, HTMLEntity: ""},
	"LAK": {Code: "LAK", IsoNumeric: 418, Name: "Lao Kip", CnName: "老挝或老挝基普", Symbol: "₭", SymbolFirst: false, AlternateSymbols: []string{"₭N"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Att", SubUnitToUnit: 100, HTMLEntity: "&#x20AD;"},
	"LBP": {Code: "LBP", IsoNumeric: 422, Name: "Lebanese Pound", CnName: "黎巴嫩镑", Symbol: "ل.ل", SymbolFirst: true, AlternateSymbols: []string{"£", "L£"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Piastre", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"LKR": {Code: "LKR", IsoNumeric: 144, Name: "Sri Lankan Rupee", CnName: "斯里兰卡卢比", Symbol: "₨", SymbolFirst: false, AlternateSymbols: []string{"රු", "ரூ", "SLRs", "/-"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x0BF9;"},
	"LRD": {Code: "LRD", IsoNumeric: 430, Name: "Liberian Dollar", CnName: "利比里亚元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"L$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"LSL": {Code: "LSL", IsoNumeric: 426, Name: "Lesotho Loti", CnName: "巴索托洛蒂", Symbol: "L", SymbolFirst: false, AlternateSymbols: []string{"M"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Sente", SubUnitToUnit: 100, HTMLEntity: ""},
	"LTL": {Code: "LTL", IsoNumeric: 440, Name: "Lithuanian Litas", CnName: "立陶宛里塔", Symbol: "Lt", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centas", SubUnitToUnit: 100, HTMLEntity: ""},
	"LVL": {Code: "LVL", IsoNumeric: 428, Name: "Latvian Lats", CnName: "拉脱维亚拉特", Symbol: "Ls", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Santīms", SubUnitToUnit: 100, HTMLEntity: ""},
	"LYD": {Code: "LYD", IsoNumeric: 434, Name: "Libyan Dinar", CnName: "利比亚第纳尔", Symbol: "ل.د", SymbolFirst: false, AlternateSymbols: []string{"LD"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Dirham", SubUnitToUnit: 1000, HTMLEntity: ""},
	"MAD": {Code: "MAD", IsoNumeric: 504, Name: "Moroccan Dirham", CnName: "摩洛哥迪拉姆", Symbol: "د.م.", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"MDL": {Code: "MDL", IsoNumeric: 498, Name: "Moldovan Leu", CnName: "摩尔多瓦列伊", Symbol: "L", SymbolFirst: false, AlternateSymbols: []string{"lei"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Ban", SubUnitToUnit: 100, HTMLEntity: ""},
	"MGA": {Code: "MGA", IsoNumeric: 969, Name: "Malagasy Ariary", CnName: "马达加斯加", Symbol: "Ar", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Iraimbilanja", SubUnitToUnit: 5, HTMLEntity: ""},
	"MKD": {Code: "MKD", IsoNumeric: 807, Name: "Macedonian Denar", CnName: "马其顿代纳尔", Symbol: "ден", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Deni", SubUnitToUnit: 100, HTMLEntity: ""},
	"MMK": {Code: "MMK", IsoNumeric: 104, Name: "Myanmar Kyat", CnName: "缅甸缅元", Symbol: "K", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Pya", SubUnitToUnit: 100, HTMLEntity: ""},
	"MNT": {Code: "MNT", IsoNumeric: 496, Name: "Mongolian Tögrög", CnName: "蒙古图格里克", Symbol: "₮", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Möngö", SubUnitToUnit: 100, HTMLEntity: "&#x20AE;"},
	"MOP": {Code: "MOP", IsoNumeric: 446, Name: "Macanese Pataca", CnName: "澳门币", Symbol: "P", SymbolFirst: false, AlternateSymbols: []string{"MOP$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Avo", SubUnitToUnit: 100, HTMLEntity: ""},
	"MRO": {Code: "MRO", IsoNumeric: 478, Name: "Mauritanian Ouguiya", CnName: "毛里塔尼亚乌吉亚", Symbol: "UM", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Khoums", SubUnitToUnit: 5, HTMLEntity: ""},
	"MTL": {Code: "MTL", IsoNumeric: 470, Name: "Maltese Lira", CnName: "马尔他里拉", Symbol: "₤", SymbolFirst: true, AlternateSymbols: []string{"Lm"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"MUR": {Code: "MUR", IsoNumeric: 480, Name: "Mauritian Rupee", CnName: "毛里求斯卢比", Symbol: "₨", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x20A8;"},
	"MVR": {Code: "MVR", IsoNumeric: 462, Name: "Maldivian Rufiyaa", CnName: "马尔代夫卢比", Symbol: "MVR", SymbolFirst: false, AlternateSymbols: []string{"MRF", "Rf", "/-", "ރ"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Laari", SubUnitToUnit: 100, HTMLEntity: ""},
	"MWK": {Code: "MWK", IsoNumeric: 454, Name: "Malawian Kwacha", CnName: "马拉维克瓦查", Symbol: "MK", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Tambala", SubUnitToUnit: 100, HTMLEntity: ""},
	"MXN": {Code: "MXN", IsoNumeric: 484, Name: "Mexican Peso", CnName: "墨西哥比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"MEX$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "$"},
	"MYR": {Code: "MYR", IsoNumeric: 458, Name: "Malaysian Ringgit", CnName: "马来西亚林吉特", Symbol: "RM", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Sen", SubUnitToUnit: 100, HTMLEntity: ""},
	"MZN": {Code: "MZN", IsoNumeric: 943, Name: "Mozambican Metical", CnName: "莫桑比克梅蒂卡尔", Symbol: "MTn", SymbolFirst: true, AlternateSymbols: []string{"MZN"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"NAD": {Code: "NAD", IsoNumeric: 516, Name: "Namibian Dollar", CnName: "纳米比亚元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"N$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"NGN": {Code: "NGN", IsoNumeric: 566, Name: "Nigerian Naira", CnName: "尼日利亚奈拉", Symbol: "₦", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Kobo", SubUnitToUnit: 100, HTMLEntity: "&#x20A6;"},
	"NIO": {Code: "NIO", IsoNumeric: 558, Name: "Nicaraguan Córdoba", CnName: "尼加拉瓜科多巴", Symbol: "C$", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: ""},
	"NOK": {Code: "NOK", IsoNumeric: 578, Name: "Norwegian Krone", CnName: "挪威克朗", Symbol: "kr", SymbolFirst: false, AlternateSymbols: []string{",-"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Øre", SubUnitToUnit: 100, HTMLEntity: "kr"},
	"NPR": {Code: "NPR", IsoNumeric: 524, Name: "Nepalese Rupee", CnName: "尼泊尔卢比", Symbol: "₨", SymbolFirst: true, AlternateSymbols: []string{"Rs", "रू"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Paisa", SubUnitToUnit: 100, HTMLEntity: "&#x20A8;"},
	"NZD": {Code: "NZD", IsoNumeric: 554, Name: "New Zealand Dollar", CnName: "纽西兰元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"NZ$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"OMR": {Code: "OMR", IsoNumeric: 512, Name: "Omani Rial", CnName: "阿曼里亚尔", Symbol: "ر.ع.", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Baisa", SubUnitToUnit: 1000, HTMLEntity: "&#xFDFC;"},
	"PAB": {Code: "PAB", IsoNumeric: 590, Name: "Panamanian Balboa", CnName: "巴拿马巴波亚", Symbol: "B/.", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centésimo", SubUnitToUnit: 100, HTMLEntity: ""},
	"PEN": {Code: "PEN", IsoNumeric: 604, Name: "Peruvian Nuevo Sol", CnName: "秘鲁索尔", Symbol: "S/.", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Céntimo", SubUnitToUnit: 100, HTMLEntity: "S/."},
	"PGK": {Code: "PGK", IsoNumeric: 598, Name: "Papua New Guinean Kina", CnName: "巴布亚新几内亚基纳", Symbol: "K", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Toea", SubUnitToUnit: 100, HTMLEntity: ""},
	"PHP": {Code: "PHP", IsoNumeric: 608, Name: "Philippine Peso", CnName: "菲律宾比索", Symbol: "₱", SymbolFirst: true, AlternateSymbols: []string{"PHP", "PhP", "P"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "&#x20B1;"},
	"PKR": {Code: "PKR", IsoNumeric: 586, Name: "Pakistani Rupee", CnName: "巴基斯坦卢比", Symbol: "₨", SymbolFirst: true, AlternateSymbols: []string{"Rs"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Paisa", SubUnitToUnit: 100, HTMLEntity: "&#x20A8;"},
	"PLN": {Code: "PLN", IsoNumeric: 985, Name: "Polish Złoty", CnName: "波兰兹罗提", Symbol: "zł", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: " ", DecimalMark: ",", SubUnit: "Grosz", SubUnitToUnit: 100, HTMLEntity: "z&#322;"},
	"PYG": {Code: "PYG", IsoNumeric: 600, Name: "Paraguayan Guaraní", CnName: "巴拉圭瓜拉尼", Symbol: "₲", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Céntimo", SubUnitToUnit: 100, HTMLEntity: "&#x20B2;"},
	"QAR": {Code: "QAR", IsoNumeric: 634, Name: "Qatari Riyal", CnName: "卡塔尔里亚尔", Symbol: "ر.ق", SymbolFirst: false, AlternateSymbols: []string{"QR"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Dirham", SubUnitToUnit: 100, HTMLEntity: "&#xFDFC;"},
	"RON": {Code: "RON", IsoNumeric: 946, Name: "Romanian Leu", CnName: "罗马尼亚列伊", Symbol: "Lei", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Bani", SubUnitToUnit: 100, HTMLEntity: ""},
	"RSD": {Code: "RSD", IsoNumeric: 941, Name: "Serbian Dinar", CnName: "", Symbol: "РСД", SymbolFirst: true, AlternateSymbols: []string{"RSD", "din", "дин"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Para", SubUnitToUnit: 100, HTMLEntity: ""},
	"RUB": {Code: "RUB", IsoNumeric: 643, Name: "Russian Ruble", CnName: "俄罗斯卢布", Symbol: "₽", SymbolFirst: false, AlternateSymbols: []string{"руб.", "р."}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Kopeck", SubUnitToUnit: 100, HTMLEntity: "&#x20BD;"},
	"RWF": {Code: "RWF", IsoNumeric: 646, Name: "Rwandan Franc", CnName: "卢旺达法郎", Symbol: "FRw", SymbolFirst: false, AlternateSymbols: []string{"RF", "R₣"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"SAR": {Code: "SAR", IsoNumeric: 682, Name: "Saudi Riyal", CnName: "亚尔", Symbol: "ر.س", SymbolFirst: true, AlternateSymbols: []string{"SR", "﷼"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Hallallah", SubUnitToUnit: 100, HTMLEntity: "&#xFDFC;"},
	"SBD": {Code: "SBD", IsoNumeric: 90, Name: "Solomon Islands Dollar", CnName: "所罗门元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"SI$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"SCR": {Code: "SCR", IsoNumeric: 690, Name: "Seychellois Rupee", CnName: "塞舌尔卢比", Symbol: "₨", SymbolFirst: false, AlternateSymbols: []string{"SRe", "SR"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x20A8;"},
	"SDG": {Code: "SDG", IsoNumeric: 938, Name: "Sudanese Pound", CnName: "苏丹镑", Symbol: "£", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Piastre", SubUnitToUnit: 100, HTMLEntity: ""},
	"SEK": {Code: "SEK", IsoNumeric: 752, Name: "Swedish Krona", CnName: "瑞典克朗", Symbol: "kr", SymbolFirst: false, AlternateSymbols: []string{":-"}, ThousandsSeparator: " ", DecimalMark: ",", SubUnit: "Öre", SubUnitToUnit: 100, HTMLEntity: ""},
	"SGD": {Code: "SGD", IsoNumeric: 702, Name: "Singapore Dollar", CnName: "新加坡元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"S$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"SHP": {Code: "SHP", IsoNumeric: 654, Name: "Saint Helenian Pound", CnName: "圣赫勒宁镑", Symbol: "£", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Penny", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"SKK": {Code: "SKK", IsoNumeric: 703, Name: "Slovak Koruna", CnName: "斯洛伐克克朗", Symbol: "Sk", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Halier", SubUnitToUnit: 100, HTMLEntity: ""},
	"SLL": {Code: "SLL", IsoNumeric: 694, Name: "Sierra Leonean Leone", CnName: "塞拉里昂利昂", Symbol: "Le", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"SOS": {Code: "SOS", IsoNumeric: 706, Name: "Somali Shilling", CnName: "索马里先令", Symbol: "Sh", SymbolFirst: false, AlternateSymbols: []string{"Sh.So"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"SRD": {Code: "SRD", IsoNumeric: 968, Name: "Surinamese Dollar", CnName: "苏里南元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"SSP": {Code: "SSP", IsoNumeric: 728, Name: "South Sudanese Pound", CnName: "南苏丹镑", Symbol: "£", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "piaster", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"STD": {Code: "STD", IsoNumeric: 678, Name: "São Tomé and Príncipe Dobra", CnName: "", Symbol: "Db", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cêntimo", SubUnitToUnit: 100, HTMLEntity: ""},
	"SVC": {Code: "SVC", IsoNumeric: 222, Name: "Salvadoran Colón", CnName: "萨尔瓦多科朗", Symbol: "₡", SymbolFirst: true, AlternateSymbols: []string{"¢"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centavo", SubUnitToUnit: 100, HTMLEntity: "&#x20A1;"},
	"SYP": {Code: "SYP", IsoNumeric: 760, Name: "Syrian Pound", CnName: "叙利亚镑", Symbol: "£S", SymbolFirst: false, AlternateSymbols: []string{"£", "ل.س", "LS", "الليرة السورية"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Piastre", SubUnitToUnit: 100, HTMLEntity: "&#x00A3;"},
	"SZL": {Code: "SZL", IsoNumeric: 748, Name: "Swazi Lilangeni", CnName: "史瓦济兰里兰吉尼", Symbol: "L", SymbolFirst: true, AlternateSymbols: []string{"E"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"THB": {Code: "THB", IsoNumeric: 764, Name: "Thai Baht", CnName: "泰铢", Symbol: "฿", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Satang", SubUnitToUnit: 100, HTMLEntity: "&#x0E3F;"},
	"TJS": {Code: "TJS", IsoNumeric: 972, Name: "Tajikistani Somoni", CnName: "塔吉克索莫尼", Symbol: "ЅМ", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Diram", SubUnitToUnit: 100, HTMLEntity: ""},
	"TMT": {Code: "TMT", IsoNumeric: 934, Name: "Turkmenistani Manat", CnName: "土库曼斯坦马纳特", Symbol: "T", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Tenge", SubUnitToUnit: 100, HTMLEntity: ""},
	"TND": {Code: "TND", IsoNumeric: 788, Name: "Tunisian Dinar", CnName: "突尼斯第纳尔", Symbol: "د.ت", SymbolFirst: false, AlternateSymbols: []string{"TD", "DT"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Millime", SubUnitToUnit: 1000, HTMLEntity: ""},
	"TOP": {Code: "TOP", IsoNumeric: 776, Name: "Tongan Paʻanga", CnName: "", Symbol: "T$", SymbolFirst: true, AlternateSymbols: []string{"PT"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Seniti", SubUnitToUnit: 100, HTMLEntity: ""},
	"TRY": {Code: "TRY", IsoNumeric: 949, Name: "Turkish Lira", CnName: "土耳其里拉", Symbol: "₺", SymbolFirst: false, AlternateSymbols: []string{"TL"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "kuruş", SubUnitToUnit: 100, HTMLEntity: ""},
	"TTD": {Code: "TTD", IsoNumeric: 780, Name: "Trinidad and Tobago Dollar", CnName: "特立尼达多巴哥元", Symbol: "$", SymbolFirst: false, AlternateSymbols: []string{"TT$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"TWD": {Code: "TWD", IsoNumeric: 901, Name: "New Taiwan Dollar", CnName: "新台币", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"NT$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"TZS": {Code: "TZS", IsoNumeric: 834, Name: "Tanzanian Shilling", CnName: "坦桑尼亚先令", Symbol: "Sh", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"UAH": {Code: "UAH", IsoNumeric: 980, Name: "Ukrainian Hryvnia", CnName: "乌克兰格里夫纳", Symbol: "₴", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Kopiyka", SubUnitToUnit: 100, HTMLEntity: "&#x20B4;"},
	"UGX": {Code: "UGX", IsoNumeric: 800, Name: "Ugandan Shilling", CnName: "乌干达先令", Symbol: "USh", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: ""},
	"USD": {Code: "USD", IsoNumeric: 840, Name: "United States Dollar", CnName: "美元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"US$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"UYU": {Code: "UYU", IsoNumeric: 858, Name: "Uruguayan Peso", CnName: "乌拉圭比索", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"$U"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Centésimo", SubUnitToUnit: 100, HTMLEntity: "&#x20B1;"},
	"UZS": {Code: "UZS", IsoNumeric: 860, Name: "Uzbekistani Som", CnName: "乌兹别克索姆", Symbol: "", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Tiyin", SubUnitToUnit: 100, HTMLEntity: ""},
	"VEF": {Code: "VEF", IsoNumeric: 937, Name: "Venezuelan Bolívar", CnName: "委内瑞拉博利瓦", Symbol: "Bs F", SymbolFirst: true, AlternateSymbols: []string{"Bs.F", "Bs"}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Céntimo", SubUnitToUnit: 100, HTMLEntity: ""},
	"VND": {Code: "VND", IsoNumeric: 704, Name: "Vietnamese Đồng", CnName: "越南盾", Symbol: "₫", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ".", DecimalMark: ",", SubUnit: "Hào", SubUnitToUnit: 1, HTMLEntity: "&#x20AB;"},
	"VUV": {Code: "VUV", IsoNumeric: 548, Name: "Vanuatu Vatu", CnName: "瓦努阿图瓦图", Symbol: "Vt", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "", SubUnitToUnit: 1, HTMLEntity: ""},
	"WST": {Code: "WST", IsoNumeric: 882, Name: "Samoan Tala", CnName: "萨摩亚塔拉", Symbol: "T", SymbolFirst: false, AlternateSymbols: []string{"WS$", "SAT", "ST"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Sene", SubUnitToUnit: 100, HTMLEntity: ""},
	"XAF": {Code: "XAF", IsoNumeric: 950, Name: "Central African Cfa Franc", CnName: "中非金融合作法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"FCFA"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"XAG": {Code: "XAG", IsoNumeric: 961, Name: "Silver (Troy Ounce)", CnName: "", Symbol: "oz t", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "oz", SubUnitToUnit: 1, HTMLEntity: ""},
	"XAU": {Code: "XAU", IsoNumeric: 959, Name: "Gold (Troy Ounce)", CnName: "", Symbol: "oz t", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "oz", SubUnitToUnit: 1, HTMLEntity: ""},
	"XCD": {Code: "XCD", IsoNumeric: 951, Name: "East Caribbean Dollar", CnName: "东加勒比元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"EC$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"XDR": {Code: "XDR", IsoNumeric: 960, Name: "Special Drawing Rights", CnName: "", Symbol: "SDR", SymbolFirst: false, AlternateSymbols: []string{"XDR"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "", SubUnitToUnit: 1, HTMLEntity: "$"},
	"XOF": {Code: "XOF", IsoNumeric: 952, Name: "West African Cfa Franc", CnName: "非共体法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"CFA"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"XPF": {Code: "XPF", IsoNumeric: 953, Name: "Cfp Franc", CnName: "太平洋法郎", Symbol: "Fr", SymbolFirst: false, AlternateSymbols: []string{"F"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Centime", SubUnitToUnit: 100, HTMLEntity: ""},
	"YER": {Code: "YER", IsoNumeric: 886, Name: "Yemeni Rial", CnName: "也门里亚尔", Symbol: "﷼", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Fils", SubUnitToUnit: 100, HTMLEntity: "&#xFDFC;"},
	"ZAR": {Code: "ZAR", IsoNumeric: 710, Name: "South African Rand", CnName: "南非兰特", Symbol: "R", SymbolFirst: true, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "&#x0052;"},
	"ZMK": {Code: "ZMK", IsoNumeric: 894, Name: "Zambian Kwacha", CnName: "赞比亚克瓦查", Symbol: "ZK", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Ngwee", SubUnitToUnit: 100, HTMLEntity: ""},
	"ZMW": {Code: "ZMW", IsoNumeric: 967, Name: "Zambian Kwacha", CnName: "赞比亚克瓦查", Symbol: "ZK", SymbolFirst: false, AlternateSymbols: []string{}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Ngwee", SubUnitToUnit: 100, HTMLEntity: ""},
	"ZWD": {Code: "ZWD", IsoNumeric: 716, Name: "Zimbabwean Dollar", CnName: "津巴布韦元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"Z$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"ZWL": {Code: "ZWL", IsoNumeric: 932, Name: "Zimbabwean Dollar", CnName: "津巴布韦元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"Z$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"ZWN": {Code: "ZWN", IsoNumeric: 942, Name: "Zimbabwean Dollar", CnName: "津巴布韦元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"Z$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
	"ZWR": {Code: "ZWR", IsoNumeric: 935, Name: "Zimbabwean Dollar", CnName: "津巴布韦元", Symbol: "$", SymbolFirst: true, AlternateSymbols: []string{"Z$"}, ThousandsSeparator: ",", DecimalMark: ".", SubUnit: "Cent", SubUnitToUnit: 100, HTMLEntity: "$"},
}

// 默认币种
var defaultCurrency *Currency  = currencies["CNY"]

func newCurrency(code string) *Currency {
	return &Currency{Code: strings.ToUpper(code)}
}


func (c *Currency) get() *Currency {
	if curr, ok := currencies[c.Code]; ok {
		return curr
	}
	return defaultCurrency
}

// 判断币种是否相同
func (c *Currency) equals(oc *Currency) bool {
	return c.Code == oc.Code
}

// 精度
func (c *Currency) precision() int32 {
	if c.SubUnitToUnit > 1 {
		if c.SubUnitToUnit < 10 {
			return 1
		}

		return int32(len(strconv.Itoa(c.SubUnitToUnit)) - 1)
	}

	return 0
}


// 取出一个币种
func GetCurrency(code string) (*Currency,bool) {
	if curr, ok := currencies[code]; ok {
		return curr,ok
	}
	return nil,false
}