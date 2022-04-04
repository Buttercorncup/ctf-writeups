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