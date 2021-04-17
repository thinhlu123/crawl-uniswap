package model

type PairInfo struct {
	PairId         string
	Name           string
	TotalLiquidity float64
}
type PageVars struct {
	Title string
	// for home page
	ListPairInfo []PairInfo

	// for detail page
	TotalLiquidity float64
	Token0USD      float64
	Token1USD      float64
	Transaction    []TokenTransaction
	Pair           Pair
}
