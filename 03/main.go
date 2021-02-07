package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	exercises3()
}

func exercises3() {
	//buffer := make([]byte, 1024)
	//rand.Reader.Read(buffer)
	//
	//bin, err := os.Create("random.dat")
	//exitIfError(err)
	//defer bin.Close()
	//bin.Write(buffer)
}

func exercises2() {
	buffer := make([]byte, 1024)
	rand.Reader.Read(buffer)

	bin, err := os.Create("random.dat")
	exitIfError(err)
	defer bin.Close()
	bin.Write(buffer)
}

func exercises1() {
	src, err := os.Open("Lenna.png")
	exitIfError(err)
	defer src.Close()

	dest, err := os.Create("CopyLenna.png")
	exitIfError(err)
	defer dest.Close()

	io.Copy(dest, src)
}

func tee() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("tee reader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = ioutil.ReadAll(teeReader)
	fmt.Println(buffer.String())
}

func multi() {
	header := bytes.NewBufferString("----HEADER----\n")
	content := bytes.NewBufferString("content\n")
	footer := bytes.NewBufferString("----FOOTER----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}

func readCSV() {
	reader := strings.NewReader(csvText)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[0], line[1:])
	}
}

const csvText = `foo1,bar1,baz1
foo2,bar2,baz2
foo3,bar3,baz3
`

func fscan() {
	const source = "123 3.14 1.0e4 test"
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v, f=%#v, g=%#v, s=%#v\n", i, f, g, s)
}

func scanText() {
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

func readText() {
	reader := bufio.NewReader(strings.NewReader(text))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

const text = `1行目
2行目
3行目
`

func writePNG() {
	file, err := os.Open("Lenna.png")
	exitIfError(err)
	defer file.Close()
	newfile, err := os.Create("Lenna2.png")
	exitIfError(err)
	defer newfile.Close()

	chunks := readChunks(file)
	io.WriteString(newfile, "\x89PNG\r\n\x1a\n")
	io.Copy(newfile, chunks[0])
	io.Copy(newfile, textChunk("foo bar baz"))
	for _, chunk := range chunks[1:] {
		io.Copy(newfile, chunk)
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func readPNG() {
	file, err := os.Open("Lenna2.png")
	exitIfError(err)
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader
	var offset int64 = 8
	file.Seek(offset, 0) // 最初の8バイトを飛ばす

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

// コマンドラインで確認する方法
// hexdump -n 60 -C Lenna2.png
func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)

	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func endian() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("BigEndian: %d\n", i)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &i)
	fmt.Printf("LittleEndian: %d\n", i)
}

func sectionRead() {
	reader := strings.NewReader("example io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}

func request() {
	conn, err := net.Dial("tcp", "example.com:80")
	exitIfError(err)

	conn.Write([]byte("GET / HTTP/1.1\nHost: example.com:80\n\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func file() {
	f, err := os.Open("file.txt")
	exitIfError(err)

	defer f.Close()
	io.Copy(os.Stdout, f)
}

func stdin() {
	for {
		buffer := make([]byte, 8)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
		}
		fmt.Printf("size=%d, input='%s'\n", size, string(buffer))
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
