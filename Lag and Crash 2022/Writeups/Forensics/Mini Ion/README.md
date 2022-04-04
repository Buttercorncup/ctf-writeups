# Mini Ion

---

## Challenge Description 
The mischievous minions were caught in the jail and you are tasked to find the message to release them!

---

## Challenge Solution
For this challenge, we were provided with a zip file challenge.zip. If we unzip it, we will be able to find a zip file `.config.7z` and two folders named `Images` and `Text`. Within the `Images` folder, we have a few images of minions, with one that cannot be opened due to errors. As for the `Text` folder, it contains 40 different folders with many text files inside each of them. Within each text file is a a long string that seems to be encoded using [Base64](https://en.wikipedia.org/wiki/Base64).

![Screenshot 2022-03-31 215904](https://user-images.githubusercontent.com/101789488/161075016-b8eb5182-a288-455e-85c6-66b1a69a7d7a.png)

![Screenshot 2022-03-31 220511](https://user-images.githubusercontent.com/101789488/161075056-c406a550-8a94-492c-87e3-4504d8afb51b.png)

![Screenshot 2022-03-31 220537](https://user-images.githubusercontent.com/101789488/161075095-0fe59a06-94ac-442d-928e-7bae2d019f9a.png)

![Screenshot 2022-03-31 221148](https://user-images.githubusercontent.com/101789488/161075697-8441f28c-03b5-4ef2-b8d9-af21290eea72.png)

My teammate first tried to find out what was wrong with the corrupted image using the [strings](https://www.ibm.com/docs/en/aix/7.2?topic=s-strings-command) command:
```
strings 6.kevin.jpg
```

From the output, we found something interesting:
```
"Look for : NzYxN"
```

If we relate this to the strings in the text files, what we need to do is to open every single text file within the folders and find the text file whereby its string contains `NzYxN`. To do this, I did some research and found this [website](https://www.digitalocean.com/community/tutorials/workflow-loop-through-files-in-a-directory) which taught me how to loop through files in a directory. Since I wanted to look through each folder, I altered the command to display the contents of the text file if it contains `NzYxN` in its string.
```
for FILE in *; do cat $FILE | grep "NzYxN" ; done 
```

With that, I was able to find the following:
```
MjAxNzYxNUlzVGhlWmlwUHdk
```

If we decode it using a [Base64 decoder](https://www.base64decode.org/), we get the password for the zip file `.config.7z`.
```
2017615IsTheZipPwd
```

From there, we unzip `.config.7z` and what we got was two files named `hello` and `PASSWORD.gif`. If we open up the gif file, we see that each frame is part of an image that was split to hide its contents. For this part, my teammate used an online tool to split the gif according to the frames and then pieced them together to get the following image. On a side note, I used [exiftool](https://exiftool.org/examples.html) and was able to obtain the password as well.

![Screenshot 2022-03-31 225005](https://user-images.githubusercontent.com/101789488/161084305-f1664f14-4ff2-4941-a5c8-4c1ea2489016.png)

![Screenshot 2022-03-31 225201](https://user-images.githubusercontent.com/101789488/161084762-231dc322-2fba-4e21-a78b-9fd8bfab12e3.png)

Since the image also included `.7z` in it, we tried renaming the `hello` file with the extension to get `hello.7z`. Even after doing this, its contents still could not be extracted. To fix this, we used a [Hex Editor](https://mh-nexus.de/en/hxd/) to edit the header to that of a 7-Zip file.

![Screenshot 2022-03-31 230411](https://user-images.githubusercontent.com/101789488/161087449-b1bfca1e-d9f1-4db8-a066-094d0d6d4e09.png)

After that, all we had to do was enter the password `F3L0n1ous_GRu` and we obtained the flag.
```
LNC2022{Th3_de5P!ca6Le_y0U}
```