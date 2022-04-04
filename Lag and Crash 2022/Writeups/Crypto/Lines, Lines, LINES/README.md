# Lines, Lines, LINES

---

## Challenge Description 
Commisoner Gordon sent me this weird picture found at the scene of a crime. What does it even mean?

Submit flag in all lowercase e.g. LNC2022{flagoeshere}

---

## Challenge Solution
For this challenge, we were provided the following image.

![lines](https://user-images.githubusercontent.com/101789488/160417960-41c1d3f5-8a80-4c41-a073-c29b4b3702d4.png)

When I saw this, the only thing I could think of doing was to search up line cipher. After browsing through images, I managed to find this [image](http://www.geocities.ws/fintan_58/writings.htm) which seemed to be of relevance to this challenge. Afterwards, I found out that this was [ogham](https://en.wikipedia.org/wiki/Ogham#:~:text=Ogham%20(%2F%CB%88%C9%92%C9%A1%C9%99m,%2C%206th%20to%209th%20centuries).).

![Screenshot 2022-03-28 221818](https://user-images.githubusercontent.com/101789488/160418347-f5530742-bbc0-423b-9238-447077cc38ec.png)

Using the image as reference, you should get the following:
```
neergnegnaro
```

From here, all you have to do is reverse the order of the characters to get the flag.
```
LNC2022{orangengreen}
```