# fone

---

## Challenge Description 
What's that noise?

---

## Challenge Solution
For this challenge, we were provided with an audio file Fone.wav. If we take a listen, it sounds like someone dialing numbers on a phone. Since I had not seen any challenge like this before, I did some research on phone number audio decoders and found a [DTMF decoder](https://github.com/ribt/dtmf-decoder).

If we go ahead and decode the audio file, we will get the following result.

![Screenshot 2022-03-31 003851](https://user-images.githubusercontent.com/101789488/160886985-40c3e8b1-bde3-4b8a-93d7-ad13be106725.png)

Looking at the string, I imagined that there was something else that we have to do in order to get the flag. After much research and malding, one of my teammates suggested submitting the output as the flag. However funny this might seem, it was actually the flag.
```
LNC2022{763*132ABD32C190##256791A}
```