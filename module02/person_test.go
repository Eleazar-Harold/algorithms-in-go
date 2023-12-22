package module02

import (
	"testing"
)

func TestSortPeople(t *testing.T) {
	TestPeopleInterface(t, SortPeople)
}

func TestLess(t *testing.T) {
	type args struct {
		i, j int
	}
	tests := []struct {
		name string
		p    people
		args args
		want bool
	}{
		{
			name: "equal ages, different last names",
			p: []Person{
				{Age: 25, FirstName: "John", LastName: "Doe"},
				{Age: 25, FirstName: "Jane", LastName: "Roe"},
			},
			args: args{0, 1},
			want: true,
		},
		{
			name: "equal ages, equal last names, different first names",
			p: []Person{
				{Age: 25, FirstName: "John", LastName: "Doe"},
				{Age: 25, FirstName: "Jane", LastName: "Doe"},
			},
			args: args{0, 1},
			want: false,
		},
		{
			name: "unequal ages, different last names",
			p: []Person{
				{Age: 45, FirstName: "John", LastName: "Doe"},
				{Age: 30, FirstName: "Jane", LastName: "Roe"},
			},
			args: args{0, 1},
			want: false,
		},
		{
			name: "unequal ages, equal last names, different first names",
			p: []Person{
				{Age: 25, FirstName: "John", LastName: "Doe"},
				{Age: 30, FirstName: "Jane", LastName: "Doe"},
			},
			args: args{0, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("people.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
