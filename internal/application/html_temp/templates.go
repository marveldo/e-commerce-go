package html_temp

import (
	"html/template"
	"strings"

	payload "github.com/marveldo/gogin/internal/application/payloads"
)

func GetEmailHtml(d *payload.EmailPayload) *strings.Builder {
     mail := `
		<h1 style="color: #336699;">Hello {{.Username}}!</h1>
		<p>Welcome to our service.</p>
		<div style="display:flex; align_items:center; justify_items:center; width:400px">
		  <p style="font-size:50px">Welcome To our Book Store</p>
		</div>
	`
	var htmlBody strings.Builder
	tmpl := template.Must(template.New("email").Parse(mail))
	tmpl.Execute(&htmlBody , d)
    return &htmlBody
}