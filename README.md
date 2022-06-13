# Mail Sender
<ul>Written in Go 1.18.2</ul>
## How to Use
Use `go build mail.go` to build the source code. The user can configure the SMTP server url & port as well as the sender email address in *mail.go*. The program reads from './text.txt' and uses the first line as the subject of the email. The rest is used as text.

To specify recipients, include them in command line arguments. For example,

`./mail example1@example.com example2@example.com`

will send the email to both addresses included.

