import React from 'react';
import { addDots } from '@/utils/numberFormats';

interface MainCounterProps {
  title: string;
  count: number;
}

export const MainCounter = ({ title, count }: MainCounterProps) => {
  return (
    <div className="bg-white/10 backdrop-blur-sm rounded-lg p-2 sm:p-3 md:p-4 border border-white/20 shadow-lg">
      <div className="flex items-center gap-2 sm:gap-3">
        <div className="text-white/80 text-xs sm:text-sm md:text-base lg:text-lg xl:text-lg font-medium">
          {title}
        </div>
        <div className="text-white text-sm sm:text-base md:text-lg lg:text-xl xl:text-2xl font-bold">
          {addDots(count)}
        </div>
      </div>
    </div>
  );
};
