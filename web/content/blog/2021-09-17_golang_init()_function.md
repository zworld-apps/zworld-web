---
title: "Golang init() function" 
description: ""
image: ""
date: 2021-09-17T11:05:00+07:00
lastmod: 2021-10-02T10:14:00+07:00
author: "Zebra"
tags: []
categories: ["blog"]
draft: false
---

# Summary
1-2 sentence summary of what this document is about and why we should work on it. 

This post is about the init function in golang and how it can be used for package initialization, avoiding filling the main function with initialization for each packet.

# Background
What is the motivation for these changes? What problems will this solve? 

The necessity to initialize a packet before main function comes when trying to reduce the number of queries to zworld's items database.



This database contains the information for all the items in the game, so it is easier to check, update and share the information t both the client and server.



When implementing chicken nest block's interaction, we had to:

- check if it was really that block
- change it to the empty version (without eggs)
- give the player the eggs collected

That means on every nest interaction, three queries will be performed.



It seems like it is just a few. But the current implementation of the find item from id name ("chicken_nest_eggs", ...) is O(n) in the number of items, checking id by id comparing strings, so probably worse. This means a very bad performance when the items database grows.

# Proposed Solution
Describe the solution to the problems outlined above and how to get to it.

We couldn't just create some globals, as the package initialization was being done in the main function, resulting in nil data when initializing the globals.



This is when Go's init function comes handy. This function allows you to run code the first time the package is initialized. The order of this initialization is in import order, following recursively imports first. And, after all, the main function will be executed.



By initializing the items database package on the init function, we can allow any package importing the items database package to initialize globals, as the database will be already initialized.



It is a very handy (and unique, as most languages don't have it) function and not so widely known.

