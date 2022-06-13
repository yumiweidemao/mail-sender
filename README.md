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

`./mail example1@example.com example2@example.com`

will send the email to both addresses included.
</li>
</ul>