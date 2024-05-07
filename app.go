package main

import (
	"bytes"
	"compress/flate"
	"context"
	"encoding/hex"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"os"
)

// App struct
type App struct {
	ctx context.Context
	rf  string
}

type Xor struct {
	First  string `json:"first"`
	Second string `json:"second"`
	Third  string `json:"third"`
	Four   string `json:"four"`
}

type XorHex struct {
	First  byte
	Second byte
	Third  byte
	Four   byte
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func decode(buffer []byte, xex XorHex) ([]byte, error) {
	buffer = buffer[4:]

	buffer[0] ^= xex.First
	buffer[1] ^= xex.Second
	buffer[2] ^= xex.Third
	buffer[3] ^= xex.Four

	reader := bytes.NewReader(buffer)
	inflated, err := ioutil.ReadAll(flate.NewReader(reader))
	if err != nil {
		return nil, err
	}

	return inflated, nil
}

func encode(buffer []byte, xex XorHex) ([]byte, error) {
	var buf bytes.Buffer
	fw, err := flate.NewWriter(&buf, flate.DefaultCompression)

	if err != nil {
		return nil, err
	}

	_, err = fw.Write(buffer)
	if err != nil {
		return nil, err
	}

	err = fw.Close()
	if err != nil {
		return nil, err
	}

	buffer = buf.Bytes()

	buffer[0] ^= xex.First
	buffer[1] ^= xex.Second
	buffer[2] ^= xex.Third
	buffer[3] ^= xex.Four

	retbuff := make([]byte, 0)
	retbuff = append(retbuff, 0x0E, 0, 0, 0)
	retbuff = append(retbuff, buffer...)

	return retbuff, nil
}

func decompress(rf string, wf string, xex XorHex) error {
	data, err := os.ReadFile(rf)
	if err != nil {
		return err
	}

	decodedData, err := decode(data, xex)
	if err != nil {
		return err
	}

	fileWr, err := os.Create(wf)
	if err != nil {
		return err
	}
	_, err = fileWr.Write(decodedData)
	if err != nil {
		return err
	}
	defer fileWr.Close()

	return nil
}

func compress(rf string, wf string, xex XorHex) error {
	data, err := os.ReadFile(rf)
	if err != nil {
		return err
	}

	encodedData, err := encode(data, xex)
	if err != nil {
		return err
	}

	fileWr, err := os.Create(wf)
	if err != nil {
		return err
	}
	_, err = fileWr.Write(encodedData)
	if err != nil {
		return err
	}
	defer fileWr.Close()

	return nil
}

func (a *App) OpenEnc() string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{Filters: []runtime.FileFilter{{DisplayName: "ENC/DEC files", Pattern: "*.enc;*.dec"}}})
	if err != nil {
		return "Error : " + err.Error()
	}

	a.rf = file

	return "Opened : " + file
}

func (a *App) DecompressEnc(xor Xor) string {
	xorFirst, _ := hex.DecodeString(xor.First)
	xorSecond, _ := hex.DecodeString(xor.Second)
	xorThird, _ := hex.DecodeString(xor.Third)
	xorFour, _ := hex.DecodeString(xor.Four)
	xorHex := XorHex{
		xorFirst[0],
		xorSecond[0],
		xorThird[0],
		xorFour[0],
	}

	if len(a.rf) == 0 {
		return "Error : Open file first"
	}

	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{Filters: []runtime.FileFilter{{DisplayName: "DEC files", Pattern: "*.dec"}}})
	if err != nil {
		return "Error : " + err.Error()
	}

	err = decompress(a.rf, file, xorHex)
	if err != nil {
		return "Error : " + err.Error()
	}

	return "Saved : " + file
}
func (a *App) CloseApp() {
	os.Exit(0)
}

func (a *App) CompressEnc(xor Xor) string {
	xorFirst, _ := hex.DecodeString(xor.First)
	xorSecond, _ := hex.DecodeString(xor.Second)
	xorThird, _ := hex.DecodeString(xor.Third)
	xorFour, _ := hex.DecodeString(xor.Four)
	xorHex := XorHex{
		xorFirst[0],
		xorSecond[0],
		xorThird[0],
		xorFour[0],
	}

	if len(a.rf) == 0 {
		return "Error : Open file first"
	}

	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{Filters: []runtime.FileFilter{{DisplayName: "ENC files", Pattern: "*.enc"}}})
	if err != nil {
		return "Error : " + err.Error()
	}

	err = compress(a.rf, file, xorHex)
	if err != nil {
		return "Error : " + err.Error()
	}

	return "Saved : " + file
}
