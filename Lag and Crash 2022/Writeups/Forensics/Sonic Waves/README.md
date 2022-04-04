# Sonic Waves

---

## Challenge Description 
Rumour has it that there is a new superhero in town, her name is Siryn. She seems to keep echoing the passcode: SCRE4M_4_M3. Find out what powers she has and how to defeat her.

---

## Challenge Solution
For this challenge, we were provided with an image file Siryn_Powers.png. If we open it, we see the following image.

![Siryn_Powers](https://user-images.githubusercontent.com/101789488/161089396-732730e5-88ac-44a3-a8e5-7c10b7767c9d.png)

My first thoughts when I tried this challenge was that steganography had to be involved and so that was the reason for the passcode. The thing was that tools such as [steghide](https://www.kali.org/tools/steghide/) would not work on png images. Additionally, when I tried submitting the image to [aperisolve](https://aperisolve.fr/), I got no results. After seeing that more people managed to solve the challenge, I thought I'd buy the hint but it wasn't as helpful as I thought it was since that was the idea I already had.

![Screenshot 2022-03-31 231922](https://user-images.githubusercontent.com/101789488/161090666-8105c3bc-8f31-424a-9f4d-d735c35a9335.png)

After much research and time, I found out about this [website](https://www.pelock.com/products/steganography-online-codec) that can help us retrieve a password encrypted message within the image. After decoding, we ended up with the following link.

![Screenshot 2022-03-31 232433](https://user-images.githubusercontent.com/101789488/161091689-134074d3-0521-434e-9bb7-031fb61c12d1.png)

Upon visiting the link, we see a audio file parad0x.wav which we then downloaded. Since the sound from the audio file was unrecognizable, we used [Audacity](https://www.audacityteam.org/download/) to open the audio file in spectrogram view and found the flag.

![Screenshot 2022-03-31 232610](https://user-images.githubusercontent.com/101789488/161092482-21363b48-7e5f-45d8-aa41-9a70a6ff160b.png)
```
LNC2022{PaR4d0x_Eff3Cts}
```