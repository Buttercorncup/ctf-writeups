# Enhanced ROT13

---

## Challenge Description 
We found a way to enhance the encryption, are you able to break it and recover the flag?

Flag (kinda): LNC2022{qrrqnururqwrbtapqrnqatqqru}

---

## Challenge Solution
For this challenge, we are provided with a python file enhanced_ROT13.py. Opening it, we see the following code.
```
lowercase_offset = ord("m") 
uppercase_offset = ord("M") 
alphabet = "abcdefghijklmnopqrxtuvwxyz"
 
## flag example is not the flag 
Flag_example = "testing" 
 
def b16(plain):  
 
    enc = ""  
 
    for a in plain:  
         binary = "{0:08b}".format(ord(a))  
         enc += alphabet[int(binary[4:], 2)]  
         enc += alphabet[int(binary[:4], 2)]  
    return enc  
 
def ROT13(b16):  
    result = ""  
    for i in b16:  
        c = ord(i)  
        if c >= ord('a') and c <= ord('z'):  
            if c > lowercase_offset:  
                c -= 13  
            else:  
                c += 13  
        elif c >= ord('A') and c <= ord('Z'):  
            if c > uppercase_offset:  
                c -= 13  
            else:  
                c += 13  
        result += chr(c)  
    return result  
 

flag=Flag_example 
 
b16 = b16(flag) 
print("Flag:  LNC2022{"+ROT13(b16)+ "}" )
```

Looking at the code, the logic is that the flag is first encrypted using the `b16` function, and then the `ROT13` function. From here, I knew that I could ignore the `ROT13` function because I could use a [ROT13 decoder](https://rot13.com/) for decryption. After decrypting `LNC2022{qrrqnururqwrbtapqrnqatqqru}`, we got `LNC2022{deedahehedjeogncdeadngddeh}`.

Since I don't have too much experience with coding, I added `print` statements and comments to help me understand the `b16` function better. So firstly, the function takes each character of the flag from the start and changes it to decimal, then binary (8-bits). After that, it splits the binary into half (4-bits for each half). Following the previous step, the second half of the binary, followed by the first half of the binary, is converted into an integer, then a character, and is finally appended to the scrambled string.
```
def b16(plain):  
 
    enc = ""  
 
    for a in plain:  
         print(a)
         print("a: " + str(ord(a)))
         binary = "{0:08b}".format(ord(a))  # Change char to decimal, then decimal to binary
         
         print("binary: " + binary)
         print(binary[4:])
         print(int(binary[4:], 2))
         print(binary[:4])
         print(int(binary[:4], 2))
         enc += alphabet[int(binary[4:], 2)]  # Convert 2nd half binary to int then char
         print("enc1: " + enc)
         enc += alphabet[int(binary[:4], 2)]  # Convert 1st half binary to int then char
         print("enc2: " + enc)
    return enc
```

From there, I knew that I had to reverse the logic of the code and take the following steps:
```
For every two letters from the back of the string "deedahehedjeogncdeadngddeh":
1. Get the character position in integer for both letters (E.g. H is 7 and E is 4)
2. Convert character position in integer to binary for both letters (E.g. 7 is 0111 and 4 is 0010)
3. Combine the two 4-bit binary numbers to form one 8-bit binary number (The letter on the right should be the first half of the binary)
4. Convert the binary number to decimal
5. Convert decimal to ASCII
6. Append to the flag
7. Repeat
8. Reverse the flag
```

Using the logic from above, I constructed the following python script:
```
from audioop import reverse

# LNC2022{deedahehedjeogncdeadngddeh}
string = ['de', 'ed', 'ah', 'eh', 'ed', 'je', 'og', 'nc', 'de', 'ad', 'ng', 'dd', 'eh'] # Separate into list since 2 chars equal 1 char in the flag
binary = ""
flag = ""

# For loop to iterate backwards
for i in reversed(string):
    for r in reversed(i):
        position = ord(r) - 97 # Get the character position in integer
        binary += "{0:04b}".format(position) # Convert character position in integer to binary and append
    if (len(binary) == 8):
        integer = int(binary, 2) # Convert binary to decimal/integer
        char = chr(integer) # Convert decimal/integer to ASCII/character
        binary = "" # Empty the variable for the next item in the list
        flag += char # Append the character to the flag

flag = flag[::-1]
print(flag) 
```

After running the python script, we managed to obtain the flag.
```
LNC2022{C4pt4In-C0m3t}
```