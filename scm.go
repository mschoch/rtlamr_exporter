//  Copyright (c) 2018 Marty Schoch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// {
//   "Time": "2018-01-238T17:08:27.723350407-05:00",
//   "Offset": 0,
//   "Length": 0,
//   "Message": {
//     "ID": 23984219,
//     "Type": 5,
//     "TamperPhy": 3,
//     "TamperEnc": 1,
//     "Consumption": 28934702,
//     "ChecksumVal": 32522
//   }
// }
type StandardConsumptionMsg struct {
	Time    string
	Offset  int
	Length  int
	Message StandardConsumptionMsgMessage
}

type StandardConsumptionMsgMessage struct {
	ID          int
	Type        int
	TamperPhy   int
	TamperEnd   int
	Consumption float64
	ChecksumVal int
}
