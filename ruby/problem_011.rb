#! /usr/bin/env ruby

# Write an efficient method that checks whether any permutation
# of an input string is a palindrome.
# You can assume the input string only contains lowercase letters.
#
#Examples:
#
# "civic" should return true
# "ivicc" should return true
# "civil" should return false
# "livci" should return false
#
# Based on the interviewcake exercises

require 'set'

def hasPalindromePermutation(msg)
  return true if msg.length < 1
  hashAux = Set.new
  
  (0..msg.length-1).each do |index|
    if hashAux.include?(msg[index])
      hashAux.delete(msg[index])
    else
      hashAux.add(msg[index])
    end
  end
  return hashAux.length <= 1
end

#
# Tests
#
def runTests
  desc = 'permutation with odd number of chars'
  result = hasPalindromePermutation('aabcbcd')
  assertTrue(result, desc)

  desc = 'permutation with even number of chars'
  result = hasPalindromePermutation('aabccbdd')
  assertTrue(result, desc)

  desc = 'no permutation with odd number of chars'
  result = hasPalindromePermutation('aabcd')
  assertFalse(result, desc)

  desc = 'no permutation with even number of chars'
  result = hasPalindromePermutation('aabbcd')
  assertFalse(result, desc)

  desc = 'empty string'
  result = hasPalindromePermutation('')
  assertTrue(result, desc)

  desc = 'one character string'
  result = hasPalindromePermutation('a')
  assertTrue(result, desc)
end

def assertTrue(value, desc)
  puts "#{desc} ... #{value ? 'PASS' : "FAIL: #{value} is not true"}"
end

def assertFalse(value, desc)
  puts "#{desc} ... #{value ? "FAIL: #{value} is not false" : 'PASS'}"
end

runTests