package gfxos

type GFX interface {
	Close()
	DrawPixel(x int, y int, r int, g int, b int) error
	SetRotation(r int) error
	InvertDisplay(i int) error
	DrawFastVLine(x int, y int, h int, r int, g int, b int) error
	DrawFastHLine(x int, y int, w int, r int, g int, b int) error
	FillRect(x int, y int, w int, h int, r int, g int, b int) error
	FillScreen(r int, g int, b int) error
	DrawLine(x0 int, y0 int, x1 int, y1 int, r int, g int, b int) error
	DrawRect(x int, y int, w int, h int, r int, g int, b int) error
	DrawCircle(x int, y int, rad int, r int, g int, b int) error
	FillCircle(x int, y int, rad int, r int, g int, b int) error
	DrawTriangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, r int, g int, b int) error
	FillTriangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, r int, g int, b int) error
	DrawChar(x int, y int, fr int, fg int, fb int, br int, bg int, bb int, size int, char int) error
	SetCursor(x int, y int) error
	SetTextColor(r int, g int, b int) error
	SetTextColorBG(fr int, fg int, fb int, br int, bg int, bb int) error
	SetTextSize(s int) error
	SetTextWrap(w int) error
	CP437(c int) error
	Print(str string) error
	PrintLn(str string) error
	SetFont(f int) error
}
