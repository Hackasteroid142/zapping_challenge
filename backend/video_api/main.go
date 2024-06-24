package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/grafov/m3u8"
)

var (
	firstCall bool       = true
	mu        sync.Mutex // Mutex para sincronizar el acceso al archivo
)

func main() {
	port := 3000

	// Iniciar la goroutine para la actualización periódica
	go func() {
		for {
			handleHLSFile()
			time.Sleep(10 * time.Second)
		}
	}()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		mu.Lock() // Asegura el acceso concurrente seguro al archivo
		defer mu.Unlock()
		res.Header().Set("Access-Control-Allow-Origin", "*")
		h := http.FileServer(http.Dir("files"))
		h.ServeHTTP(res, req)
	})

	fmt.Printf("Starting server on %v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHLSFile() {
	mu.Lock() // Asegura el acceso concurrente seguro al archivo
	defer mu.Unlock()
	fmt.Println("Starting handleHLSFile")
	file, err := os.Open("files/segment.m3u8")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	playlist, listType, err := m3u8.DecodeFrom(file, true)

	if err != nil {
		fmt.Println("Error al decodificar el archivo:", err)
		return
	}

	if listType != m3u8.MEDIA {
		fmt.Println("El archivo no es una lista de reproducción de medios")
		return
	}

	mediaPlaylist := playlist.(*m3u8.MediaPlaylist)
	mediaSequence := mediaPlaylist.SeqNo
	fmt.Printf("EXT-X-MEDIA-SEQUENCE: %d\n", mediaSequence)

	newPlaylist, err := m3u8.NewMediaPlaylist(3, 10)
	if err != nil {
		fmt.Println("Error al crear la nueva lista de reproducción:", err)
		return
	}

	// Añadir segmentos a la nueva lista de reproducción
	if firstCall {
		// Se agregan primeros segment para primer llamado de API
		newPlaylist.Append(fmt.Sprintf("segment%d.ts", mediaSequence), 10.0, "")
		newPlaylist.Append(fmt.Sprintf("segment%d.ts", mediaSequence+1), 10.0, "")
		newPlaylist.Append(fmt.Sprintf("segment%d.ts", mediaSequence+2), 10.0, "")
		firstCall = false
	} else if mediaSequence >= 61 {
		// Se agrega primeros segment cuando se comienza a llegar al limite de segment disponibles
		newPlaylist = mediaPlaylist
		newPlaylist.Append(fmt.Sprintf("segment%d.ts", mediaSequence-61), 10.0, "")
		newPlaylist.Remove()
		if mediaSequence == 63 {
			// Si se alcanza la cantidad limite de segment, se renicia el media-sequence
			newPlaylist.SeqNo = 0
		}
	} else {
		// En caso de que no se cumpla ninguno de los casos anteriores se agrega el segment siguiente
		newPlaylist = mediaPlaylist
		newPlaylist.Append(fmt.Sprintf("segment%d.ts", mediaSequence+3), 10.0, "")
		newPlaylist.Remove()
	}

	// Guardar la nueva lista de reproducción en un archivo
	newFile, err := os.Create("files/segment.m3u8")
	if err != nil {
		fmt.Println("Error al crear el archivo nuevo:", err)
		return
	}
	defer newFile.Close()

	// Se escribe el nuevo archivo con los segment
	newFile.Write(newPlaylist.Encode().Bytes())

}
