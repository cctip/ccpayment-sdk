FROM php:8.1-cli

# Install required PHP extensions
RUN apt-get update && apt-get install -y \
    git \
    unzip \
    libcurl4-openssl-dev \
    pkg-config \
    && docker-php-ext-install curl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /app

# Copy SDK files
COPY composer.json ./
COPY src/ ./src/
COPY test.php ./

# Install dependencies
RUN composer install --no-dev --optimize-autoloader

# Set environment variables (will be overridden at runtime)
ENV CCPAYMENT_APP_ID=""
ENV CCPAYMENT_APP_SECRET=""
ENV CCPAYMENT_PROXY_URL=""

# Run test
CMD ["php", "test.php"]
