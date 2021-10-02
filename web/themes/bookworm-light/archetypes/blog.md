---
title: "{{ .Title }}" 
description: "{{ .Description }}"
image: "{{ .Banner }}"
date: {{ .CreationDate.Format "2006-01-02T15:04:05+07:00" }}
lastmod: {{ .LastModified.Format "2006-01-02T15:04:05+07:00" }}
author: "{{ .Author }}"
tags: [{{ range $i, $v := .Tags }}{{if (eq $i 0)}}"{{$v.Name}}"{{else}},"{{$v.Name}}"{{end}}{{end}}]
categories: [{{ range $i, $v := .Categories }}{{if (eq $i 0)}}"{{$v.Name}}"{{else}},"{{$v.Name}}"{{end}}{{end}}]
draft: false
---
