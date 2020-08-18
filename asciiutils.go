package main

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func TrimLeftSpace(buf []byte) []byte {
	for i := 0; len(buf) > 0; i++ {
		if asciiSpace[buf[0]] != 1 {
			break
		}
		buf = buf[1:]
	}
	return buf
}

func TrimRightSpace(buf []byte) []byte {
	for i := 0; len(buf) > 0; i++ {
		if asciiSpace[buf[len(buf)-1]] != 1 {
			break
		}
		buf = buf[:len(buf)-1]
	}
	return buf
}

func TrimSpace(buf []byte) []byte {
	return TrimLeftSpace(TrimRightSpace(buf))
}

func Split2Space(buf []byte) ([]byte, []byte) {
	for i := 0; i < len(buf); i++ {
		if asciiSpace[buf[i]] == 1 {
			return buf[:i], buf[i+1:]
		}
	}
	return buf, []byte{}
}
