# Scripts Klasörü

Bu klasör çeşitli yardımcı scriptleri içerir.

## 📋 Mevcut Scriptler

### 1. `validate_migration.go` - Migration Doğrulama Scripti

Bu script, migration dosyasındaki ülke kodlarını Go constants dosyası ile karşılaştırır.

#### 🎯 Ne Yapar?

- **Go Constants**: `constants/countries.go` dosyasından ülke kodlarını okur
- **CHECK Constraint**: Migration dosyasındaki CHECK constraint'teki ülke kodlarını okur  
- **INSERT Satırları**: Migration dosyasındaki INSERT satırlarındaki ülke kodlarını okur
- **Karşılaştırma**: Üç kaynak arasında eksik/fazla ülke kodlarını bulur

#### 🚀 Nasıl Çalıştırılır?

```bash
# Makefile ile (önerilen)
make validate-migration

# Veya doğrudan
cd scripts
go run validate_migration.go
```

#### 📊 Çıktı Örneği

```
🔍 Migration Validation Script - Ülke Kodları Kontrolü
============================================================
✅ Go constants: 195 ülke kodu bulundu
✅ CHECK constraint: 195 ülke kodu bulundu
✅ INSERT satırları: 195 ülke kodu bulundu

============================================================
📊 KARŞILAŞTIRMA SONUÇLARI
============================================================

🔍 Go Constants vs CHECK Constraint:
   ✅ Go Constants ve CHECK Constraint tamamen eşleşiyor

🔍 Go Constants vs INSERT Satırları:
   ✅ Go Constants ve INSERT Satırları tamamen eşleşiyor

🔍 CHECK Constraint vs INSERT Satırları:
   ✅ CHECK Constraint ve INSERT Satırları tamamen eşleşiyor

============================================================
📋 GENEL ÖZET
============================================================
🎉 TÜM ÜLKE KODLARI EŞLEŞİYOR!
   Toplam: 195 ülke kodu
```

#### ⚠️ Hata Durumları

Eğer uyumsuzluk varsa:

```
❌ CHECK Constraint'de eksik olanlar (2): ['XX', 'YY']
❌ INSERT Satırları'nda fazla olanlar (1): ['ZZ']
```

#### 🔧 Teknik Detaylar

- **Regex Pattern'ler**: Ülke kodlarını otomatik olarak bulur
- **Set Karşılaştırması**: O(n) performans ile hızlı karşılaştırma
- **Sıralama**: Alfabetik sıralama ile düzenli çıktı
- **Hata Yönetimi**: Dosya bulunamazsa uygun hata mesajları

#### 📁 Dosya Yapısı

```
scripts/
├── validate_migration.go    # Ana validation scripti
├── README.md               # Bu dosya
└── log-monitor.sh         # Log monitoring scripti
```

## 🚀 Gelecek Scriptler

- Database migration validation
- API endpoint testing
- Performance monitoring
- Data consistency checks
