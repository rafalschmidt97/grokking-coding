# Problem 5: Zigzag Iterator (medium)

## Problem Statement

Given two 1d vectors, implement an iterator to return their elements
alternately.

For example, given two 1d vectors:

```text
v1 = [1, 2]
v2 = [3, 4, 5, 6]
```

By calling next() repeatedly until hasNext() returns false, the order of
elements returned by next should be: [1, 3, 2, 4, 5, 6].

## Example

```text
Example 1:

i = ZigzagIterator([1, 2], [3, 4, 5, 6])
print(i.next())  # returns 1
print(i.next())  # returns 3
print(i.next())  # returns 2
print(i.next())  # returns 4
print(i.next())  # returns 5
print(i.next())  # returns 6
print(i.hasNext())  # returns False
Example 2:

i = ZigzagIterator([1, 2, 3, 4], [5, 6])
print(i.next())  # returns 1
print(i.next())  # returns 5
print(i.next())  # returns 2
print(i.next())  # returns 6
print(i.next())  # returns 3
print(i.next())  # returns 4
print(i.hasNext())  # returns False
Example 3:

i = ZigzagIterator([1, 2], [])
print(i.next())  # returns 1
print(i.next())  # returns 2
print(i.hasNext())  # returns False
```

## Constraints

```text
0 <= v1.length, v2.length <= 1000
1 <= v1.length + v2.length <= 2000
-2^31 <= v1[i], v2[i] <= 2^31 - 1
```
