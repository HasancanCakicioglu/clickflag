import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  metadataBase: new URL('https://clickflag.com'),
  title: {
    default: 'ClickFlag – Most Clicked Flags and Countries',
    template: '%s | ClickFlag'
  },
  description: 'Explore the most popular country flags with real-time clicks.',
  alternates: {
    canonical: '/'
  },
  openGraph: {
    type: 'website',
    locale: 'en_US',
    url: 'https://clickflag.com/',
    siteName: 'ClickFlag',
    title: 'ClickFlag – Most Clicked Flags and Countries',
    description: 'Explore the most popular country flags with real-time clicks.'
  },
  twitter: {
    card: 'summary_large_image',
    title: 'ClickFlag – Most Clicked Flags and Countries',
    description: 'Explore the most popular country flags with real-time clicks.'
  },
  robots: {
    index: true,
    follow: true
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        {children}
        <script type="application/ld+json" suppressHydrationWarning>
          {JSON.stringify({
            '@context': 'https://schema.org',
            '@type': 'WebSite',
            name: 'ClickFlag',
            url: 'https://clickflag.com/',
            potentialAction: {
              '@type': 'SearchAction',
              target: 'https://clickflag.com/?q={search_term_string}',
              'query-input': 'required name=search_term_string'
            }
          })}
        </script>
      </body>
    </html>
  );
}
