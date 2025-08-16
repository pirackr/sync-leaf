# Task List

This file tracks the high-level tasks required to build the sync application.

## Initial Plan
- [x] Define the overall project architecture for a folder synchronization app using WebRTC.
- [ ] Implement a basic, cross-platform file watcher with a pluggable interface for future enhancements.
- [ ] Establish a WebRTC data channel to sync files between peers.
- [ ] Add basic authentication for connecting peers.
- [ ] Provide a configuration mechanism to specify folders to sync.
- [ ] Integrate `go vet` and `staticcheck` into the CI pipeline.
- [ ] Set up test coverage reporting and add a coverage badge to the README.
- [ ] Document setup, usage, and development workflows.

> Tests should accompany each step and be written alongside features.

## Architecture Breakdown

### Clients (Peers)
- [ ] File watcher and metadata engine to monitor changes and store hashes, timestamps, versions, sizes, and permissions.
- [ ] Sync manager to orchestrate synchronization, concurrency, conflict resolution, and version history.
- [ ] WebRTC layer with STUN/TURN support and separate control and data channels.
- [ ] Chunking, compression, and encryption pipeline for large file transfers.
- [ ] UI or CLI for selecting folders, viewing sync status, and resolving conflicts.

### Signaling Service
- [ ] Lightweight REST/WebSocket server for peer discovery, SDP/ICE exchange, authentication, and connection metadata.
- [ ] Optional relay or TURN credential support for peers that cannot connect directly.

### Synchronization Flow
- [ ] Client initialization registers with the signaling service and builds a local manifest.
- [ ] Exchange manifests, negotiate folder state, and apply conflict resolution rules.
- [ ] Transfer required file chunks over the data channel with checksum verification and retries.
- [ ] Update local metadata stores after transfers complete.
- [ ] Maintain continuous sync via file watchers for incremental updates.

### Persistence & Recovery
- [ ] Persist metadata in a local store to enable restart recovery.
- [ ] Optional versioning or snapshots to allow rollback.

### Security
- [ ] Ensure WebRTC DTLS/SRTP encryption for all connections.
- [ ] Authenticate peers using tokens or public-key infrastructure.
- [ ] Provide optional end-to-end encryption for file payloads.

### Scalability
- [ ] Support multi-peer synchronization via mesh, spanning tree, or relay strategies.
- [ ] Implement selective sync to limit bandwidth and storage usage.
- [ ] Design for extensibility to plug in different storage backends or transports.
