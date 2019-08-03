#! /usr/bin/env ruby

# Topic: all about functional programming
# cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns 
# the first and last element of that pair.
# For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.
# 
# Given this implementation of cons:
# 
# def cons(a, b):
#     def pair(f):
#         return f(a, b)
#     return pair
#
# Implement car and cdr.

# @param {int, int} a,b
# @return {lambda}
def cons(a,b)
    pair = -> (method) {method.call(a,b)}
    return pair
end

# @param {lambda} method
# @return {int}
def car(method)
    left = -> (a,b) {return a}
    return method.call(left)
end 

# @param {lambda} method
# @return {int}
def cdr(method)
    right = -> (a,b) {return b}
    return method.call(right)
end 

raise "Test failed" unless car(cons(3,4)) == 3
raise "Test failed" unless cdr(cons(3,4)) == 4
raise "Test failed" unless car(cons(1,1)) == 1
raise "Test failed" unless cdr(cons(1,1)) == 1