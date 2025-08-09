# Development Stage
# FROM node:18-alpine AS development

# WORKDIR /app

# COPY package*.json ./

# RUN npm ci

# COPY . .

# EXPOSE 3000


# CMD ["npm", "run", "dev"]

# Builder Stage
FROM node:18-alpine 

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 3000

RUN npm run build
CMD ["npm", "run","start"]

# Production Stage 

# FROM node:18-alpine AS production

# WORKDIR /app

# Copy the built artifacts from the builder stage
# COPY --from=builder /app/.next/standalone ./
# COPY --from=builder /app/.next/static ./.next/static

# Set the environment variables (if needed)

