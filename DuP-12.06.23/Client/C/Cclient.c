#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <winsock2.h>

#define BUFFER_SIZE 1024

int main() {
    WSADATA wsaData;
    SOCKET clientSocket;
    struct sockaddr_in serverAddress;
    char serverIP[] = "127.0.0.1"; // Server IP address
    int serverPort = 1401; // Server port number
    char message[BUFFER_SIZE];
    int bytesRead;

    // Initialize Winsock
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        printf("WSAStartup failed.\n");
        return 1;
    }

    // Create a socket for the client
    if ((clientSocket = socket(AF_INET, SOCK_STREAM, 0)) == INVALID_SOCKET) {
        printf("Failed to create socket.\n");
        return 1;
    }

    // Set up server address
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_addr.s_addr = inet_addr(serverIP);
    serverAddress.sin_port = htons(serverPort);

    // Connect to the server
    if (connect(clientSocket, (struct sockaddr*)&serverAddress, sizeof(serverAddress)) < 0) {
        printf("[ERR] Connect: %s:%d.\n", serverIP, serverPort);
        closesocket(clientSocket);
        return 1;
    }

    printf("[INFO] Connected: %s:%d\n", serverIP, serverPort);

    while (1) {
        // Read user input
        printf("DuP$: ");
        fgets(message, BUFFER_SIZE, stdin);

        // Remove trailing newline character
        message[strcspn(message, "\n")] = '\0';

        // Send the message to the server
        send(clientSocket, message, strlen(message), 0);

        if (strcmp(message, "exit") == 0) {
            break;
        }

        // Receive the server's response
        bytesRead = recv(clientSocket, message, BUFFER_SIZE, 0);
        if (bytesRead > 0) {
            // Null-terminate the received message
            message[bytesRead] = '\0';

            printf("DuP#: %s\n", message);
        } else if (bytesRead == 0) {
            printf("[ERR] Server closed the connection.\n");
            break;
        } else {
            printf("[ERR] Receive data.\n");
            break;
        }
    }

    // Cleanup Winsock
    closesocket(clientSocket);
    WSACleanup();

    return 0;
}
