"""CCPayment SDK Models"""

from dataclasses import dataclass, field
from typing import Optional, List, Dict, Any


# ==================== Basic Info Models ====================

@dataclass
class NetworkInfo:
    chain: str
    chain_full_name: str
    contract: str
    precision: int
    can_deposit: bool
    can_withdraw: bool
    minimum_deposit_amount: str
    minimum_withdraw_amount: str
    maximum_withdraw_amount: str
    is_support_memo: bool
    protocol: str


@dataclass
class Coin:
    coin_id: int
    symbol: str
    coin_full_name: str
    logo_url: str
    status: str
    networks: Dict[str, NetworkInfo]
    price: str


@dataclass
class Fiat:
    fiat_id: int
    symbol: str
    logo_url: str
    mark: str
    usd_rate: str


@dataclass
class Chain:
    chain: str
    chain_full_name: str
    explorer: str
    logo_url: str
    status: str
    confirmations: int
    base_url: str
    is_evm: bool
    support_memo: bool


@dataclass
class ChainSimple:
    chain: str
    chain_full_name: str
    explorer: str
    logo_url: str
    status: str
    confirm_num: int
    is_evm: bool


@dataclass
class Fee:
    coin_id: int
    coin_symbol: str
    amount: str


# ==================== Merchant Assets Models ====================

@dataclass
class Asset:
    coin_id: int
    coin_symbol: str
    available: str


# ==================== Merchant Deposit Models ====================

@dataclass
class DepositRecord:
    record_id: str
    reference_id: str
    order_id: str
    coin_id: int
    coin_symbol: str
    chain: str
    contract: str
    coin_usd_price: str
    from_address: str
    to_address: str
    to_memo: Optional[str]
    amount: str
    service_fee: str
    tx_id: str
    tx_index: int
    status: str
    arrived_at: int
    is_flagged_as_risky: Optional[bool]


# ==================== Merchant Withdraw Models ====================

@dataclass
class WithdrawFee:
    coin_id: int
    coin_symbol: str
    amount: str


@dataclass
class WithdrawRecord:
    record_id: str
    withdraw_type: str
    app_id: str
    coin_id: int
    coin_symbol: str
    chain: str
    from_address: str
    to_address: str
    cwallet_user: str
    order_id: str
    tx_id: Optional[str]
    to_memo: Optional[str]
    status: str
    amount: str
    fee: WithdrawFee
    reason: Optional[str]
    coin_usd_price: str


@dataclass
class AutoWithdrawRecord:
    record_id: str
    coin_id: int
    coin_symbol: str
    chain: str
    order_id: str
    to_address: str
    to_memo: Optional[str]
    amount: str
    tx_id: Optional[str]
    status: str
    fee: WithdrawFee
    service_fee: str
    created_at: int


@dataclass
class RiskyRefundRecord:
    record_id: str
    coin_id: int
    coin_symbol: str
    chain: str
    order_id: str
    to_address: str
    to_memo: Optional[str]
    amount: str
    tx_id: Optional[str]
    status: str
    fee: WithdrawFee
    created_at: int
    from_deposit: List[str]


# ==================== Merchant Batch Withdraw Models ====================

@dataclass
class AddressInfo:
    address: str
    seq: int
    memo: Optional[str] = None
    codes: Optional[List[int]] = None


@dataclass
class AddressInfoResult:
    seq: int
    codes: List[int]


@dataclass
class BatchWithdrawOrder:
    seq: int
    address: str
    amount: str
    memo: Optional[str] = None
    remark: Optional[str] = None
    record_id: Optional[str] = None
    order_id: Optional[str] = None
    status: Optional[str] = None
    network_fee: Optional[str] = None
    tx_id: Optional[str] = None
    reason: Optional[str] = None
    created_at: Optional[int] = None
    updated_at: Optional[int] = None


@dataclass
class BatchWithdrawStats:
    total: int
    succeeded: int
    failed: int
    canceled: int
    processing: int
    exec_seq: int


@dataclass
class BatchWithdrawCoin:
    coin_id: int
    coin_symbol: str
    coin_price: str


# ==================== User Deposit Models ====================

@dataclass
class UserDepositRecord:
    user_id: str
    record_id: str
    coin_id: int
    chain: str
    contract: str
    coin_symbol: str
    tx_id: str
    coin_usd_price: str
    from_address: str
    to_address: str
    to_memo: Optional[str]
    amount: str
    service_fee: str
    status: str
    arrived_at: int
    is_flagged_as_risky: Optional[bool]


# ==================== User Withdraw Models ====================

@dataclass
class UserWithdrawRecord:
    user_id: str
    record_id: str
    withdraw_type: str
    coin_id: int
    chain: str
    order_id: str
    tx_id: Optional[str]
    coin_symbol: str
    from_address: str
    to_address: str
    cwallet_user: str
    to_memo: Optional[str]
    amount: str
    status: str
    fee: WithdrawFee
    coin_usd_price: str


# ==================== User Transfer Models ====================

@dataclass
class ToUser:
    user_id: str
    amount: str


@dataclass
class TransferRecord:
    record_id: str
    coin_id: int
    coin_symbol: str
    order_id: str
    from_user_id: str
    to_user_id: str
    amount: str
    status: str
    remark: Optional[str]
    coin_usd_price: str


@dataclass
class BatchTransferRecord:
    record_id: str
    user_id: str
    coin_id: int
    coin_symbol: str
    order_id: str
    to_users: List[ToUser]
    status: str
    remark: Optional[str]
    coin_usd_price: str


# ==================== Orders Models ====================

@dataclass
class PaidInfo:
    record_id: str
    from_address: str
    amount: str
    service_fee: str
    txid: str
    status: str
    arrived_at: int
    rate: str
    is_flagged_as_risky: bool


@dataclass
class InvoicePaidInfo:
    record_id: str
    coin_id: int
    coin_symbol: str
    chain: str
    from_address: str
    to_address: str
    to_memo: str
    paid_amount: str
    service_fee: str
    txid: str
    status: str
    arrived_at: int
    rate: str
    paid_value: str
    is_flagged_as_risky: bool


# ==================== Checkout Models ====================

@dataclass
class CheckoutCoin:
    coin_id: int
    symbol: str
    logo_url: str
    chains: List[str]


# ==================== Swap Models ====================

@dataclass
class SwapCoin:
    coin_id: int
    symbol: str
    logo_url: str
    chains: List[str]


@dataclass
class SwapEstimate:
    from_coin_id: int
    from_coin_symbol: str
    from_amount: str
    to_coin_id: int
    to_coin_symbol: str
    to_amount: str
    rate: str


@dataclass
class SwapRecord:
    record_id: str
    order_id: str
    from_coin_id: int
    from_coin_symbol: str
    from_amount: str
    to_coin_id: int
    to_coin_symbol: str
    to_amount: str
    rate: str
    status: str
    created_at: int
