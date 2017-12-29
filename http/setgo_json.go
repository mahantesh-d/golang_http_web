package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	b := stringinput1()
	heading1("Set Json")
	subheading1(b)
	if b == "stock_adjustment" {
		a := stock_adjustment()
		fmt.Print(a)
	} else if b == "obtain_detail" {
		a := obtain_detail()
		fmt.Print(a)
	} else if b == "sub_stock_detail_transfer" {
		a := sub_stock_detail_transfer()
		fmt.Print(a)
	} else if b == "sub_stock_daily_detail" {
		a := sub_stock_daily_detail()
		fmt.Print(a)
	} else if b == "transfer_out_mismatch" {
		a := transfer_out_mismatch()
		fmt.Print(a)
	} else if b == "request_goods" {
		a := request_goods()
		fmt.Print(a)
	} else if b == "order_transfer" {
		a := order_transfer()
		fmt.Print(a)
	} else if b == "sale_out_detail" {
		a := sale_out_detail()
		fmt.Print(a)
	} else if b == "check_out_stock_detail" {
		a := check_out_stock_detail()
		fmt.Print(a)
	} else {
		fmt.Print("err")
	}

}

func stringinput1() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter below option: \t" +
		"> stock_adjustment \t" +
		"> obtain_detail \t" +
		"> sub_stock_detail_transfer \t" +
		"> sub_stock_daily_detail \t" +
		"> transfer_out_mismatch \t" +
		"> request_goods \t" +
		"> order_transfer \t" +
		"> sale_out_detail \t" +
		"> check_out_stock_detail")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func stock_adjustment() []string {
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

func obtain_detail() []string {
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

func sub_stock_daily_detail() []string {
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

func transfer_out_mismatch() []string {
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

func request_goods() []string {
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

func order_transfer() []string {
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

func sale_out_detail() []string {
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

func check_out_stock_detail() []string {
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
func heading1(b string) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(b)
	fmt.Println("-----------------------------------------------------------")
}

func subheading1(c string) {
	fmt.Println("-----" + c + "---------------------------------------------")

}
