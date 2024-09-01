# Sort Vowels in a String (medium)

## Problem Statement

Given a string s, return an updated string t such that all consonants in the
string s stay in their original positions while any vowels in the string are
reordered according to their ASCII values.

The vowels are 'A', 'E', 'I', 'O', and 'U'. These vowels can appear in lowercase
or uppercase. All other letters except vowels are consonants.

## Example

```text
Example 1:
Input: "gamE"
Expected Output: "gEma"
Justification: The vowels in "gamE" are 'a' and 'E'. Sorting these by ASCII values,
'E' comes before 'a'. Hence, they are placed in the order 'E', 'a' in the output,
while consonants remain unchanged.

Example 2:
Input: "aEiOu"
Expected Output: "EOaiu"
Justification: The input consists only of vowels. When we sort them according
to their ASCII values, the uppercase vowels come first and then the lowercase
vowel comes in 'aiou' order.
```

## Constraints

```text
1 <= s.length <= 10^5
s consists only of letters of the English alphabet in uppercase and lowercase.
```
