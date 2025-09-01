import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface MyClicksStore {
  clicks: Record<string, number>;  // TR: 23, US: 56
  total: number;                   // 79 (ayrı tutuyoruz)
  
  increment: (countryCode: string) => void;
  set: (countryCode: string, count: number) => void;
  getTotal: () => number;
}

export const useMyClicksStore = create<MyClicksStore>()(
  persist(
    (set, get) => ({
      clicks: {},
      total: 0,
      
      increment: (countryCode) => set((state) => ({
        clicks: {
          ...state.clicks,
          [countryCode]: (state.clicks[countryCode] || 0) + 1
        },
        total: state.total + 1
      })),
      
      set: (countryCode, count) => set((state) => {
        const oldCount = state.clicks[countryCode] || 0;
        const difference = count - oldCount;
        
        return {
          clicks: { ...state.clicks, [countryCode]: count },
          total: state.total + difference
        };
      }),
      
      getTotal: () => get().total
    }),
    {
      name: 'my-clicks-storage', // localStorage key'i
    }
  )
);
