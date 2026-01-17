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
  }
}

export default nextConfig
