package gfxos

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"log"
	"sync"
)

type Matrix struct {
	SPort *io.ReadWriteCloser
	Wmux sync.Mutex
}


func Open(portName string) (*Matrix, error) {
	// Set up options.
	options := serial.OpenOptions{
		PortName: "COM30",
		BaudRate: 19200,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		return nil, err
	}

	// Create new Matrix
	newMatrix := Matrix{}
	newMatrix.SPort = &port

	return &newMatrix, nil
}

func (m Matrix) Close() () {
	port := *m.SPort
	port.Close()
}
func (m *Matrix) DrawPixel(x int, y int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("01%02x%02x%04x", x, y, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) InvertDisplay(i int) error {
	err := m.write(fmt.Sprintf("03%02x", i))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) SetRotation(r int) error {
	err := m.write(fmt.Sprintf("02%02x", r))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) write(text string) error {
	b := []byte(text + "\n")
	fmt.Println(b)

	// Write bytes to the port.
	m.Wmux.Lock()
	port := *m.SPort
	_, err := port.Write(b)
	m.Wmux.Unlock()

	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	return nil
}

func color888(r int, g int, b int) int {
	new_r := mapRange(r, 0, 255, 0, 31)
	new_g := mapRange(g, 0, 255, 0, 63)
	new_b := mapRange(b, 0, 255, 0, 31)

	new_g = new_g << 5
	new_r = new_r << 11

	color := new_r | new_g | new_b
	return color
}

func mapRange(value int, fromLow int, fromHigh int, toLow int, toHigh int) int {
	return (((value - fromLow) * (toHigh - toLow)) / (fromHigh - fromLow)) + toLow
}
