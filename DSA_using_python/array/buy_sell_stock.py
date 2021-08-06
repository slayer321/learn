
# Brute force solution

def maxProfit(self, prices: List[int]) -> int:
    max_profit = 0
    n = len(prices)
    for i in range(n-1):
        for j in range(i+1, n):
            profit = prices[j] - prices[i]
            if profit > max_profit:
                max_profit = profit
    return max_profit


# Optimal solution

def maxProfit(self, prices: List[int]) -> int:
    profit = 0
    min_prc = float("inf")
    n = len(prices)
    for i in range(n):
        min_prc = min(min_prc, prices[i])
        profit = max(profit, prices[i] - min_prc)
    return profit


arr = [100, 180, 260, 310, 40, 535, 695]
print(stockBuySell(arr, len(arr)))
