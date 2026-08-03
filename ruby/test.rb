require_relative 'lib/ccpayment'

# Get credentials from environment variables
app_id = ENV['CCPAYMENT_APP_ID'] || ''
app_secret = ENV['CCPAYMENT_APP_SECRET'] || ''
proxy_url = ENV['CCPAYMENT_PROXY_URL'] || ''

if app_id.empty? || app_secret.empty?
  puts 'Error: Please set CCPAYMENT_APP_ID and CCPAYMENT_APP_SECRET environment variables'
  exit 1
end

puts "API Key: #{app_id}"
puts "Using Proxy: #{!proxy_url.empty?}"
puts '=' * 50

# Create client
client = CCPayment::Client.new(app_id, app_secret)

# Configure HTTP proxy if provided
if !proxy_url.empty?
  client.set_proxy("http://#{proxy_url}")
  puts "Proxy configured: #{proxy_url}"
end

# Test 1: Get coin list
puts '\n[Test 1] Get Coin List...'
begin
  response = client.basic_info.get_coin_list
  coins = response['coins'] || []
  puts "✓ Successfully retrieved coin list, total #{coins.length} coins"
  if coins.length > 0
    puts "  First coin: #{coins[0]['symbol']} (ID: #{coins[0]['coinId']})"
  end
rescue CCPayment::APIError => e
  puts "✗ API Error: code=#{e.code}, message=#{e.message}"
rescue => e
  puts "✗ Error: #{e.message}"
end

# Test 2: Get merchant assets
puts '\n[Test 2] Get Merchant Assets...'
begin
  response = client.merchant_assets.get_app_coin_asset_list
  assets = response['assets'] || []
  puts "✓ Successfully retrieved merchant assets, total #{assets.length} assets"
  if assets.length > 0
    assets[0..4].each do |asset|
      puts "  #{asset['coinSymbol']}: #{asset['available']}"
    end
    if assets.length > 5
      puts "  ... and #{assets.length - 5} more"
    end
  end
rescue CCPayment::APIError => e
  puts "✗ API Error: code=#{e.code}, message=#{e.message}"
rescue => e
  puts "✗ Error: #{e.message}"
end

puts '\n' + '=' * 50
puts 'Test completed'
