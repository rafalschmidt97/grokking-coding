# Problem 6: Simplify Path (medium)

## Problem Statement

Given an absolute file path in a Unix-style file system, simplify it by
converting ".." to the previous directory and removing any "." or multiple
slashes. The resulting string should represent the shortest absolute path.

## Example

```text
1.
   Input: "/a//b////c/d//././/.."
   Output: "/a/b/c"

2.
   Input: "/../"
   Output: "/"

3.
   Input: "/home//foo/"
   Output: "/home/foo"
```

## Constraints

```text
1 <= path.length <= 3000
path consists of English letters, digits, period '.', slash '/' or '_'.
path is a valid absolute Unix path.
```
