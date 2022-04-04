# Norse_God

---

## Challenge Description 
Hmm where did Thor came from? What the old gods religion?

NOTE: Flag is not in flag format, input flag as LNC2022{FLAG}!

---

## Challenge Solution
For this challlenge, we were provided with a image file Party_Thor.png. Since the challenge description asked two questions, I thought that searching them up will give me some useful information to find the flag. From searching, all I found was the answer [Norse Mythology](https://en.wikipedia.org/wiki/Norse_mythology#:~:text=Norse%20or%20Scandinavian%20mythology%20is,folklore%20of%20the%20modern%20period.).

From there I was thinking maybe we can get some other information out of the image file. I first tried the [strings](https://www.ibm.com/docs/en/aix/7.2?topic=s-strings-command) command:
```
strings Party_Thor.png
```

Since the strings command was not able to find anything useful, I tried again using exiftool:
```
exiftool Party_Thor.png
```

Looking at the ouput, I noticed that there was a strange string that seemed to be encoded using [Base64](https://en.wikipedia.org/wiki/Base64) encoding.

![Screenshot 2022-03-29 171313](https://user-images.githubusercontent.com/101789488/160578142-ee25cf1d-cc7d-4ef5-b6fb-f29b6085caa4.png)

```
4ZuW4Zqo4ZuK4ZuBX+GavjDhmrHhm4rhm5Zf4Zqx4Zqi4Zq+4ZuW4ZuK
```

Using a [Base64 decoder](https://www.base64decode.org/), I decoded the string and got the following output:
```
ᛖᚨᛊᛁ_ᚾ0ᚱᛊᛖ_ᚱᚢᚾᛖᛊ
```

At this point, I thought that the output was just letters but displayed in a weird or distorted way, so I tried figuring out what it actually says. After trying to decipher the message, I realised that even if it was distorted, it would not make any sense. Remembering the answer I found from searching up the questions in the challenge description, I decided to search up norse decoder and found this [website](https://www.dcode.fr/elder-futhark).

It seemed that the message was actually a Elder Futhark ciphertext which the website would be able to decode. Decoding the ciphertext, we get the flag.

![Screenshot 2022-03-29 172357](https://user-images.githubusercontent.com/101789488/160579625-aa7624d0-8eb3-4e09-a088-a198b1bd8323.png)
```
LNC2022{easi_n0rse_runes}
```