# Spelling Bee
A version of the spelling bee game based on a client-server model.

## Prerequisites
Before running, ensure you have the latest version of Go (1.25.3).

## Installation
1. **Clone the Repository**
```bash
git clone https://github.com/lewcol/spelling-bee-game
cd spelling-bee-game
```

2. **Install Go Dependencies**
```bash
go mod tidy 
```
3. **Add Wordlist**
The game requires you to provide a list of valid words.
Create the wordlist as a JSON file of the following format:
```json
{
  "word": 1,
  "wordtwo": 1
}
```
Name this file "words_dictionary.json" and place it in the server/wordlists directory.
```bash
mv words_dictionary.json server/wordlists
```

## Running the Application
This repository comprises 3 executable applications with different purposes.
Currently, each application only works when executed within its own directory.
This will be fixed in a later version.

"generate_pangrams/generate_pangrams.go" creates the pangrams.json file of starting words from the wordlist in words_dictionary.json
```bash
go run generate_pangrams.go
```
This application must be run once to generate the wordlists required for the game.

"server/server.go" runs the server for the Spelling Bee game.
```bash
go run server.go
```

"client/client.go" runs the server for the Spelling Bee game.
```bash
go run client.go
```
A running server instance must exist for clients to connect.

## Design Patterns Used
### Singleton
Used for structs which only should to be instantiated once by the application, e.g. the Manager game manager, or the dictionary lookup struct.
Functions can call the package's get function for these structs to instantiate them and afterwards retrieve the existing instance.

The manager struct uses a mutex to enforce a lock on its resources while its methods run, ensuring clients attempting to access its methods simultaneously
do not affect the consistency of other clients. (e.g. Two clients create session at once, both are created before the id counter
is incremented leading to one client overwriting another client's session).

Implementing this struct as a singleton ensures clients cannot bypass the mutex lock of the manager by calling the method in
another instance.

The dictionary struct holds the lookup maps of words and pangrams in its memory.
These objects are require a relatively large amount of memory compared to other objects in the application.
Implementing this struct as a singleton ensures only a single copy of these large objects exists in memory at a time.

### Factory
Used by open_json_as_map.json. The function WordMapFactory creates one of two types of map depending
on its argument.

The two wordlist files each have a different format which are therefore read and stored as maps with different key and value types.
WordMapFactory allows one function to be used to create the different types of map. This encapsulates the map creation logic,
and allows new types to be added without changing existing code elsewhere.

### Proxy
Used by the client. ClientProxy acts as a proxy between the game client and the ManagerClient which communicates to the server.
It implements the expected interface of ManagerClient and controls its creation. 
It creates ManagerClient during its own creation, and extends its methods while still calling and returning
from the real method.

The proxy may provide a slight optimisation by delaying creation of the real ManagerClient until it is needed.
The proxy is able to control access to the server, enforcing rate limiting to prevent clients
overwhelming the server with rapid calls. It may also intercept requests and perform input validation on their fields
to prevent unexpected inputs from the user crashing the server.