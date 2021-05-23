package main

import "testing"

func TestApp_isVaccineAvailable(t *testing.T) {
	type fields struct {
		Cowin        CoWin
		PinCode      int
		Age          int
		ScanDays     int
		StartDate    string
		ScanInterval int
		DoseNumber   int
		VaccineName  string
	}
	type args struct {
		session *Sessions
		date    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should return false when no dose is available",
			args: args{
				session: &Sessions{
					MinAgeLimit:            45,
					AvailableCapacityDose1: 0,
					AvailableCapacityDose2: 0,
					Date:                   "02-01-2021",
				},
				date: "02-01-2021",
			},
			fields: fields{
				Age: 45,
			},
			want: false,
		},
		{
			name: "Should return true when dose-1 is available",
			args: args{
				session: &Sessions{
					MinAgeLimit:            45,
					AvailableCapacityDose1: 10,
					AvailableCapacityDose2: 0,
					Date:                   "02-01-2021",
				},
				date: "02-01-2021",
			},
			fields: fields{
				DoseNumber: 1,
				Age:        45,
			},
			want: true,
		},
		{
			name: "Should return true when dose-2 is available",
			args: args{
				session: &Sessions{
					MinAgeLimit:            45,
					AvailableCapacityDose1: 0,
					AvailableCapacityDose2: 10,
					Date:                   "02-01-2021",
				},
				date: "02-01-2021",
			},
			fields: fields{
				DoseNumber: 2,
				Age:        45,
			},
			want: true,
		},
		{
			name: "Should return false lower age group is selected",
			args: args{
				session: &Sessions{
					MinAgeLimit:            45,
					AvailableCapacityDose1: 0,
					AvailableCapacityDose2: 10,
				},
				date: "02-01-2021",
			},
			fields: fields{
				DoseNumber: 2,
				Age:        18,
			},
			want: false,
		},
		{
			name: "Should return true for COVISHIELD",
			args: args{
				session: &Sessions{
					MinAgeLimit:            18,
					AvailableCapacityDose1: 0,
					Date:                   "02-01-2021",
					Vaccine:                "COVISHIELD",
					AvailableCapacityDose2: 10,
				},
				date: "02-01-2021",
			},
			fields: fields{
				VaccineName: "COVISHIELD",
				DoseNumber:  2,
				Age:         18,
			},
			want: true,
		},
		{
			name: "Should return true for ANY Vaccine",
			args: args{
				session: &Sessions{
					MinAgeLimit:            18,
					AvailableCapacityDose1: 0,
					Date:                   "02-01-2021",
					Vaccine:                "COVISHIELD",
					AvailableCapacityDose2: 10,
				},
				date: "02-01-2021",
			},
			fields: fields{
				VaccineName: "ANY",
				DoseNumber:  2,
				Age:         18,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				Cowin:        tt.fields.Cowin,
				PinCode:      tt.fields.PinCode,
				Age:          tt.fields.Age,
				ScanDays:     tt.fields.ScanDays,
				StartDate:    tt.fields.StartDate,
				ScanInterval: tt.fields.ScanInterval,
				DoseNumber:   tt.fields.DoseNumber,
				VaccineName:  tt.fields.VaccineName,
			}
			if got := app.isVaccineAvailable(tt.args.session, tt.args.date); got != tt.want {
				t.Errorf("isVaccineAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
