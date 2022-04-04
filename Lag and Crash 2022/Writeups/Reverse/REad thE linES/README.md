# REad thE linES

---

## Challenge Description 
Supervillain Reece gave us this file. He believes that we can't change what he did. Prove him wrong!

---

## Challenge Solution
For this challenge, we are provided with a binary executable file named `REadthElinES`. For the first part, my teammate did some analysis on the file which lead him to use the [strings](https://www.ibm.com/docs/en/aix/7.2?topic=s-strings-command) and [grep](https://www.ibm.com/docs/en/aix/7.2?topic=g-grep-command) command:
```
strings REadthElinES | grep "LNC2022"
```

The output gave us the first half of the flag:
```
LNC2022{D1d_y0u88553s!!
```

Next, we are going to use the [chmod](https://www.howtogeek.com/437958/how-to-use-the-chmod-command-on-linux/) linux command to change the access permissions of the file and make it executable:

```
chmod +x REadthElinES 
```

If we execute the file, it will prompt us for a password.

![Screenshot 2022-03-28 173750](https://user-images.githubusercontent.com/101789488/160370684-6da241c6-c6ce-4b9c-94a9-110bcb13caf8.png)

Next, we want to make use of a debugger like [GDB](https://www.tutorialspoint.com/gnu_debugger/what_is_gdb.htm). I used [GEF (GDB Enhanced Features)](https://github.com/hugsy/gef) for this part. Starting off, I used the command to pull up the gef interface:
```
gdb REadthElinES
```
![Screenshot 2022-03-28 181419](https://user-images.githubusercontent.com/101789488/160377113-eddc16bb-beb7-4c3d-883c-5b5692f3bc9a.png)

From here, the next step is to set a breakpoint in the `main` function of the program using the command:
```
entry-break
```
![Screenshot 2022-04-01 010404](https://user-images.githubusercontent.com/101789488/161110849-ca76632c-2f9a-4fa3-bc3e-7c308c0a4332.png)

Now what we want to do is look through the assembly instructions and see if there is anything interesting. It would be easier if you have some understanding of [x86 Assembly](https://en.wikipedia.org/wiki/X86_assembly_language). So, I first [disassembled](https://visualgdb.com/gdbreference/commands/disassemble) `main` using the command:
```
disassemble main
```
![Screenshot 2022-03-28 170817](https://user-images.githubusercontent.com/101789488/160365262-b2c22e17-d8ec-407e-a1ca-c795070069ef.png)

If we take a look at the following lines, we can see a `__isoc99_scanf@plt` function which takes in user input. From there, we can infer that the `strcmp@plt` function makes a string comparison between the user input and a value, returning 0 only if they are equal. The value returned is then stored in `eax`. Following that, the program executes a [test](https://en.wikipedia.org/wiki/TEST_(x86_instruction)) instruction which sets the [Zero Flag](https://en.wikipedia.org/wiki/Zero_flag) only if eax is 0. The [je (Jump if equal)](http://unixwiz.net/techtips/x86-jumps.html) instruction after that is only executed if the Zero Flag is set. Looking at the rest of the code, we can see some interesting functions `reee100` and `reee200`.
```
0x0000555555555c66 <+160>:   call   0x555555555060 <strcmp@plt>
0x0000555555555c6b <+165>:   test   eax,eax
0x0000555555555c6d <+167>:   je     0x555555555c7b <main+181>
0x0000555555555c6f <+169>:   mov    eax,0x0
0x0000555555555c74 <+174>:   call   0x55555555575f <reee100>
0x0000555555555c79 <+179>:   jmp    0x555555555c85 <main+191>
0x0000555555555c7b <+181>:   mov    eax,0x0
0x0000555555555c80 <+186>:   call   0x555555555bb6 <reee200>
0x0000555555555c85 <+191>:   mov    eax,0x0
0x0000555555555c8a <+196>:   leave  
0x0000555555555c8b <+197>:   ret 
```

If we disassemble both the `reee100` and `reee200` functions, we notice more of such `reee` functions nested within. If we were to disassemble all of them, we find one more `strcmp@plt` function under `reee70`.

![Screenshot 2022-03-28 173102](https://user-images.githubusercontent.com/101789488/160369469-0246269c-69ee-4f9d-8f94-97a3c03e1c57.png)

From the challenge description, we know that Supervillain Reece believes that we can't change what he did. So maybe if we change the way the program operates, we can get the flag. For example, the test instruction will cause the program to not execute the je instruction in `main` and hence not take the jump. So what we want to do is take the jump. 

To do that, we will first enter `ni` which means next instruction, until we hit the je instruction. From the image, we can see a `NOT taken [Reason: !(Z)]`. We can also see that after I displayed the flags, `zero` is not bolded. What this basically means is that the jump is not taken because the Zero Flag is not set. 

![Screenshot 2022-03-28 174301](https://user-images.githubusercontent.com/101789488/160371661-ab78e83a-0938-4d37-a716-9cd4150613c3.png)

To set the Zero Flag, we just enter the following:
```
flags +ZERO
```

Now that the Zero Flag is set, going to the next instruction will bring us to the `reee200` function instead of the `reee100` function. Upon reaching the `reee200` function, what we want to do is step into the function to examine it using `si` which means step into. Doing that will show us that there is a `reee78` function nested inside.

![Screenshot 2022-03-28 175007](https://user-images.githubusercontent.com/101789488/160372906-f1199884-cc55-499f-885f-b0f2a6781a02.png)

Next, we want to do the exact same thing as before, which is to go to the next instruction until we hit `reee78`, and then step into the function. Doing that will show us a `reee70` function.

![Screenshot 2022-03-28 175300](https://user-images.githubusercontent.com/101789488/160373434-988eff42-bd3a-478d-8f0e-e433d6cbf3e3.png)

Once again, we want to step into the `reee70` function. Doing so will bring us to the second `strcmp@plt` function. Going to the [jne (Jump if not equal)](http://unixwiz.net/techtips/x86-jumps.html) instruction, we can see that once again the Zero Flag is not set, hence the jump is taken.

![Screenshot 2022-03-28 175441](https://user-images.githubusercontent.com/101789488/160373766-410202d9-e4fe-47f3-91af-1f451804206c.png)

Since we want to change what the program is meant to do, we want to not take the jump. To do that, we set the Zero Flag:
```
flags +ZERO
```

Now that we have already bypassed all the `strcmp@plt` functions, we have changed the logic of the program. Now lets try continuing using:
```
continue
```

With that, we were able to get the program to print the second half of the flag.

![Screenshot 2022-03-28 175852](https://user-images.githubusercontent.com/101789488/160374622-de3b50ad-7171-41e3-8486-70b889ff47bb.png)
```
_F1nD_M3}
```

Combining the two parts of the flag, we get:
```
LNC2022{D1d_y0u88553s!!_F1nD_M3}
```