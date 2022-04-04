# Invisibility

---

## Challenge Description 
We have been keeping an eye on this evil lair for some time now. It would be useful if we could look up though.

---

## Challenge Solution
For this challenge, we were provided with a image file lair.bmp. If we open the file, we see the following image.

![Screenshot 2022-03-31 204120](https://user-images.githubusercontent.com/101789488/161057026-9624d25c-4a10-44c9-bb3d-40cded4d41d6.png)

From the challenge description, notice it hints us that it would be useful if we could look up. This lead me to think that maybe the upper portion of the image is what we want to focus on. Since the challenge is called `Invisibility`, my first thought is that maybe they hid some information inside the image. To see if any steganography was used, I submitted the image on [aperisolve](https://aperisolve.fr/).

From the results, we got nothing much except this image that tells us we are looking at the wrong place.

![Screenshot 2022-03-31 205005](https://user-images.githubusercontent.com/101789488/161058586-88141c6a-16c9-46b7-92a7-c78354488001.png)

Coincidentally enough, one of my teammates searched up the image and found the original image as shown. From this, we can infer that the image the challenge provided us with was cropped. Maybe what they meant by looking up is to somehow revert the image to its original size and view the upper portion?

![Screenshot 2022-03-31 205144](https://user-images.githubusercontent.com/101789488/161059191-8ca1d83e-9e0a-48c2-863f-ee183267ae9b.png)

With the assumption that we had, I did some research and found this [writeup](https://www.aperikube.fr/docs/pragyan_2018/pictorial_mess/) whereby they altered the height of the image to find what was hidden. Since the challenge in the writeup featured a png image, I had to figure out how we were going to do the same with a bmp image.

For this next part, I opened lair.bmp in a [Hex Editor](https://mh-nexus.de/en/hxd/) to try and alter its height in an attempt to reveal the upper portion of the image. Since I did not know what each byte represented, it took me some time before I figured out which byte I needed to change. If we take a look at the second row, we see a value `FA` which represents `250` in hexadecimal. Since the dimensions of lair.bmp are 631 for the width and 250 for the height, we can assume that the byte `FA` represents the height of the file.

![Screenshot 2022-03-31 210632](https://user-images.githubusercontent.com/101789488/161061833-e1d5ab0b-e033-4a39-89ee-8c06ae398449.png)

At this point, I was extremely lost because the only thing I knew to do was to try increasing the size from `FA` to `FF` since I did not know how changing the bytes on the left or right will affect the file. In the end, I thought of an idea which was to download sample bmp files from the internet and compare the bytes. Doing this, I was able to find out that changing the size from `FA` to `DD 01` would give us the upper portion of the image.

![Screenshot 2022-03-31 211935](https://user-images.githubusercontent.com/101789488/161065089-2037a0cd-8d52-470a-8bad-c42894fac931.png)

![Screenshot 2022-03-31 211901](https://user-images.githubusercontent.com/101789488/161066187-f20d599a-0baf-4fb0-8bcb-4fb310e15e58.png)

Finally, just convert the result from [hexadecimal to text](http://www.unit-conversion.info/texttools/hexadecimal/) and we've got the flag.
```
LNC2022{h1dd3n_1n_pl41n_s1ght_271021}
```