#!/usr/bin/env ruby

# A meeting is stored as an array of integers [start_time, end_time]. 
# These integers represent the number of 30-minute blocks past 9:00am. 
# Write a method mergeRanges() that takes an array of multiple meeting time 
# ranges and returns an array of condensed ranges. 
# For example, given: 
# [[0, 1], [3, 5], [4, 8], [10, 12], [9, 10]]
#
# your method would return: 
# [[0, 1], [3, 8], [9, 12]]
#
# Based on the interviewcake exercises

def mergeRanges(meetings)
  sortedMeetings = meetings.sort
  mergedMeetings = [sortedMeetings[0]]

  sortedMeetings[0..-1].each do |startTime, endTime|
    startTimeMerged, endTimeMerged = mergedMeetings[-1]
    if startTime <= endTimeMerged
      endTimeMerged = [endTimeMerged,endTime].max
      mergedMeetings[-1] = [startTimeMerged, endTimeMerged]
    else
      mergedMeetings.push([startTime,endTime])
    end
  end

  return mergedMeetings
end

meetings = [[0, 1], [3, 5], [4, 8], [10, 12], [9, 10]]
raise "Test failed" unless mergeRanges(meetings) == [[0, 1], [3, 8], [9, 12]]