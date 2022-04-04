#  Sussy Web

---

## Challenge Description 
Mr Web has been seen loitering around his favourite restraunt in New York City. He seems to be full of mystery... Check out what you can do to uncover his true identity.

http://c1.lagncrash.com:9009

---

## Challenge Solution
Upon visiting the link provided, you will see the following webpage.

![Screenshot 2022-03-28 023757](https://user-images.githubusercontent.com/101789488/160295737-44b7a58b-b94d-413c-8c26-f7b9ff223ba6.png)

If we inspect the webpage, we can see a secret passphrase in the html source:
```
Secret Passphrase: superspooderman
```

I also tried looking at `style.css` under the page's `Sources` and searched for `LNC2022`, which gave me the first half of the flag as a comment:

![Screenshot 2022-03-28 024254](https://user-images.githubusercontent.com/101789488/160295873-4402607a-05be-4f66-be19-32d8cb70d103.png)
```
LNC2022{zeR0
```

If we look under `Console`, we see that the server failed to load a resource named `spooder-background.jpg`. As such, I tried accessing `/photo/spooder-background.jpg` and found the following image:

![spooder-background](https://user-images.githubusercontent.com/101789488/160296049-2b4a6e6d-2163-420f-a748-dbb0569956f6.jpg)

From here, the only thing that could possibly make sense was [steganography](https://en.wikipedia.org/wiki/Steganography). I tried uploading the image to [aperisolve](https://aperisolve.fr/) but found nothing. At this point, I was just trying everything I knew of and so happened to upload the spiderman image from the main webpage onto aperisolve and saw that someone posted the second half of the flag as a common password.

After doing another challenge `Sonic Waves`, I realised that this [tool](https://www.pelock.com/products/steganography-online-codec) was the key to getting the second half of the flag. All you have to do is upload the spiderman image from the webpage and enter the secret passphrase to obtain the decrypted text message which contains the second part of the flag.

![Screenshot 2022-03-28 030548](https://user-images.githubusercontent.com/101789488/160296660-a6a27ac8-0b0b-48a1-b4f4-69c076a652f2.png)
```
_N1gHtZ}
```

Combining both parts of the flag, we get:
```
LNC2022{zeR0_N1gHtZ}
```