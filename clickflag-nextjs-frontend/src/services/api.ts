import { API_CONFIG, API_ENDPOINTS, HTTP_METHODS } from '@/constants';

// API Response Types
interface ApiResponse<T> {
  success: boolean;
  message: string;
  data: T;
}

interface CountriesData {
  [countryCode: string]: number;
}

interface CountryClickData {
  country_code: string;
}

// API Service Class
class ApiService {
  private baseUrl: string;
  private timeout: number;

  constructor() {
    this.baseUrl = API_CONFIG.BASE_URL;
    this.timeout = API_CONFIG.TIMEOUT;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const url = `${this.baseUrl}${endpoint}`;
    
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), this.timeout);

    try {
      const response = await fetch(url, {
        ...options,
        signal: controller.signal,
        headers: {
          'Content-Type': 'application/json',
          ...options.headers,
        },
      });

      clearTimeout(timeoutId);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data: ApiResponse<T> = await response.json();
      
      if (!data.success) {
        throw new Error(data.message || 'API request failed');
      }

      return data;
    } catch (error) {
      clearTimeout(timeoutId);
      
      if (error instanceof Error) {
        if (error.name === 'AbortError') {
          throw new Error('Request timeout');
        }
        throw error;
      }
      
      throw new Error('Unknown error occurred');
    }
  }

  // Get countries data
  async getCountries(): Promise<CountriesData> {
    const response = await this.request<CountriesData>(API_ENDPOINTS.COUNTRIES, {
      method: HTTP_METHODS.GET,
    });
    
    return response.data;
  }

  // Post country click (fire and forget)
  postCountryClick(countryCode: string): void {
    const data: CountryClickData = { country_code: countryCode };
    
    fetch(`${this.baseUrl}${API_ENDPOINTS.COUNTRIES}`, {
      method: HTTP_METHODS.POST,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    }).catch((error) => {
      // Hata olsa bile devam et
      console.warn('Country click POST failed:', error);
    });
  }
}

// Export singleton instance
export const apiService = new ApiService();
