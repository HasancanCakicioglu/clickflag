-- Migration 002: Replace country code 'TW' with 'SS' and update CHECK constraint
-- Safe table rebuild pattern for SQLite

BEGIN IMMEDIATE;

-- 1) Yeni tablo: CHECK listesi 'TW' çıkarıldı, 'SS' eklendi
CREATE TABLE IF NOT EXISTS countries_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    country_code VARCHAR(3) NOT NULL UNIQUE,
    value INTEGER NOT NULL DEFAULT 0,
    CHECK (country_code IN (
        'AF','AL','DZ','AD','AO','AG','AR','AM','AU','AT','AZ','BS','BH','BD','BB','BY','BE','BZ','BJ','BT','BO','BA','BW','BR','BN','BG','BF','BI','KH','CM','CA','CV','CF','TD','CL','CN','CO','KM','CG','CD','CR','CI','HR','CU','CY','CZ','DK','DJ','DM','DO','EC','EG','SV','GQ','ER','EE','ET','FJ','FI','FR','GA','GM','GE','DE','GH','GR','GD','GT','GN','GW','GY','HT','HN','HU','IS','IN','ID','IR','IQ','IE','IL','IT','JM','JP','JO','KZ','KE','KI','KP','KR','KW','KG','LA','LV','LB','LS','LR','LY','LI','LT','LU','MK','MG','MW','MY','MV','ML','MT','MH','MR','MU','MX','FM','MD','MC','MN','ME','MA','MZ','MM','NA','NR','NP','NL','NZ','NI','NE','NG','NO','OM','PK','PW','PS','PA','PG','PY','PE','PH','PL','PT','QA','RO','RU','RW','KN','LC','VC','WS','SM','ST','SA','SN','RS','SC','SL','SG','SK','SI','SB','SO','ZA','ES','LK','SD','SR','SZ','SE','CH','SY','SS','TJ','TZ','TH','TL','TG','TO','TT','TN','TR','TM','TV','UG','UA','AE','GB','US','UY','UZ','VA','VU','VE','VN','YE','ZM','ZW'
    ))
);

-- 2) Veriyi taşırken TW -> SS dönüşümü
INSERT INTO countries_new (id, country_code, value)
SELECT
    id,
    CASE WHEN country_code = 'TW' THEN 'SS' ELSE country_code END AS country_code,
    value
FROM countries;

-- 3) Eski tablo ile değiştir
DROP TABLE countries;
ALTER TABLE countries_new RENAME TO countries;

-- 4) İndeksi yeniden oluştur
CREATE INDEX IF NOT EXISTS idx_countries_country_code ON countries(country_code);

COMMIT;