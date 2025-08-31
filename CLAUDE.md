# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Tennis Match Logger is an MVP mobile app for logging tennis matches using **audio narration + manual button taps**. The core concept is to reduce manual data entry during live matches by combining continuous voice commentary with simple touch inputs, then using AI to suggest detailed match data from the audio.

## Architecture

## Code repo reference
Refer to the following code repo /Users/chunyaozou/Documents/c/Go_flutter/Chapter13/13-11/ about how the project is archtectured, APIs design, data model layer, business logic layer, API layer, middleware layer etc.

This is a **Flutter frontend + Go backend** project with a local-first approach:

- **Frontend**: Flutter app (`/frontend/`) with cross-platform mobile support
- **Backend**: Go HTTP server (`/backend/`) with REST API and SQLite database  
- **Design System**: Comprehensive UX wireframes in `/frontend/design/`

The MVP focuses on 3 core screens:
1. **Match Setup**: Player names + match format selection
2. **Live Match**: Audio recording + winner buttons + AI suggestions  
3. **Match Summary**: Final score + basic stats + audio playback



## Key Technical Decisions

- **Local-first storage**: SQLite for match data, local file storage for audio
- **No cloud dependencies**: Everything works offline 
- **Audio-centric UX**: Continuous recording with speech-to-text processing
- **Simple AI**: Basic NLP for tennis terms (ace, winner, error) with confidence scoring
- **One-handed operation**: Large touch targets optimized for use during matches

## Common Commands

### Frontend (Flutter)
```bash
cd frontend
flutter pub get          # Install dependencies
flutter run              # Run on connected device/emulator
flutter build apk        # Build Android APK
flutter build ios        # Build iOS (requires Mac/Xcode)
flutter test             # Run tests
flutter analyze          # Static analysis
```

### Backend (Go)
```bash
cd backend
go mod tidy              # Clean up dependencies  
go run main.go           # Start development server on :8080
go build -o tennis-api main.go  # Build binary
go test ./...            # Run tests
```

## Data Models

### Database Schema (SQLite)
- **matches**: Core match data (players, format, winner, score, audio path)
- **points**: Individual point records (winner, type, transcript, confidence)

### Planned Flutter Structure
```
lib/
├── main.dart              # App entry point
├── screens/               # Main UI screens
│   ├── match_setup.dart   # Player/format selection
│   ├── live_match.dart    # Core logging interface  
│   └── match_summary.dart # Post-match review
├── models/                # Data models
│   ├── match.dart        # Match data structure
│   ├── point.dart        # Point data structure
│   └── player.dart       # Player data structure  
├── services/              # Business logic
│   ├── audio_service.dart # Recording/speech-to-text
│   ├── database_service.dart # SQLite operations
│   └── speech_service.dart   # NLP processing
└── widgets/               # Reusable UI components
    ├── score_display.dart # Score visualization
    ├── audio_indicator.dart # Recording status
    └── point_button.dart  # Winner/type buttons
```

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

## Development Phases

1. **Week 1**: Basic Flutter app structure and navigation
2. **Week 2**: Audio recording and speech-to-text integration  
3. **Week 3**: Live match UI and manual button logging
4. **Week 4**: AI integration and suggestion system
5. **Week 5**: Match summary, export, and testing

## Design Resources

Comprehensive UX designs are available in `/frontend/design/`:
- **MVP wireframes**: `/frontend/design/mvp/` (simplified scope)
- **Full wireframes**: `/frontend/design/mockups/` (future features)
- **Interactive HTML prototypes**: Open `.html` files in browser

## Performance Targets

- App startup: <3 seconds
- Point logging response: <500ms  
- Audio processing: <2 seconds per point
- Storage per match: <50MB
- Battery life: Acceptable for 2-hour matches
- One-handed operation with tennis gloves

## Backend API Structure

Current Go backend has placeholder routes:
- `GET /v1/health` - Health check
- `GET /v1/matches` - List matches  
- `POST /v1/matches` - Create match
- `GET /v1/matches/:id` - Get specific match
- `PUT /v1/matches/:id` - Update match
- Player routes also planned

Uses httprouter + alice middleware for logging and panic recovery.

