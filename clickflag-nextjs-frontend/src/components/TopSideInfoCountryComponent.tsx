import React, { useState, useEffect, useRef, useCallback } from 'react';
import { useTotalClicksStore, useMyClicksStore } from '@/store';
import { addDots } from '@/utils/numberFormats';
import { getCountryName } from '@/constants/countries';
import flags from '@/constants/flags';

export const TopSideInfoCountryComponent: React.FC = () => {
  const [currentCountry, setCurrentCountry] = useState<string | null>(null);
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const currentCountryRef = useRef<string | null>(null);
  
  // Store'dan bilgileri al
  const totalClicks = useTotalClicksStore((state) => state.clicks);
  const myClicks = useMyClicksStore((state) => state.clicks);

  // currentCountry değiştiğinde ref'i güncelle
  useEffect(() => {
    currentCountryRef.current = currentCountry;
  }, [currentCountry]);

  // Click handler'ı useCallback ile optimize et - dependency yok
  const handleClick = useCallback((e: MouseEvent) => {
    const target = e.target as HTMLElement;
    const flagCode = target.closest('[data-flag-code]')?.getAttribute('data-flag-code');
    
    if (flagCode) {
      // Eğer aynı bayrak zaten gösteriliyorsa, sadece timeout yenile
      const current = currentCountryRef.current;
      
      if (current === flagCode) {
        // Aynı bayrak - timeout yenile
      } else {
        // Farklı bayrak - güncelle
        setCurrentCountry(flagCode);
      }
      
      // Önceki timeout'u temizle
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      
      // Yeni timeout başlat
      timeoutRef.current = setTimeout(() => {
        setCurrentCountry(null);
        timeoutRef.current = null;
      }, 3000);
    }
  }, []); // ← Hiç dependency yok!

  // Event listener'ı sadece bir kez ekle
  useEffect(() => {
    document.addEventListener('click', handleClick);
    return () => {
      document.removeEventListener('click', handleClick);
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
    };
  }, [handleClick]);

  // Görünür değilse hiçbir şey render etme
  if (!currentCountry) {
    return null;
  }

  const countryTotalClicks = totalClicks[currentCountry] || 0;
  const countryMyClicks = myClicks[currentCountry] || 0;
  const countryName = getCountryName(currentCountry);
  
  // Bayrak bileşenini al
  const FlagComponent = flags[currentCountry as keyof typeof flags];

  return (
    <div className="fixed top-4 left-4 z-50 animate-fade-in">
      <div className="bg-white/10 backdrop-blur-md rounded-xl p-3 sm:p-4 md:p-5 border border-white/20 shadow-2xl max-w-[200px] sm:max-w-xs md:max-w-sm">
        {/* Bayrak */}
        <div className="flex justify-center mb-2 sm:mb-3 md:mb-4">
          <div className="w-10 h-10 sm:w-14 sm:h-14 md:w-16 md:h-16">
            {FlagComponent ? <FlagComponent /> : <div className="w-full h-full bg-red-500 rounded flex items-center justify-center text-white font-bold text-sm sm:text-base md:text-lg">🇹🇷</div>}
          </div>
        </div>

        {/* Ülke Bilgileri */}
        <div className="space-y-1.5 sm:space-y-2 md:space-y-3">
          {/* Ülke Adı */}
          <div className="text-center">
            <h3 className="text-white font-bold text-xs sm:text-base md:text-lg truncate">
              {countryName}
            </h3>
            <p className="text-white/60 text-xs sm:text-sm md:text-base font-mono">
              {currentCountry}
            </p>
          </div>

          {/* İstatistikler */}
          <div className="space-y-1 sm:space-y-1.5 md:space-y-2 pt-1 sm:pt-2 md:pt-3">
            {/* Toplam Tıklama */}
            <div className="flex justify-between items-center text-xs sm:text-sm md:text-base">
              <span className="text-white/80 mr-2">Total Clicks</span>
              <span className="text-white font-semibold">
                {addDots(countryTotalClicks)}
              </span>
            </div>

            {/* Benim Tıklamalarım */}
            <div className="flex justify-between items-center text-xs sm:text-sm md:text-base">
              <span className="text-white/80 mr-2">My Clicks</span>
              <span className="text-white font-semibold">
                {addDots(countryMyClicks)}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
