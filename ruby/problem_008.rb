#!/usr/bin/env ruby

# Your team is scrambling to decipher a recent message, 
# worried it's a plot to break into a major European National Cake Vault. 
# The message has been mostly deciphered, but all the words are backward! 
# Your colleagues have handed off the last step to you.

# Write a method reverseWords() that takes a message as a string 
# and reverses the order of the words in place.
#
# Based on the interviewcake exercises


def reverseSomething(msg, leftIndex, rightIndex)
  while leftIndex < rightIndex
    msg[leftIndex], msg[rightIndex] = msg[rightIndex], msg[leftIndex]
    leftIndex +=1
    rightIndex -=1
  end
  return msg
end

def reverseWords(message)
    message = reverseSomething(message, 0, message.length-1)
    
    counter = 0
    mostLeftIndex = 0

    message.each_char { |letter|
      if letter == " " && (message.length-1 != counter)
        reverseSomething(message, mostLeftIndex, counter-1)
        mostLeftIndex = counter + 1
      end
      if counter == message.length-1
        reverseSomething(message, mostLeftIndex, counter)
      end
      counter += 1
    }

    return message
end

#
# Tests
#
def runTests
  desc = 'one word'
  message = 'vault'
  reverseWords(message)
  expected = 'vault'
  assertEqual(message, expected, desc)

  desc = 'two words'
  message = 'thief cake'
  reverseWords(message)
  expected = 'cake thief'
  assertEqual(message, expected, desc)

  desc = 'three words'
  message = 'one another get'
  reverseWords(message)
  expected = 'get another one'
  assertEqual(message, expected, desc)

  desc = 'multiple words same length'
  message = 'rat the ate cat the'
  reverseWords(message)
  expected = 'the cat ate the rat'
  assertEqual(message, expected, desc)

  desc = 'multiple words different lengths'
  message = 'yummy is cake bundt chocolate'
  reverseWords(message)
  expected = 'chocolate bundt cake is yummy'
  assertEqual(message, expected, desc)

  desc = 'empty string'
  message = ''
  reverseWords(message)
  expected = ''
  assertEqual(message, expected, desc)
end

def assertEqual(a, b, desc)
  puts "#{desc} ... #{a == b ? 'PASS' : "FAIL: #{a.inspect} != #{b.inspect}"}"
end

runTests