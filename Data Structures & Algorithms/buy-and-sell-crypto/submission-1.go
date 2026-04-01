func maxProfit(prices []int) int {
	profit := 0
	buyPrice := 101
	for i := range prices {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		}

		if profit < (prices[i] - buyPrice) {
			profit = prices[i] - buyPrice
		} 
	}
	return profit	
}
