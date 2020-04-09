package geometry

func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil

	} else {
		return 0, errors.New("Zero edges not allowed")
	}

}
