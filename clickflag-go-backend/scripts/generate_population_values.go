package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// CountryPopulation represents a country with its population
type CountryPopulation struct {
	Code       string
	Name       string
	Population int
	Value      int
}

// Population data - real population numbers
func getPopulationData() map[string]CountryPopulation {
	return map[string]CountryPopulation{
		// A
		"AD": {"AD", "Andorra", 77265, 0},
		"AE": {"AE", "United Arab Emirates", 10081715, 0},
		"AF": {"AF", "Afghanistan", 38928346, 0},
		"AG": {"AG", "Antigua and Barbuda", 97929, 0},
		"AL": {"AL", "Albania", 2877797, 0},
		"AM": {"AM", "Armenia", 2963243, 0},
		"AO": {"AO", "Angola", 32866272, 0},
		"AR": {"AR", "Argentina", 45195774, 0},
		"AT": {"AT", "Austria", 9006398, 0},
		"AU": {"AU", "Australia", 25499884, 0},
		"AZ": {"AZ", "Azerbaijan", 10139177, 0},

		// B
		"BA": {"BA", "Bosnia and Herzegovina", 3280819, 0},
		"BB": {"BB", "Barbados", 287375, 0},
		"BD": {"BD", "Bangladesh", 164689383, 0},
		"BE": {"BE", "Belgium", 11589623, 0},
		"BF": {"BF", "Burkina Faso", 20903273, 0},
		"BG": {"BG", "Bulgaria", 6948445, 0},
		"BH": {"BH", "Bahrain", 1701575, 0},
		"BI": {"BI", "Burundi", 11890784, 0},
		"BJ": {"BJ", "Benin", 12123200, 0},
		"BN": {"BN", "Brunei", 437479, 0},
		"BO": {"BO", "Bolivia", 11673021, 0},
		"BR": {"BR", "Brazil", 212559417, 0},
		"BS": {"BS", "Bahamas", 393244, 0},
		"BT": {"BT", "Bhutan", 771608, 0},
		"BW": {"BW", "Botswana", 2351627, 0},
		"BY": {"BY", "Belarus", 9449323, 0},
		"BZ": {"BZ", "Belize", 397628, 0},

		// C
		"CA": {"CA", "Canada", 37742154, 0},
		"CD": {"CD", "DR Congo", 89561403, 0},
		"CF": {"CF", "Central African Republic", 4829767, 0},
		"CG": {"CG", "Congo", 5518087, 0},
		"CH": {"CH", "Switzerland", 8654622, 0},
		"CI": {"CI", "Côte d'Ivoire", 26378274, 0},
		"CL": {"CL", "Chile", 19116201, 0},
		"CM": {"CM", "Cameroon", 26545863, 0},
		"CN": {"CN", "China", 1439323776, 0},
		"CO": {"CO", "Colombia", 50882891, 0},
		"CR": {"CR", "Costa Rica", 5094118, 0},
		"CU": {"CU", "Cuba", 11326616, 0},
		"CV": {"CV", "Cape Verde", 555987, 0},
		"CY": {"CY", "Cyprus", 1207359, 0},
		"CZ": {"CZ", "Czech Republic", 10689209, 0},

		// D
		"DZ": {"DZ", "Algeria", 44616624, 0},
		"DE": {"DE", "Germany", 83783942, 0},
		"DJ": {"DJ", "Djibouti", 988000, 0},
		"DK": {"DK", "Denmark", 5792202, 0},
		"DM": {"DM", "Dominica", 71986, 0},
		"DO": {"DO", "Dominican Republic", 10847910, 0},

		// E
		"EC": {"EC", "Ecuador", 17643054, 0},
		"EE": {"EE", "Estonia", 1326535, 0},
		"EG": {"EG", "Egypt", 102334404, 0},
		"ER": {"ER", "Eritrea", 3546421, 0},
		"ES": {"ES", "Spain", 46754778, 0},
		"ET": {"ET", "Ethiopia", 114963588, 0},

		// F
		"FI": {"FI", "Finland", 5540720, 0},
		"FJ": {"FJ", "Fiji", 896444, 0},
		"FM": {"FM", "Micronesia", 115023, 0},
		"FR": {"FR", "France", 65273511, 0},

		// G
		"GA": {"GA", "Gabon", 2225734, 0},
		"GD": {"GD", "Grenada", 112523, 0},
		"GE": {"GE", "Georgia", 3989167, 0},
		"GH": {"GH", "Ghana", 31072940, 0},
		"GM": {"GM", "Gambia", 2416668, 0},
		"GN": {"GN", "Guinea", 13132795, 0},
		"GQ": {"GQ", "Equatorial Guinea", 1402985, 0},
		"GR": {"GR", "Greece", 10423054, 0},
		"GT": {"GT", "Guatemala", 17915568, 0},
		"GW": {"GW", "Guinea-Bissau", 1968001, 0},
		"GY": {"GY", "Guyana", 786552, 0},

		// H
		"HN": {"HN", "Honduras", 9904607, 0},
		"HR": {"HR", "Croatia", 4105267, 0},
		"HT": {"HT", "Haiti", 11402528, 0},
		"HU": {"HU", "Hungary", 9660351, 0},

		// I
		"ID": {"ID", "Indonesia", 273523615, 0},
		"IL": {"IL", "Israel", 9291000, 0},
		"IN": {"IN", "India", 1380004385, 0},
		"IQ": {"IQ", "Iraq", 40462701, 0},
		"IR": {"IR", "Iran", 83992949, 0},
		"IE": {"IE", "Ireland", 4937786, 0},
		"IT": {"IT", "Italy", 60461826, 0},

		// J
		"JM": {"JM", "Jamaica", 2961167, 0},
		"JO": {"JO", "Jordan", 10203134, 0},
		"JP": {"JP", "Japan", 126476461, 0},

		// K
		"KE": {"KE", "Kenya", 53771296, 0},
		"KG": {"KG", "Kyrgyzstan", 6524195, 0},
		"KH": {"KH", "Cambodia", 16718965, 0},
		"KI": {"KI", "Kiribati", 119449, 0},
		"KM": {"KM", "Comoros", 869601, 0},
		"KN": {"KN", "Saint Kitts and Nevis", 53199, 0},
		"KP": {"KP", "North Korea", 25778816, 0},
		"KR": {"KR", "South Korea", 51269185, 0},
		"KW": {"KW", "Kuwait", 4270571, 0},
		"KZ": {"KZ", "Kazakhstan", 18776707, 0},

		// L
		"LA": {"LA", "Laos", 7275560, 0},
		"LB": {"LB", "Lebanon", 6825445, 0},
		"LC": {"LC", "Saint Lucia", 183627, 0},
		"LI": {"LI", "Liechtenstein", 38128, 0},
		"LK": {"LK", "Sri Lanka", 21413249, 0},
		"LR": {"LR", "Liberia", 5057681, 0},
		"LS": {"LS", "Lesotho", 2142249, 0},
		"LT": {"LT", "Lithuania", 2722289, 0},
		"LU": {"LU", "Luxembourg", 625978, 0},
		"LV": {"LV", "Latvia", 1886198, 0},
		"LY": {"LY", "Libya", 6871292, 0},

		// M
		"MA": {"MA", "Morocco", 36910560, 0},
		"MC": {"MC", "Monaco", 39242, 0},
		"MD": {"MD", "Moldova", 2617820, 0},
		"ME": {"ME", "Montenegro", 621718, 0},
		"MG": {"MG", "Madagascar", 27691018, 0},
		"MH": {"MH", "Marshall Islands", 59190, 0},
		"MK": {"MK", "North Macedonia", 2083374, 0},
		"ML": {"ML", "Mali", 20250833, 0},
		"MM": {"MM", "Myanmar", 54409800, 0},
		"MN": {"MN", "Mongolia", 3278290, 0},
		"MR": {"MR", "Mauritania", 4649658, 0},
		"MT": {"MT", "Malta", 441543, 0},
		"MU": {"MU", "Mauritius", 1271768, 0},
		"MV": {"MV", "Maldives", 540544, 0},
		"MW": {"MW", "Malawi", 19129952, 0},
		"MX": {"MX", "Mexico", 128932753, 0},
		"MY": {"MY", "Malaysia", 32365999, 0},
		"MZ": {"MZ", "Mozambique", 31255435, 0},

		// N
		"NA": {"NA", "Namibia", 2540905, 0},
		"NE": {"NE", "Niger", 24206644, 0},
		"NG": {"NG", "Nigeria", 206139589, 0},
		"NI": {"NI", "Nicaragua", 6624554, 0},
		"NL": {"NL", "Netherlands", 17134872, 0},
		"NO": {"NO", "Norway", 5421241, 0},
		"NP": {"NP", "Nepal", 29136808, 0},
		"NR": {"NR", "Nauru", 10824, 0},

		// O
		"OM": {"OM", "Oman", 5106626, 0},

		// P
		"PA": {"PA", "Panama", 4314767, 0},
		"PE": {"PE", "Peru", 32971854, 0},
		"PG": {"PG", "Papua New Guinea", 8947024, 0},
		"PH": {"PH", "Philippines", 109581078, 0},
		"PK": {"PK", "Pakistan", 220892340, 0},
		"PL": {"PL", "Poland", 37846611, 0},
		"PS": {"PS", "Palestine", 5101414, 0},
		"PT": {"PT", "Portugal", 10196709, 0},
		"PW": {"PW", "Palau", 18094, 0},
		"PY": {"PY", "Paraguay", 7132538, 0},

		// Q
		"QA": {"QA", "Qatar", 2881053, 0},

		// R
		"RO": {"RO", "Romania", 19237691, 0},
		"RS": {"RS", "Serbia", 8737371, 0},
		"RU": {"RU", "Russia", 145912025, 0},
		"RW": {"RW", "Rwanda", 12952218, 0},

		// S
		"SA": {"SA", "Saudi Arabia", 34813871, 0},
		"SB": {"SB", "Solomon Islands", 686884, 0},
		"SC": {"SC", "Seychelles", 98347, 0},
		"SD": {"SD", "Sudan", 43849260, 0},
		"SE": {"SE", "Sweden", 10099265, 0},
		"SG": {"SG", "Singapore", 5850342, 0},
		"SI": {"SI", "Slovenia", 2078938, 0},
		"SK": {"SK", "Slovakia", 5459642, 0},
		"SL": {"SL", "Sierra Leone", 7976983, 0},
		"SM": {"SM", "San Marino", 33931, 0},
		"SN": {"SN", "Senegal", 16743927, 0},
		"SO": {"SO", "Somalia", 15893222, 0},
		"SR": {"SR", "Suriname", 586634, 0},
		"ST": {"ST", "Sao Tome and Principe", 219159, 0},
		"SV": {"SV", "El Salvador", 6486205, 0},
		"SY": {"SY", "Syria", 17500658, 0},
		"SZ": {"SZ", "Eswatini", 1160164, 0},
		"WS": {"WS", "Samoa", 198414, 0},

		// T
		"TD": {"TD", "Chad", 16425864, 0},
		"TG": {"TG", "Togo", 8278724, 0},
		"TH": {"TH", "Thailand", 69799978, 0},
		"TJ": {"TJ", "Tajikistan", 9537645, 0},
		"TL": {"TL", "Timor-Leste", 1318445, 0},
		"TM": {"TM", "Turkmenistan", 6031200, 0},
		"TN": {"TN", "Tunisia", 11818619, 0},
		"TO": {"TO", "Tonga", 105695, 0},
		"TR": {"TR", "Turkey", 84339067, 0},
		"TT": {"TT", "Trinidad and Tobago", 1399488, 0},
		"TV": {"TV", "Tuvalu", 11792, 0},
		"TW": {"TW", "Taiwan", 23816775, 0},
		"TZ": {"TZ", "Tanzania", 59734218, 0},

		// U
		"UA": {"UA", "Ukraine", 44134693, 0},
		"UG": {"UG", "Uganda", 45741007, 0},
		"US": {"US", "United States", 331002651, 0},
		"UY": {"UY", "Uruguay", 3473730, 0},
		"UZ": {"UZ", "Uzbekistan", 33469203, 0},

		// V
		"VA": {"VA", "Vatican City", 825, 0},
		"VC": {"VC", "Saint Vincent and the Grenadines", 110940, 0},
		"VE": {"VE", "Venezuela", 28435943, 0},
		"VN": {"VN", "Vietnam", 97338579, 0},
		"VU": {"VU", "Vanuatu", 307145, 0},

		// Y
		"YE": {"YE", "Yemen", 29825964, 0},

		// Z
		"ZA": {"ZA", "South Africa", 59308690, 0},
		"ZM": {"ZM", "Zambia", 18383955, 0},
		"ZW": {"ZW", "Zimbabwe", 14862924, 0},
	}
}

func main() {
	fmt.Println("🌍 Nüfus Oranlı Gerçek Rastgele Dağılım Sistemi")
	fmt.Println(strings.Repeat("=", 60))

	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Get population data
	countries := getPopulationData()

	// Total budget to distribute
	totalBudget := 3532414 // 3,532,414

	fmt.Printf("🎯 Toplam Bütçe: %d\n", totalBudget)
	fmt.Printf("🌍 Toplam Ülke: %d\n", len(countries))
	fmt.Printf("📊 Toplam Dünya Nüfusu: %d\n", getTotalPopulation(countries))

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🔄 Nüfus Oranlı Rastgele Dağılım Başlıyor...")
	fmt.Println(strings.Repeat("=", 60))

	// Create weighted country list based on population
	var weightedCountries []string
	totalPopulation := getTotalPopulation(countries)

	// Her ülke için nüfus oranında şans ver
	for code, country := range countries {
		// Nüfus oranında tekrar ekle (büyük nüfus = daha fazla şans)
		weight := (country.Population * 1000) / totalPopulation // 1000 ile çarp ki küçük ülkeler de şans bulsun
		if weight < 1 {
			weight = 1 // Minimum 1 şans
		}

		for i := 0; i < weight; i++ {
			weightedCountries = append(weightedCountries, code)
		}
	}

	fmt.Printf("📈 Ağırlıklı ülke listesi oluşturuldu: %d eleman\n", len(weightedCountries))
	fmt.Printf("🎲 Her ülkenin nüfusuna göre şansı hesaplandı\n\n")

	// Main distribution loop - 3.5 milyon döngü
	distributionCount := 0
	maxDistributions := totalBudget

	fmt.Println("🚀 Dağıtım başlıyor...")

	for distributionCount < maxDistributions {
		// Rastgele bir ülke seç (nüfus oranlı şans ile)
		randomIndex := rand.Intn(len(weightedCountries))
		selectedCountry := weightedCountries[randomIndex]

		// Seçilen ülkeye +1 ekle
		country := countries[selectedCountry]
		country.Value++
		countries[selectedCountry] = country
		distributionCount++

		// Her 100,000 dağıtımda progress göster
		if distributionCount%100000 == 0 {
			fmt.Printf("📊 %d / %d dağıtım tamamlandı (%.1f%%)\n",
				distributionCount, maxDistributions,
				float64(distributionCount)/float64(maxDistributions)*100)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("✅ Dağıtım Tamamlandı!")
	fmt.Println(strings.Repeat("=", 60))

	// Sort countries by value (descending)
	var countryList []CountryPopulation
	for _, country := range countries {
		countryList = append(countryList, country)
	}
	sort.Slice(countryList, func(i, j int) bool {
		return countryList[i].Value > countryList[j].Value
	})

	// Display results
	fmt.Printf("%-4s %-25s %-15s %-10s %-10s\n", "Kod", "Ülke Adı", "Nüfus", "Değer", "Oran")
	fmt.Println(strings.Repeat("-", 80))

	totalValue := 0
	for _, country := range countryList {
		totalValue += country.Value

		// Format population
		popStr := formatPopulation(country.Population)

		// Calculate percentage of total budget
		percentage := float64(country.Value) / float64(totalBudget) * 100

		fmt.Printf("%-4s %-25s %-15s %-10d %-10.2f%%\n",
			country.Code,
			country.Name,
			popStr,
			country.Value,
			percentage)
	}

	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("🎯 Toplam Dağıtılan: %d\n", totalValue)
	fmt.Printf("📊 Toplam Bütçe: %d\n", totalBudget)
	fmt.Printf("📈 Fark: %d\n", totalBudget-totalValue)

	if totalValue == totalBudget {
		fmt.Println("✅ Mükemmel! Tüm bütçe dağıtıldı!")
	} else {
		fmt.Println("⚠️  Bütçe tam dağıtılamadı!")
	}

	// Generate SQL INSERT statement
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📝 SQL INSERT STATEMENT:")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("INSERT OR IGNORE INTO countries (country_code, value) VALUES")

	// Group by alphabet
	groups := make(map[string][]CountryPopulation)
	for _, country := range countryList {
		firstLetter := string(country.Code[0])
		groups[firstLetter] = append(groups[firstLetter], country)
	}

	// Sort groups alphabetically
	var groupKeys []string
	for key := range groups {
		groupKeys = append(groupKeys, key)
	}
	sort.Strings(groupKeys)

	// Print grouped INSERT values
	for i, key := range groupKeys {
		groupCountries := groups[key]

		fmt.Printf("    -- %s\n", key)
		for j, country := range groupCountries {
			if i == len(groupKeys)-1 && j == len(groupCountries)-1 {
				fmt.Printf("    ('%s', %d);\n", country.Code, country.Value)
			} else {
				fmt.Printf("    ('%s', %d), ", country.Code, country.Value)
			}
		}
	}
}

// Helper functions
func getTotalPopulation(countries map[string]CountryPopulation) int {
	total := 0
	for _, country := range countries {
		total += country.Population
	}
	return total
}

func formatPopulation(population int) string {
	if population >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(population)/1000000)
	} else if population >= 1000 {
		return fmt.Sprintf("%.1fK", float64(population)/1000)
	}
	return fmt.Sprintf("%d", population)
}
