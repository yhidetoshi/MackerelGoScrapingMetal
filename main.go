package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/PuerkitoBio/goquery"
	"github.com/mackerelio/mackerel-client-go"
)

var (
	targetUrl = "https://gold.tanaka.co.jp/commodity/souba/index.php"
	mkrKey    = os.Getenv("MKRKEY")
	client    = mackerel.NewClient(mkrKey)
)

const (
	serviceName = "Metal"
	timezone    = "Asia/Tokyo"
	offset      = 9 * 60 * 60
)

func main() {
	lambda.Start(Handler)
}

// func main() {
func Handler() {
	var goldRetailTax, goldPurchaseTax, platinumRetailTax, platinumPurchaseTax string

	doc, err := goquery.NewDocument(targetUrl)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch gold and platinum price
	doc.Find("#metal_price_sp").Each(func(_ int, s *goquery.Selection) {
		// Gold
		goldRetailTax = s.Children().Find("td.retail_tax").First().Text()
		goldPurchaseTax = s.Children().Find("td.purchase_tax").First().Text()
		// Platinum
		platinumRetailTax = s.Children().Find("td.retail_tax").Eq(1).Text()
		platinumPurchaseTax = s.Children().Find("td.purchase_tax").Eq(1).Text()
	})

	// Format
	strGoldRetailTax := strings.Replace(goldRetailTax[0:5], ",", "", -1)
	strGoldPurchaseTax := strings.Replace(goldPurchaseTax[0:5], ",", "", -1)
	strPlatinumRetailTax := strings.Replace(platinumRetailTax[0:5], ",", "", -1)
	strPlatinumPurchaseTax := strings.Replace(platinumPurchaseTax[0:5], ",", "", -1)

	// Convert string to int
	intGoldRetailTax, _ := strconv.Atoi(strGoldRetailTax)
	intGoldPurchaseTax, _ := strconv.Atoi(strGoldPurchaseTax)
	intPlatinumRetailTax, _ := strconv.Atoi(strPlatinumRetailTax)
	intPlatinumPurchaseTax, _ := strconv.Atoi(strPlatinumPurchaseTax)

	jst := time.FixedZone(timezone, offset)
	nowTime := time.Now().In(jst)

	mkrErr := PostValuesToMackerel(intGoldRetailTax, intGoldPurchaseTax, intPlatinumRetailTax, intPlatinumPurchaseTax, nowTime)
	if mkrErr != nil {
		fmt.Println(mkrErr)
	}
}

// PostValuesToMackerel Post Metrics to Mackerel
func PostValuesToMackerel(goldRetailTax int, goldPurchaseTax int, platinumRetailTax int, platinumPurchaseTax int, nowTime time.Time) error {
	// Post Gold metrics
	errGold := client.PostServiceMetricValues(serviceName, []*mackerel.MetricValue{
		&mackerel.MetricValue{
			Name:  "Gold.retail_tax",
			Time:  nowTime.Unix(),
			Value: goldRetailTax,
		},
		{
			Name:  "Gold.purchase_tax",
			Time:  nowTime.Unix(),
			Value: goldPurchaseTax,
		},
	})
	if errGold != nil {
		fmt.Println(errGold)
	}

	// Post Platinum metrics
	errPlatinum := client.PostServiceMetricValues(serviceName, []*mackerel.MetricValue{
		&mackerel.MetricValue{
			Name:  "Platinum.retail_tax",
			Time:  nowTime.Unix(),
			Value: platinumRetailTax,
		},
		{
			Name:  "Platinum.purchase_tax",
			Time:  nowTime.Unix(),
			Value: platinumPurchaseTax,
		},
	})
	if errPlatinum != nil {
		fmt.Println(errPlatinum)
	}

	return nil
}
