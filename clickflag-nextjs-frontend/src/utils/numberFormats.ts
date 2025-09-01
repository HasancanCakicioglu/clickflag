// Sayıyı kısaltma fonksiyonu
export const formatNumber = (num: number): string => {
    if (num < 1000) return num.toString();
    if (num < 1000000) return (num / 1000).toFixed(1).replace(/\.0$/, '') + 'K';
    if (num < 1000000000) return (num / 1000000).toFixed(1).replace(/\.0$/, '') + 'M';
    return (num / 1000000000).toFixed(1).replace(/\.0$/, '') + 'B';
  };
  
  // Her 3 hanede bir nokta koyan basit fonksiyon (büyük sayılar için güvenli)
  export const addDots = (num: number): string => {
    // Çok büyük sayılar için kısaltma kullan
    if (num > 1e15) {
      return formatNumber(num);
    }
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
  };
  