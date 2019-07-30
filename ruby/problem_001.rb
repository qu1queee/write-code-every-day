#!/usr/bin/env ruby

# Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
# For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
#
# Bonus: Can you do this in one pass?

# @param {int, list} k, list 
# @result {boolean}
def addUpToArg(k, list)
    unsortedList = []

    list.each do |num|
        unsortedList.push(k-num)
        if unsortedList.include?(num)
            return true
        end
    end

    return false

end


raise "Test Fail" unless addUpToArg(17, [10, 15, 3, 7])
raise "Test Fail" unless !addUpToArg(17, [10, 15, 3, 6])