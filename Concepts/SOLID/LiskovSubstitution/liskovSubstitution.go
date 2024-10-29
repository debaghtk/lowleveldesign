package liskovsubstitution

import "fmt"

type MediaFile interface {
	Play()
	GetDuration() int // in seconds
}

type VideoMediaFile interface {
	Play()
	GetDuration() int // in seconds
	GetResolution() Resolution
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

type GifFile struct {
	filename   string
	resolution Resolution
	isAnimated bool
	duration   int // Will be 0 for static GIFs
}

func (g *GifFile) Play() {
	if g.isAnimated {
		fmt.Printf("Playing animated GIF: %s\n", g.filename)
	} else {
		fmt.Printf("Displaying static GIF: %s\n", g.filename)
	}
}

func (g *GifFile) GetDuration() int {
	if !g.isAnimated {
		return 0 // Static GIFs have no duration
	}
	return g.duration
}

func (g *GifFile) GetResolution() Resolution {
	return g.resolution
}

// Optional: Add a specific method for GIF-specific behavior
func (g *GifFile) IsAnimated() bool {
	return g.isAnimated
}
