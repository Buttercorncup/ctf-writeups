# ezgo

---

## Challenge Description 
Go get the flag (no pun intended)

---

## Challenge Solution
For this challenge, we were provided with the main program in both Linux and Windows, as well as the program's source code. Firstly, I executed the program to see what it was and got the following output:
```
[================================]
Please access http://localhost:8080/view/arandomstring
For example: http://localhost:8080/view/ANewPage
if there's any error, just change the random string
[================================]
```

Upon visiting the link, I see the following webpage.

![Screenshot 2022-03-28 042422](https://user-images.githubusercontent.com/101789488/160299689-4b24444d-b081-489f-bbad-290d15a64f76.png)

I tested the program and realised that it displays the whatever the user inputs. This made me think that the webpage could be vulnerable to [Server Side Template Injection (SSTI)](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection). Similar to other challenges, I tried the following payload:
```
{{7*7}}
```

However, it failed and displayed:
```
This page isn't working
```

Honestly at this point I was extremely lost and consulted a friend that told me about [SSTI in GO](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection#ssti-in-go). After more research, I understood how SSTI in GO worked from this [webpage](https://www.onsecurity.io/blog/go-ssti-method-research/).

Since I knew that the objective was to call a function, I took a look at the source code given and noticed the `Getflag()` function.

![Screenshot 2022-03-28 042334](https://user-images.githubusercontent.com/101789488/160300015-bb5614e1-50f2-4a9b-a3f3-258619abb7bf.png)

From there, I went back to the webpage and tried the following payload:
```
{{.Getflag}}
```

With the payload, I managed to obtain the flag.

![Screenshot 2022-03-28 043953](https://user-images.githubusercontent.com/101789488/160300123-bff13d5f-89e2-43eb-a251-f7a3dfde275f.png)

```
LNC2022{G0_gO_g0Oo}
```