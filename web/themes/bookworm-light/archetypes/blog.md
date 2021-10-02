---
title: "{{ .Title }}" 
description: "{{ .Description }}"
image: "{{ .Banner }}"
date: {{ .CreationDate.Format "2006-01-02T15:04:05+07:00" }}
lastmod: {{ .LastModified.Format "2006-01-02T15:04:05+07:00" }}
author: "{{ .Author }}"
tags:{{ range .Tags }}
    - "{{ .Name }}"{{end}}
categories:{{ range .Categories }}
    - "{{ .Name }}"{{end}}
draft: false
---

