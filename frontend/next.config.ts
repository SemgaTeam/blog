import type { NextConfig } from "next"

const nextConfig: NextConfig = {
  reactStrictMode: true,

  // если Go API на другом домене
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "**"
      }
    ]
  },
    async rewrites() {
    return [
      {
        source: '/api/:path*',            
        destination: 'http://localhost:8080/api/:path*', 
      },
    ];
  },
  
}

export default nextConfig
