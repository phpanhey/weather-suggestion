package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	req, err := http.NewRequest("GET", "https://www.ach-du-schan.de/", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/140.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	temp, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(doc.Find(`td[style*="font-size: 3.5em"][style*="color: #FF0000"]`).First().Text()), "°C", ""), ",", "."), 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.2f °C ", temp)
	fmt.Print(calculateSuggestion(temp))
}

func calculateSuggestion(temp float64) string {
	if temp < 5 {
		return "🥶 🧥"
	}

	if temp < 12 {
		return "🧥"
	}

	if temp < 18 {
		return "👕 🍂"
	}

	if temp < 23 {
		return "👕"
	}

	return "🩳 ☀️"
}
