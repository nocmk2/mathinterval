
# mathinterval

parse math interval with golang

## Install
go get github.com/nocmk2/mathinterval

## Example
``` golang
func Example() {
	tests := []struct {
		expr  string
		value float64
		want  bool
	}{
		{"[1,100)", 1.0009, true},
		{"[1,100)", -1, false},
		{"(1,100)", 1, false},
		{"[9,198)", 9, true},
		{"(9,198)", 9, false},
		{"[1,inf)", 10000, true},
		{"[100,200]", 200, true},
		{"(1,123.98]", 123, true},
		{"(-inf,100000)", -3955.955, true},
	}

	for _, tt := range tests {
		fmt.Println(Get(tt.expr).Hit(tt.value))
	}

	//Output:
	//true
	//false
	//false
	//true
	//false
	//true
	//true
	//true
	//true

}
```