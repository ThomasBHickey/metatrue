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
    "math/big"
    "testing"
    //"fmt"
)

func TestX07(t *testing.T) {
    dig[0] = 1
    dig[1] = 2
    dig[2] = 5
    rat := (*big.Rat)(round_decimals(3))
    //fmt.Println("sc .125?", rat.String())
    if rat.String()!="1/8" {
        t.Fail()
    }
    //fmt.Println("denom", rat.Denom())
    //fmt.Println("FloatString on .125 rat", rat.FloatString(5))
    //rat2 := big.NewRat(1, 3)
    //fmt.Println("FloatString on 1/3", rat2.FloatString(10))
}
