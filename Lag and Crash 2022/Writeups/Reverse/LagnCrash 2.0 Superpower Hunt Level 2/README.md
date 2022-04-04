# LagnCrash 2.0 Superpower Hunt Level 2

---

## Challenge Description 
A superhero origin story where you set out to get superpowers while you mourn the death of your cat, Lord Fuzzball. Part 2 of the series!

(The SuperpowerHunt.zip file used is the same as Level 1)

---

## Challenge Solution
For this challenge, we were provided with a zip file containing a game as well as its related files. Upon opening the game, this is what we see.

![Screenshot 2022-03-28 161810](https://user-images.githubusercontent.com/101789488/160356089-decd245f-235d-46f8-816e-564f17318060.png)

Since this challenge made use of the same game as `LagnCrash 2.0 Superpower Hunt Level 1`, you would have already obtained the flag if you did that challenge. Else make use of the [cat](https://www.tecmint.com/13-basic-cat-command-examples-in-linux/) and [grep](https://www.ibm.com/docs/en/aix/7.2?topic=g-grep-command) command:
```
cat level0 | grep -a "LNC2022"
```

The output is as shown:
```
ף=o�:�?�?�?5!�B�?�u�B�?�?�?????ff�>���@�A�@??�5?�5?�L�B�?�?�?????�z?�«A�A�@??�?�?�?�?�????t���
                                                                                              ��CΪ�B??�?�?�?�?�????t���
                                                                                                                       ��CΪ�B??!�?���d5�?�?�?�?�????"�/hC�A??��?�?�?�? LNC2022{pUtt1ng_7he_p4sT_b3h1nd}M�����?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�Pu����� A A��A�B�����>���?�?�?�?LNC2022{4m_v3rY_h1Gh}M��?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�?�Pu�����@@@@��A�B�����>���?�?�?�?G(�?Hey.
(�?LNC2022{0h_h3y_y0u_n0tic3d}E~�j A@@�������>@
```

Just like that, we have obtained the flag.
```
LNC2022{pUtt1ng_7he_p4sT_b3h1nd}
```