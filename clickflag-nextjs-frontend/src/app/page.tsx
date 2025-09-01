"use client";

import { SortedFlagGrid } from '@/components/SortedFlagGrid';
import { TopFlags } from '@/components/TopFlags';
import { MyClicksCounter } from '@/components/MyClicksCounter';
import { TotalClicksCounter } from '@/components/TotalClicksCounter';
import { TopSideInfoCountryComponent } from '@/components/TopSideInfoCountryComponent';
import { useTotalClicksSync } from '@/hooks/useTotalClicksSync';
import flags from '@/constants/flags';

export default function Home() {
  //console.log('🟢 Page.tsx render oldu'); // Test için

  // API sync hook'u - re-render etmez
  useTotalClicksSync();

  const sampleFlags = Object.entries(flags).map(([code, Flag]) => ({
    code,
    Flag: Flag as React.ComponentType<React.SVGProps<SVGSVGElement>>
  }));

  return (
    <div className="min-h-screen bg-[#1e1f23] p-4 relative overflow-hidden">
      <div className="fixed top-4 right-4 flex gap-2">
        <MyClicksCounter />
        <TotalClicksCounter />
      </div>
      
      {/* Country Info Component - Sol üstte */}
      <TopSideInfoCountryComponent />
      
      <div className="pt-12">
        <TopFlags flagComponents={sampleFlags} />
        <SortedFlagGrid flagComponents={sampleFlags} />
      </div>
    </div>
  );
}
