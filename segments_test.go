// Copyright 2016 Andrew O'Neill

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package choices

import "testing"

func TestSegmentContains(t *testing.T) {
	tests := []struct {
		seg  segments
		n    uint64
		want bool
	}{
		{
			seg:  segments{},
			n:    0,
			want: false,
		},
		{
			seg:  segments{},
			n:    1,
			want: false,
		},
		{
			seg:  segments{1},
			n:    0,
			want: true,
		},
		{
			seg:  segments{1 << 7},
			n:    7,
			want: true,
		},
		{
			seg:  segments{0, 1<<7 | 1<<4},
			n:    12,
			want: true,
		},
		{
			seg:  segments{0, 1 << 7},
			n:    14,
			want: false,
		},
	}
	for _, test := range tests {
		got := test.seg.contains(test.n)
		if test.want != got {
			t.Errorf("%v.contains(%v) = %v, want %v",
				test.seg,
				test.n,
				got,
				test.want)
		}
	}
}

func TestSegmentsAvailable(t *testing.T) {
	tests := []struct {
		seg  segments
		want []int
	}{
		{seg: segments{}, want: []int{}},
		{seg: segments{1}, want: []int{0}},
		{seg: segments{1, 1}, want: []int{0, 8}},
		{seg: segments{255, 1}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{seg: segments{1, 255}, want: []int{0, 8, 9, 10, 11, 12, 13, 14, 15}},
		{seg: segments{1, 127}, want: []int{0, 8, 9, 10, 11, 12, 13, 14}},
		{
			seg:  segments{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 72, 80, 88, 96, 104, 112, 120},
		},
		{
			seg:  segments{1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7, 1 << 7},
			want: []int{7, 15, 23, 31, 39, 47, 55, 63, 71, 79, 87, 95, 103, 111, 119, 127},
		},
	}
	for _, test := range tests {
		got := test.seg.available()
		if len(got) != len(test.want) {
			t.Errorf("%v.available() = %v, want %v", test.seg, got, test.want)
			t.FailNow()
		}
		for i, v := range got {
			if v != test.want[i] {
				t.Errorf("%v.available() = %v, want %v", test.seg, got, test.want)
				t.FailNow()
			}
		}

	}
}

func TestSegmentsSet(t *testing.T) {
	tests := []struct {
		seg   segments
		index int
		value bit
		want  segments
	}{
		{
			seg:   segments{},
			index: 0,
			value: one,
			want:  segments{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			seg:   segments{},
			index: 13,
			value: one,
			want:  segments{0, 1 << 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			seg:   SegmentsAll,
			index: 15,
			value: zero,
			want:  segments{255, 127, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		},
	}

	for _, test := range tests {
		test.seg.set(test.index, test.value)
		for i, v := range test.seg {
			if v != test.want[i] {
				t.Errorf("test.set(%v, %v) = %v, want %v", test.index, test.value, test.seg, test.want)
			}
		}
	}
}

func TestSegmentsSample(t *testing.T) {
	tests := []struct {
		seg  segments
		num  int
		want segments
	}{
		{
			seg:  SegmentsAll,
			num:  0,
			want: segments{},
		},
		{
			seg:  SegmentsAll,
			num:  1,
			want: segments{0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		got := test.seg.sample(test.num)
		for i, v := range got {
			if v != test.want[i] {
				t.Errorf("test.sample() = %v %v, want %v", got, test.seg, test.want)
			}
		}
	}
}
