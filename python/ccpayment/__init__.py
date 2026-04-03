"""
CCPayment Python SDK
Official Python SDK for CCPayment API v2.
"""

from .client import Client
from .models import *
from .exceptions import APIError

__version__ = "1.0.0"
__all__ = ["Client", "APIError"]
