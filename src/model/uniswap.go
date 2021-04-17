package model

type (
	Token struct {
		TokenId string `bson:"token_id,omitempty" mapstructure:"id,omitempty"`
		Symbol  string `bson:"symbol,omitempty" mapstructure:"symbol,omitempty"`
		Name    string `bson:"name,omitempty" mapstructure:"name,omitempty"`
		//Decimals           string `bson:"decimals,omitempty" mapstructure:"decimals,omitempty"`
		//TradeVolume        string `bson:"trade_volume,omitempty" mapstructure:"trade_volume,omitempty"`
		//TradeVolumeUSD     string `bson:"trade_volume_usd,omitempty" mapstructure:"trade_volume_usd,omitempty"`
		//UntrackedVolumeUSD string `bson:"untracked_volume_usd,omitempty" mapstructure:"untracked_volume_usd,omitempty"`
		//TxCount            string `bson:"tx_count,omitempty" mapstructure:"tx_count,omitempty"`
		TotalLiquidity string `bson:"total_liquidity,omitempty" mapstructure:"totalLiquidity,omitempty"`
		DerivedETH     string `bson:"derived_eth,omitempty" mapstructure:"derivedETH,omitempty"`
	}
	Pair struct {
		PairId string `bson:"pair_id,omitempty" mapstructure:"id,omitempty"`
		Token0 Token  `bson:"token_0,omitempty" mapstructure:"token0,omitempty"`
		Token1 Token  `bson:"token_1,omitempty" mapstructure:"token1,omitempty"`
		//Reserve0           string `bson:"reserve_0,omitempty" mapstructure:"reserve0,omitempty"`
		//Reserve1           string `bson:"reserve_1,omitempty" mapstructure:"reserve1,omitempty"`
		//TotalSupply        string `bson:"total_supply,omitempty" mapstructure:"totalSupply,omitempty"`
		//ReserveETH         string `bson:"reserve_eth,omitempty" mapstructure:"reserveEth,omitempty"`
		//ReserveUSD         string `bson:"reserve_usd,omitempty" mapstructure:"reserveUsd,omitempty"`
		//TrackedReserveETH  string `bson:"tracked_reserve_eth,omitempty" mapstructure:"trackedReserveEth,omitempty"`
		Token0Price string `bson:"token_0_price,omitempty" mapstructure:"token0Price,omitempty"`
		Token1Price string `bson:"token_1_price,omitempty" mapstructure:"token1Price,omitempty"`
		//VolumeToken0       string `bson:"volume_token_0,omitempty" mapstructure:"volumeToken0,omitempty"`
		//VolumeToken1       string `bson:"volume_token_1,omitempty" mapstructure:"volumeToken1,omitempty"`
		//VolumeUSD          string `bson:"volume_usd,omitempty" mapstructure:"volumeUsd,omitempty"`
		//UntrackedVolumeUSD string `bson:"untracked_volume_usd,omitempty" mapstructure:"untrackedVolumeUsd,omitempty"`
		//TxCount            string `bson:"tx_count,omitempty" mapstructure:"txCount,omitempty"`
	}
	TokenTransaction struct {
		Type      string `bson:"type,omitempty"`
		Timestamp string `bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`
		//SwapId     string `bson:"swap_id,omitempty" mapstructure:"id,omitempty"`
		//Pair       Pair   `bson:"pair,omitempty" mapstructure:"pair,omiempty"`
		PairId     string `bson:"pair_id,omitempty"`
		Sender     string `bson:"sender,omitempty"  mapstructure:"sender,omiempty"`
		Amount0In  string `bson:"amount_0_in,omitempty" mapstructure:"amount0In,omiempty"`
		Amount1In  string `bson:"amount_1_in,omitempty" mapstructure:"amount1In,omiempty"`
		Amount0Out string `bson:"amount_0_out,omitempty" mapstructure:"amount0Out,omiempty"`
		Amount1Out string `bson:"amount_1_out,omitempty" mapstructure:"amount1Out,omiempty"`
		To         string `bson:"to,omitempty" mapstructure:"to,omiempty"`
		AmountUSD  string `bson:"amount_usd,omitempty" mapstructure:"amountUSD,omiempty"`
		// for burn, mint
		Amount0 string `bson:"amount_0,omitempty" mapstructure:"amount0,omiempty"`
		Amount1 string `bson:"amount_1,omitempty" mapstructure:"amount1,omiempty"`
	}
	Bundle struct {
		Type     string `bson:"type"`
		ETHPrice string `bson:"eth_price,omitempty" mapstructure:"ethPrice,omiempty"`
	}
)
