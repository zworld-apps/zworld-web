+++
title = "Improved UI system"
date = 2019-11-04
lastmod = 2020-01-31T23:40:58+00:00
categories = ["version", "0.7.7"]
draft = false
description = "Creating UI components now is easier. Support for themes."
author = "Zebra"
type = "post"
+++

UI system needed to be rewritten. Everytime a new UI component had to be
added, there was a long process of tweaking variables, rewritting things, adding
other attributes to main components...

In order to solve all of this problems, I came to the simple but efficient
solution of creating a code generator according to [Raylib GUI layout](https://raylibtech.itch.io/rguilayout) file
formats.

Also rewritting the entire UI system made creating new themes a possibility.
Themes are stored in **_assets/_** folder.
