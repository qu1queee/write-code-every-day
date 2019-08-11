#! /usr/bin/env ruby
#
# Given an array of integers, find the highest product you can get 
# from three of the integers. 
#
# Based on the interviewcake exercises

def highestProductOfThree(arrayOfInts)
  if arrayOfInts.length < 3
    raise ArgumentError, "highest product of 3, requires 3 integers"    
  end  

  highest = [arrayOfInts[0],arrayOfInts[1]].max
  lowest = [arrayOfInts[0],arrayOfInts[1]].min
  highestOfTwo = arrayOfInts[0] * arrayOfInts[1]
  lowestOfTwo = arrayOfInts[0] * arrayOfInts[1]
  highestOfThree = arrayOfInts[0] * arrayOfInts[1] * arrayOfInts[2]

  arrayOfInts[2..-1].each_with_index do |value, index|
    highestOfThree = [
      highestOfThree, 
      value * highestOfTwo,
      value * lowestOfTwo,
    ].max
    
    highestOfTwo = [
      highestOfTwo,
      value * lowest,
      value * highest,
    ].max

    lowestOfTwo = [
      lowestOfTwo,
      value * lowest,
      value * highest,
    ].min
    
    highest = [highest,value].max
    
    lowest = [lowest,value].min
    
  end
  return highestOfThree
end

#
# Tests
#
def run_tests
  actual = highestProductOfThree([1, 2, 3, 4])
  expected = 24
  assertEqual(actual, expected, 'short array')

  actual = highestProductOfThree([6, 1, 3, 5, 7, 8, 2])
  expected = 336
  assertEqual(actual, expected, 'longer array')

  actual = highestProductOfThree([-5, 4, 8, 2, 3])
  expected = 96
  assertEqual(actual, expected, 'array has one negative')

  actual = highestProductOfThree([-10, 1, 3, 2, -10])
  expected = 300
  assertEqual(actual, expected, 'array has two negatives')

  actual = highestProductOfThree([-5, -1, -3, -2])
  expected = -6
  assertEqual(actual, expected, 'array is all negatives')

  assertRaises('empty array raises error') do
    highestProductOfThree([])
  end

  assertRaises('one number raises error') do
    highestProductOfThree([1])
  end

  assertRaises('two numbers raises error') do
    highestProductOfThree([1, 1])
  end
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

run_tests