| **Concept**  | **Description**                                                                                                                      |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------|
| **RPC**      | A request-response protocol initiated by the client, which sends a request to a remote server to execute a procedure with parameters. The server responds, and the application continues. |
| **JSON-RPC** | An RPC protocol encoded in JSON.                                                                                                      |
| **gRPC**     | A cross-platform, open-source, high-performance RPC framework using HTTP/2 for transport and Protocol Buffers for interface description. It includes features like authentication, bidirectional streaming, flow control, blocking/nonblocking bindings, and cancellation/timeouts. |

***Differences***


| **Characteristic**       | **JSON-RPC**                                                                 | **gRPC**                                                                                          |
|--------------------------|------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| **Serialization Format** | JSON (human-readable, text-based)                                            | Protocol Buffers (protobuf, compact binary format)                                                |
| **Transport Protocol**   | HTTP/HTTPS (web-based applications)                                          | Multiple protocols including HTTP/2 (supports multiplexing and server push, efficient for real-time) |
| **Service Definition**   | No standardized method definition; requires manual documentation and agreement | Protocol Buffers (language-agnostic, allows automatic code generation)                            |
| **Streaming Support**    | No native support for streaming; relies on multiple requests/responses       | Built-in support for bidirectional and server-side streaming, reducing latency and overhead       |
