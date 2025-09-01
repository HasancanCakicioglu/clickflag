import React, { memo, useMemo } from 'react';
import { useTotalClicksStore } from '@/store';
import { formatNumber } from '@/utils/numberFormats';

interface TopFlagsProps {
  flagComponents: { code: string; Flag: React.ComponentType<React.SVGProps<SVGSVGElement>> }[];
}

// Tek bir top flag component'i
const TopFlagItem = memo<{
  code: string;
  Flag: React.ComponentType<React.SVGProps<SVGSVGElement>>;
  clickCount: number;
  rank: number;
  isFirst: boolean;
}>(({ Flag, clickCount, rank, isFirst }) => {
  // Tailwind responsive sınıfları - mobil hariç daha büyük
  const sizeClass = isFirst 
    ? 'w-16 h-16 sm:w-24 sm:h-24 md:w-28 md:h-28 lg:w-32 lg:h-32 xl:w-36 xl:h-36' 
    : 'w-12 h-12 sm:w-20 sm:h-20 md:w-24 md:h-24 lg:w-28 lg:h-28 xl:w-32 xl:h-32';
  
  const clickCountClass = isFirst 
    ? 'text-sm sm:text-lg md:text-xl lg:text-2xl xl:text-2xl font-bold' 
    : 'text-xs sm:text-base md:text-lg lg:text-xl xl:text-xl font-medium';

  const rankClass = 'text-white/60 text-xs sm:text-sm md:text-sm lg:text-sm xl:text-sm';

  return (
    <div className="flex flex-col items-center">
      {/* Taç emojisi - sadece 1. için */}
      {isFirst && (
        <div className="text-3xl sm:text-5xl md:text-6xl lg:text-7xl xl:text-7xl mb-1 sm:mb-3 md:mb-4 lg:mb-4 xl:mb-4">👑</div>
      )}
      
      {/* Bayrak */}
      <div className={`${sizeClass} flex-shrink-0 mb-1 sm:mb-3 md:mb-3 lg:mb-3 xl:mb-3`}>
        <Flag />
      </div>
      
      {/* Tıklama sayısı */}
      <div className={`text-white/90 ${clickCountClass}`}>
        {formatNumber(clickCount)}
      </div>
      
      {/* Sıra numarası */}
      <div className={`${rankClass} mt-0.5 sm:mt-1 md:mt-1 lg:mt-1 xl:mt-1`}>
        #{rank}
      </div>
    </div>
  );
});

TopFlagItem.displayName = 'TopFlagItem';

export const TopFlags = memo<TopFlagsProps>(({ flagComponents }) => {
  //console.log('🏆 TopFlags render oldu');
  
  const totalClicks = useTotalClicksStore((state) => state.clicks);

  const top3Flags = useMemo(() => {
    // Tüm bayrakları tıklama sayısına göre sırala
    const sortedFlags = [...flagComponents]
      .map(({ code, Flag }) => ({
        code,
        Flag,
        clickCount: totalClicks[code] || 0
      }))
      .sort((a, b) => b.clickCount - a.clickCount)
      .slice(0, 3); // İlk 3'ü al

    return sortedFlags;
  }, [flagComponents, totalClicks]);

  // Eğer hiç tıklama yoksa gösterme
  if (top3Flags.length === 0 || top3Flags[0].clickCount === 0) {
    return null;
  }

           return (
      <div className="mt-12 sm:mt-12 md:mt-0 lg:mt-4 xl:mt-4">
                 <div className="flex justify-center items-end gap-4 sm:gap-8 md:gap-10 lg:gap-12 xl:gap-16 mb-2 sm:mb-4 md:mb-6 lg:mb-8 xl:mb-8">
         {/* 2. sıra - sol */}
         {top3Flags[1] && (
           <TopFlagItem
             code={top3Flags[1].code}
             Flag={top3Flags[1].Flag}
             clickCount={top3Flags[1].clickCount}
             rank={2}
             isFirst={false}
           />
         )}
         
         {/* 1. sıra - ortada */}
         <div className="flex flex-col items-center -mt-4 sm:-mt-8 md:-mt-10 lg:-mt-12 xl:-mt-12">
           <TopFlagItem
             code={top3Flags[0].code}
             Flag={top3Flags[0].Flag}
             clickCount={top3Flags[0].clickCount}
             rank={1}
             isFirst={true}
           />
         </div>
         
         {/* 3. sıra - sağ */}
         {top3Flags[2] && (
           <TopFlagItem
             code={top3Flags[2].code}
             Flag={top3Flags[2].Flag}
             clickCount={top3Flags[2].clickCount}
             rank={3}
             isFirst={false}
           />
         )}
        </div>
        
                 {/* Ayraç */}
         <div className="flex justify-center pt-1 sm:pt-2 md:pt-3 lg:pt-4 xl:pt-4 pb-4 sm:pb-6 md:pb-8 lg:pb-10 xl:pb-12">
          <div className="w-11/12 sm:w-10/12 md:w-9/12 lg:w-8/12 xl:w-8/12 h-0.5 sm:h-1 md:h-1 lg:h-1 xl:h-1 bg-gradient-to-r from-transparent via-gray-400 to-transparent"></div>
        </div>
      </div>
   );
 });

TopFlags.displayName = 'TopFlags';
