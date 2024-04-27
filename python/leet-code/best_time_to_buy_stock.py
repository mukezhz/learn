'''
Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),

Note: You cannot sell a stock before you buy one.

Example:
Input: [7, 1, 5, 3, 6, 4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6 - 1 = 5
Not 7 - 1 = 6, as selling price needs to be larger than buying price.
'''


input = [7, 1, 5, 3, 6, 4]
buy_pointer = 0
sell_pointer = 1
max_profit = 0

while sell_pointer < len(input):
    if input[buy_pointer] < input[sell_pointer]:
        max_profit = max(max_profit, input[sell_pointer] - input[buy_pointer])
    else:
        buy_pointer = sell_pointer
    sell_pointer += 1
    
print(max_profit)