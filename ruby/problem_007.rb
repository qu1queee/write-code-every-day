#! /usr/bin/env ruby

# Write a method that takes a string and reverses the letters in place. ↴ 
#
# Based on the interviewcake exercises

def reverse(str)

  leftIndex = 0
  rightIndex = str.length-1
  
  while leftIndex < rightIndex
    str[leftIndex], str[rightIndex] = str[rightIndex], str[leftIndex]
    leftIndex += 1
    rightIndex -= 1
  end
  
  return str

end

#
# Tests
#
def runTests
  desc = 'empty string'
  string = ''
  reverse(string)
  expected = ''
  assertEqual(string, expected, desc)

  desc = 'single character string'
  string = 'A'
  reverse(string)
  expected = 'A'
  assertEqual(string, expected, desc)

  desc = 'longer string'
  string = 'ABCDE'
  reverse(string)
  expected = 'EDCBA'
  assertEqual(string, expected, desc)
end

def assertEqual(a, b, desc)
  puts "#{desc} ... #{a == b ? 'PASS' : "FAIL: #{a.inspect} != #{b.inspect}"}"
end

runTests