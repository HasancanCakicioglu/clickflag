// App Configuration
export const APP_CONFIG = {
  NAME: 'ClickFlag',
  VERSION: '1.0.0',
  ENVIRONMENT: process.env.NODE_ENV || 'development',
} as const;

// Page Titles
export const PAGE_TITLES = {
  HOME: 'ClickFlag - Home',
  ABOUT: 'ClickFlag - About',
} as const;
