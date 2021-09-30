+++
title = "Linux updater"
date = 2019-11-26
lastmod = 2020-01-31T23:40:58+00:00
categories = ["version", "0.7.8"]
draft = false
description = "Linux updater fixed and new keybindings."
author = "Zebra"
type = "post"
+++

Until now, when you tried to update the client in a Linux machine, it wouldn't
succeed. That was because accessing API's /latestrelease page returned a Windows
version.

After updating the API and the client, Linux releases could be finally
downloaded by the updater.

In this hotfix, some keybindings were added:

-   Toggle display entity names <F2>
-   Zoom without keypad <+/->
