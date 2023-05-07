package main

import (
	"os"
	"text/template"
)

func main() {
	nginxTemplate := `
server {
    listen   80;
    server_name localhost;
    location / {
        root   /usr/share/nginx/html;
        index  index.html;
    }
}
`

	tmpl, err := template.New("nginx.conf").Parse(nginxTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("nginx.conf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, nil)
	if err != nil {
		panic(err)
	}

	println("Generated nginx.conf successfully")
}
