# SimplePasswordCache

SPwC (Simple Password Cache) is a tool written in Go to have all the benefits of <a href = "https://linux.die.net/man/1/pass"> pass(1)</a> but written in a compiled
language.

# Installing
Download the pre-release binary and move it to your bin folder or run `export PATH=$PATH:</path/to/file>` then `source ~/.bashrc or ~/.zshrc`
<h3>Build from source</h3>

```
git clone git@gitlab.com:grumbulon/spwc.git

cd spwc

go build
```

With the above steps you will have a functional binary named SimplePasswordCache

<h3>Recommendation</h3>
Since the name SimplePasswordCache is quite long you should alias it to something shorter like spwc or whatever you want.

# Usage 
Spwc has the following complete commands

* init
* list
* insert
* version

Spwc is adding the following commands in a future release

* show
* insert with passphrase
* delete

First create an armor file if you do not have one from a public key

`gpg2 -a --export <PubKey ID> > /var/tmp/pubkey.asc`

Next, you will want to initialize your password cache with the above created armored file

`spwc init /var/tmp/pubkey.asc`

Your password cache has been created and the program will store them in ~/.passwordcache for your convenience. 

To view existing password you will run
`spwc list`

This will output a your keys in the .passwordcache folder using tree.

Currently since the show command is not functional (Golang makes encryption easy but decryption a hassle) you will need to use gpg2 to decrypt your files. Simply run

`gpg2 -d ~/.passwordcache/<gpgFileName>`
and it will output your password in plain text. 
