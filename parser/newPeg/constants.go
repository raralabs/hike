package newPeg

type src uint8

const(
	FILEJOB	src = iota+1
	FAKEJOB
	BRANCHJOB
	SECSOURCE
)

type trns uint8

const(
	AGGJOB trns = iota +1
	DOJOB
	MAPJOB
	JOINJOB
)

type aggregator uint8

const(
	COUNT aggregator = iota +1
	MAX
)

type dotype uint8

const(
	FILTER dotype = iota +1
)

type snk uint8

const(
	STDOUT snk = iota+1
	BLACKHOLE
	PLOT
	INTO
)
