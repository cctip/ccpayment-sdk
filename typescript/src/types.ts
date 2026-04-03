export interface NetworkInfo {
  chain: string;
  chainFullName: string;
  contract: string;
  precision: number;
  canDeposit: boolean;
  canWithdraw: boolean;
  minimumDepositAmount: string;
  minimumWithdrawAmount: string;
  maximumWithdrawAmount: string;
  isSupportMemo: boolean;
  protocol: string;
}

export interface Coin {
  coinId: number;
  symbol: string;
  coinFullName: string;
  logoUrl: string;
  status: string;
  networks: Record<string, NetworkInfo>;
  price: string;
}

export interface Fiat {
  fiatId: number;
  symbol: string;
  logoUrl: string;
  mark: string;
  usdRate: string;
}

export interface Chain {
  chain: string;
  chainFullName: string;
  explorer: string;
  logoUrl: string;
  status: string;
  confirmations: number;
  baseUrl: string;
  isEVM: boolean;
  supportMemo: boolean;
}

export interface ChainSimple {
  chain: string;
  chainFullName: string;
  explorer: string;
  logoUrl: string;
  status: string;
  confirmNum: number;
  isEVM: boolean;
}

export interface Fee {
  coinId: number;
  coinSymbol: string;
  amount: string;
}

export interface Asset {
  coinId: number;
  coinSymbol: string;
  available: string;
}

export interface DepositRecord {
  recordId: string;
  referenceId: string;
  orderId: string;
  coinId: number;
  coinSymbol: string;
  chain: string;
  contract: string;
  coinUSDPrice: string;
  fromAddress: string;
  toAddress: string;
  toMemo?: string;
  amount: string;
  serviceFee: string;
  txId: string;
  txIndex: number;
  status: string;
  arrivedAt: number;
  isFlaggedAsRisky?: boolean;
}

export interface WithdrawRecord {
  recordId: string;
  withdrawType: string;
  appId: string;
  coinId: number;
  coinSymbol: string;
  chain: string;
  fromAddress: string;
  toAddress: string;
  cwalletUser: string;
  orderId: string;
  txId?: string;
  toMemo?: string;
  status: string;
  amount: string;
  fee: Fee;
  reason?: string;
  coinUSDPrice: string;
}

export interface AddressInfo {
  address: string;
  seq: number;
  memo?: string;
  codes?: number[];
}

export interface BatchWithdrawOrder {
  seq: number;
  address: string;
  amount: string;
  memo?: string;
  remark?: string;
  recordId?: string;
  orderId?: string;
  status?: string;
  networkFee?: string;
  txId?: string;
  reason?: string;
  createdAt?: number;
  updatedAt?: number;
}

export interface TransferRecord {
  recordId: string;
  coinId: number;
  coinSymbol: string;
  orderId: string;
  fromUserId: string;
  toUserId: string;
  amount: string;
  status: string;
  remark?: string;
  coinUSDPrice: string;
}

export interface PaidInfo {
  recordId: string;
  fromAddress: string;
  amount: string;
  serviceFee: string;
  txid: string;
  status: string;
  arrivedAt: number;
  rate: string;
  isFlaggedAsRisky: boolean;
}

export interface SwapEstimate {
  fromCoinId: number;
  fromCoinSymbol: string;
  fromAmount: string;
  toCoinId: number;
  toCoinSymbol: string;
  toAmount: string;
  rate: string;
}

export interface SwapRecord {
  recordId: string;
  orderId: string;
  fromCoinId: number;
  fromCoinSymbol: string;
  fromAmount: string;
  toCoinId: number;
  toCoinSymbol: string;
  toAmount: string;
  rate: string;
  status: string;
  createdAt: number;
}

export interface AutoWithdrawRecord {
  recordId: string;
  orderId: string;
  coinId: number;
  coinSymbol: string;
  amount: string;
  toAddress: string;
  chain: string;
  status: string;
  createdAt: number;
  txHash?: string;
}

export interface RiskyRefundRecord {
  recordId: string;
  originalRecordId: string;
  coinId: number;
  coinSymbol: string;
  amount: string;
  fromAddress: string;
  toAddress: string;
  chain: string;
  status: string;
  createdAt: number;
  reason: string;
}

export interface WithdrawFee {
  coinId: number;
  coinSymbol: string;
  chain: string;
  fee: string;
  feeCoin: string;
}

// API Response types
export interface APIResponse<T> {
  code: number;
  msg: string;
  data: T;
}
