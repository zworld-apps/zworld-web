---
title: "Establishing connection" 
description: ""
image: ""
date: 2021-09-06T09:58:00+07:00
lastmod: 2021-10-02T10:14:00+07:00
author: "Zebra"
tags: []
categories: ["articles"]
draft: false
---

# Summary
1-2 sentence summary of what this document is about and why we should work on it. 

As the communication between client and server is being done through TCP and UDP, we need a way to relate UDP packets to the TCP connection (as they are linked to the same user session).

# Background
What is the motivation for these changes? What problems will this solve? Include graphs, metrics, etc. if relevant. 

This is now done by letting the client tell the server which port it will be using. This can cause several problems:

- Port spoofing: someone with bad intentions could impersonate the session of the user and send malicious packets.
- NAT and proxied connections failure: this problem appeared when adding Docker support for the server. The client tells the server that he will be using some UDP port and the server tries to reach that address at that port, but Docker can map the ports, so the connection will be unreachable to that port (as the server needs to use the mapped port). 

# Goals 
What are the outcomes that will result from these changes? How will we evaluate success for the proposed changes? 

- More reliable packet authority as the packets would not be spoofed or MitM.
- Possibility of containerization of both ends, but specially the server.

# Proposed Solution
Describe the solution to the problems outlined above. Include enough detail to allow for productive discussion and comments from readers.

Use  encryption to ensure a secure communication between client an server. Use an authentication token (not encrypted) to identify the client author of the packet, and ensure its authenticity by decrypting the packet and parsing a known header.

# Things to consider 🤔
- **This system would allow two connections with the same IP and port. Would this matter? **Previously, unique connection by IP and port was achieving by storing the address as key of the connections pool. This is irrelevant as what we want to disallow is login in twice. There is a possible flaw that two mapped ports could be the same, but that wouldn't make sense (and it is a network problem, not game).
