// scripts/create-flag-components.js
// SVG dosyalarını React bileşenlerine dönüştürür

import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const FLAGS_DIR = path.join(__dirname, '../src/flags');
const OUTPUT_FILE = path.join(__dirname, '../src/app/constants/flags.tsx');

// SVG'yi React bileşenine dönüştür - En basit ve güvenli yaklaşım
function svgToComponent(svgContent, countryCode) {
  // SVG'yi temizle ve React attribute'larına çevir
  let cleanSvg = svgContent
    // React için gerekli dönüşümler
    .replace(/xmlns:xlink="[^"]*"/gi, 'xmlnsXlink="http://www.w3.org/1999/xlink"')
    .replace(/xlink:href="/gi, 'xlinkHref="')
    .replace(/xml:space="/gi, 'xmlSpace="')
    .replace(/class="/gi, 'className="')
    
    // Kebab-case'den camelCase'e dönüştür
    .replace(/fill-rule="/gi, 'fillRule="')
    .replace(/stroke-width="/gi, 'strokeWidth="')
    .replace(/stroke-linecap="/gi, 'strokeLinecap="')
    .replace(/stroke-linejoin="/gi, 'strokeLinejoin="')
    .replace(/stroke-opacity="/gi, 'strokeOpacity="')
    .replace(/stroke-miterlimit="/gi, 'strokeMiterlimit="')
    .replace(/stroke-dasharray="/gi, 'strokeDasharray="')
    .replace(/fill-opacity="/gi, 'fillOpacity="')
    .replace(/clip-path="/gi, 'clipPath="')
    .replace(/clip-rule="/gi, 'clipRule="')
    .replace(/stop-opacity="/gi, 'stopOpacity="')
    .replace(/stop-color="/gi, 'stopColor="')
    .replace(/marker-mid="/gi, 'markerMid="')
    .replace(/marker-start="/gi, 'markerStart="')
    .replace(/marker-end="/gi, 'markerEnd="')
    .replace(/text-anchor="/gi, 'textAnchor="')
    .replace(/font-family="/gi, 'fontFamily="')
    .replace(/font-size="/gi, 'fontSize="')
    .replace(/font-weight="/gi, 'fontWeight="')
    .replace(/letter-spacing="/gi, 'letterSpacing="')
    .replace(/word-spacing="/gi, 'wordSpacing="')
    
    // Problemli attribute'ları kaldır
    .replace(/\s+style="[^"]*"/gi, '') // style kaldır
    .replace(/\s+overflow="[^"]*"/gi, '') // overflow kaldır
    .replace(/\s+width="[^"]*"/gi, '') // width kaldır 
    .replace(/\s+height="[^"]*"/gi, '') // height kaldır
    
    // Boş attribute'ları temizle (boolean problemine sebep olanlar)
    .replace(/\s+stroke(?=\s|>)/g, '') // değersiz stroke
    .replace(/\s+fill(?=\s|>)/g, '') // değersiz fill
    .replace(/\s+marker(?=\s|>)/g, '') // değersiz marker
    
    // Çoklu boşlukları temizle
    .replace(/\s+/g, ' ')
    .replace(/>\s+</g, '><')
    .trim();

  // SVG tag'ini çıkar ve içeriği al
  const svgContentMatch = cleanSvg.match(/<svg[^>]*>(.*)<\/svg>/s);
  const svgContentOnly = svgContentMatch ? svgContentMatch[1].trim() : cleanSvg;
  
  return `  ${countryCode}: (props: React.SVGProps<SVGSVGElement>) => (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 640 480"
      {...props}
    >
      ${svgContentOnly}
    </svg>
  ),`;
}

async function main() {
  try {
    // SVG dosyalarını oku
    const files = fs.readdirSync(FLAGS_DIR).filter(file => file.endsWith('.svg'));
    
    if (files.length === 0) {
      console.log('SVG dosyası bulunamadı!');
      return;
    }

    // Dosyaları sırala
    files.sort();

    let components = 'import * as React from "react";\n\nconst flags = {\n';

    for (const file of files) {
      const countryCode = file.replace('.svg', '').toUpperCase();
      const filePath = path.join(FLAGS_DIR, file);
      const svgContent = fs.readFileSync(filePath, 'utf8');
      
      const component = svgToComponent(svgContent, countryCode);
      components += component + '\n';
      
      console.log(`✓ ${countryCode} dönüştürüldü`);
    }

    components += '};\n\n';
    components += 'export default flags;\n';

    // Dosyayı yaz
    fs.writeFileSync(OUTPUT_FILE, components);
    
    console.log(`\n✅ ${files.length} bayrak bileşeni oluşturuldu: ${OUTPUT_FILE}`);
    console.log('\nKullanım:');
    console.log('import flags from \'./flags\'');
    console.log('const Flag = flags[countryCode];');
    console.log('<Flag />');

  } catch (error) {
    console.error('Hata:', error);
  }
}

main();