import { create } from 'zustand';

interface TotalClicksStore {
  clicks: Record<string, number>;  // TR: 45, US: 89
  total: number;                   // 134 (ayrı tutuyoruz)
  
  increment: (countryCode: string) => void;
  set: (countryCode: string, count: number) => void;
  setBulk: (data: Record<string, number>) => void; // Toplu set etme
  getTotal: () => number;
}

export const useTotalClicksStore = create<TotalClicksStore>((set, get) => ({
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
  
  setBulk: (data) => set((state) => {
    // Önceki total ile yeni total'i karşılaştır
    const oldTotal = state.total;
    const newDataTotal = Object.values(data).reduce((sum, count) => sum + count, 0);
    
    // Yeni veri daha azsa, farkı total'e ekle (veri kaybını önle)
    const totalDifference = Math.max(0, oldTotal - newDataTotal);
    
    return {
      clicks: data,
      total: newDataTotal + totalDifference
    };
  }),
  
  getTotal: () => get().total
}));
