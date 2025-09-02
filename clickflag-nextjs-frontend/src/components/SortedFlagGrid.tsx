import React, { memo, useMemo, useEffect, useRef } from 'react';
import { FlagComponent } from '@/components/FlagComponent';
import { useTotalClicksStore, useMyClicksStore } from '@/store';
import { apiService } from '@/services/api';
import { useEmojiAnimation } from '@/hooks/useEmojiAnimation';

interface SortedFlagGridProps {
  flagComponents: { code: string; Flag: React.ComponentType<React.SVGProps<SVGSVGElement>> }[];
}

// Sadece sıralama yapan component
const FlagSorter = memo<{ 
  flagComponents: { code: string; Flag: React.ComponentType<React.SVGProps<SVGSVGElement>> }[];
  children: (sortedFlags: { code: string; Flag: React.ComponentType<React.SVGProps<SVGSVGElement>> }[]) => React.ReactNode;
}>(({ flagComponents, children }) => {
  const totalClicks = useTotalClicksStore((state) => state.clicks);

  const sortedFlags = useMemo(() => {
    return [...flagComponents].sort((a, b) => {
      const aClicks = totalClicks[a.code] || 0;
      const bClicks = totalClicks[b.code] || 0;
      return bClicks - aClicks; // En çok tıklanan üstte
    });
  }, [flagComponents, totalClicks]);

  return <>{children(sortedFlags)}</>;
});

FlagSorter.displayName = 'FlagSorter';

// Ana grid component - hiç değişmez
const FlagGridRenderer = memo<{ 
  flagComponents: { code: string; Flag: React.ComponentType<React.SVGProps<SVGSVGElement>> }[];
  onFlagClick: (countryCode: string) => void;
}>(({ flagComponents, onFlagClick }) => {
  //console.log('🟣 FlagGridRenderer render oldu');
  
  return (
    <div className="px-4 sm:px-6 md:px-8 lg:px-12 xl:px-16">
      <div className="grid grid-cols-6 sm:grid-cols-8 md:grid-cols-10 lg:grid-cols-12 xl:grid-cols-16 gap-3 sm:gap-4 md:gap-5 lg:gap-6 xl:gap-8">
        {flagComponents.map(({ code, Flag }) => (
          <FlagComponent
            key={code}
            code={code}
            Flag={Flag}
            onClick={onFlagClick}
          />
        ))}
      </div>
    </div>
  );
});

FlagGridRenderer.displayName = 'FlagGridRenderer';

export const SortedFlagGrid = memo<SortedFlagGridProps>(({ flagComponents }) => {
  //console.log('🟣 SortedFlagGrid render oldu');
  
  const incrementMyClick = useMyClicksStore((state) => state.increment);
  const incrementTotalClick = useTotalClicksStore((state) => state.increment);
  const { createEmoji } = useEmojiAnimation();

  // Tıklama sesi
  const clickAudioRef = useRef<HTMLAudioElement | null>(null);
  useEffect(() => {
    const audio = new Audio('/click-sound.mp3');
    audio.preload = 'auto';
    audio.volume = 0.4; 
    clickAudioRef.current = audio;
    return () => {
      clickAudioRef.current = null;
    };
  }, []);

  const handleFlagClick = (countryCode: string) => {
    //console.log('🟣 Flag tıklandı:', countryCode);
    
    // Store'ları güncelle
    incrementMyClick(countryCode);    // MyClicks artır
    incrementTotalClick(countryCode); // TotalClicks artır
    
    // API'ye POST isteği gönder (fire and forget)
    apiService.postCountryClick(countryCode);
    
    // Tıklama sesini çal
    try {
      const audio = clickAudioRef.current;
      if (audio) {
        // Art arda tıklamalarda başa sar
        audio.currentTime = 0;
        void audio.play();
      }
    } catch {}

    // Emoji oluştur
    //console.log('🟣 createEmoji çağrılıyor...');
    createEmoji();
    //console.log('🟣 createEmoji çağrıldı');
  };

  return (
    <FlagSorter flagComponents={flagComponents}>
      {(sortedFlags) => (
        <FlagGridRenderer 
          flagComponents={sortedFlags} 
          onFlagClick={handleFlagClick}
        />
      )}
    </FlagSorter>
  );
});

SortedFlagGrid.displayName = 'SortedFlagGrid';
