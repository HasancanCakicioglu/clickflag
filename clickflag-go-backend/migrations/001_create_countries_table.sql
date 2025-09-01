-- Migration: Create countries table
-- Date: 2024-01-01

CREATE TABLE IF NOT EXISTS countries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    country_code VARCHAR(3) NOT NULL UNIQUE,
    value INTEGER NOT NULL DEFAULT 0,
    CHECK (country_code IN ('AF', 'AL', 'DZ', 'AD', 'AO', 'AG', 'AR', 'AM', 'AU', 'AT', 'AZ', 'BS', 'BH', 'BD', 'BB', 'BY', 'BE', 'BZ', 'BJ', 'BT', 'BO', 'BA', 'BW', 'BR', 'BN', 'BG', 'BF', 'BI', 'KH', 'CM', 'CA', 'CV', 'CF', 'TD', 'CL', 'CN', 'CO', 'KM', 'CG', 'CD', 'CR', 'CI', 'HR', 'CU', 'CY', 'CZ', 'DK', 'DJ', 'DM', 'DO', 'EC', 'EG', 'SV', 'GQ', 'ER', 'EE', 'ET', 'FJ', 'FI', 'FR', 'GA', 'GM', 'GE', 'DE', 'GH', 'GR', 'GD', 'GT', 'GN', 'GW', 'GY', 'HT', 'HN', 'HU', 'IS', 'IN', 'ID', 'IR', 'IQ', 'IE', 'IL', 'IT', 'JM', 'JP', 'JO', 'KZ', 'KE', 'KI', 'KP', 'KR', 'KW', 'KG', 'LA', 'LV', 'LB', 'LS', 'LR', 'LY', 'LI', 'LT', 'LU', 'MK', 'MG', 'MW', 'MY', 'MV', 'ML', 'MT', 'MH', 'MR', 'MU', 'MX', 'FM', 'MD', 'MC', 'MN', 'ME', 'MA', 'MZ', 'MM', 'NA', 'NR', 'NP', 'NL', 'NZ', 'NI', 'NE', 'NG', 'NO', 'OM', 'PK', 'PW', 'PS', 'PA', 'PG', 'PY', 'PE', 'PH', 'PL', 'PT', 'QA', 'RO', 'RU', 'RW', 'KN', 'LC', 'VC', 'WS', 'SM', 'ST', 'SA', 'SN', 'RS', 'SC', 'SL', 'SG', 'SK', 'SI', 'SB', 'SO', 'ZA', 'ES', 'LK', 'SD', 'SR', 'SZ', 'SE', 'CH', 'SY', 'TW', 'TJ', 'TZ', 'TH', 'TL', 'TG', 'TO', 'TT', 'TN', 'TR', 'TM', 'TV', 'UG', 'UA', 'AE', 'GB', 'US', 'UY', 'UZ', 'VA', 'VU', 'VE', 'VN', 'YE', 'ZM', 'ZW'))
);

-- Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_countries_country_code ON countries(country_code);

-- Insert initial data for all valid country codes (195 countries from constants)
INSERT OR IGNORE INTO countries (country_code, value) VALUES
    ('AF', 0), ('AL', 0), ('DZ', 0), ('AD', 0), ('AO', 0),
    ('AG', 0), ('AR', 0), ('AM', 0), ('AU', 0), ('AT', 0),
    ('AZ', 0), ('BS', 0), ('BH', 0), ('BD', 0), ('BB', 0),
    ('BY', 0), ('BE', 0), ('BZ', 0), ('BJ', 0), ('BT', 0),
    ('BO', 0), ('BA', 0), ('BW', 0), ('BR', 0), ('BN', 0),
    ('BG', 0), ('BF', 0), ('BI', 0), ('KH', 0), ('CM', 0),
    ('CA', 0), ('CV', 0), ('CF', 0), ('TD', 0), ('CL', 0),
    ('CN', 0), ('CO', 0), ('KM', 0), ('CG', 0), ('CD', 0),
    ('CR', 0), ('CI', 0), ('HR', 0), ('CU', 0), ('CY', 0),
    ('CZ', 0), ('DK', 0), ('DJ', 0), ('DM', 0), ('DO', 0),
    ('EC', 0), ('EG', 0), ('SV', 0), ('GQ', 0), ('ER', 0),
    ('EE', 0), ('ET', 0), ('FJ', 0), ('FI', 0), ('FR', 0),
    ('GA', 0), ('GM', 0), ('GE', 0), ('DE', 0), ('GH', 0),
    ('GR', 0), ('GD', 0), ('GT', 0), ('GN', 0), ('GW', 0),
    ('GY', 0), ('HT', 0), ('HN', 0), ('HU', 0), ('IS', 0),
    ('IN', 0), ('ID', 0), ('IR', 0), ('IQ', 0), ('IE', 0),
    ('IL', 0), ('IT', 0), ('JM', 0), ('JP', 0), ('JO', 0),
    ('KZ', 0), ('KE', 0), ('KI', 0), ('KP', 0), ('KR', 0),
    ('KW', 0), ('KG', 0), ('LA', 0), ('LV', 0), ('LB', 0),
    ('LS', 0), ('LR', 0), ('LY', 0), ('LI', 0), ('LT', 0),
    ('LU', 0), ('MK', 0), ('MG', 0), ('MW', 0), ('MY', 0),
    ('MV', 0), ('ML', 0), ('MT', 0), ('MH', 0), ('MR', 0),
    ('MU', 0), ('MX', 0), ('FM', 0), ('MD', 0), ('MC', 0),
    ('MN', 0), ('ME', 0), ('MA', 0), ('MZ', 0), ('MM', 0),
    ('NA', 0), ('NR', 0), ('NP', 0), ('NL', 0), ('NZ', 0),
    ('NI', 0), ('NE', 0), ('NG', 0), ('NO', 0), ('OM', 0),
    ('PK', 0), ('PW', 0), ('PS', 0), ('PA', 0), ('PG', 0),
    ('PY', 0), ('PE', 0), ('PH', 0), ('PL', 0), ('PT', 0),
    ('QA', 0), ('RO', 0), ('RU', 0), ('RW', 0), ('KN', 0),
    ('LC', 0), ('VC', 0), ('WS', 0), ('SM', 0), ('ST', 0),
    ('SA', 0), ('SN', 0), ('RS', 0), ('SC', 0), ('SL', 0),
    ('SG', 0), ('SK', 0), ('SI', 0), ('SB', 0), ('SO', 0),
    ('ZA', 0), ('ES', 0), ('LK', 0), ('SD', 0), ('SR', 0),
    ('SZ', 0), ('SE', 0), ('CH', 0), ('SY', 0), ('TW', 0),
    ('TJ', 0), ('TZ', 0), ('TH', 0), ('TL', 0), ('TG', 0),
    ('TO', 0), ('TT', 0), ('TN', 0), ('TR', 0), ('TM', 0),
    ('TV', 0), ('UG', 0), ('UA', 0), ('AE', 0), ('GB', 0),
    ('US', 0), ('UY', 0), ('UZ', 0), ('VA', 0), ('VU', 0),
    ('VE', 0), ('VN', 0), ('YE', 0), ('ZM', 0), ('ZW', 0);
