import React from 'react';
import { MainCounter } from './CounterComponent';
import { useTotalClicksStore } from '@/store';

export const TotalClicksCounter = () => {
  //console.log('🟡 TotalClicksCounter render oldu');
  const total = useTotalClicksStore((state) => state.total);
  
  return <MainCounter title="Total Clicks" count={total} />;
};
