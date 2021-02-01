package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() { //kaynak kodumuzu aldık siteden
	resp, err := http.Get("http://realtime.paragaranti.com/asp/xml/icpiyasa.asp")
	if err != nil {
		log.Fatal("HATA: %s", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data))// siteden alinan verilerin html kodlarını gormek isterseniz calistirin

	var s icpiyasa
	xml.Unmarshal(data, &s)
	fmt.Println(s.Stocks)

}

type stock struct { //xml piyasamizi tanimladik
	XMLName xml.Name `xml:"STOCK"`
	SYMBOL  string   `xml:"SYMBOL"`
	DESC    string   `xml:"DESC"`
	LAST    string   `xml:"LAST"`
	PERNC   string   `xml:"PERNC"`
}
type icpiyasa struct {
	XMLName xml.Name `xml:"ICPIYASA"`
	Stocks  []stock  `xml:"STOCK"`
}

func (s stock) String() string {
	return fmt.Sprintf("\t Sembol:%s - Acıklama:%s - Son Deger:%s - Degisim:%s - \n", s.SYMBOL, s.DESC, s.LAST, s.PERNC)
}
