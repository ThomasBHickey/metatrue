//	Copyright 2014 Thomas B. Hickey (thomasbhickey@gmail.com)
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//		http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
//
// MetaTrue is to a great extent a reimplementation of Donald Knuth's MetaFont

package metatrue

import (
	//"math/big"
	"testing"
	"fmt"
)

func TestX07(t *testing.T) {
	dig[0] = 1
	dig[1] = 2
	dig[2] = 5
	sc := round_decimals(3)
	if sc.floatString() != "0.125" {
		t.Error("expected 0.125, got", sc.floatString())
	}

}

func TestS101(t *testing.T) {
	if half_unit != 2*quarter_unit ||
		three_quarter_unit != 3*quarter_unit ||
		unity != 2*half_unit ||
		two != 2*unity ||
		three != 3*unity {
		//fmt.Printf("1/4:%x, 1/2:%x, 3/4:%x, 1:%x",quarter_unit, half_unit, three_quarter_unit, unity)
		t.Error()
	}
}

func TestS095(t *testing.T) {
	bi := (int64(1) << 31) - 1
	if el_gordo != bi {
		t.Error("El Gordo not right?", el_gordo, bi)
	}
}

func TestS103(t *testing.T) {
	sc := (scaled)(unity / 3)
	tsThird := sc.floatString()
	//fmt.Println("floatString of 1/3", tsThird)
	if tsThird != "0.33333" {
		t.Error("Expected", "0.33333", "got", tsThird)
	}
	ts1 := (scaled)(unity).floatString()
	//fmt.Println("floatString of 1", ts1)
	if ts1 != "1" {
		t.Error("Expected", "1", "got", ts1)
	}
	ts2Thirds := (scaled)(two / 3).floatString()
	if ts2Thirds != "0.66666" {
		t.Error("Expected 0.66666, got", ts2Thirds)
	}
	tsneg := (scaled)(-two / 3).floatString()
	if tsneg != "-0.66666" {
		t.Error("expected -0.66666, got", tsneg)
	}
}

func TestS105(t *testing.T) {
	if fraction_one != fraction_half*2 ||
		fraction_two != fraction_one*2 ||
		fraction_three != fraction_one+fraction_two ||
		fraction_four != 2*fraction_two {
		t.Error("fractions", fraction_half>>28, fraction_one>>28, fraction_two>>28, fraction_three>>28, fraction_four>>28)
	}
}

func TestS106(t *testing.T) {
	if forty_five_deg*2 != ninety_deg {
		t.Error("2*45 != 90?", forty_five_deg>>20, ninety_deg>>20)
	}
	if one_eighty_deg != 2*ninety_deg ||
		three_sixty_deg != 4*ninety_deg {
		t.Error("Problem with degrees", one_eighty_deg>>20, three_sixty_deg>>20)
	}
}

func TestS107(t *testing.T) {
	if make_fraction(1, 2) != fraction_half {
		t.Error("1/2 fraction not working")
	}
}

func TestS109(t *testing.T) {
	if take_fraction(2, fraction_half) != 1 {
		t.Error("take_fraction problem", take_fraction(2, fraction_half))
	}
}

func TestS112(t *testing.T) {
	if take_scaled(1, unity) != 1 {
		t.Error("take_scaled", take_scaled(1, unity))
	}
	if take_scaled(4, quarter_unit) != 1 {
		t.Error("take_scaled expected 4", take_scaled(1, quarter_unit))
	}
}

func TestS114(t *testing.T){
    if make_scaled(1,2) != half_unit {
        fmt.Printf("half_unit    %x\n", half_unit)
        fmt.Printf("returned val %x\n", make_fraction(1,2))
        t.Error("make_fraction", make_fraction(1,2))
    }
}
