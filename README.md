# Task 1. Separate Text
Write a function/method that accepts a list of string tokens and a separate text string as an input and checks if the latter string can be represented as a concatenation of any subset of tokens from the first argument, where each token can be used multiple times. Examples:

```
[ "ab", "bc", "a" ] vs. "abc" = true
[ "a", "ab", "bc" ] vs. "ab" = true
[ "a", "ab", "bc" ] vs. "" = true
[ "ab", "bc" ] <-> "abc" = false
[ "ab", "bc", "c" ] <-> "abbcbc" = true
```
