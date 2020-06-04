package fileid

func Query(fname string) (uint64, error) {
	return queryFilenameById(fname)
}
