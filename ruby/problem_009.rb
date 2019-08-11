#! /usr/bin/env ruby

#
# In order to win the prize for most cookies sold, 
# my friend Alice and I are going to merge our Girl Scout Cookies orders and enter as one unit.
#
# Each order is represented by an "order id" (an integer). We have our lists of orders sorted numerically already, 
# in arrays. Write a method to merge our arrays of orders into one sorted array.
#
# For example:
#
# myArray     = [3, 4, 6, 10, 11, 15]
# alicesArray = [1, 5, 8, 12, 14, 19]
# puts mergeArrays(myArray, alicesArray)
# prints [1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19]
#
# Based on the interviewcake exercises

def mergeArrays(myArray, alicesArray)
  
  return myArray if alicesArray.length < 1
  return alicesArray if myArray.length < 1

  mergedArray = Array.new(myArray.length+alicesArray.length)  

  currentIndex = 0
  myArrayIndex = 0
  aliceArrayIndex = 0

  while currentIndex <  mergedArray.length
  
    exhaustedMyArray = myArrayIndex >= myArray.length
    exhaustedAlice = aliceArrayIndex >= alicesArray.length
    
    if !exhaustedMyArray && (exhaustedAlice || (myArray[myArrayIndex] < alicesArray[aliceArrayIndex]))
        mergedArray[currentIndex] = myArray[myArrayIndex]
        myArrayIndex += 1
    else
        mergedArray[currentIndex] = alicesArray[aliceArrayIndex]
        aliceArrayIndex += 1
    end
    
    currentIndex += 1
  end

  return  mergedArray
end

#
# Tests
#
def runTests
  desc = 'both arrays are empty'
  actual = mergeArrays([], [])
  expected = []
  assertEqual(actual, expected, desc)

  desc = 'first array is empty'
  actual = mergeArrays([], [1, 2, 3])
  expected = [1, 2, 3]
  assertEqual(actual, expected, desc)

  desc = 'second array is empty'
  actual = mergeArrays([5, 6, 7], [])
  expected = [5, 6, 7]
  assertEqual(actual, expected, desc)

  desc = 'both arrays have some numbers'
  actual = mergeArrays([2, 4, 6], [1, 3, 7])
  expected = [1, 2, 3, 4, 6, 7]
  assertEqual(actual, expected, desc)

  desc = 'arrays are different lengths'
  actual = mergeArrays([2, 4, 6, 8], [1, 7])
  expected = [1, 2, 4, 6, 7, 8]
  assertEqual(actual, expected, desc)
end

def assertEqual(a, b, desc)
  puts "#{desc} ... #{a == b ? 'PASS' : "FAIL: #{a.inspect} != #{b.inspect}"}"
end

runTests

