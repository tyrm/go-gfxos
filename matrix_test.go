package gfxos

//MockReadWriteCloser - fake read write closer object
type MockReadWriteCloser struct {
	BytesRead    []byte
	BytesWritten []byte
	ReadErr      error
	WriteErr     error
	CloseErr     error
}

//Read - satisfies reader interface
func (r *MockReadWriteCloser) Read(p []byte) (n int, err error) {

	if err = r.ReadErr; err == nil {
		r.BytesRead = p
		n = len(p)
	}
	return
}

//Close - satisfies closer interface
func (r *MockReadWriteCloser) Close() (err error) {
	err = r.CloseErr
	return
}

//Write - satisfies writer interface
func (r *MockReadWriteCloser) Write(p []byte) (n int, err error) {

	if err = r.WriteErr; err != nil {
		r.BytesWritten = p
		n = len(p)
	}
	return
}
