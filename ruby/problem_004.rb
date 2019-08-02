#!/usr/bin/env ruby

# Given an array of integers, find the first missing positive integer in linear time and constant space. 
# In other words, find the lowest positive integer that does not exist in the array. The array can contain 
# duplicates and negative numbers as well.
# 
# For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
# 
# You can modify the input array in-place.

# @param {list} list
# @return {int}
def findLowestPositive(list)
    listSize = list.length
    (0...listSize).each do |i|
        while (list[i] > 0) && (list[i] <= listSize) && (list[i] != list[list[i]-1])
            t = list[list[i]-1]
            list[list[i]-1] = list[i]
            list[i] = t
        end
    end

    list.each.with_index do |e,i|
        if e != i+1
         return i+1
        end
    end

    return listSize+1
end

raise "Test Fail" unless findLowestPositive([3, 4, -1, 1]) == 2
raise "Test Fail" unless findLowestPositive([1, 2, 0]) == 3
raise "Test Fail" unless findLowestPositive([1, 2, 4]) == 3
raise "Test Fail" unless findLowestPositive([1, 2, 3]) == 4