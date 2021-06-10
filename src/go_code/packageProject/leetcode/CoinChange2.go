package leetcode

type CoinChange2 struct {

}

func Change(amount int, coins []int) int {
	 var dp = make([]int, amount+1)
	 //init
	 dp[0] = 1;
	for _, coin := range coins{
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
