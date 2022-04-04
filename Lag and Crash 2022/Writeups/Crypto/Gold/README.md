# Gold

---

## Challenge Description 
The golden bug seems to be mischievous...

NOTE: Flag is not in flag format, input flag as LNC2022{FLAG} (all uppercase)!

---

## Challenge Solution
For this challenge, we were provided a text file gold.txt. Upon opening it, you should see the following.
```
:5:3‡0†-6.48(†8-‡†8†
```

Since I have seen a challenge similar to this in `Lag and Crash 2021`, I knew that this had to be related to the gold bug cipher. With that information, I searched up gold bug decoder and found this [decoder](https://www.dcode.fr/gold-bug-poe). Using the decoder, I simply pasted the ciphertext from gold.txt and decrypted it to get the flag.

![Screenshot 2022-03-28 225743](https://user-images.githubusercontent.com/101789488/160426778-6eade4eb-691e-4c26-ba15-0cc0dcd53000.png)
```
LNC2022{YAYGOLDCIPHERDECODED}
```