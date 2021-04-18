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
	TotalLiquidity string
	Token0USD      string
	Token1USD      string
	Transaction    []TokenTransaction
	Pair           Pair
}
