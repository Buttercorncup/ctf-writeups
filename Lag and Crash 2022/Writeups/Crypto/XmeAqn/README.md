# XmeAqn

---

## Challenge Description 
Write a function that takes two equal-length buffers and produces their XOR combination.

value1 = "73757065726d616e20686174652062617473"

value2 = "6261746d616e206c6f766573207375706572"

Use produced combination to fill the blanks LNC2022{C_ _ _r_ _ _} (no spaces in-between)

1st character use 24th character of the produced combination - 2nd character use 18th character of the produced combination - 3rd character use 20th character of the produced combination - 4th character use 4th character of the produced combination - 5th character use 5th character of the produced combination - 6th character use 6th character of the produced combination

---

## Challenge Solution
Since the challenge description has hinted us to write a XOR function, we can simply search up a [XOR Calculator](https://xor.pw/) to save time. Taking the two values, we XOR them together to get the following output.

![Screenshot 2022-03-28 220847](https://user-images.githubusercontent.com/101789488/160416411-1a48184c-3b84-408c-abd2-0fbaaadcfec5.png)
```
11140408130341024f1e0407455317111101
```

Now, for the second part, the 6 characters that the description refers to are the 6 blanks in the incomplete flag. Basically all we have to do is fill it in using the XOR-ed result we got and we have the flag.
```
LNC2022{C7fer404}
```