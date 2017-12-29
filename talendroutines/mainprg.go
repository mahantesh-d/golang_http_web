package main

import (
	"fmt"
	"math/rand"
	"bufio"
	"os"
	"time"
	"strconv"
	"container/list"
	"github.com/satori/go.uuid"
)

func main(){
	a := userinput()
	if a == "1" {
		stringrandomint()
	}else if a == "2" {
		randomnumber()
	}else if a == "3" {
		randomstring()
	}else if a == "4" {
		mapgo()
	}else if a == "5" {
		setgo()
	}else if a == "6" {
		listgo()
	}else if a == "7" {
		set_jsongo()
	}else if a == "8" {
		uuidgo()
	}else {
		fmt.Print("err")
	}
}

func userinput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Number: \n 1>String Random Integer \n 2>Random Number \n 3>Random String \n 4>Map " +
		"\n 5>Set \n 6>List \n 7>Set Json \n 8>Random UUID")
	scanner.Scan()
	text := scanner.Text()
	return text
}

//**********************************************************************************************************************

func stringrandomint(){
	b := input()
	heading("String Random")
	subheading(b)
	if b == "RT" || b == "TS" || b == "SS" || b == "RW" || b == "OS"  {
		a := test13(b)
		fmt.Println(a)
	}else if b == "IN" {
		a := test14(b)
		fmt.Println(a)
	}else if b == "Pu" || b == "C"{
		a := test10(b)
		fmt.Println(a)
	}else {
		fmt.Println("err")
	}

}
func input() string  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text: \t RT \t TS \t SS \t RW \t OS \t IN \t Pu \t C")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func test13(i string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(10000000000000)
	c := strconv.Itoa(a)
	f := 13 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return i+h
}

func test14(i string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(100000000000000)
	c := strconv.Itoa(a)
	f := 14 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return i+h
}

func test10(i string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(10000000000)
	c := strconv.Itoa(a)
	f := 10 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return i+h
}

//**********************************************************************************************************************

func randomnumber(){
	b := numinput()
	heading("Random Number")
	subheading(b)
	if b == "10" {
		a := num10()
		fmt.Println(a)
	}else if b == "6"{
		a := num6()
		fmt.Println(a)
	}else if b == "4"{
		a := num4()
		fmt.Println(a)
	}else if b == "3"{
		a := num3()
		fmt.Println(a)
	}else {
		fmt.Println("err")
	}

}

func numinput() string  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Number: \t 10 \t 6 \t 4 \t 3")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func num10() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(10000000000)
	c := strconv.Itoa(a)
	f := 10 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return h
}

func num6() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(1000000)
	c := strconv.Itoa(a)
	f := 6 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return h
}

func num4() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(10000)
	c := strconv.Itoa(a)
	f := 4 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return h
}

func num3() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(1000)
	c := strconv.Itoa(a)
	f := 3 - len(c)
	g := "0"
	h := c
	for j :=0; j<f; j++ {
		h = g + h
	}
	return h
}

//**********************************************************************************************************************

func randomstring(){
	b := stringinput()
	heading("Random String of Arrays")
	if b == "1" {
		subheading("Company")
		a := company()
		fmt.Println(a)
	}else if b == "2"{
		subheading("Location Code")
		a := locationcode()
		fmt.Println(a)
	}else if b == "3"{
		subheading("Brand")
		a := brand()
		fmt.Println(a)
	}else if b == "4" {
		subheading("Model")
		a := model()
		fmt.Println(a)
	}else if b == "5" {
		subheading("Mat Description")
		a := matdescription()
		fmt.Println(a)
	}else if b == "6" {
		subheading("Mat Code")
		a := matcode()
		fmt.Println(a)
	}else {
		fmt.Println("err")
	}

}

func stringinput() string  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Number: \t" +
		"1 > Company \t" +
		"2 > Locationcode \t" +
		"3 > Brand \t" +
		"4 > Model \t" +
		"5 > Matdescription \t" +
		"6 > Matcode")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func company() string {
	b := []string{"AIS", "AWN", "WDS"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func locationcode() string  {
	b := []string{"1177", "1771"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func brand() string  {
	b := []string{"ASUS", "SAMSUNG", "LAVA", "APPLE", "APPLE", "ZTE", "HUAWEI", "CHEERO_CHEERO",
		"BRAND_WANSMORN", "ONLINE", "D-LINK"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func model() string {
	b := []string {"ZC550KL",
		"N920X",
		"T715",
		"T815",
		"SS CASE",
		"W3",
		"IRIS600",
		"IPHONESE16GB" ,
		"IPHONESE64GB",
		"IRIS560",
		"IRIS755",
		"IRIS870",
		"T620",
		"T285",
		"IRIS550",
		"E8377S-153",
		"MF910L",
		"A710F",
		"J500G",
		"T715Y",
		"BACK UP BATTERY",
		"IRIS360",
		"W3",
		"IRIS600",
		"J105B",
		"IRIS550",
		"IRIS510",
		"IRIS750",
		"IRIS600",
		"IPHONESE16GB",
		"IPHONESE64GB", "DHP-600AV", "J500G", "G930F", "G935F", "IRIS515", "T116N", "G318", "S6E32"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]

}

func matdescription() string {
	b := []string{"NEW LTE SAM S6 32GB-BLACK 01DEMO" ,
		"SIM 3PP 3in1 BOS Zeed Shop - C3" ,
		"NEW 3G SAM GALAXY TAB3V-EBONY BLACK 01" ,
		"SIM 3PO Half3in1PostpaidLucky&WealthS-C3" ,
		"SIM 3PO Half3in1PostpaidHealth&WorkS- C3" ,
		"SIM 3PO Half3in1PostpaidMoney&LoveS - C3" ,
		"SIM 3PO ReplugPre2PostWork&LearnS - C309" ,
		"NEW 3G SAM GALAXY J1-WHITE 01" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - N3" ,
		"SIM 3PP Half 3in1 Traveller 299Shop - C3" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3" ,
		"CASH CARD 100B *999 (Twin) v.2" ,
		"SIM 3PP Half 3in1 2Fly 199B Shop - C3" ,
		"SIM 3PP 3in1 2FLY Shop 50B -C3" ,
		"NEW 3G SAM GALAXY TAB3V-CREAM WHITE 01" ,
		"NEW 3G SAM GALAXY TAB3V-EBONY BLACK 01" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3" ,
		"SIM 3PO Replug PostJameBPLPool Shop-N320" ,
		"SIM 3PP Half 3in1 Traveller 49 - C3" ,
		"SIM 3PP Half 3in1 Traveller 49 - C3" ,
		"SIM 3PP Half 3in1 Traveller 49 - C3" ,
		"SIM 3PP Half 3in1 Traveller 49 - C3" ,
		"SIM 3PP Half 3in1 Traveller 49 - C3" ,
		"NEW 3G SAM GALAXY J1-WHITE 01" ,
		"SIM 3PO 3in1 Pre2Post Instant - X331" ,
		"TWZ Power Tube Speaker_Green" ,
		"SIM 3PP Half 3in1 2Fly 199B Shop - C3" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3" ,
		"SIM 3PO Half3in1Pre2PostMoney&Love - C3" ,
		"SIM 3PO Half3in1Pre2PostHealth&Work - C3" ,
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3" ,
		"หมอนผ้าห่มน้องอุ่นใจ 14 นิ้ว" ,
		"Aunjai Power Bank 8,000 mAh with Box" ,
		"แก้วทำน้ำแข็งใส" ,
		"กระเป๋าผ้าสีเขียว พิมพ์น้องอุ่นใจ" ,
		"NEW 3G LAV IRIS708-GRAY 03 C3" ,
		"NEW 3G LAV IRIS600-BLACK 01" ,
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)" ,
		"NEW 3G LAV IRIS810-BLACK 01 X3" ,
		"NEW 3G LAV IRIS510-GREY 02 (SIM LOCK)" ,
		"NEW 3G LAV IRIS708-GRAY 02 (LOCK SIM)" ,
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)" ,
		"NEW 3G LAV IRIS600-BLACK 01" ,
		"NEW LTE LAV IRIS750-GREY 03 (SS/SL)" ,
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)" ,
		"NEW 3G LAV IRIS600-BLACK 01" ,
		"NEW 3G LAV IRIS510-GREY 02 (SIM LOCK)" ,
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)" ,
		"NEW 3G LAV IRIS600-BLACK 01" ,
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)" ,
		"NEW 3G LAV IRIS510-GREY 02 (SIM LOCK)" ,
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func matcode() string {
	b := []string{"NEWSAM0S632-BK01D" ,
		"C300SIM3PPTBZEEDS" ,
		"NEWSAMT116N-EB01" ,
		"C300SIM3POH3IN1LWS" ,
		"C300SIM3POH3IN1HWS" ,
		"C300SIM3POH3IN1MLS" ,
		"C309SIM3POREPPWLS" ,
		"NEWSAM0J100-WH01" ,
		"N300SIM3PP4GH3FDS" ,
		"C300SIM3PPHTT299SH" ,
		"C300SIM3PP4GH3FDS" ,
		"8858839100666" ,
		"C300SIM3PPHT2F199S" ,
		"C300SIM3PPT2FLYS50" ,
		"NEWSAMT116N-CW01" ,
		"NEWSAMT116N-EB01" ,
		"C300SIM3PP4GH3FDS" ,
		"N320SIM3POREPOJBPS" ,
		"C300SIM3PPHTTRAV49" ,
		"NEWSAM0J100-WH01" ,
		"X331SIM3POTPPO" ,
		"52001273" ,
		"C300SIM3PPHT2F199S" ,
		"C309SIM3POH3PPML" ,
		"C309SIM3POH3PPHW" ,
		"C300SIM3PP4GH3FDS" ,
		"52001259" ,
		"52001280" ,
		"52002130" ,
		"52002134" ,
		"NEWLAV0I600-BK01" ,
		"NEWLAV000W3-BK01" ,
		"NEWLAV0I810-BK01X3" ,
		"NEWLAV0I510-GY02" ,
		"NEWLAV0I708-GY02" ,
		"NEWLAV000W3-BK01" ,
		"NEWLAV0I600-BK01" ,
		"NEWLAV0I750-GY03" ,
		"NEWLAV000W3-BK01" ,
		"NEWLAV0I600-BK01" ,
		"NEWLAV0I510-GY02" ,
		"NEWLAV000W3-BK01" ,
		"NEWLAV0I600-BK01" ,
		"NEWLAV0I510-GY02" ,
		"NEWLAV000W3-BK01"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

//**********************************************************************************************************************

func mapgo(){
	heading("Map")
	b := "Map"
	subheading(b)
	fmt.Print(mapg())
}

func mapg() map[string]string{
	a := make(map[string]string)
	a["A"] = "AIS"
	a["U"] = "Unotech"
	a["E"] = "Entornica"
	return a
}

//**********************************************************************************************************************

func setgo(){
	heading("Set")
	b := "Set"
	subheading(b)
	a := setg()
	fmt.Print(a)
}

func setg() []string{
	s := make([]string, 3)
	s[0] = "Asus"
	s[1] = "Device"
	s[2] = "Handset"
	return s
}

//**********************************************************************************************************************

func listgo() {
	heading("List")
	b := "List"
	subheading(b)
	l := list.New()
	e3 := l.PushBack("Mumbai")
	l.PushFront("Unotech")
	l.InsertBefore("Software", e3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

//**********************************************************************************************************************

func set_jsongo()  {
	b := stringinput1()
	heading("Set Json")
	if b == "1"  {
		subheading("stock_adjustment")
		a := stock_adjustment()
		fmt.Print(a)
	}else if b == "2" {
		subheading("obtain_detail")
		a := obtain_detail()
		fmt.Print(a)
	}else if b == "3" {
		subheading("sub_stock_detail_transfer")
		a := sub_stock_detail_transfer()
		fmt.Print(a)
	}else if b == "4" {
		subheading("sub_stock_daily_detail")
		a := sub_stock_daily_detail()
		fmt.Print(a)
	}else if b == "5" {
		subheading("transfer_out_mismatch")
		a := transfer_out_mismatch()
		fmt.Print(a)
	}else if b == "6" {
		subheading("request_goods")
		a := request_goods()
		fmt.Print(a)
	}else if b == "7" {
		subheading("order_transfer")
		a := order_transfer()
		fmt.Print(a)
	}else if b == "8"{
		subheading("sale_out_detail")
		a := sale_out_detail()
		fmt.Print(a)
	}else if b == "9" {
		subheading("check_out_stock_detail")
		a := check_out_stock_detail()
		fmt.Print(a)
	}else {
		fmt.Print("err")
	}

}

func stringinput1() string  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Number: \t" +
		"1 > stock_adjustment \t" +
		"2 > obtain_detail \t" +
		"3 > sub_stock_detail_transfer \t" +
		"4 > sub_stock_daily_detail \t" +
		"5 > transfer_out_mismatch \t" +
		"6 > request_goods \t" +
		"7 > order_transfer \t" +
		"8 > sale_out_detail \t" +
		"9 > check_out_stock_detail")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func stock_adjustment() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"matType\": \"serial\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"itemNo\": \"0002\"," +
		"\"fromPlant\": \"1123\"," +
		"\"fromStorage\": \"4351\"," +
		"\"qty\": \"2\"," +
		"\"serialList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"reason\": \"\"," +
		"\"subStock\": \"Sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"reason\": \"text\"," +
		"\"subStock\": \"Sale@Shop\"" +
		"}]" +
		"}"
	return s
}

func obtain_detail() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"salesCode\": \"3660\"," +
		"\"salesType\": \"DirectSales\"," +
		"\"salesName\": \"Suporn Roddon\"," +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"matType\": \"serial\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"company\": \"AWN\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"itemNo\": \"0002\"," +
		"\"qty\": \"10\"," +
		"\"serialList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"reason\": \"\"," +
		"\"fromSubStock\": \"Sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"reason\": \"text\"," +
		"\"fromSubStock\": \"Sale@Shop\"" +
		"}]" +
		"}"
	return s
}

func sub_stock_detail_transfer() []string {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"matType\": \"serial\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"outQty\": \"2\"," +
		"\"unit\": \"set\"," +
		"\"serialList\": [{" +
		"\"fromSubStock\": \"Sale@Shop\"," +
		"\"serial\": \"000000000000001\"," +
		"\"reason\": \"\"" +
		"}, {" +
		"\"fromSubStock\": \"Sale@Shop\"," +
		"\"serial\": \"000000000000002\"," +
		"\"reason\": \"text\"" +
		"}]" +
		"}"
	return s
}

func sub_stock_daily_detail() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"matType\": \"serial\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"serialList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"reason\": \"\"," +
		"\"statusBySerial\": \"Sold\"," +
		"\"fromSubStock\": \"sale@shop\"," +
		"\"qtyByStatus\": \"1\"," +
		"\"qtyByPickup\": \"1\"," +
		"\"createSerialBy\": \"somza\"," +
		"\"createSerialDateTime\": \"2016-09-09 15:37:50+0700\"," +
		"\"salesName\": \"somza\"," +
		"\"receiptNo\": \"S01609090001\"," +
		"\"msisdn\": \"66890009878\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"reason\": \"\"," +
		"\"statusBySerial\": \"Sold\"," +
		"\"fromSubStock\": \"sale@shop\"," +
		"\"qtyByStatus\": \"1\"," +
		"\"qtyByPickup\": \"1\"," +
		"\"createBy\": \"somza\"" +
		"}]" +
		"}"
	return s
}

func transfer_out_mismatch() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWAPP0SE16-GD01\"," +
		"\"matType\": \"serial\"," +
		"\"brand\": \"APPLE\"," +
		"\"model\": \"IPHONESE16GB\"," +
		"\"color\": \"GOLD\"," +
		"\"commercialName\": \"NEW LTE APP IPHONESE 16GB-GOLD 01\"," +
		"\"company\": \"AWN\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"overQty \": \"10\"," +
		"\"missingQty\": \"10\"" +
		"}"
	return s
}

func request_goods() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"division\": \"11text\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"reqQty\": \"10\"," +
		"\"unit\": \"SET\"," +
		"\"quotaFlag\": \"Y\"," +
		"\"quotaMstId\": \"\"," +
		"\"quotaDtlId\": \"\"," +
		"\"itemNo\": \"0001\"," +
		"\"plantFrom\": \"1111\"," +
		"\"storageFrom\": \"2222\"," +
		"\"plantTo\": \"3333\"," +
		"\"storageTo\": \"4444\"" +
		"}"
	return s
}

func order_transfer() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"matType\": \"serial\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"transferQty\": \"10\"," +
		"\"unit\": \"SET\"," +
		"\"itemNo\": \"0001\"," +
		"\"fromPlant\": \"1111\"," +
		"\"fromStorage\": \"2222\"," +
		"\"toPlant\": \"3333\"," +
		"\"toStorage\": \"4444\"," +
		"\"serialList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"toPlant\": \"\"," +
		"\"toStorage\": \"\"," +
		"\"reason\": \"\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"toPlant\": \"\"," +
		"\"toStorage\": \"\"," +
		"\"reason\": \"abcdefg\"" +
		"}]" +
		"}"
	return s
}

func sale_out_detail() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"matCode\": \"NEWASUZC550-BK01\"," +
		"\"matType\": \"serial\"," +
		"\"commercialName\": \"ASUS 16G\"," +
		"\"matDescription\": \"NEW 4G ASU MAX 16/2GB,5000mAh-BLACK 01\"," +
		"\"brand\": \"ASUS\"," +
		"\"model\": \"ZC550KL\"," +
		"\"color\": \"BLACK\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"salesQty\": \"1\"," +
		"\"stockBeforeQty\": \"5\"," +
		"\"stockAfterQty\": \"4\"," +
		"\"serialList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"serialStatus\": \"sale\"," +
		"\"serialQty\": \"1\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"serialStatus\": \"sale\"," +
		"\"serialQty\": \"1\"" +
		"}]" +
		"" +
		"}"
	return s
}

func check_out_stock_detail() []string  {
	s := make([]string, 1)
	s[0] = "{" +
		"\"company\": \"\"," +
		"\"matCode\": \"\"," +
		"\"matType\": \"\"," +
		"\"brand\": \"\"," +
		"\"model\": \"\"," +
		"\"color\": \"\"," +
		"\"commercialName\": \"ASUS16G\"," +
		"\"matDescription\": \"NEW4GASUMAX16/2GB, 5000mAh-BLACK01\"," +
		"\"productType\": \"DEVICE\"," +
		"\"productSubType\": \"HANDSET\"," +
		"\"totalScanQty\": \"\"," +
		"\"diffQty\": \"0\"," +
		"\"matchQty\": \"5\"," +
		"\"serialMatchList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000003\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000004\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000005\"," +
		"\"substock\": \"sale@Shop\"" +
		"}]," +
		"\"notFoundQty\": \"\"," +
		"\"serialNotFoundList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000003\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000004\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000005\"," +
		"\"substock\": \"sale@Shop\"" +
		"}]," +
		"\"notScanQty\": \"\"," +
		"\"serialNotScanList\": [{" +
		"\"serial\": \"000000000000001\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000002\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000003\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000004\"," +
		"\"substock\": \"sale@Shop\"" +
		"}, {" +
		"\"serial\": \"000000000000005\"," +
		"\"substock\": \"sale@Shop\"" +
		"}]" +
		"" +
		"}"
	return s

}

//**********************************************************************************************************************

func uuidgo(){
	heading("Random UUID")
	b := "UUID"
	subheading(b)
	a := uuid.NewV4()
	c := a.String()
	fmt.Print(c)

}
func heading(b string){
	fmt.Println("\n-----------------------------------------------------------")
	fmt.Println("HEADING >> " +b)
	fmt.Println("-----------------------------------------------------------")
}

func subheading(c string){
	fmt.Println("-----SUBHEADING >> " + c + "---------------------------------------------")

}