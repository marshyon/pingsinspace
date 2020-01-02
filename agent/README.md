# Pings in Space Client

The client may be run within a docker container - see the Dockerfile in this directory so Dockter needs
to be installed / enabled on your platform in order to run this.

### Manual Testing

to run a test manually open a termainal ( standard bash shell in Linux / Mac or Powershell in Linux or
Powershell in Windows ) and change to this directory

to build the image :

```
docker build --rm  -t pingsinspaceapi:latest  .
```

run the container in an interactive shell :

```
docker run -it --name agent --rm pingsinspaceapi:latest
```

then run the agent manually with

```
cd /root/agent
./agent
```

you should see something like the following

```
f825141d-c1dd-4661-8fd2-9b6f738f1eb7	0	command completed with errors exit status 2 HTTP CRITICAL: HTTP/1.1 302 Found - string 'yahoo' not found on 'https://www.yahoo.com:443/' - 770 bytes in 0.099 second response time |time=0.098905s;;;0.000000;10.000000 size=770B;;;0	0	2	2020-01-02 11:41:50+002020-01-02 11:41:50+00	0
448e42cb-54e4-4af8-83c5-31ae0694fba9	2	HTTP OK: HTTP/1.1 200 OK - 12776 bytes in 0.114 second response time |time=0.113987s;;;0.000000;10.000000 size=12776B;;;0	1	0	2020-01-02 11:41:50+00	2020-01-02 11:41:50+00	2
```
