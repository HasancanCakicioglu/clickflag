import { useEffect } from 'react';
import { useTotalClicksStore } from '@/store';
import { apiService } from '@/services/api';

export const useTotalClicksSync = () => {
  const setBulk = useTotalClicksStore((state) => state.setBulk);

  useEffect(() => {
    // İlk veriyi al
    const fetchData = async () => {
      try {
        const data = await apiService.getCountries();
        setBulk(data);
        //console.log('🔄 UTC sync: Veri güncellendi');
      } catch (error) {
        console.error('UTC sync failed:', error);
      }
    };

    // İlk veriyi hemen al
    fetchData();
    //console.log('🔄 İlk veri hemen alınıyor...');

    // UTC'ye göre bir sonraki 5 saniyelik interval'i hesapla
    const getNextInterval = () => {
      const now = Date.now();
      const intervalMs = 5000; // 5 saniye
      // Şu anki 5 saniyelik periyodun sonuna kadar bekle + 10ms buffer
      return Math.floor(now / intervalMs) * intervalMs + intervalMs + 10;
    };

    // İlk interval'e kadar bekle, sonra senkronize et
    const nextInterval = getNextInterval();
    const initialDelay = Math.max(0, nextInterval - Date.now());
    
    //console.log(`🔄 UTC sync: ${initialDelay}ms sonra başlayacak`);
    
    const initialTimer = setTimeout(() => {
      // Sonraki veriyi al
      fetchData();
      
      // Sonra her 5 saniyede bir devam et
      const interval = setInterval(fetchData, 5000);
      
      // Cleanup
      return () => clearInterval(interval);
    }, initialDelay);

    // Cleanup
    return () => {
      clearTimeout(initialTimer);
    };
  }, [setBulk]);
};
