# Go-Deer

**Go-Deer** is an independently developed, self-hosted personal cloud and file management system, built primarily in Go.

The goal is: build a private cloud infrastructure for personal use, providing file storage, management, and access through a custom-built backend and web interface.

---

## Project Goals

- Build a lightweight HTTP server in Go.
- Implement a custom backend and REST API.
- Provide secure file upload and download functionality.
- Manage files and directories through a web interface.
- Implement user authentication and access control.
- Store file metadata using SQLite.
- Support multiple users and shared storage.
- Eventually support automatic photo uploads from mobile devices.
- Run independently on dedicated Linux server hardware.

## Planned Architecture

| Component | Technology |
|---|---|
| Backend | Go |
| HTTP Server | Go / POSIX sockets |
| Database | SQLite |
| Frontend | Not decided.. most likely Angular |
| Storage | Local filesystem |

## Planned Features

### File Management
- Upload files
- Download files
- Create directories
- Rename and delete files
- Browse directory structures

### User Management
- User authentication
- Individual user storage
- Shared directories
- Access permissions

### Photo Management
- Photo uploads
- Image previews and thumbnails
- Albums and organization
- Mobile photo synchronization

### Server Management
- Storage monitoring
- Server health information
- File system statistics
- User and storage management

## Development Roadmap

- [ ] Initialize project structure
- [ ] Implement basic Go HTTP server
- [ ] Implement HTTP request parsing
- [ ] Create API routing
- [ ] Implement file upload
- [ ] Implement file download
- [ ] Integrate SQLite
- [ ] Implement authentication
- [ ] Build web interface
- [ ] Implement file management
- [ ] Add multi-user support
- [ ] Implement photo management
- [ ] Prepare dedicated server deployment
- [ ] Implement backup and storage management

## Security

Security is a fundamental part of the project.

Planned considerations include:

- Input validation
- Safe filesystem operations
- Path traversal protection
- Secure password hashing
- User permission isolation
- TLS for network communication
- Secure file storage and backup strategies
