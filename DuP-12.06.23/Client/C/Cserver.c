#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <winsock2.h>

#define MAX_CLIENTS 10
#define BUFFER_SIZE 1024

void handleClient(SOCKET clientSocket) {
    char buffer[BUFFER_SIZE];
    int bytesRead;

    // Receive and process client messages
    while ((bytesRead = recv(clientSocket, buffer, BUFFER_SIZE, 0)) > 0) {
        // Null-terminate the received message
        buffer[bytesRead] = '\0';

        printf("Received message from client: %s\n", buffer);

        // Send a response back to the client
        send(clientSocket, buffer, bytesRead, 0);
    }

    if (bytesRead == 0) {
        printf("Client disconnected.\n");
    } else {
        printf("Receive failed with error.\n");
    }

    // Close the client socket
    closesocket(clientSocket);
}

int main() {
    WSADATA wsaData;
    SOCKET serverSocket, clientSocket;
    struct sockaddr_in serverAddress, clientAddress;
    int clientAddressSize;

    // Initialize Winsock
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        printf("WSAStartup failed.\n");
        return 1;
    }

    // Create a socket for the server
    if ((serverSocket = socket(AF_INET, SOCK_STREAM, 0)) == INVALID_SOCKET) {
        printf("Failed to create socket.\n");
        return 1;
    }

    // Set up server address
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_addr.s_addr = INADDR_ANY;
    serverAddress.sin_port = htons(8080);

    // Bind the server socket
    if (bind(serverSocket, (struct sockaddr*)&serverAddress, sizeof(serverAddress)) == SOCKET_ERROR) {
        printf("Bind failed.\n");
        closesocket(serverSocket);
        return 1;
    }

    // Listen for incoming connections
    if (listen(serverSocket, MAX_CLIENTS) == SOCKET_ERROR) {
        printf("Listen failed.\n");
        closesocket(serverSocket);
        return 1;
    }

    printf("Server is listening on port 8080...\n");

    while (1) {
        // Accept a client connection
        clientAddressSize = sizeof(clientAddress);
        clientSocket = accept(serverSocket, (struct sockaddr*)&clientAddress, &clientAddressSize);
        if (clientSocket == INVALID_SOCKET) {
            printf("Accept failed.\n");
            closesocket(serverSocket);
            return 1;
        }

        printf("Client connected: %s:%d\n", inet_ntoa(clientAddress.sin_addr), ntohs(clientAddress.sin_port));

        // Handle the client in a separate function
        handleClient(clientSocket);
    }

    // Cleanup Winsock
    closesocket(serverSocket);
    WSACleanup();

    return 0;
}
