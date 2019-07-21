package gfxos

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type Matrix struct {
	SPort *io.ReadWriteCloser
	Wmux sync.Mutex
}


func Open(port *io.ReadWriteCloser) (*Matrix, error) {
	// Create new Matrix
	newMatrix := Matrix{}
	newMatrix.SPort = port

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

func (m *Matrix) SetRotation(r int) error {
	err := m.write(fmt.Sprintf("02%02x", r))
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

func (m *Matrix) drawFastVLine(x int, y int, h int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("04%02x%02x%02x%04x", x, y, h, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) drawFastHLine(x int, y int, w int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("05%02x%02x%02x%04x", x, y, w, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) FillRect(x int, y int, w int, h int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("06%02x%02x%02x%02x%04x", x, y, w, h, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) FillScreen(r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("07%04x", color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) DrawLine(x0 int, y0 int, x1 int, y1 int, h int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("08%02x%02x%02x%02x%04x", x0, y0, x1, y1, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) DrawRect(x int, y int, w int, h int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("09%02x%02x%02x%02x%04x", x, y, w, h, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) DrawCircle(x int, y int, rad int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("0a%02x%02x%02x%04x", x, y, rad, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) FillCircle(x int, y int, rad int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("0c%02x%02x%02x%04x", x, y, rad, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) DrawTriangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, h int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("0e%02x%02x%02x%02x%02x%02x%04x", x0, y0, x1, y1, x2, y2, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) FillTriangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, h int, r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("0f%02x%02x%02x%02x%02x%02x%04x", x0, y0, x1, y1, x2, y2, color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) DrawChar(x int, y int, fr int, fg int, fb int, br int, bg int, bb int, size int, char int) error {
	foreground_color := color888(fr, fg, fb)
	background_color := color888(br, bg, bb)

	err := m.write(fmt.Sprintf("10%02x%02x%04x%04x%02x%02x", x, y, foreground_color, background_color, size, char))
	if err != nil {
		return err
	}
	return nil
}

func (m *Matrix) SetCursor(x int, y int, w int) error {
	err := m.write(fmt.Sprintf("11%02x%02x", x, y))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) SetTextColor(r int, g int, b int) error {
	color := color888(r, g, b)

	err := m.write(fmt.Sprintf("12%04x", color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) setTextColorBG(fr int, fg int, fb int, br int, bg int, bb int) error {
	foreground_color := color888(fr, fg, fb)
	background_color := color888(br, bg, bb)

	err := m.write(fmt.Sprintf("13%04x%04x", foreground_color, background_color))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) SetTextSize(s int) error {
	err := m.write(fmt.Sprintf("14%02x", s))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) SetTextWrap(w int) error {
	err := m.write(fmt.Sprintf("15%02x", w))
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix) CP437(c int) error {
	err := m.write(fmt.Sprintf("15%02x", c))
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

