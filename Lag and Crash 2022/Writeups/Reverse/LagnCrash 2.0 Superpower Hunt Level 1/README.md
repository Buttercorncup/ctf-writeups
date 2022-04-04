# LagnCrash 2.0 Superpower Hunt Level 1

---

## Challenge Description 
A superhero origin story where you set out to get superpowers while you mourn the death of your cat, Lord Fuzzball. Part 1 of the series!

---

## Challenge Solution
For this challenge, we were provided with a zip file containing a game as well as its related files. Upon opening the game, this is what we see.

![Screenshot 2022-03-28 161810](https://user-images.githubusercontent.com/101789488/160356089-decd245f-235d-46f8-816e-564f17318060.png)

Looking at the game, I have absolutely no idea what to do. At this point, I just went through the different game related files trying to find the flag with the [cat](https://www.tecmint.com/13-basic-cat-command-examples-in-linux/) and [grep](https://www.ibm.com/docs/en/aix/7.2?topic=g-grep-command) command. The command that gave us what we wanted was:
```
cat level0 | grep -a "LNC2022"
```

With the command, I was able to get the following output:
```
ף=o�:�?�?�?5!�B�?�u�B�?�?�?????ff�>���@�A�@??�5?�5?�L�B�?�?�?????�z?�«A�A�@??�?�?�?�?�????t���
                                                                                              ��CΪ�B??�?�?�?�?�????t���
                                                                                                                       ��CΪ�B??!�?���d5�?�?�?�?�????"�/hC�A??��?�?�?�? LNC2022{pUtt1ng_7he_p4sT_b3h1nd}M�����?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�Pu����� A A��A�B�����>���?�?�?�?LNC2022{4m_v3rY_h1Gh}M��?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�Pu�����@@@@��A�B�����>���?�?�?�?G(�?Hey.
(�?LNC2022{0h_h3y_y0u_n0tic3d}E~�j A@@�������>@
```

Just like that, we have obtained the flag.
```
LNC2022{4m_v3rY_h1Gh}
```