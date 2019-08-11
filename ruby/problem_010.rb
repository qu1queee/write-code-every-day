#! /usr/bin/env ruby

require 'set'
#
#
# Users on longer flights like to start a second movie right when their first one ends,
# but they complain that the plane usually lands before they can see the ending. 
# So you're building a feature for choosing two movies whose total runtimes will equal the exact flight length.
# Write a method that takes an integer flightLength (in minutes) and an array of integers movieLengths (in minutes) 
# and returns a boolean indicating whether there are two numbers in movieLengths whose sum equals flightLength.
#
# When building your method:
#   Assume your users will watch exactly two movies
#   Dont make your users watch the same movie twice
#   Optimize for runtime over memory
#
# Based on the interviewcake exercises

def twoMoviesInFlight(movieLengths, flightLength)

  # Determine if two movie runtimes add up to the flight length
  someList = Set.new
  movieLengths[0..-1].each do |durationofMovie|
    return true if someList.include?(flightLength-durationofMovie)
    someList.add(durationofMovie)
  end
  return false
end

#
# Tests
#
def runTests
  desc = 'short flight'
  result = twoMoviesInFlight([2, 4], 1)
  assertFalse(result, desc)

  desc = 'long flight'
  result = twoMoviesInFlight([2, 4], 6)
  assertTrue(result, desc)

  desc = 'one movie half flight length'
  result = twoMoviesInFlight([3, 8], 6)
  assertFalse(result, desc)

  desc = 'two movies half flight length'
  result = twoMoviesInFlight([3, 8, 3], 6)
  assertTrue(result, desc)

  desc = 'lots of possible pairs'
  result = twoMoviesInFlight([1, 2, 3, 4, 5, 6], 7)
  assertTrue(result, desc)

  desc = 'not using first movie'
  result = twoMoviesInFlight([4, 3, 2], 5)
  assertTrue(result, desc)

  desc = 'only one movie'
  result = twoMoviesInFlight([6], 6)
  assertFalse(result, desc)

  desc = 'no movies'
  result = twoMoviesInFlight([], 2)
  assertFalse(result, desc)
end

def assertTrue(value, desc)
  puts "#{desc} ... #{value ? 'PASS' : "FAIL: #{value} is not true"}"
end

def assertFalse(value, desc)
  puts "#{desc} ... #{value ? "FAIL: #{value} is not false" : 'PASS'}"
end

runTests