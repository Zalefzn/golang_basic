package latihan_test

import (
	"testing"

	"go-latihan-1/latihan"
)

func TestLatihan(t *testing.T){
	if latihan.BestFunc() != "Latihan Dasar" {
		t.Fatal('Cant Access this file!');
	}
		
}