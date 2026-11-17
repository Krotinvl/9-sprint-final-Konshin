package main

import "testing"

// Пишите тесты в этом файле

func TestGenerate(t *testing.T) {
	test := []struct {
		name      string
		size      int
		wantLen   int
		expectNil bool
	}{
		{
			name:      "Отрицательный размер",
			size:      -5,
			wantLen:   0,
			expectNil: true,
		},
		{
			name:      "Нулевой размер",
			size:      0,
			wantLen:   0,
			expectNil: true,
		},
		{
			name:      "1 элемент",
			size:      1,
			wantLen:   1,
			expectNil: false,
		},
		{
			name:      "10 элементов",
			size:      10,
			wantLen:   10,
			expectNil: false,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			generatedSlice := generateRandomElements(tc.size)

			if tc.expectNil && generatedSlice != nil {
				t.Errorf("Ожидался nil, получено %v", generatedSlice)
			}

			if !tc.expectNil && generatedSlice == nil {
				t.Errorf("Ожидался не nil, получено nil")
			}

			if len(generatedSlice) != tc.wantLen {
				t.Errorf("Ожидалась длина %d, получена длина %d", tc.wantLen, len(generatedSlice))
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	test := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "Не созданный срез",
			data: nil,
			want: 0,
		},
		{
			name: "Пустой срез",
			data: []int{},
			want: 0,
		},
		{
			name: "Срез из одного элемента",
			data: []int{10},
			want: 10,
		},
		{
			name: "Обычный срез",
			data: []int{10, 7, 20, 80, 34},
			want: 80,
		},
		{
			name: "Срез из отрицательных чисел",
			data: []int{-1, -2, -3, -10},
			want: -1,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			maxInt := maximum(tc.data)

			if maxInt != tc.want {
				t.Errorf("Ожидалось %d, получили %d", tc.want, maxInt)
			}
		})
	}
}
