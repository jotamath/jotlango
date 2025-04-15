# WebSocket Library

The `websocket` library provides WebSocket server and client functionality for real-time bidirectional communication.

## Import

```jt
import "websocket"
```

## Classes

### Connection

```jt
class Connection {
    prop string Id
    prop map<string, string> Headers

    fn send(message string) : void
    fn close() : void
}
```

Represents a WebSocket connection.

### Server

```jt
class Server {
    prop int Port
    prop string Path
    prop fn(Connection) onConnect
    prop fn(Connection, string) onMessage
    prop fn(Connection) onDisconnect

    fn listen() : void
}
```

WebSocket server implementation.

## Examples

### Chat Server

```jt
import "websocket"
import "io"

class ChatRoom {
    prop map<string, websocket.Connection> Clients
    prop list<string> MessageHistory

    fn broadcast(message string) : void {
        MessageHistory.add(message)
        for id, client in Clients {
            client.send(message)
        }
    }

    fn addClient(conn websocket.Connection) : void {
        Clients[conn.Id] = conn
        broadcast("User {conn.Id} joined the chat")
    }

    fn removeClient(conn websocket.Connection) : void {
        Clients.remove(conn.Id)
        broadcast("User {conn.Id} left the chat")
    }
}

// Create chat room
chatRoom = new ChatRoom {
    Clients = {},
    MessageHistory = []
}

// Create WebSocket server
server = new websocket.Server {
    Port = 8080,
    Path = "/chat",

    fn onConnect(conn websocket.Connection) : void {
        chatRoom.addClient(conn)
        
        // Send message history
        for message in chatRoom.MessageHistory {
            conn.send(message)
        }
    },

    fn onMessage(conn websocket.Connection, message string) : void {
        chatRoom.broadcast("User {conn.Id}: {message}")
    },

    fn onDisconnect(conn websocket.Connection) : void {
        chatRoom.removeClient(conn)
    }
}

// Start server
io.println("WebSocket Chat Server started on ws://localhost:8080/chat")
call server.listen()
```

### Client Example (JavaScript)

```javascript
// Connect to WebSocket server
const ws = new WebSocket('ws://localhost:8080/chat');

// Handle connection open
ws.onopen = () => {
    console.log('Connected to chat server');
};

// Handle incoming messages
ws.onmessage = (event) => {
    console.log('Received:', event.data);
};

// Handle connection close
ws.onclose = () => {
    console.log('Disconnected from chat server');
};

// Send message
function sendMessage(message) {
    ws.send(message);
}
```

## Best Practices

1. Handle connection errors gracefully
2. Implement reconnection logic in clients
3. Clean up resources when connections close
4. Use appropriate message formats (JSON, text, etc.)
5. Implement heartbeat mechanism for long-lived connections
6. Handle connection timeouts
7. Validate messages before broadcasting
8. Implement proper error handling
9. Consider implementing message queuing for reliability
10. Document your WebSocket API 