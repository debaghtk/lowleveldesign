# Liskov Substitution Principle Problem Statement

## Background
A media player application needs to handle different types of media files (audio and video). The application should be able to play these files and display their metadata.

## Problem
The current implementation violates LSP:

```go
type MediaFile interface {
    Play()
    GetDuration() int    // in seconds
    GetResolution() Resolution // This causes LSP violation
}

type Resolution struct {
    Width  int
    Height int
}

type VideoFile struct {
    filename   string
    duration   int
    resolution Resolution
}

func (v *VideoFile) Play() {
    fmt.Printf("Playing video: %s\n", v.filename)
}

func (v *VideoFile) GetDuration() int {
    return v.duration
}

func (v *VideoFile) GetResolution() Resolution {
    return v.resolution
}

type AudioFile struct {
    filename string
    duration int
}

func (a *AudioFile) Play() {
    fmt.Printf("Playing audio: %s\n", a.filename)
}

func (a *AudioFile) GetDuration() int {
    return a.duration
}

// This violates LSP because AudioFile doesn't truly have resolution
func (a *AudioFile) GetResolution() Resolution {
    // Forced to return meaningless data or panic
    return Resolution{0, 0} // or panic("Audio files don't have resolution")
}
```

## Issues
1. The `MediaFile` interface forces audio files to implement `GetResolution()` which doesn't make sense for audio
2. Audio files either need to return meaningless data or panic
3. Code that works with `MediaFile` can't reliably use all interface methods
4. The hierarchy violates "Is-Substitutable-For" relationship

## Challenge
Refactor the media player system to follow the Liskov Substitution Principle so that:
1. All types that implement an interface can be used interchangeably
2. Subtypes don't need to implement methods that don't make sense for them
3. Client code can work with interfaces without knowing concrete types
4. No methods should panic or return meaningless data

## Requirements
1. Support both audio and video file playback
2. Maintain ability to get duration for all media types
3. Allow access to resolution only for video files
4. Ensure all interface methods make sense for implementing types
5. Make the code testable and maintainable

## Bonus Challenge
Add support for a new media type (e.g., GIF files) to demonstrate how your solution maintains LSP while handling special cases.

## Expected Usage
The solution should allow code like this to work naturally:

```go
func PlayMedia(media MediaFile) {
    media.Play()
    fmt.Printf("Duration: %d seconds\n", media.GetDuration())
}

// And if it's a video, should be able to safely access resolution
func DisplayVideoInfo(video VideoMedia) {
    res := video.GetResolution()
    fmt.Printf("Resolution: %dx%d\n", res.Width, res.Height)
}
```
