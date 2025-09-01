import React, { memo } from 'react';
import { useTotalClicksStore } from '@/store';
import { formatNumber } from '@/utils/numberFormats';

interface FlagComponentProps {
  code: string;
  Flag: React.ComponentType<React.SVGProps<SVGSVGElement>>;
  onClick?: (countryCode: string) => void;
}

// Bayrak SVG'si - hiç değişmez
const FlagSVG = memo<{ Flag: React.ComponentType<React.SVGProps<SVGSVGElement>> }>(({ Flag }) => {
  return (
    <div className="w-full aspect-[4/3] flex-shrink-0">
      <Flag />
    </div>
  );
});

FlagSVG.displayName = 'FlagSVG';

// Loading placeholder
const FlagPlaceholder = memo(() => {
  return (
    <div className="w-full aspect-[4/3] flex-shrink-0 bg-gray-800 rounded animate-pulse"></div>
  );
});

FlagPlaceholder.displayName = 'FlagPlaceholder';

// Sadece sayı - değişir
const ClickCount = memo<{ code: string }>(({ code }) => {
  const clickCount = useTotalClicksStore((state) => state.clicks[code] || 0);
  const isLoading = useTotalClicksStore((state) => Object.keys(state.clicks).length === 0);
  
  return (
    <div className="mt-1 text-white/80 text-xs font-medium h-3 select-none pointer-events-none">
      {isLoading ? null : formatNumber(clickCount)}
    </div>
  );
});

ClickCount.displayName = 'ClickCount';

export const FlagComponent = memo<FlagComponentProps>(({ 
  code, 
  Flag,
  onClick
}) => {
  const isLoading = useTotalClicksStore((state) => Object.keys(state.clicks).length === 0);
  
  const handleClick = () => {
    onClick?.(code);
  };

  return (
    <div 
      className="w-full max-w-[80px] flex flex-col items-center justify-start cursor-pointer hover:scale-105 transition-transform select-none"
      onClick={handleClick}
      data-flag-code={code}
    >
      {isLoading ? <FlagPlaceholder /> : <FlagSVG Flag={Flag} />}
      <ClickCount code={code} />
    </div>
  );
});

FlagComponent.displayName = 'FlagComponent';
