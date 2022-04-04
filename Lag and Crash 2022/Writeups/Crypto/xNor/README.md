# xNor

---

## Challenge Description 
Not so commonly used algorithm..

So many weird characters.... I think I should change them into something else..

---

## Challenge Solution
For this challenge, we are provided with a text file xNor.txt. Opening it, we see the following.
```
a = "ØÝÚ¬¼©¨éÔðÿñÄÓÿü®â.«p."
b = "klfasdemsanaisdf2r¢1ë¬"
```

When I first saw xNor, I went and searched up on [XNOR Calculator](https://codebeautify.org/xnor-calculator). Since the challenge description also hinted us to change the weird characters into something else, converting [ASCII to binary](https://www.binaryhexconverter.com/ascii-text-to-binary-converter) should do the trick considering it is an XNOR operation. When converting, it is necessary to remove the padding or else it will interfere with the XNOR operation.

![Screenshot 2022-03-28 215159](https://user-images.githubusercontent.com/101789488/160413199-6e1a4040-5729-4def-b01f-8ae60a2c69f2.png)

![Screenshot 2022-03-28 215219](https://user-images.githubusercontent.com/101789488/160413214-90273d23-6d58-403b-a98a-0855c258684a.png)

Using the two binary values obtained, we will feed them into the XNOR Calculator to obtain the result.

![Screenshot 2022-03-28 220145](https://user-images.githubusercontent.com/101789488/160414966-71487d83-ce52-48d7-9525-d0d0f1028f27.png)

Converting the binary back to ASCII will give us the flag.

![Screenshot 2022-03-28 215520](https://user-images.githubusercontent.com/101789488/160413662-84d3fde3-7cab-4969-ad96-d292867bc1de.png)

On a side note, here's a python script I created that gives you the flag but I can't explain it fully.
```
# Converted each string to decimal and seperated them within a list
a = [216, 221, 218, 172, 188, 169, 168, 233, 212, 240, 255, 241, 196, 211, 255, 252, 174, 226, 46, 171, 112, 46]
b = [107, 108, 102, 97, 115, 100, 101, 109, 115, 97, 110, 97, 105, 115, 100, 102, 50, 114, 162, 49, 235, 172]

flag = ""
i = 0

# Perform Bitwise XOR and NOT operation between each item in list a and b and add to flag
while i < 22:
    result = (a[i] ^~ b[i]) + 256
    result = chr(result)
    flag += (str(result))
    i+=1

print(flag)
```

Flag:
```
LNC2022{XnnoR_decosed}
```