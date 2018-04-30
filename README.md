# How do I test it using Docker? #

Create Docker network to run nodes on it:

```bash
docker network create --driver=bridge p2p-fh-network
```

Build containers for your nodes

```bash
docker build -t node1 .
docker build -t node2 .
docker build -t node2 .
```

After the previous steps are completed, you may run your nodes.
First node:

```bash
docker run -i -p 1337:8000 --net=p2p-fh-network node1
```

Then another nodes that connects to first:
```bash
# you may use get_ip script for your convenience

docker run -i -p 8000:1338 --net=p2p-fh-network node2 /go/bin/app `./get_ip node1`:8000

# after previous step you can use either node1 or node 2 as node to connect

docker run -i -p 8000:1339 --net=p2p-fh-network node3 /go/bin/app `./get_ip node2`:8000
```


# TODO #
~~basic server~~  
~~basic peer-to-peer logic~~  
~~make server to handle multiple connections simultaneously~~  
sending files  
retrivieng files  
encryption  
storing files (logs)  