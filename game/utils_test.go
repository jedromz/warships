package game

import "testing"

func TestValidateBoardPlacement(t *testing.T) {
	type args struct {
		coords []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid placement",
			args: args{
				coords: []string{"A1", "A2", "A3", "A4"},
			},
			want: true,
		},
		{
			name: "Invalid placement - diagonal ships",
			args: args{
				coords: []string{"A1", "B2", "C3", "D4"},
			},
			want: false,
		},
		{
			name: "Invalid placement - adjacent ships",
			args: args{
				coords: []string{"A1", "A2", "B2", "B3"},
			},
			want: false,
		},
		{
			name: "Invalid placement - ships touching edges",
			args: args{
				coords: []string{"A1", "A2", "A3", "A8"},
			},
			want: false,
		},
		{
			name: "Invalid placement - ships touching edges",
			args: args{
				coords: []string{
					"A1", "A2", "A3", "A4", // battleship
					"B1", "B2", "B3", // cruiser 1
					"C1", "C2", "C3", // cruiser 2
					"D1", "D2", // destroyer 1
					"E1", "E2", // destroyer 2
					"F1", "F2", // destroyer 3
					"G1",  // submarine 1
					"H1",  // submarine 2
					"I1",  // submarine 3
					"J10", // submarine 4 - invalid
				},
			},
			want: false,
		},
		{
			name: "Valid placement - all ships",
			args: args{
				coords: []string{
					"A1", "A2", "A3", "A4", // battleship
					"C1", "C2", "C3", // cruiser 1
					"E1", "E2", "E3", // cruiser 2
					"G1", "G2", // destroyer 1
					"I1", "I2", // destroyer 2
					"A6", "A7", // destroyer 3
					"C6", // submarine 1
					"E6", // submarine 2
					"G6", // submarine 3
					"I6", // submarine 4
				},
			},
			want: true,
		},

		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateBoardPlacement(tt.args.coords); got != tt.want {
				t.Errorf("ValidateBoardPlacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
