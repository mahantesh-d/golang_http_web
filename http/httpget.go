package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://10.138.32.76:9001/v1/alltrade/Stock_Adjustment?filters=(adjustStockNo=AJ0000726538072)(locationCode=1771)")
	if err != nil {
		log.Fatal("Could not read the host")
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	fmt.Println(buf.String())
}
