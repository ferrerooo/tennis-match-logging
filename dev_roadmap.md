# Tennis Match Logger - Development Roadmap

## Architecture Overview

Based on the reference architecture from Chapter 13-11, this project follows a **stateless frontend + stateful backend** approach:

- **Frontend (Flutter)**: Audio capture, real-time display of backend data, UI controls
- **Backend (Go)**: STT processing, tennis NLP, score calculation, AI analysis, data persistence
- **Communication**: WebSocket for real-time updates, REST API for match management

## Reference Architecture Patterns

From `/Users/chunyaozou/Documents/c/Go_flutter/Chapter13/13-11/`:

- **Clean separation**: `main.go` → `app.go` → `api.go` → models/handlers
- **Dependency injection**: Config → Models → API → Routes  
- **Middleware stack**: Alice middleware chain with logging, auth, rate limiting
- **Database layer**: PostgreSQL with sqlx, proper connection pooling
- **Structured config**: YAML-based configuration with environment support

## Implementation Phases

### **Phase 1: Core Backend Infrastructure (Week 1)**

```
Dependencies: None
└── Backend Core
    ├── 1. Database Schema (SQLite for MVP)
    │   ├── matches table
    │   ├── points table  
    │   ├── audio_segments table
    │   └── match_insights table
    ├── 2. Models Layer
    │   ├── match.go
    │   ├── point.go
    │   └── audio.go
    ├── 3. Config & App Structure
    │   ├── config.yaml (following reference pattern)
    │   ├── main.go
    │   └── app/app.go
    └── 4. Basic HTTP Routes
        ├── POST /v1/matches
        ├── GET /v1/matches/:id
        └── GET /v1/health
```

**Deliverables:**
- Working HTTP server with basic match CRUD
- Database schema and models
- Configuration system
- Middleware stack (logging, recovery, rate limiting)

### **Phase 2: Audio Processing Pipeline (Week 2)**

```
Dependencies: Phase 1 complete
└── Audio Engine
    ├── 1. Audio Upload Handler
    │   └── POST /v1/matches/:id/audio
    ├── 2. STT Integration
    │   ├── WhisperAPI/Google STT
    │   └── Audio chunking strategy
    ├── 3. Tennis NLP Engine
    │   ├── Term recognition (ace, winner, error)
    │   ├── Player name detection
    │   └── Score parsing logic
    └── 4. Background Processing
        ├── Queue system for audio
        └── Worker goroutines
```

**Deliverables:**
- Audio upload and storage
- STT processing pipeline
- Tennis-specific term recognition
- Background job processing
- Score calculation from transcript

### **Phase 3: Real-time WebSocket Layer (Week 3)**

```
Dependencies: Phase 1 & 2 complete  
└── Real-time Communication
    ├── 1. WebSocket Upgrade Handler
    │   └── GET /v1/matches/:id/live
    ├── 2. Message Broadcasting
    │   ├── Score updates
    │   ├── Transcript updates
    │   └── AI insights
    ├── 3. Connection Management
    │   ├── Client registry
    │   └── Graceful disconnection
    └── 4. State Synchronization
        ├── Current match state
        └── Historical data replay
```

**Deliverables:**
- WebSocket connection handling
- Real-time message broadcasting
- Client connection management
- Live match state synchronization

### **Phase 4: Flutter Frontend (Week 4)**

```
Dependencies: Phase 3 complete
└── Flutter App  
    ├── 1. Core Architecture
    │   ├── models/ (match, score, transcript)
    │   ├── services/ (websocket, audio)
    │   └── screens/ (setup, live, summary, list)
    ├── 2. WebSocket Client
    │   ├── Connection management
    │   ├── Message handling
    │   └── Automatic reconnection
    ├── 3. Audio Recording
    │   ├── Continuous recording
    │   ├── Audio streaming to backend  
    │   └── Permission handling
    └── 4. UI Implementation
        ├── Live match screen
        ├── Match summary
        └── Match list
```

**Deliverables:**
- Complete Flutter app structure
- WebSocket client implementation
- Audio recording and streaming
- All UI screens (setup, live, summary, list)
- Cross-platform builds (iOS/Android)

### **Phase 5: AI Analysis & Polish (Week 5)**

```
Dependencies: All previous phases
└── Advanced Features
    ├── 1. AI Insights Engine
    │   ├── Match momentum analysis
    │   ├── Performance recommendations
    │   └── Statistical summaries
    ├── 2. Export Functionality
    │   ├── PDF reports
    │   ├── CSV data export
    │   └── Audio playback
    └── 3. Error Handling & Edge Cases
        ├── Network failures
        ├── Audio processing errors
        └── Manual corrections
```

**Deliverables:**
- AI-powered match analysis
- Export functionality (PDF, CSV)
- Comprehensive error handling
- Manual correction capabilities
- Performance optimization

## Critical Dependencies

1. **Phase 1 → Phase 2**: Database models must exist before audio processing
2. **Phase 2 → Phase 3**: STT pipeline needed before real-time updates  
3. **Phase 3 → Phase 4**: WebSocket API must be stable before Flutter implementation
4. **Phase 4 → Phase 5**: Basic functionality required before AI enhancements

## Recommended Development Order

```
1. Start with backend core (database + basic REST API)
2. Add audio upload endpoint (test with Postman)
3. Integrate STT processing (verify audio→text pipeline)
4. Implement WebSocket broadcasting (test with browser)
5. Build Flutter UI screens (connect to WebSocket)  
6. Add AI analysis layer (enhance existing data)
```

## WebSocket API Specification

### Connection
```
GET /v1/matches/{matchId}/live
Upgrade: websocket
```

### Message Types (Backend → Frontend)

#### Score Update
```json
{
  "type": "score_update",
  "timestamp": "2025-01-15T10:30:00Z",
  "data": {
    "match_id": "uuid",
    "set_scores": [6, 4],
    "game_scores": [3, 2], 
    "point_score": "30-15",
    "serving_player": 1
  }
}
```

#### Transcript Update
```json
{
  "type": "transcript_update", 
  "timestamp": "2025-01-15T10:30:00Z",
  "data": {
    "text": "Ace down the T",
    "confidence": 94,
    "audio_timestamp": "2:15:30",
    "recognized_terms": ["ace"]
  }
}
```

#### AI Insight
```json
{
  "type": "ai_insight",
  "timestamp": "2025-01-15T10:30:00Z",
  "data": {
    "category": "performance",
    "message": "Player 1 serving exceptionally well this set",
    "confidence": 92,
    "statistics": {
      "first_serve_percentage": 85,
      "aces_this_set": 3
    }
  }
}
```

#### Match Status
```json
{
  "type": "match_status",
  "timestamp": "2025-01-15T10:30:00Z",
  "data": {
    "status": "in_progress" | "completed" | "paused",
    "duration": "1:45:30",
    "total_points": 142
  }
}
```

#### Match Completion
```json
{
  "type": "match_completed",
  "timestamp": "2025-01-15T10:30:00Z",
  "data": {
    "winner": "Player 1",
    "final_score": "6-4, 6-2",
    "duration": "1:45:30",
    "summary": {
      "total_winners": [24, 16],
      "unforced_errors": [12, 18],
      "aces": [8, 3]
    }
  }
}
```

### Message Types (Frontend → Backend)

#### Audio Chunk
```json
{
  "type": "audio_chunk",
  "data": {
    "audio_data": "base64_encoded_audio",
    "timestamp": "2:15:30",
    "format": "wav",
    "sample_rate": 44100
  }
}
```

#### Manual Correction
```json
{
  "type": "correction",
  "data": {
    "transcript_id": "uuid",
    "corrected_text": "Double fault, second serve",
    "score_correction": {
      "player1_score": "30",
      "player2_score": "30"
    }
  }
}
```

## Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: httprouter + alice middleware
- **Database**: SQLite (MVP) → PostgreSQL (production)
- **WebSocket**: gorilla/websocket
- **STT**: OpenAI Whisper API or Google Speech-to-Text
- **Config**: YAML-based configuration

### Frontend
- **Framework**: Flutter 3.x
- **State Management**: Provider or Riverpod
- **WebSocket**: web_socket_channel package
- **Audio**: flutter_sound for recording
- **HTTP**: dio for REST calls
- **UI**: Material Design 3

### Infrastructure
- **Local Development**: Docker Compose
- **Audio Storage**: Local filesystem (MVP)
- **Logging**: Structured JSON logging
- **Monitoring**: Built-in metrics endpoint

## MVP Scope

**IN SCOPE:**
- Audio recording during matches (2+ hours continuous)
- Manual winner buttons ([Player 1 Won] / [Player 2 Won])  
- Basic point types (Ace, Winner, Error)
- Speech-to-text with tennis term recognition
- AI suggestions with confidence scoring
- Real-time score tracking
- Match summary with basic statistics
- Local data storage and simple export

**OUT OF SCOPE** (Future features):
- Advanced shot analysis (forehand/backhand details)
- Cloud sync and social features  
- Complex analytics and trends
- Video integration
- Multi-language support

## Performance Targets

- App startup: <3 seconds
- Point logging response: <500ms  
- Audio processing: <2 seconds per point
- Storage per match: <50MB
- Battery life: Acceptable for 2-hour matches
- WebSocket latency: <100ms
- Concurrent matches: 10+ simultaneous

## Testing Strategy

### Backend Testing
- Unit tests for models and business logic
- Integration tests for API endpoints
- WebSocket connection testing
- Load testing with multiple clients

### Frontend Testing  
- Widget tests for UI components
- Integration tests for audio recording
- WebSocket reconnection testing
- Cross-platform compatibility testing

---

*Last Updated: January 2025*