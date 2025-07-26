/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        serverActions: {
            bodySizeLimit: '5mb',  // config the body size limit because the default is 1mb only
        },
    },
    reactStrictMode: false
};
export default nextConfig;