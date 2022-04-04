# HandsomeHerooooo

---

## Challenge Description 
A cast from At a Distance, Spring is Green!!!! Help him uncover the hidden file.

---

## Challenge Solution
For this challenge, we were provided with an image file b.jpg. Reading the challenge description, a hidden file hints us that [steganography](https://www.comptia.org/blog/what-is-steganography#:~:text=Steganography%20is%20the%20practice%20of,to%20friends%20using%20invisible%20ink.) was used to conceal the flag. If we open the image, we see a picture of an actor from "At a Distance, Spring is Green".

![b](https://user-images.githubusercontent.com/101789488/160914161-36685845-26a5-42c9-a64a-f515b0337923.jpg)

If we look up the cast for the series, we will be able to find the actor's name Bae In-hyuk. However, in the event that it is difficult to tell, you can always do a reverse image search.

![Screenshot 2022-03-31 032337](https://user-images.githubusercontent.com/101789488/160914565-f226d6e6-22f3-45ec-826c-638099b020ed.png)

Since we know that steganography was used to conceal a hidden file within the image file, we want to use a tool such as [steghide](https://www.kali.org/tools/steghide/) to recover it. Now, something to know is that when you try to recover a hidden file concealed by steganography, a password will usually be required. For this specific challenge, this part was extremely guessy because we are not sure what exactly the password could be. Since the actor's name contains a hyphen in it, it makes it all the more confusing. After trying multiple passwords, we managed to recover the hidden file using the following command and password:
```
steghide extract -sf b.jpg
```
```
baeinhyuk
```

After executing the command, the hidden file recovered was flag.txt. If we open it up, we see the that the flag seems to be in hexadecimal format.
```
a = 67 75 6D 69 68 6F

flag = LNC2022{a}
```

If we simply convert the string from [hexadecimal to text](http://www.unit-conversion.info/texttools/hexadecimal/), we will get the flag.
```
LNC2022{gumiho}
```