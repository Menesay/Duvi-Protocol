import socket

def main():
    # Connect to TCP server
    server_address = ('localhost', 1401)
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect(server_address)
        print("Connected to server:", server_address)
    except ConnectionRefusedError:
        print("Failed to connect to the server.")
        return

    while True:
        # Read user input
        message = input("Enter a message: ")

        try:
            # Send message to server
            #sock.sendall(message)
            sock.sendall(message.encode())
            
            print(message)

            if message.lower() == "exit":
                break

            # Receive server response
            response = sock.recv(1024).decode()
            print("Server response:", response)
        except ConnectionError:
            print("Failed to send/receive data. The connection may have been closed.")
            break

    # Close the socket
    sock.close()

if __name__ == '__main__':
    main()
