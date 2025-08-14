package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

func main() {
	// Comprobar si un archivo existe.
	path := "./fundamentos/files/example.txt"
	exists, err := fileExists(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El archivo existe:", exists)

	// Crear un archivo.
	fmt.Println("Creando el archivo:", createFile(path))

	// Crear un archivo si no existe.
	fmt.Println("Creando el archivo si no existe:", createFileIfNotExists(path))

	// Escribir en un archivo.
	fmt.Println("Escribiendo en el archivo:", writeFile(path, []byte("Hello, world!")))

	// Agregar nuevo contenido a un archivo sin sobreescribirlo.
	fmt.Println("Agregar contenido al archivo:", appendFile(path, []byte("Este es un nuevo contenido.")))

	// Leer el contenido de un archivo.
	content, err := readFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Leyendo el contenido del archivo:", string(content))

	// Estadísticas de un archivo.
	size, modTime, err := fileStats(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tamaño del archivo:", size)
	fmt.Println("Fecha de modificación del archivo:", modTime)

	// Eliminar un archivo.
	fmt.Println("Eliminando el archivo:", deleteFile(path))

	// Lista de archivos en un directorio.
	files, err := listDirList("./articulos")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Archivos en el directorio:", files)

	for _, listedFile := range files {
		fmt.Println(listedFile.Name())
	}
}

func fileExists(path string) (bool, error) {
	// Recupera información del archivo.
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func createFile(path string) error {

	// Si el archivo ya existe, elimina su contenido y lo crea nuevamente.
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// Importante cerrar el archivo para liberar o limpiar recursos.
	// Cada vez que abres un archivo en Go, consumes recursos del sistema como memoria y file descriptors.
	// Evita fugas de memoria.
	defer file.Close()

	return nil
}

func createFileIfNotExists(path string) error {
	// Abre un archivo con indicadores predefinidos.
	// Como quieres interactuar con el archivo.
	// El operador | (pipe) permite combinar varios indicadores.
	// O_RDWR: Permite leer y escribir en el archivo.
	// O_CREATE: Crea el archivo si no existe.
	file, err := os.OpenFile(path, os.O_RDWR|os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

// Sobreescribir el contenido de un archivo
func writeFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func appendFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	return err
}

// Para leer archivos pequeños.
// Para archivos grandes mejor usar buffer.
func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return data, nil

	// return os.ReadFile(path)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

// int64: Tamaño del archivo en bytes.
// time.Time: Fecha de creación del archivo o modificación del archivo.
func fileStats(path string) (int64, time.Time, error) {
	fileInfo, err := os.Lstat(path)

	if err != nil {
		return -1, time.Time{}, err
	}

	size := fileInfo.Size()
	modTime := fileInfo.ModTime()

	return size, modTime, nil
}

// Listar todos los archivos no directorios.
func listDirList(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	listedFiles := make([]fs.DirEntry, 0, len(entries))

	for _, entry := range entries {
		if !entry.IsDir() {
			listedFiles = append(listedFiles, entry)
		}
	}

	return listedFiles, nil
}
