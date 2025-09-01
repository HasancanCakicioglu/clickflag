import { useCallback } from 'react';
import { RANDOM_ICONS} from '@/constants/emojis';

export const useEmojiAnimation = () => {
  const createEmoji = useCallback(() => {
    //console.log('🎯 Emoji oluşturuluyor...');
    
    const randomIcon = RANDOM_ICONS[Math.floor(Math.random() * RANDOM_ICONS.length)];
    const randomX = Math.random() * (window.innerWidth - 100);
    const randomY = Math.random() * (window.innerHeight - 100);
    
    // Ekran boyutuna göre emoji boyutu hesapla
    const getEmojiSize = () => {
      const screenWidth = window.innerWidth;
      const screenHeight = window.innerHeight;
      const screenArea = screenWidth * screenHeight;
      
      // Ekran alanına göre boyut belirle
      if (screenArea < 500000) { // Küçük ekran (mobil)
        return '32px';
      } else if (screenArea < 1500000) { // Orta ekran (tablet)
        return '40px';
      } else if (screenArea < 3000000) { // Büyük ekran (laptop)
        return '48px';
      } else { // Çok büyük ekran (desktop)
        return '56px';
      }
    };
    
    const emojiSize = getEmojiSize();
    //console.log('🎯 Emoji:', randomIcon, 'Boyut:', emojiSize, 'Pozisyon:', randomX, randomY);
    
    // Emoji elementi oluştur - rastgele pozisyon
    const emoji = document.createElement('div');
    emoji.textContent = randomIcon;
    emoji.style.position = 'fixed';
    emoji.style.left = randomX + 'px';
    emoji.style.top = randomY + 'px';
    emoji.style.fontSize = emojiSize;
    emoji.style.pointerEvents = 'none';
    emoji.style.zIndex = '999999';
    emoji.style.color = 'white';
    emoji.style.textShadow = '2px 2px 4px rgba(0,0,0,0.8)';
    emoji.style.animation = 'bounce 2s ease-out forwards';
    
    // CSS animasyon ekle (eğer yoksa)
    if (!document.getElementById('emoji-animation-style')) {
      const style = document.createElement('style');
      style.id = 'emoji-animation-style';
      style.textContent = `
        @keyframes bounce {
          0% { transform: translateY(0) scale(1); opacity: 1; }
          50% { transform: translateY(-50px) scale(1.2); opacity: 0.8; }
          100% { transform: translateY(-100px) scale(0.8); opacity: 0; }
        }
      `;
      document.head.appendChild(style);
    }
    
    document.body.appendChild(emoji);
    //console.log('🎯 Emoji DOM\'a eklendi');
    
    // 2 saniye sonra kaldır
    setTimeout(() => {
      if (document.body.contains(emoji)) {
        document.body.removeChild(emoji);
        //console.log('🎯 Emoji kaldırıldı');
      }
    }, 2000);
  }, []);

  return { createEmoji };
};
