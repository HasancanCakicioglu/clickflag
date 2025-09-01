import React from 'react';
import { MainCounter } from './CounterComponent';
import { useMyClicksStore } from '@/store';

export const MyClicksCounter = () => {
  //console.log('🔵 MyClicksCounter render oldu');
  const total = useMyClicksStore((state) => state.total);
  
  return <MainCounter title="My Clicks" count={total} />;
};
