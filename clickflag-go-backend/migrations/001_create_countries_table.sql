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

-- Insert initial data for all valid country codes with population-based values (192 countries)
INSERT OR IGNORE INTO countries (country_code, value) VALUES
    -- A
    ('AF', 17454), ('AR', 17438), ('AO', 13897), ('AU', 10324), ('AL', 3505),
    ('AG', 3497), ('AE', 3475), ('AZ', 3448), ('AT', 3436), ('AM', 3420),
    ('AD', 3369),
    -- B
    ('BR', 93997), ('BD', 72715), ('BF', 7043), ('BW', 3563), ('BB', 3533),
    ('BO', 3506), ('BT', 3505), ('BI', 3497), ('BZ', 3489), ('BN', 3480),
    ('BJ', 3477), ('BY', 3474), ('BH', 3474), ('BA', 3453), ('BE', 3419),
    ('BG', 3418), ('BS', 3415),
    -- C
    ('CN', 649991), ('CD', 38357), ('CO', 20836), ('CA', 13839), ('CM', 10564),
    ('CI', 10432), ('CL', 6937), ('CR', 3581), ('CF', 3539), ('CZ', 3508),
    ('CH', 3480), ('CG', 3469), ('CY', 3434), ('CV', 3422), ('CU', 3414),
    -- D
    ('DE', 34747), ('DZ', 17296), ('DJ', 3552), ('DO', 3517), ('DK', 3443),
    ('DM', 3438),
    -- E
    ('ET', 48630), ('EG', 45256), ('ES', 20721), ('EC', 6842), ('EE', 3504),
    ('ER', 3368),
    -- F
    ('FR', 27882), ('FM', 3590), ('FI', 3557), ('FJ', 3425),
    -- G
    ('GH', 13851), ('GT', 6969), ('GY', 3531), ('GQ', 3517), ('GR', 3517),
    ('GD', 3512), ('GN', 3500), ('GM', 3493), ('GE', 3491), ('GA', 3488),
    ('GW', 3469),
    -- H
    ('HN', 3467), ('HU', 3458), ('HR', 3440), ('HT', 3421),
    -- I
    ('IN', 622395), ('ID', 121123), ('IR', 34706), ('IT', 23939), ('IQ', 17328),
    ('IE', 3621), ('IL', 3358), ('IS', 0),
    -- J
    ('JP', 55587), ('JO', 3519), ('JM', 3446),
    -- K
    ('KR', 20907), ('KE', 20727), ('KP', 10527), ('KH', 7018), ('KZ', 6955),
    ('KI', 3537), ('KG', 3532), ('KM', 3507), ('KW', 3483), ('KN', 3405),
    -- L
    ('LK', 6809), ('LV', 3531), ('LB', 3524), ('LA', 3487), ('LC', 3463),
    ('LU', 3457), ('LT', 3421), ('LI', 3421), ('LR', 3395), ('LY', 3389),
    ('LS', 3359),
    -- M
    ('MX', 55556), ('MM', 24318), ('MY', 13941), ('MZ', 13935), ('MA', 13889),
    ('MG', 10473), ('ML', 6992), ('MW', 6956), ('MH', 3547), ('MK', 3523),
    ('MR', 3523), ('MT', 3473), ('MD', 3473), ('MN', 3472), ('MC', 3440),
    ('MV', 3431), ('ME', 3427), ('MU', 3425),
    -- N
    ('NG', 90146), ('NP', 10423), ('NE', 10322), ('NL', 7055), ('NA', 3508),
    ('NR', 3444), ('NI', 3398), ('NO', 3361), ('NZ', 0),
    -- O
    ('OM', 3417),
    -- P
    ('PK', 97452), ('PH', 48825), ('PE', 14124), ('PL', 13819), ('PS', 3607),
    ('PY', 3507), ('PW', 3505), ('PG', 3491), ('PT', 3456), ('PA', 3456),
    -- Q
    ('QA', 3452),
    -- R
    ('RU', 62715), ('RO', 6883), ('RS', 3557), ('RW', 3489),
    -- S
    ('SD', 17467), ('SA', 13696), ('SY', 7046), ('SO', 6920), ('SN', 6826),
    ('SI', 3647), ('SZ', 3634), ('SV', 3539), ('SB', 3538), ('SE', 3518),
    ('SK', 3512), ('SG', 3471), ('SM', 3468), ('ST', 3459), ('SR', 3454),
    ('SL', 3423), ('SC', 3322),
    -- T
    ('TR', 34945), ('TH', 31342), ('TZ', 24166), ('TW', 10489), ('TD', 6994),
    ('TT', 3619), ('TJ', 3565), ('TN', 3540), ('TM', 3490), ('TO', 3462),
    ('TL', 3455), ('TG', 3436), ('TV', 3384),
    -- U
    ('US', 149930), ('UG', 17507), ('UA', 17218), ('UZ', 13772), ('UY', 3596),
    ('GB', 0),
    -- V
    ('VN', 42199), ('VE', 10631), ('VC', 3560), ('VA', 3542), ('VU', 3494),
    -- W
    ('WS', 3523),
    -- Y
    ('YE', 10450),
    -- Z
    ('ZA', 24427), ('ZM', 6846), ('ZW', 3394);
