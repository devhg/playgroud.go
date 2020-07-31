package m_interface

import "testing"

func TestAnimalDo(t *testing.T) {
	type args struct {
		a Animal
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCat_Eat(t *testing.T) {
	type fields struct {
		ID   int
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cat{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := c.Eat(); got != tt.want {
				t.Errorf("Eat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCat_Run(t *testing.T) {
	type fields struct {
		ID   int
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cat{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := c.Run(); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDog_Eat(t *testing.T) {
	type fields struct {
		ID   int
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dog{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := d.Eat(); got != tt.want {
				t.Errorf("Eat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDog_Run(t *testing.T) {
	type fields struct {
		ID   int
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dog{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := d.Run(); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_(t *testing.T) {
	TestInterface()
}
