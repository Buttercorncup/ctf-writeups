# Found Disk

---

## Challenge Description 
A group of citizens were kidnapped, with only a file left behind... could you find out who did it?

---

## Challenge Solution
For this challenge, we were provided with a docx file disk.docx. If we try to open it, we get an error that says the file contains unreadable content.

![Screenshot 2022-03-31 025855](https://user-images.githubusercontent.com/101789488/160910546-fce25b34-2ad8-470f-bfd3-8729163e6dae.png)

Since trying to recover the contents brought us nowhere, my teammate used the [file command](https://www.geeksforgeeks.org/file-command-in-linux-with-examples/) to check the file type of disk.docx. From there, we discovered that the file extension was actually changed and the file is actually an audio file.
```
file disk.docx
```

![Screenshot 2022-03-31 024102](https://user-images.githubusercontent.com/101789488/160908020-eba1d5a4-e959-4cd9-b268-82b37101e5d9.png)

After changing the file extension and listening to the audio file, it seemed to be some unrecognizable sound. My teammate then opened the audio file in [Audacity](https://www.audacityteam.org/download/) and switched to spectrogram view which gave us the following link.
```
bit.ly/3no0rWY
```

![Screenshot 2022-03-31 024730](https://user-images.githubusercontent.com/101789488/160908816-7cb388a1-3ef4-4fcf-b774-ce0ca23250df.png)

If we visit the link, it brings us to a page that allows us to download an image file miXtapefORmusic.jpg. Looking at the name of the file, we knew that it had to be related to the XOR operation since the respective letters were capitalized. Since we needed more information in order to move on, I decided to make use of the [exiftool command](https://exiftool.org/examples.html) to extract the metadata from the file.

From the results, it seemed to me that there were two pieces of information that would be crucial to obtaining the flag. Under `Artist` and `Owner Name`, we see that there are two strings in hexadecimal that are of the same length. Isn't it a coincidence that we're trying to find two values to perform a XOR operation?
```
Artist: 1c2a31423d212c4c9e12541a0d13253252216e2c112c7e
Owner Name: 506472700d131e37d36761736e5e115c63522c18724703
```

![Screenshot 2022-03-31 025122](https://user-images.githubusercontent.com/101789488/160909817-b98724c1-bc87-451e-99e9-058331cdc2f2.png)

With that being said, I used a [XOR calculator](https://xor.pw/#) to calculate the result and then converted it from [hexadecimal to text](http://www.unit-conversion.info/texttools/hexadecimal/) to obtain the flag.
```
LNC2022{Mu5icM4n1sB4ck}
```