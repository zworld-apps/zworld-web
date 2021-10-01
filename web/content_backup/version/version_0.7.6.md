+++
title = "Local item data"
date = 2019-09-11
lastmod = 2020-01-31T23:40:58+00:00
categories = ["version", "0.7.6"]
draft = false
description = "Item info stored locally. No need to update the whole game to modify items."
author = "Zebra"
type = "post"
+++

In the past versions, it was difficult to modify item properties and there was a
need for recompiling. I decided to store item properties in the game database.
If there was a need to update items, it will be as simple as exporting the
database objects to a _JSON_ file.

In further versions, I expect to have separate updating methods. One for the
base game and another for the assets (items' properties included).
