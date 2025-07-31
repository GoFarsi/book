// Workbox configuration for Hugo static site
// See: https://developer.chrome.com/docs/workbox/

module.exports = {
  globDirectory: "static/",
  globPatterns: [
    "**/*.{html,js,css,json,svg,png,jpg,jpeg,webp,woff2,woff,ttf,ico,xml,md}"
  ],
  swDest: "static/sw-workbox.js",
  clientsClaim: true,
  skipWaiting: true,
  runtimeCaching: [
    // Fonts: CacheFirst for best performance, long expiration
    {
      urlPattern: /\.(?:woff2|woff|ttf)$/,
      handler: "CacheFirst",
      options: {
        cacheName: "font-cache",
        expiration: {
          maxEntries: 20,
          maxAgeSeconds: 365 * 24 * 60 * 60 // 1 year
        }
      }
    },
    // Static assets: StaleWhileRevalidate for fast loads and updates
    {
      urlPattern: /\/.+\.(?:js|css|png|jpg|jpeg|svg|webp|ico|xml|json|md)$/,
      handler: "StaleWhileRevalidate",
      options: {
        cacheName: "assets-cache",
        expiration: {
          maxEntries: 150,
          maxAgeSeconds: 60 * 24 * 60 * 60 // 60 Days
        }
      }
    },
    // HTML: NetworkFirst for freshness, fallback to cache
    {
      urlPattern: /\/.+\.(?:html)$/,
      handler: "NetworkFirst",
      options: {
        cacheName: "html-cache",
        expiration: {
          maxEntries: 80,
          maxAgeSeconds: 14 * 24 * 60 * 60 // 14 Days
        }
      }
    }
  ]
};
