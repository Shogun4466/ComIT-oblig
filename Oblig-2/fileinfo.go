// Oppgave 1A

package main 

import "fmt"
        "log" 
        "os"

/* Oppgavens innhold skal ha 
Information about file <filnavn>:
Size: X in bytes, KB, MB and GB
Is/Is not a directory
Is/Is not a regular file
Has Unix permission bits: -rwxrwxrwx
Is/Is not append only
Is/Is not a device file
Is/Is not a symbolic link
funksjonsuttryk for variablene med os 
	var bytes int64 = info.Size()
	var kb int64 = (bytes / 1024)
	var mg float64 = (float64)(kb / 1024)
	var gb float64 = (mg / 1024)
	var modePerm = mode & os.ModePerm
	var append = mode & os.ModeAppend
	var device = mode & os.ModeDevice
	var charDevice = mode & os.ModeCharDevice
	var unixBlock = mode & os.ModeDevice
	var symLink = mode & os.ModeSymlink
Bruk av kode og eksempler fra GoLang kopieres og redigeres etter behov under
*/ 


// golang for bruk av FileMode
const (
        // The single letters are the abbreviations
        // used by the String method's formatting.
        ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
        ModeAppend                                     // a: append-only
        ModeExclusive                                  // l: exclusive use
        ModeTemporary                                  // T: temporary file; Plan 9 only
        ModeSymlink                                    // L: symbolic link
        ModeDevice                                     // D: device file
        ModeNamedPipe                                  // p: named pipe (FIFO)
        ModeSocket                                     // S: Unix domain socket
        ModeSetuid                                     // u: setuid
        ModeSetgid                                     // g: setgid
        ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        ModeSticky                                     // t: sticky

        // Mask for the type bits. For regular files, none will be set.
        ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

        ModePerm FileMode = 0777 // Unix permission bits
)

// eksempel på en funksjonsargument for oppgaven 
func main() {
	fi, err := os.Lstat("some-filename")
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}

/*
Kode for å vise typer av file interface: 
type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}
*/

/* Oppgave 1B
Må inneholde en directory som gjør at +build kan kjøre. Eksempel finner man på Golang-packages 
*/
