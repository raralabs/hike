package newPeg

type stage uint8

const(
	SOURCE stage = iota+1
	TRANSFORM
	SINK
)
