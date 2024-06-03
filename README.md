# Duvi Server

## Overview
Duvi Server is a simple TCP server written in Go. It handles multiple client connections, processes various commands, and performs basic file system operations such as listing directories and reading files.

## Features
- **Client Connection Handling:** Manages multiple client connections concurrently.
- **Command Processing:** Supports a set of predefined commands from clients.
  - `exit`: Disconnects the client.
  - `duvi`: Responds with "basil".
  - `basil`: Responds with "duvi".
  - `.`: Lists the contents of the current directory.
  - `..`: Changes to the parent directory.
  - `... <dir>`: Changes to a specified subdirectory.
  - `cat <file>`: Reads and returns the content of a specified file.

## Usage
1. **Clone the Repository:**
   ```sh
   git clone https://github.com/yourusername/duvi-server.git
   cd duvi-server
   ```

2. **Run the Server:**
   ```sh
   go run main.go
   ```

3. **Specify the Port:**
   When prompted, enter the port number you want the server to listen on.

## Code Structure
- **`banner()`**: Displays the server banner.
- **`listDir(path string) (arr []string, err error)`**: Lists files and directories in the specified path.
- **`handle(c net.Conn)`**: Handles incoming client connections and processes commands.
- **`serve()`**: Initiates the server, listens for incoming connections, and starts a new goroutine for each client.
- **`main()`**: Entry point of the application. Displays the banner and starts the server.

## Example Commands
- **Listing Directory Contents:**
  ```
  > .
  C:\path\to\directory#
  D: subdirectory1
  F: file1.txt
  F: file2.txt
  ```
- **Changing Directory:**
  ```
  > ... subdirectory1
  C:\path\to\directory\subdirectory1#
  ```
- **Reading a File:**
  ```
  > cat file1.txt
  File content goes here...
  ```
  
---

With these instructions, you should be able to set up and run the Duvi Server, as well as understand its basic functionality and structure. Enjoy exploring the code and extending its capabilities!
