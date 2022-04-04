# Inverse Null

---

## Challenge Description 
It seems that the text is in a pretty funny way... I wonder what order it is.

NOTE: Submit flag in all caps

---

## Challenge Solution
For this challenge, we are provided with a text file inverse_null.txt. Opening it, we see the following.
```
Sm9objogV2hhdOKAmVMgVGhhdD8gU2FtdWVsOiBMb2wgQW4gT3Bwb3J0dW5pc3RpYyAxMjIyIDEwMDAgMjIxMiA0MTMyIFBlcnNvbnsgSXMgU29sdmluZyBUaGlzIEVhc3kgTnVsbCBDdGYgQ2hhbGxlbmdlfSBKb2huOldvdyBJdCBJcyBSZWFsbHkgRWFzeS4=
```

Looking at it, experience tells me that it has been encoded using [Base64](https://en.wikipedia.org/wiki/Base64) encoding. From there we simply use a [Base64 decoder](https://www.base64decode.org/) to obtain the following message:
```
John: What’S That? Samuel: Lol An Opportunistic 1222 1000 2212 4132 Person{ Is Solving This Easy Null Ctf Challenge} John:Wow It Is Really Easy.
```

Initially, I knew that the flag had to be within the curly braces but I did not know how it was hidden. After staring at the message for a long time, I realised that combining the last digit of the numbers give us `2022`. As such, I tried taking the last character of each word in Samuel's dialogue. This gave me:
```
LNC2022n{sgsylfe}
```

At this point I was wondering why there was a letter `n` there and why did they put the word `Person` between `2022` and `{`. I was also wondering why there was a spacing between `{` and `Is`. 

After awhile, it made more sense to me that the flag was the last character of each item separated by a space rather than the last character of each word. This meant that `{` was the last character of `Person{` and `}` was the last character of `Challenge}`. With that information, I was able to construct the flag.
```
LNC2022{SGSYLF}
```