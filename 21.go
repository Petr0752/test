package main

import "fmt"

// Интерфейс для воспроизведения музыки
type MusicPlayer interface {
	PlayMusic(file string)
}

// Интерфейс для воспроизведения видео
type MediaPlayer interface {
	PlayVideo(file string)
}

// Реализация воспроизведения музыки через MusicPlayer
type MyMusicPlayer struct{}

func (p *MyMusicPlayer) PlayMusic(file string) {
	fmt.Printf("Playing music file: %s\n", file)
}

// Адаптер, который позволяет использовать MyMusicPlayer как MediaPlayer
type MusicPlayerAdapter struct {
	musicPlayer MusicPlayer
}

func (a *MusicPlayerAdapter) PlayVideo(file string) {
	fmt.Printf("Playing video file: %s\n", file)
	a.musicPlayer.PlayMusic(file)
}

// Пример использования адаптера
func main() {
	musicPlayer := &MyMusicPlayer{}
	mediaPlayer := &MusicPlayerAdapter{musicPlayer: musicPlayer}
	mediaPlayer.PlayVideo("my_video.mp4")
}
