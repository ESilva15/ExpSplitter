package expenses

// Do the test like this
// func TestSplitByPercent(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		total       string
// 		shares      []Share
// 		ownerIdx    int
// 		ownerAbsorb bool
// 		want        []int64 // cents
// 	}{
// 		{
// 			name:  "Even split, no remainder",
// 			total: "120.00",
// 			shares: []Share{
// 				{ID: "P1", Percent: d("50")},
// 				{ID: "P2", Percent: d("50")},
// 			},
// 			ownerIdx:    0,
// 			ownerAbsorb: false,
// 			want:        []int64{6000, 6000},
// 		},
// 		{
// 			name:  "Three way split with remainder (largest remainder mode)",
// 			total: "111.11",
// 			shares: []Share{
// 				{ID: "P1", Percent: d("34")},
// 				{ID: "P2", Percent: d("33")},
// 				{ID: "P3", Percent: d("33")},
// 			},
// 			ownerIdx:    0,
// 			ownerAbsorb: false,
// 			// One possible: P1=37.78, P2=36.67, P3=36.66
// 			want: []int64{3778, 3667, 3666},
// 		},
// 		{
// 			name:  "Three way split with remainder (owner absorbs mode)",
// 			total: "111.11",
// 			shares: []Share{
// 				{ID: "P1", Percent: d("34")},
// 				{ID: "P2", Percent: d("33")},
// 				{ID: "P3", Percent: d("33")},
// 			},
// 			ownerIdx:    0,
// 			ownerAbsorb: true,
// 			// Owner takes both extra cents
// 			want: []int64{3779, 3666, 3666},
// 		},
// 		{
// 			name:  "Unequal split with rounding",
// 			total: "10.00",
// 			shares: []Share{
// 				{ID: "P1", Percent: d("70")},
// 				{ID: "P2", Percent: d("30")},
// 			},
// 			ownerIdx:    0,
// 			ownerAbsorb: false,
// 			want:        []int64{700, 300}, // perfect
// 		},
// 		{
// 			name:  "Edge case: all to one person",
// 			total: "99.99",
// 			shares: []Share{
// 				{ID: "P1", Percent: d("100")},
// 			},
// 			ownerIdx:    0,
// 			ownerAbsorb: true,
// 			want:        []int64{9999},
// 		},
// 		{
// 			name:  "Stuff",
// 			total: "111.11",
// 			shares: []Share{
// 				{ID: "P1", Percent: d("0.34")},
// 				{ID: "P2", Percent: d("0.33")},
// 				{ID: "P3", Percent: d("0.33")},
// 			},
// 			ownerIdx:    0,
// 			ownerAbsorb: true,
// 			want:        []int64{3779, 3666, 3666},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := SplitByPercent(d(tt.total), tt.shares, tt.ownerIdx, tt.ownerAbsorb)
// 			if err != nil {
// 				t.Fatalf("unexpected error: %v", err)
// 			}
// 			if len(got) != len(tt.want) {
// 				t.Fatalf("length mismatch: got %d, want %d", len(got), len(tt.want))
// 			}
// 			for i := range got {
// 				if got[i].Cents != tt.want[i] {
// 					t.Errorf("person %s: got %d, want %d", tt.shares[i].ID, got[i].Cents, tt.want[i])
// 				}
// 			}
// 		})
// 	}
// }
