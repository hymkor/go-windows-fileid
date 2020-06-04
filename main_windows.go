package fileid

import (
	"golang.org/x/sys/windows"
)

func queryFilenameById(fname string) (uint64, error) {
	_fname, err := windows.UTF16PtrFromString(fname)
	if err != nil {
		return 0, err
	}
	handle, err := windows.CreateFile(_fname,
		windows.GENERIC_READ,
		0,
		nil,
		windows.OPEN_EXISTING,
		0,
		0)

	if err != nil {
		return 0, err
	}
	defer windows.CloseHandle(handle)

	var data windows.ByHandleFileInformation

	if err = windows.GetFileInformationByHandle(handle, &data); err != nil {
		return 0, err
	}

	return (uint64(data.FileIndexHigh) << 32) | uint64(data.FileIndexLow), nil
}
