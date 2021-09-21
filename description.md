## What is it

 - Test application

### Task

 - read a dictionary from a file 
 - print the longest word from dictionary values

### Dictionary preparation

 - generate a dictionary with random number of records with random keys and values
 - save it to txt, yaml, json files

### Layers

 - fs,storage: check file status, existense, permissions
 - netcom: sockets, protocols, connections, status, recovery
 - io: open/close file/socket/connection, search content, read/write content 
 - config: read command line options, env variables, config files
 - data: define structures, methods
 - transform: data conversion, encoding/decoding, validation
 - compute: data analysis, aggregation, filtering, searching, combining, comparison, compliance 

### Implementation priority

- fs,storage: 
  - local file system
- netcom: 
  - none
- io: 
  - open/close/read file
- config: 
  1. default options
  2. read command line options
- data: 
  - structures:
    - dictionary
    - strings
  - methods: 
    - search dictionary for a sequence of letter
      1. predefined [a-zA-Z]
      2. given in arguments
- transform:
  1. Marshalling from plaintext/yaml/json to a dictionary
- compute: 
  - search dictionary for the longest letter sequence
  - output the result



