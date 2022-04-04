# Tom's Descend

---

## Challenge Description 
Upon delivery a rather light package to a faceless man, Tom was in awe to discover his computer in the center of the pavilion. Upon getting closer to the monitor, There the company's new website is finally up and Tom is not only eager to find out the standings on the only old browser in the computer.

http://c1.lagncrash.com:9010

---

## Challenge Solution
Upon visiting the link, we see the following webpage.

![Screenshot 2022-03-28 031118](https://user-images.githubusercontent.com/101789488/160296874-959b58ee-f95e-43d0-9cb4-85865e5a9a90.png)

I first tried entering random inputs into the `Username` and `Password` fields but nothing happened. However, when I submitted an input under the `Team Search`, I noticed a table popped up.

![Screenshot 2022-03-28 031430](https://user-images.githubusercontent.com/101789488/160296971-435bad14-103b-48df-b41f-86de40e3d03d.png)

From the challenge description, my brain somehow linked `a rather light package` to [sqlite](https://www.sqlite.org/index.html). As such, I knew that the webpage was vulnerable to [SQL Injection](https://www.w3schools.com/sql/sql_injection.asp). Following this, I searched up for [SQL Injection Payloads](https://github.com/payloadbox/sql-injection-payload-list) and inserted them one by one until I obtained some results. One of the payloads that worked was:
```
admin' or 1=1/*
```

Using the payload above, I was able to get the webpage to display a table with half of the flag in it.

![Screenshot 2022-03-28 031912](https://user-images.githubusercontent.com/101789488/160297156-399a9a1f-6067-4c07-8ed6-56efa149245e.png)
```
First half: LNC2022{T0m_Is
```

At this point, I was clueless as to where to get the second half of the flag, so my teammate and I researched on [SQL Injection using UNION](https://portswigger.net/web-security/sql-injection/union-attacks) as well as [SQLite Schema Table](https://www.sqlite.org/schematab.html) in an attempt to find the other existing tables within the database. After some time, my teammate managed to come up with the following payload:
```
' OR 1=1 UNION SELECT name, tbl_name FROM sqlite_master --
```

With the payload, we obtained the following results.

![Screenshot 2022-03-28 032242](https://user-images.githubusercontent.com/101789488/160297320-912d5fc9-78bf-4e62-88fe-96d67360e15d.png)

From here, I knew that there were other tables in the SQL database with the names `teams` and `users`. As such, I tried the following payloads:
```
' OR 1=1 UNION SELECT * FROM teams --
' OR 1=1 UNION SELECT * FROM users --
```

Although we found no useful information from the `teams` table, the `users` table showed us what seemed to be the second half of the flag.

![Screenshot 2022-03-28 032506](https://user-images.githubusercontent.com/101789488/160297382-6c0c4940-b51e-440f-8d10-9f79fead10a6.png)

After joining the two parts together, we get:
```
LNC2022{T0m_IsReDY_De4d}
```

To be very honest, after realising that this was not the flag, we knew that the flag had to be something along the lines of "Tom is already dead". With that assumption, we tried many possibilities until we were able to identify the correct flag. However, that was not the intended way to solve this challenge.

From the challenge description, we can see that there were hints of an `old browser`. We also see the following if we inspect the html source code.

![Screenshot 2022-03-28 035337](https://user-images.githubusercontent.com/101789488/160298347-4975a0b0-a9c7-471b-828b-de53f83ad411.png)

This means that the method to get the second half of the flag would be to change your user agent to `Mozilla/4.0 (compatible; MSIE 5.5; Windows 95; BCD2000)`. This can be done by using [curl](https://developer.ibm.com/articles/what-is-curl-command/#:~:text=cURL%2C%20which%20stands%20for%20client,data%20you%20want%20to%20send.). Steps are as shown:
```
Inspect the webpage > Navigate to Network > Copy as cURL(cmd) > Paste in linux terminal > Modify the user agent > Send
```
![Screenshot 2022-03-28 034151](https://user-images.githubusercontent.com/101789488/160297996-93e8162a-5620-46e2-9ec9-7d534c000ce7.png)

```
Second half: _AlLReDY_De4d}
```

So combining both parts of the flag, we get:
```
LNC2022{T0m_Is_AlLReDY_De4d}
```