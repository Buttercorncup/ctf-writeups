# Side-kick

---

## Challenge Description 
Behind every Superhero is a great Side-Kick! Or is it beside? I dunno just go figure out what this means ;p

---

## Challenge Solution
For this challenge, we were provided an audio file sidekick.wav. For the first step, it might require some experience with audio sounds. Listening to the audio file, we narrow down the sound to [Slow-Scan Television (SSTV)](https://en.wikipedia.org/wiki/Slow-scan_television). Since SSTV is a picture transmission method, we want to decode the audio file using an [SSTV decoder](https://github.com/colaclanth/sstv) to obtain a picture. To do so, we use the following command:
```
sstv -d sidekick.wav -o result.png
```

After successfully decoding the audio file, we obtain the following image.

![Screenshot 2022-03-30 011901](https://user-images.githubusercontent.com/101789488/160669023-b455d7dd-45b3-4ca3-940b-73cae3d58474.png)

If we take a closer look at the image, we can tell that it resembles a QR code. For this part, one of my teammates stretched the image and added the missing components to repair the QR code. After doing so, we obtained the following QR code.

![Screenshot 2022-03-30 012132](https://user-images.githubusercontent.com/101789488/160669416-e04f9f0b-2bb0-45d6-b99a-c9ca0f41218f.png)

After scanning the QR code, we successfully obtained the flag.
```
LNC2022{s1d3k1ck5_r_kewl2}
```