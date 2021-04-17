package model

type PairInfo struct {
	Name           string
	TotalLiquidity int
}
type PageVars struct {
	Title string
	// for home page
	ListPairInfo []PairInfo

	// for detail page
	TotalLiquidity int
	Transaction    []TokenTransaction
	ETHPrice       int
	Pair
}
