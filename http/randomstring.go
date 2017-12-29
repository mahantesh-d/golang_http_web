package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	b := stringinput()
	heading1("Random String of Arrays")
	subheading1(b)
	if b == "Company" {
		a := company()
		fmt.Println(a)
	} else if b == "Locationcode" {
		a := locationcode()
		fmt.Println(a)
	} else if b == "Brand" {
		a := brand()
		fmt.Println(a)
	} else if b == "Model" {
		a := model()
		fmt.Println(a)
	} else if b == "Matdescription" {
		a := matdescription()
		fmt.Println(a)
	} else if b == "Matcode" {
		a := matcode()
		fmt.Println(a)
	} else {
		fmt.Println("err")
	}

}

func stringinput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter below option: \n" +
		"> Company \n" +
		"> Locationcode \n" +
		"> Brand \n" +
		"> Model \n" +
		"> Matdescription \n" +
		"> Matcode")
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

func locationcode() string {
	b := []string{"1177", "1771"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func brand() string {
	b := []string{"ASUS", "SAMSUNG", "LAVA", "APPLE", "APPLE", "ZTE", "HUAWEI", "CHEERO_CHEERO",
		"BRAND_WANSMORN", "ONLINE", "D-LINK"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func model() string {
	b := []string{"ZC550KL",
		"N920X",
		"T715",
		"T815",
		"SS CASE",
		"W3",
		"IRIS600",
		"IPHONESE16GB",
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
	b := []string{"NEW LTE SAM S6 32GB-BLACK 01DEMO",
		"SIM 3PP 3in1 BOS Zeed Shop - C3",
		"NEW 3G SAM GALAXY TAB3V-EBONY BLACK 01",
		"SIM 3PO Half3in1PostpaidLucky&WealthS-C3",
		"SIM 3PO Half3in1PostpaidHealth&WorkS- C3",
		"SIM 3PO Half3in1PostpaidMoney&LoveS - C3",
		"SIM 3PO ReplugPre2PostWork&LearnS - C309",
		"NEW 3G SAM GALAXY J1-WHITE 01",
		"SIM 3PP 4G Half 3in1 Freedom Shop - N3",
		"SIM 3PP Half 3in1 Traveller 299Shop - C3",
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3",
		"CASH CARD 100B *999 (Twin) v.2",
		"SIM 3PP Half 3in1 2Fly 199B Shop - C3",
		"SIM 3PP 3in1 2FLY Shop 50B -C3",
		"NEW 3G SAM GALAXY TAB3V-CREAM WHITE 01",
		"NEW 3G SAM GALAXY TAB3V-EBONY BLACK 01",
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3",
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3",
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3",
		"SIM 3PO Replug PostJameBPLPool Shop-N320",
		"SIM 3PP Half 3in1 Traveller 49 - C3",
		"SIM 3PP Half 3in1 Traveller 49 - C3",
		"SIM 3PP Half 3in1 Traveller 49 - C3",
		"SIM 3PP Half 3in1 Traveller 49 - C3",
		"SIM 3PP Half 3in1 Traveller 49 - C3",
		"NEW 3G SAM GALAXY J1-WHITE 01",
		"SIM 3PO 3in1 Pre2Post Instant - X331",
		"TWZ Power Tube Speaker_Green",
		"SIM 3PP Half 3in1 2Fly 199B Shop - C3",
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3",
		"SIM 3PO Half3in1Pre2PostMoney&Love - C3",
		"SIM 3PO Half3in1Pre2PostHealth&Work - C3",
		"SIM 3PP 4G Half 3in1 Freedom Shop - C3",
		"หมอนผ้าห่มน้องอุ่นใจ 14 นิ้ว",
		"Aunjai Power Bank 8,000 mAh with Box",
		"แก้วทำน้ำแข็งใส",
		"กระเป๋าผ้าสีเขียว พิมพ์น้องอุ่นใจ",
		"NEW 3G LAV IRIS708-GRAY 03 C3",
		"NEW 3G LAV IRIS600-BLACK 01",
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)",
		"NEW 3G LAV IRIS810-BLACK 01 X3",
		"NEW 3G LAV IRIS510-GREY 02 (SIM LOCK)",
		"NEW 3G LAV IRIS708-GRAY 02 (LOCK SIM)",
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)",
		"NEW 3G LAV IRIS600-BLACK 01",
		"NEW LTE LAV IRIS750-GREY 03 (SS/SL)",
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)",
		"NEW 3G LAV IRIS600-BLACK 01",
		"NEW 3G LAV IRIS510-GREY 02 (SIM LOCK)",
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)",
		"NEW 3G LAV IRIS600-BLACK 01",
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)",
		"NEW 3G LAV IRIS510-GREY 02 (SIM LOCK)",
		"NEW 3G LAV W3-BLACK 01 (SIM LOCK)"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func matcode() string {
	b := []string{"NEWSAM0S632-BK01D",
		"C300SIM3PPTBZEEDS",
		"NEWSAMT116N-EB01",
		"C300SIM3POH3IN1LWS",
		"C300SIM3POH3IN1HWS",
		"C300SIM3POH3IN1MLS",
		"C309SIM3POREPPWLS",
		"NEWSAM0J100-WH01",
		"N300SIM3PP4GH3FDS",
		"C300SIM3PPHTT299SH",
		"C300SIM3PP4GH3FDS",
		"8858839100666",
		"C300SIM3PPHT2F199S",
		"C300SIM3PPT2FLYS50",
		"NEWSAMT116N-CW01",
		"NEWSAMT116N-EB01",
		"C300SIM3PP4GH3FDS",
		"N320SIM3POREPOJBPS",
		"C300SIM3PPHTTRAV49",
		"NEWSAM0J100-WH01",
		"X331SIM3POTPPO",
		"52001273",
		"C300SIM3PPHT2F199S",
		"C309SIM3POH3PPML",
		"C309SIM3POH3PPHW",
		"C300SIM3PP4GH3FDS",
		"52001259",
		"52001280",
		"52002130",
		"52002134",
		"NEWLAV0I600-BK01",
		"NEWLAV000W3-BK01",
		"NEWLAV0I810-BK01X3",
		"NEWLAV0I510-GY02",
		"NEWLAV0I708-GY02",
		"NEWLAV000W3-BK01",
		"NEWLAV0I600-BK01",
		"NEWLAV0I750-GY03",
		"NEWLAV000W3-BK01",
		"NEWLAV0I600-BK01",
		"NEWLAV0I510-GY02",
		"NEWLAV000W3-BK01",
		"NEWLAV0I600-BK01",
		"NEWLAV0I510-GY02",
		"NEWLAV000W3-BK01"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	return b[a]
}

func heading1(b string) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(b)
	fmt.Println("-----------------------------------------------------------")
}

func subheading1(c string) {
	fmt.Println("-----" + c + "---------------------------------------------")

}
