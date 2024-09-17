# Meeting Rooms II (medium)

## Problem Statement

Given a list of time intervals during which meetings are scheduled, determine
the minimum number of meeting rooms that are required to ensure that none of the
meetings overlap in time.

## Example

```text
Example 1:
Input: [[10, 15], [20, 25], [30, 35]]
Expected Output: 1
Justification: There are no overlapping intervals in the given list.
So, only 1 meeting room is enough for all the meetings.

Example 2:
Input: [[10, 20], [15, 25], [24, 30]]
Expected Output: 2
Justification: The first and second intervals overlap, and the second
and third intervals overlap as well. So, we need 2 rooms.

Example 3:
Input: [[10, 20], [20, 30]]
Expected Output: 1
Justification: The end time of the first meeting is the same as the
start time of the second meeting. So, one meeting can be scheduled
right after the other in the same room.
```

## Constraints

```text
1 <= intervals.length <= 10^4
0 <= starti < endi <= 10^6
```
