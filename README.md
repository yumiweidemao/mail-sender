# Mail Sender
<ul>Written in Go 1.18.2</ul>

## How to Use

<ul>Use `go build mail.go` to build the source code</ul>
<ul>The user can configure the SMTP server url & port as well as the sender email address in *mail.go*</ul>
<ul>The program reads from './text.txt' and uses the first line as the subject of the email. The rest is used as text.</ul>

<ul>To specify recipients, include them in command line arguments. For example,

`./mail example1@example.com example2@example.com`

will send the email to both addresses included.</ul>
