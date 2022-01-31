package audio

type Buffer []int16

func BufferToF64(b *Buffer) (out []float64) {
	out = make([]float64, len(*b))
	for i, x := range *b {
		out[i] = float64(x)
	}
	return out
}
