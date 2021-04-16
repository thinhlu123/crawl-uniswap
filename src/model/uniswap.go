package model

type (
	UniswapFactory struct {
		PairCount         string `bson:"pair_count,omitempty" mapstructure:"pair_count,omitempty"`
		TotalVolumeUSD    string `bson:"total_volume_usd,omitempty" mapstructure:"total_volume_usd,omitempty"`
		TotalVolumeETH    string `bson:"total_volume_eth,omitempty" mapstructure:"total_volume_eth,omitempty"`
		TotalLiquidityUSD string `bson:"total_liquidity_usd,omitempty" mapstructure:"total_liquidity_usd,omitempty"`
		TotalLiquidityETH string `bson:"total_liquidity_eth,omitempty" mapstructure:"total_liquidity_eth,omitempty"`
		TxCount           string `bson:"tx_count,omitempty" mapstructure:"tx_count,omitempty"`
	}
	Token struct {
		Symbol             string `bson:"symbol,omitempty" mapstructure:"symbol,omitempty"`
		Name               string `bson:"name,omitempty" mapstructure:"name,omitempty"`
		Decimals           string `bson:"decimals,omitempty" mapstructure:"decimals,omitempty"`
		TradeVolume        string `bson:"trade_volume,omitempty" mapstructure:"trade_volume,omitempty"`
		TradeVolumeUSD     string `bson:"trade_volume_usd,omitempty" mapstructure:"trade_volume_usd,omitempty"`
		UntrackedVolumeUSD string `bson:"untracked_volume_usd,omitempty" mapstructure:"untracked_volume_usd,omitempty"`
		TxCount            string `bson:"tx_count,omitempty" mapstructure:"tx_count,omitempty"`
		TotalLiquidity     string `bson:"total_liquidity,omitempty" mapstructure:"total_liquidity,omitempty"`
		DerivedETH         string `bson:"derived_eth,omitempty" mapstructure:"derived_eth,omitempty"`
	}
	Pair struct {
		//Factory UniswapFactory
		Token0             Token  `bson:"token_0,omitempty" mapstructure:"token0,omitempty"`
		Token1             Token  `bson:"token_1,omitempty" mapstructure:"token1,omitempty"`
		Reserve0           string `bson:"reserve_0,omitempty" mapstructure:"reserve0,omitempty"`
		Reserve1           string `bson:"reserve_1,omitempty" mapstructure:"reserve1,omitempty"`
		TotalSupply        string `bson:"total_supply,omitempty" mapstructure:"totalSupply,omitempty"`
		ReserveETH         string `bson:"reserve_eth,omitempty" mapstructure:"reserveEth,omitempty"`
		ReserveUSD         string `bson:"reserve_usd,omitempty" mapstructure:"reserveUsd,omitempty"`
		TrackedReserveETH  string `bson:"tracked_reserve_eth,omitempty" mapstructure:"trackedReserveEth,omitempty"`
		Token0Price        string `bson:"token_0_price,omitempty" mapstructure:"token0Price,omitempty"`
		Token1Price        string `bson:"token_1_price,omitempty" mapstructure:"token1Price,omitempty"`
		VolumeToken0       string `bson:"volume_token_0,omitempty" mapstructure:"volumeToken0,omitempty"`
		VolumeToken1       string `bson:"volume_token_1,omitempty" mapstructure:"volumeToken1,omitempty"`
		VolumeUSD          string `bson:"volume_usd,omitempty" mapstructure:"volumeUsd,omitempty"`
		UntrackedVolumeUSD string `bson:"untracked_volume_usd,omitempty" mapstructure:"untrackedVolumeUsd,omitempty"`
		TxCount            string `bson:"tx_count,omitempty" mapstructure:"txCount,omitempty"`
	}
)
