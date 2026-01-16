package smscredit

var normalSegments = []segmentRule{
	{MaxLen: 160, Credits: 1},
	{MaxLen: 306, Credits: 2},
	{MaxLen: 459, Credits: 3},
	{MaxLen: 612, Credits: 4},
	{MaxLen: 765, Credits: 5},
	{MaxLen: 918, Credits: 6},
}

var turkishSegments = []segmentRule{
	{MaxLen: 155, Credits: 1},
	{MaxLen: 294, Credits: 2},
	{MaxLen: 441, Credits: 3},
	{MaxLen: 588, Credits: 4},
	{MaxLen: 735, Credits: 5},
	{MaxLen: 882, Credits: 6},
}

var unicodeSegments = []segmentRule{
	{MaxLen: 70, Credits: 1},
	{MaxLen: 133, Credits: 2},
	{MaxLen: 200, Credits: 3},
	{MaxLen: 267, Credits: 4},
	{MaxLen: 334, Credits: 5},
	{MaxLen: 401, Credits: 6},
}
