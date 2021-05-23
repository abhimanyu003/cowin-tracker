package main

import "testing"

func Test_validateDate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should return error for non-valid date",
			args: args{
				input: "02-444-2021",
			},
			wantErr: true,
		},
		{
			name: "Should return nil for valid date",
			args: args{
				input: "02-11-2021",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDate(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("validateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should return err for non-valid number",
			args: args{
				input: "cowin",
			},
			wantErr: true,
		},
		{
			name: "Should return nil for valid number",
			args: args{
				input: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateNumber(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("validateNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
