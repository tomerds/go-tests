package main

import "testing"

func TestFizzBuzz(t *testing.T) {

	got := FizzBuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestIterator(t *testing.T) {

	got := Iterator(100)
	want := `1
2
3
4
Buzz
6
7
8
9
Buzz
11
12
13
14
FizzBuzz
16
17
18
19
Buzz
21
22
23
24
Buzz
26
27
28
29
FizzBuzz
31
32
33
34
Buzz
36
37
38
39
Buzz
41
42
43
44
FizzBuzz
46
47
48
49
Buzz
51
52
53
54
Buzz
56
57
58
59
FizzBuzz
61
62
63
64
Buzz
66
67
68
69
Buzz
71
72
73
74
FizzBuzz
76
77
78
79
Buzz
81
82
83
84
Buzz
86
87
88
89
FizzBuzz
91
92
93
94
Buzz
96
97
98
99
Buzz
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
