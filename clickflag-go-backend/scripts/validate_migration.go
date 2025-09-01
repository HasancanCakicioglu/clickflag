package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println("🔍 Migration Validation Script - Ülke Kodları Kontrolü")
	fmt.Println(strings.Repeat("=", 60))

	// 1. Go constants dosyasından ülke kodlarını oku
	goConstants := readGoConstants()
	fmt.Printf("✅ Go constants: %d ülke kodu bulundu\n", len(goConstants))

	// 2. Migration dosyasından CHECK constraint'teki ülke kodlarını oku
	checkConstraint := readCheckConstraint()
	fmt.Printf("✅ CHECK constraint: %d ülke kodu bulundu\n", len(checkConstraint))

	// 3. Migration dosyasından INSERT satırlarındaki ülke kodlarını oku
	insertCodes := readInsertCodes()
	fmt.Printf("✅ INSERT satırları: %d ülke kodu bulundu\n", len(insertCodes))

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 KARŞILAŞTIRMA SONUÇLARI")
	fmt.Println(strings.Repeat("=", 60))

	// 4. Go constants vs CHECK constraint karşılaştırması
	fmt.Println("\n🔍 Go Constants vs CHECK Constraint:")
	compareCountryCodes(goConstants, checkConstraint, "Go Constants", "CHECK Constraint")

	// 5. Go constants vs INSERT satırları karşılaştırması
	fmt.Println("\n🔍 Go Constants vs INSERT Satırları:")
	compareCountryCodes(goConstants, insertCodes, "Go Constants", "INSERT Satırları")

	// 6. CHECK constraint vs INSERT satırları karşılaştırması
	fmt.Println("\n🔍 CHECK Constraint vs INSERT Satırları:")
	compareCountryCodes(checkConstraint, insertCodes, "CHECK Constraint", "INSERT Satırları")

	// 7. Genel özet
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📋 GENEL ÖZET")
	fmt.Println(strings.Repeat("=", 60))

	if len(goConstants) == len(checkConstraint) && len(goConstants) == len(insertCodes) {
		fmt.Println("🎉 TÜM ÜLKE KODLARI EŞLEŞİYOR!")
		fmt.Printf("   Toplam: %d ülke kodu\n", len(goConstants))
	} else {
		fmt.Println("⚠️  ÜLKE KODLARINDA UYUMSUZLUK VAR!")
		fmt.Printf("   Go Constants: %d\n", len(goConstants))
		fmt.Printf("   CHECK Constraint: %d\n", len(checkConstraint))
		fmt.Printf("   INSERT Satırları: %d\n", len(insertCodes))
	}
}

// readGoConstants: Go constants dosyasından ülke kodlarını okur
func readGoConstants() []string {
	file, err := os.Open("../constants/countries.go")
	if err != nil {
		log.Fatal("❌ Go constants dosyası açılamadı:", err)
	}
	defer file.Close()

	var codes []string
	scanner := bufio.NewScanner(file)

	// Ülke kodu regex pattern'i
	codePattern := regexp.MustCompile(`"([A-Z]{2})"`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := codePattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				codes = append(codes, match[1])
			}
		}
	}

	sort.Strings(codes)
	return codes
}

// readCheckConstraint: Migration dosyasından CHECK constraint'teki ülke kodlarını okur
func readCheckConstraint() []string {
	file, err := os.Open("../migrations/001_create_countries_table.sql")
	if err != nil {
		log.Fatal("❌ Migration dosyası açılamadı:", err)
	}
	defer file.Close()

	var codes []string
	scanner := bufio.NewScanner(file)

	// CHECK constraint regex pattern'i
	checkPattern := regexp.MustCompile(`CHECK \(country_code IN \(([^)]+)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := checkPattern.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Ülke kodlarını ayır ve temizle
			codesStr := matches[1]
			codesStr = strings.ReplaceAll(codesStr, "'", "")
			codesStr = strings.ReplaceAll(codesStr, " ", "")
			codes = strings.Split(codesStr, ",")
			break
		}
	}

	sort.Strings(codes)
	return codes
}

// readInsertCodes: Migration dosyasından INSERT satırlarındaki ülke kodlarını okur
func readInsertCodes() []string {
	file, err := os.Open("../migrations/001_create_countries_table.sql")
	if err != nil {
		log.Fatal("❌ Migration dosyası açılamadı:", err)
	}
	defer file.Close()

	var codes []string
	scanner := bufio.NewScanner(file)

	// INSERT satırları regex pattern'i - artık herhangi bir sayı olabilir
	insertPattern := regexp.MustCompile(`\('([A-Z]{2})', \d+\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := insertPattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				codes = append(codes, match[1])
			}
		}
	}

	sort.Strings(codes)
	return codes
}

// compareCountryCodes: İki ülke kodu listesini karşılaştırır
func compareCountryCodes(list1, list2 []string, name1, name2 string) {
	// Set oluştur
	set1 := make(map[string]bool)
	set2 := make(map[string]bool)

	for _, code := range list1 {
		set1[code] = true
	}
	for _, code := range list2 {
		set2[code] = true
	}

	// Eksik olanları bul
	var missingIn2 []string
	for _, code := range list1 {
		if !set2[code] {
			missingIn2 = append(missingIn2, code)
		}
	}

	// Fazla olanları bul
	var extraIn2 []string
	for _, code := range list2 {
		if !set1[code] {
			extraIn2 = append(extraIn2, code)
		}
	}

	// Sonuçları yazdır
	if len(missingIn2) == 0 && len(extraIn2) == 0 {
		fmt.Printf("   ✅ %s ve %s tamamen eşleşiyor\n", name1, name2)
	} else {
		if len(missingIn2) > 0 {
			fmt.Printf("   ❌ %s'de eksik olanlar (%d): %v\n", name2, len(missingIn2), missingIn2)
		}
		if len(extraIn2) > 0 {
			fmt.Printf("   ❌ %s'de fazla olanlar (%d): %v\n", name2, len(extraIn2), extraIn2)
		}
	}
}
