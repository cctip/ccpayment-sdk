#!/usr/bin/env python3
import os
import sys

# Add parent directory to path
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

from ccpayment import Client
from ccpayment.exceptions import APIError


def main():
    # Get credentials from environment variables
    app_id = os.environ.get('CCPAYMENT_APP_ID')
    app_secret = os.environ.get('CCPAYMENT_APP_SECRET')
    proxy_url = os.environ.get('CCPAYMENT_PROXY_URL')

    if not app_id or not app_secret:
        print("Error: Please set CCPAYMENT_APP_ID and CCPAYMENT_APP_SECRET environment variables")
        sys.exit(1)

    print(f"API Key: {app_id}")
    print(f"Using Proxy: {bool(proxy_url)}")
    print("=" * 50)

    # Create client
    client = Client(app_id, app_secret)

    # Configure HTTP proxy if provided
    if proxy_url:
        client.set_proxy(f"http://{proxy_url}")
        print(f"Proxy configured: {proxy_url}")

    # Test 1: Get coin list
    print("\n[Test 1] Get Coin List...")
    try:
        response = client.basic_info.get_coin_list()
        print(f"DEBUG: response type={type(response)}, value={response}")
        if isinstance(response, list):
            print(f"✓ Successfully retrieved coin list, total {len(response)} coins")
            if response:
                print(f"  First coin: {response[0].get('symbol', 'N/A')} (ID: {response[0].get('coinId', 'N/A')})")
        elif isinstance(response, dict):
            coins = response.get('coins', [])
            print(f"✓ Successfully retrieved coin list, total {len(coins)} coins")
            if coins:
                print(f"  First coin: {coins[0]['symbol']} (ID: {coins[0]['coinId']})")
        else:
            print(f"✗ Unexpected response type: {type(response)}")
    except APIError as e:
        print(f"✗ API Error: code={e.code}, message={str(e)}")
    except Exception as e:
        print(f"✗ Error: {e}")

    # Test 2: Get merchant assets
    print("\n[Test 2] Get Merchant Assets...")
    try:
        response = client.merchant_assets.get_app_coin_asset_list()
        print(f"DEBUG: response type={type(response)}, value={response[:100] if isinstance(response, list) and len(response) > 0 else response}")
        if isinstance(response, list):
            print(f"✓ Successfully retrieved merchant assets, total {len(response)} assets")
            if response:
                for i, asset in enumerate(response[:5]):
                    print(f"  {asset.get('coinSymbol', 'N/A')}: {asset.get('available', 'N/A')}")
                if len(response) > 5:
                    print(f"  ... and {len(response) - 5} more")
        elif isinstance(response, dict):
            assets = response.get('assets', [])
            print(f"✓ Successfully retrieved merchant assets, total {len(assets)} assets")
            if assets:
                for i, asset in enumerate(assets[:5]):
                    print(f"  {asset['coinSymbol']}: {asset['available']}")
                if len(assets) > 5:
                    print(f"  ... and {len(assets) - 5} more")
        else:
            print(f"✗ Unexpected response type: {type(response)}")
    except APIError as e:
        print(f"✗ API Error: code={e.code}, message={str(e)}")
    except Exception as e:
        import traceback
        print(f"✗ Error: {e}")
        traceback.print_exc()

    # Test 3: Health check
    print("\n[Test 3] Health Check...")
    try:
        response = client.utilities.health()
        print(f"✓ Health check passed: {response}")
    except APIError as e:
        print(f"✗ API Error: code={e.code}, message={str(e)}")
    except Exception as e:
        print(f"✗ Error: {e}")

    print("\n" + "=" * 50)
    print("Test completed")


if __name__ == '__main__':
    main()
