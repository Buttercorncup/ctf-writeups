# Plumber

---

## Challenge Description 
I wonder whats within all this output, maybe there is a flag? I wonder if there is a way to output and obtain just the flag ...

`nc c1.lagncrash.com 9006`

---

## Challenge Solution
For this challenge, we are provided the command used to obtain the output. First we want to see what we get if we run the command. After running the command, it seems like the server constantly outputs unrecognized characters until the user stops it.

![Screenshot 2022-03-30 005838](https://user-images.githubusercontent.com/101789488/160666359-a22c3ada-a6ef-40a9-a2a9-54e372cd1b96.png)

From the challenge description, we are hinted that the way to obtain the flag is to ignore the useless output and find just the flag. With that in mind, the way to solve this would be to use the [grep](https://www.ibm.com/docs/en/aix/7.2?topic=g-grep-command) command. Note that if we do not specify `-a`, the output will not be displayed. To connect to the server and output just the flag, we will use the following command:
```
nc c1.lagncrash.com 9006 | grep -a "LNC2022"
```

As simple as that, we managed to obtain the flag.

![Screenshot 2022-03-30 010739](https://user-images.githubusercontent.com/101789488/160667131-2c6436da-0697-43ad-b52b-e15897c828de.png)
```
LNC2022{w3lc0me_t0_th3_aBy$s}
```