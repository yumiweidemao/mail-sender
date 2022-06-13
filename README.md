# Mail Sender

<ul>
    <li>Written in Go 1.18.2</li>
</ul>

## How to Use

<ul>
<li>
Use 

`go build mail.go`

to build the source code.
</li>
<li>
The user can configure the SMTP server URL & port as well as the sender email address in 'mail.go'.
</li>
<li>
The program reads from './text.txt' and uses the first line as the subject of the email. The rest is used as text.
</li>
<li>
To specify recipients, include them in command line arguments. For example,
=======
<ul>Written in Go 1.18.2</ul>

## How to Use

<ul>Use `go build mail.go` to build the source code</ul>
<ul>The user can configure the SMTP server url & port as well as the sender email address in *mail.go*</ul>
<ul>The program reads from './text.txt' and uses the first line as the subject of the email. The rest is used as text.</ul>
>>>>>>> 5f54489b617c1a37c749d2683a13f8301b6da011

<ul>To specify recipients, include them in command line arguments. For example,

<<<<<<< HEAD
will send the email to both addresses included.
</li>
</ul>
=======
`./mail example1@example.com example2@example.com`

will send the email to both addresses included.</ul>
>>>>>>> 5f54489b617c1a37c749d2683a13f8301b6da011
