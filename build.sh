

# Using Docker
#!/usr/bin/env/ bash
#docker build . -t Chatserver --no-cache

# Below are the manual steps for running the chat server

# Chatserver
Chat Server using go language

# STEPS for build the project:

$ go build .
$ ./main.exe
$ ./client.exe
OR
 $ ./pkg.exe

# CURL POST REQUEST:

$ curl --header "Content-Type: application/json" --request POST --data '{"from":"Amit Gupta","content":"Hi Wiz,Welcome in My Chat ROOM. Hope you are doing good !!  Yesterday we had a good discussion really it was nice connect with you. Happy Weekend !!"}' http://localhost:8000/

# CURL GET REQUEST:
$ curl --header "Content-Type: application/json" --request GET http://localhost:8000/


