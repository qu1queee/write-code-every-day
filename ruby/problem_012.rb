#! /usr/bin/env ruby

# Write an efficient method that takes stockPrices and 
# returns the best profit I could have made from one purchase 
# and one sale of one share of Apple stock yesterday.
#
# For example:
#
# stockPrices = [10, 7, 5, 8, 11, 9]
#
# getMaxProfit(stockPrices)
# returns 6 (buying for $5 and selling for $11)
#
# Based on the interviewcake exercises

def getMaxProfit(stockPrices)

  if stockPrices.length < 2
    raise ArgumentError, "getting a profit requires at least 2 prices"    
  end

  buyPrice = stockPrices[0]
  maxProfit = stockPrices[1] - stockPrices[0] 
  
  stockPrices[1..-1].each do |price|
    potentialProfit = price - buyPrice
    maxProfit = [potentialProfit, maxProfit].max
    buyPrice = [buyPrice, price].min
  end
  
  return maxProfit
end

#
# Tests
#
def runTests
  desc = 'price goes up then down'
  actual = getMaxProfit([1, 5, 3, 2])
  expected = 4
  assertEqual(actual, expected, desc)

  desc = 'price goes down then up'
  actual = getMaxProfit([7, 2, 8, 9])
  expected = 7
  assertEqual(actual, expected, desc)

  desc = 'price goes up all day'
  actual = getMaxProfit([1, 6, 7, 9])
  expected = 8
  assertEqual(actual, expected, desc)

  desc = 'price goes down all day'
  actual = getMaxProfit([9, 7, 4, 1])
  expected = -2
  assertEqual(actual, expected, desc)

  desc = 'price stays the same all day'
  actual = getMaxProfit([1, 1, 1, 1])
  expected = 0
  assertEqual(actual, expected, desc)

  desc = 'error with empty prices'
  assertRaises(desc) {
    getMaxProfit([])
  }

  desc = 'error with one price'
  assertRaises(desc) {
    getMaxProfit([1])
  }
end

def assertEqual(a, b, desc)
  puts "#{desc} ... #{a == b ? 'PASS' : "FAIL: #{a.inspect} != #{b.inspect}"}"
end

def assertRaises(desc)
  yield
  puts "#{desc} ... FAIL"
rescue
  puts "#{desc} ... PASS"
end

runTests