# Seek

---

## Challenge Description 
What do you seek

`http://c1.lagncrash.com:9008`

---

## Challenge Solution
Upon visiting the link given, we see the following page.

![Screenshot 2022-03-28 011147](https://user-images.githubusercontent.com/101789488/160292675-7e63830f-430e-4643-9540-088a776f0ed6.png)

Since clicking the submit button did not do anything, I decided to try [curl](https://developer.ibm.com/articles/what-is-curl-command/#:~:text=cURL%2C%20which%20stands%20for%20client,data%20you%20want%20to%20send.) on the URL:
```
curl http://c1.lagncrash.com:9008/
```

With that, we were able to get the html source from the server:
```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Belch</title>
</head>
<body>
  <h3>What do you seek</h3>
  <form method="POST">
    <input type="submit" placeholder="The flag"/>
  </form>
</body>
</html> 
```

To be very honest at this point, I was guessing that since the form allows the `POST` HTTP method, maybe I can retrieve some information by adding the argument to my command. The modified command I used is shown below:
```
curl -X POST http://c1.lagncrash.com:9008
```

The output I got from that was:
```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>Redirecting...</title>
<h1>Redirecting...</h1>
<p>You should be redirected automatically to target URL: <a href="/ahhhhhhhhhhhhhhhhhhhhhhhhhhhh">/ahhhhhhhhhhhhhhhhhhhhhhhhhhhh</a>. If not click the link. 
```

Not knowing what the output was, I decided to try adding the href string to the URL but nothing happened. So once again I edited the command and tried:
```
curl -X POST http://c1.lagncrash.com:9008/ahhhhhhhhhhhhhhhhhhhhhhhhhhhh
```

The output I got was:
```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>Redirecting...</title>
<h1>Redirecting...</h1>
<p>You should be redirected automatically to target URL: <a href="/">/</a>. If not click the link. 
```

I once again edited my command and used:
```
curl -X POST http://c1.lagncrash.com:9008/ahhhhhhhhhhhhhhhhhhhhhhhhhhhh/
```

This time, the output I got was slightly different:
```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>405 Method Not Allowed</title>
<h1>Method Not Allowed</h1>
<p>The method is not allowed for the requested URL.</p>
```

Since the output stated that the `POST` HTTP method was not allowed, I tried the other methods and found that the `GET` HTTP method was allowed. The output I got was:
```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
```

At this point, I was at a loss for what to do. After some consultation, I found out that my curl command was not displaying any HTTP headers. After modifying my command once more, I tried the following:
```
curl -vvv GET http://c1.lagncrash.com:9008/ahhhhhhhhhhhhhhhhhhhhhhhhhhhh/
```

This time the output I got stated that nothing was found:
```
*   Trying 34.143.164.90:9008...
* Connected to c1.lagncrash.com (34.143.164.90) port 9008 (#0)
> GET /ahhhhhhhhhhhhhhhhhhhhhhhhhhhh/ HTTP/1.1
> Host: c1.lagncrash.com:9008
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 404 NOT FOUND
< Content-Length: 232
< Content-Type: text/html; charset=utf-8
< Date: Sun, 27 Mar 2022 17:30:55 GMT
< Server: waitress
< 
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
* Connection #0 to host c1.lagncrash.com left intact
```

Once again, I tried the same command without the `/` at the back of the command:
```
curl -vvv GET http://c1.lagncrash.com:9008/ahhhhhhhhhhhhhhhhhhhhhhhhhhhh
```

Finally we see something interesting in the output:
```
*   Trying 34.143.164.90:9008...
* Connected to c1.lagncrash.com (34.143.164.90) port 9008 (#0)
> GET /ahhhhhhhhhhhhhhhhhhhhhhhhhhhh HTTP/1.1
> Host: c1.lagncrash.com:9008
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 302 FOUND
< Content-Length: 208
< Content-Type: text/html; charset=utf-8
< Date: Sun, 27 Mar 2022 17:31:05 GMT
< Location: http://c1.lagncrash.com:9008/
< Server: waitress
< Token: 4c4e43323032327b7333336b5f346e645f7330756768745f34667433727d
< 
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>Redirecting...</title>
<h1>Redirecting...</h1>
* Connection #0 to host c1.lagncrash.com left intact
<p>You should be redirected automatically to target URL: <a href="/">/</a>. If not click the link. 
```

If we take the Token value and convert it from [hexadecimal to text](http://www.unit-conversion.info/texttools/hexadecimal/), we get the flag.
```
LNC2022{s33k_4nd_s0ught_4ft3r}	
```