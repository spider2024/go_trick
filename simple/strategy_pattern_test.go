package simple

import "testing"

func Test_add_do(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		a    *add
		args args
		want int
	}{
		{"add", &add{}, args{1, 1}, 2},
		{"add", &add{}, args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &add{}
			if got := a.do(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add.do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reduce_do(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		r    *reduce
		args args
		want int
	}{
		{"reduce", &reduce{}, args{1, 1}, 0},
		{"reduce", &reduce{}, args{1, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reduce{}
			if got := r.do(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("reduce.do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperator_calculate(t *testing.T) {
	type fields struct {
		strategy IStrategy
	}
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"add", fields{strategy: &add{}}, args{1, 1}, 2},
		{"add", fields{strategy: &reduce{}}, args{1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &Operator{
				strategy: tt.fields.strategy,
			}
			if got := op.calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Operator.calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
