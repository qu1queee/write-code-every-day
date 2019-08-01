#!/usr/bin/env ruby

# Given the root to a binary tree, implement serialize(root), 
# which serializes the tree into a string, and deserialize(s), 
# which deserializes the string back into the tree.
#
# For example, given the following Node class
#
# class Node:
#     def __init__(self, val, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
# The following test should pass:
#
# node = Node('root', Node('left', Node('left.left')), Node('right'))
# assert deserialize(serialize(node)).left.left.val == 'left.left'

class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

# Encodes tree to a single string.
#
# @param {TreeNode} root
# @return {string}
def serialize(root, serialized = [])
    if root.nil?
        serialized.push(nil)
    else
        serialized.push(root.val)
        serialized = serialize(root.left,serialized)
        serialized = serialize(root.right,serialized)
    end
    return serialized
end

# Decodes encoded data to tree.
#
# @param {string} data
# @return {TreeNode}
def deserialize(data)
    if data[0].nil?
        return nil
    else
        val = data.shift 
        root = TreeNode.new(val)
        root.left = deserialize(data)
        data.shift
        root.right = deserialize(data)
    end
    return root   
end


data = TreeNode.new(5)
data.left = TreeNode.new(2)
data.left.left = TreeNode.new(1)
data.right = TreeNode.new(7)

raise "Test Fail" unless deserialize(serialize(data)).left.left.val == data.left.left.val
raise "Test Fail" unless deserialize(serialize(data)).right.val == data.right.val