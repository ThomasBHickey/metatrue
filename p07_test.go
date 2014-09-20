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
	//"fmt"
)

func TestX07(t *testing.T) {
	dig[0] = 1
	dig[1] = 2
	dig[2] = 5
	sc := round_decimals(3)
	//fmt.Println("sc .125?", sc.floatString())
	if sc.floatString() != "0.125" {
		t.Error("expected 0.125, got", sc.floatString())
	}
	//fmt.Println("denom", rat.Denom())
	//fmt.Println("FloatString on .125 rat", rat.FloatString(5))
	//rat2 := big.NewRat(1, 3)
	//fmt.Println("FloatString on 1/3", rat2.FloatString(10))
}

func TestS103(t *testing.T) {
	sc := (scaled)(unity/3)
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
	ts2Thirds := (scaled)(two/3).floatString()
	if ts2Thirds != "0.66666" {
		t.Error("Expected 0.66666, got", ts2Thirds)
	}
	tsneg := (scaled)(-two/3).floatString()
	if tsneg != "-0.66666" {
	    t.Error("expected -0.66666, got", tsneg)
	}
}
