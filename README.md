# tpb

The Pirate Bay CLI

## Usage

To search for torrents simply run:

```
tpb <search term>
```

There is no need to escape the spaces between the words in the search term.

After that, you will get a list of responses from The Pirate Bay API.
You can get more information (including the magnet link) about a response using its ID:

```
tpb <ID>
```

## Installing

Clone the repository

```
git clone github.com/Daniel-Zimmer/tpb
```

Make sure you have **go** and **go-dep** installed

cd into the repository and run dep init

```
cd tpb/src
dep init
```

Build

```
go build
```

Add executable to PATH. One easy way is to copy the executable to /bin

```
sudo cp src /bin/tpb
```
