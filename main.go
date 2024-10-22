package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Kelimeyi komut satırından al
	word := os.Args[1]
	words := os.Args[2]
	// \n kaçış dizisini gerçek satır sonu karakterine dönüştür
	word = strings.ReplaceAll(word, "\n", "\n")
	// Dosyayı oku
	var fileContent string
	if words == "standard" {
		file, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else if words == "shadow" {
		file, err := os.ReadFile("shadow.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else if words == "thinkertoy" {
		file, err := os.ReadFile("thinkertoy.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else {
		fmt.Println("Geçersiz kelime grubu:", words)
		return
	}
	// Dosya içeriğini satırlara ayır
	lines := strings.Split(fileContent, "\n")
	// Her kelime için ASCII sanatını yazdır
	for i, line := range strings.Split(word, "\n") {
		if line == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}
		// Her bir karakteri işleyerek ASCII sanatını oluştur
		for h := 1; h < 9; h++ {
			for _, char := range line {
				printAsciiArtForCharacter(char, h, lines)
			}
			fmt.Println()
		}
	}
}

func printAsciiArtForCharacter(char rune, lineIndex int, lines []string) {
	// ASCII karakterinin indeksini hesapla
	index := (int(char) - 32) * 9
	// ASCII sanatını yazdır
	if index >= 0 && index+8 <= len(lines) {
		fmt.Print(lines[index+lineIndex]) // Burada fmt.Println yerine fmt.Print kullanıldı.
	}
}
