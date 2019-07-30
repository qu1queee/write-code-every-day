#!/usr/bin/env ruby

# Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
# For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].


# Following two functions rely on division
#
# @param {list} list
# @result {int}
def productOfNumbers(list)
    result = 1
    list.each do |i|
        result *= i
    end
    return result
end

# @param {list} list
# @result {list}
def newArray(list)
    newList = []
    total = productOfNumbers(list)
    list.each do |n|
        newList.push(total/n)
    end
    return newList
end

raise "Test Fail" unless newArray([1, 2, 3, 4, 5]) == [120, 60, 40, 30, 24]
raise "Test Fail" unless newArray([3,2,1]) == [2,3,6]
raise "Test Fail" unless newArray([1,2,3,4]) == [24,12,8,6]

# Follow-up: what if you can't use division?
#
# New factor array, see above description
# It does not use division, keeping a small
# time complexity
#
# @param {list} list
# @return {list}
def generateFactorList(list)
    lefList = []
    rightList = []
    finalList = []

    lefList[0] = 1
    for i in 1..list.length-1
        lefList.push(lefList[i-1]*list[i-1])
    end

    rightList[0] = 1
    for i in 1..list.length-1
        rightList.push(rightList[i-1]*list.reverse[i-1])
    end

    list.each_with_index do |n,i|
        finalList.push(lefList[i]*rightList.reverse[i])
    end
    return finalList
end


raise "Test Fail" unless generateFactorList([1, 2, 3, 4, 5]) == [120, 60, 40, 30, 24]
raise "Test Fail" unless generateFactorList([3,2,1]) == [2,3,6]
raise "Test Fail" unless generateFactorList([1,2,3,4]) == [24,12,8,6]