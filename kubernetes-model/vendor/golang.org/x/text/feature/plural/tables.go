/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package plural

// CLDRVersion is the CLDR version from which the tables in this package are derived.
const CLDRVersion = "31"

var ordinalRules = []pluralCheck{ // 58 elements
	0:  {cat: 0x2f, setID: 0x4},
	1:  {cat: 0x3a, setID: 0x5},
	2:  {cat: 0x22, setID: 0x1},
	3:  {cat: 0x22, setID: 0x6},
	4:  {cat: 0x22, setID: 0x7},
	5:  {cat: 0x2f, setID: 0x8},
	6:  {cat: 0x3c, setID: 0x9},
	7:  {cat: 0x2f, setID: 0xa},
	8:  {cat: 0x3c, setID: 0xb},
	9:  {cat: 0x2d, setID: 0xc},
	10: {cat: 0x2d, setID: 0xd},
	11: {cat: 0x2f, setID: 0xe},
	12: {cat: 0x35, setID: 0x3},
	13: {cat: 0xc5, setID: 0xf},
	14: {cat: 0x2, setID: 0x1},
	15: {cat: 0x5, setID: 0x3},
	16: {cat: 0xd, setID: 0x10},
	17: {cat: 0x22, setID: 0x1},
	18: {cat: 0x2f, setID: 0x11},
	19: {cat: 0x3d, setID: 0x12},
	20: {cat: 0x2f, setID: 0x13},
	21: {cat: 0x3a, setID: 0x14},
	22: {cat: 0x2f, setID: 0x15},
	23: {cat: 0x3b, setID: 0x16},
	24: {cat: 0x2f, setID: 0xa},
	25: {cat: 0x3c, setID: 0xb},
	26: {cat: 0x22, setID: 0x1},
	27: {cat: 0x23, setID: 0x17},
	28: {cat: 0x24, setID: 0x18},
	29: {cat: 0x22, setID: 0x19},
	30: {cat: 0x23, setID: 0x2},
	31: {cat: 0x24, setID: 0x18},
	32: {cat: 0xf, setID: 0x13},
	33: {cat: 0x1a, setID: 0x14},
	34: {cat: 0xf, setID: 0x15},
	35: {cat: 0x1b, setID: 0x16},
	36: {cat: 0xf, setID: 0x1a},
	37: {cat: 0x1d, setID: 0x1b},
	38: {cat: 0xa, setID: 0x1c},
	39: {cat: 0xa, setID: 0x1d},
	40: {cat: 0xc, setID: 0x1e},
	41: {cat: 0xe4, setID: 0x0},
	42: {cat: 0x5, setID: 0x3},
	43: {cat: 0xd, setID: 0xc},
	44: {cat: 0xd, setID: 0x1f},
	45: {cat: 0x22, setID: 0x1},
	46: {cat: 0x23, setID: 0x17},
	47: {cat: 0x24, setID: 0x18},
	48: {cat: 0x25, setID: 0x20},
	49: {cat: 0x22, setID: 0x21},
	50: {cat: 0x23, setID: 0x17},
	51: {cat: 0x24, setID: 0x18},
	52: {cat: 0x25, setID: 0x20},
	53: {cat: 0x21, setID: 0x22},
	54: {cat: 0x22, setID: 0x1},
	55: {cat: 0x23, setID: 0x2},
	56: {cat: 0x24, setID: 0x23},
	57: {cat: 0x25, setID: 0x24},
} // Size: 140 bytes

var ordinalIndex = []uint8{ // 20 elements
	0x00, 0x00, 0x02, 0x03, 0x04, 0x05, 0x07, 0x09,
	0x0d, 0x0e, 0x11, 0x14, 0x1a, 0x1d, 0x20, 0x26,
	0x2d, 0x31, 0x35, 0x3a,
} // Size: 44 bytes

var ordinalLangToIndex = []uint8{ // 754 elements
	// Entry 0 - 3F
	0x00, 0x0d, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x11, 0x11, 0x00, 0x00, 0x00, 0x00,
	0x0f, 0x00, 0x00, 0x0f, 0x0f, 0x00, 0x00, 0x05,
	0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 40 - 7F
	0x00, 0x00, 0x11, 0x11, 0x11, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x12, 0x12, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 80 - BF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	// Entry C0 - FF
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 100 - 13F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02, 0x00,
	0x00, 0x00, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	// Entry 140 - 17F
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
	0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x10,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x03, 0x02,
	0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 180 - 1BF
	0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09,
	0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
	0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 1C0 - 1FF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x0e,
	0x00, 0x00, 0x00, 0x00, 0x0c, 0x0c, 0x02, 0x02,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 200 - 23F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x04, 0x04,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 240 - 27F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x02, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 280 - 2BF
	0x00, 0x0a, 0x0a, 0x0a, 0x0a, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x06, 0x06, 0x00, 0x00,
	// Entry 2C0 - 2FF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00,
} // Size: 778 bytes

var ordinalInclusionMasks = []uint64{ // 100 elements
	// Entry 0 - 1F
	0x0000000400004009, 0x00000002120800d3, 0x0000000010a10195, 0x0000000842810581,
	0x0000000841030081, 0x0000001210010041, 0x0000001100011001, 0x0000000614010001,
	0x0000000614018001, 0x0000000600012001, 0x0000000200014001, 0x0000000010198031,
	0x0000000010610331, 0x0000000040010f01, 0x0000000040070001, 0x0000000010010001,
	0x0000000000011001, 0x000000001c010001, 0x000000001c010001, 0x0000000000012001,
	0x0000000020014001, 0x0000000010080011, 0x0000000010200111, 0x0000000040000501,
	0x0000000040020001, 0x0000000010000001, 0x0000000000001001, 0x0000000014000001,
	0x0000000014000001, 0x0000000000002001, 0x0000000000004001, 0x0000000010080011,
	// Entry 20 - 3F
	0x0000000010200111, 0x0000000040000501, 0x0000000040020001, 0x0000000010000001,
	0x0000000000001001, 0x0000000014000001, 0x0000000014000001, 0x0000000000002001,
	0x0000000080014001, 0x0000000010080011, 0x0000000010200111, 0x0000000040000501,
	0x0000000040020001, 0x0000000010000001, 0x0000000000001001, 0x0000000014000001,
	0x0000000014000001, 0x0000000000002001, 0x0000000020004001, 0x0000000010080011,
	0x0000000010200111, 0x0000000040000501, 0x0000000040020001, 0x0000000010000001,
	0x0000000000001001, 0x0000000014000001, 0x0000000014000001, 0x0000000000002001,
	0x0000000080014001, 0x0000000010080011, 0x0000000010200111, 0x0000000040000501,
	// Entry 40 - 5F
	0x0000000040020001, 0x0000000010000001, 0x0000000000001001, 0x0000000014000001,
	0x0000000014000001, 0x0000000000002001, 0x0000000020004001, 0x0000000010080011,
	0x0000000010200111, 0x0000000040000501, 0x0000000040020001, 0x0000000010000001,
	0x0000000000001001, 0x0000000014000001, 0x0000000014000001, 0x0000000000002001,
	0x000000002001c001, 0x0000000010080011, 0x0000000010200111, 0x0000000040000501,
	0x0000000040020001, 0x0000000010000001, 0x0000000000001001, 0x0000000014000001,
	0x0000000014000001, 0x0000000000002001, 0x0000000080004001, 0x0000000010080011,
	0x0000000010200111, 0x0000000040000501, 0x0000000040020001, 0x0000000010000001,
	// Entry 60 - 7F
	0x0000000000001001, 0x0000000014000001, 0x0000000014000001, 0x0000000000002001,
} // Size: 824 bytes

// Slots used for ordinal: 3A of 0xFF rules; 14 of 0xFF indexes; 37 of 64 sets

var cardinalRules = []pluralCheck{ // 166 elements
	0:   {cat: 0x2, setID: 0x3},
	1:   {cat: 0x22, setID: 0x1},
	2:   {cat: 0x2, setID: 0x4},
	3:   {cat: 0x2, setID: 0x4},
	4:   {cat: 0x7, setID: 0x1},
	5:   {cat: 0x62, setID: 0x3},
	6:   {cat: 0x22, setID: 0x4},
	7:   {cat: 0x7, setID: 0x3},
	8:   {cat: 0x42, setID: 0x1},
	9:   {cat: 0x22, setID: 0x4},
	10:  {cat: 0x22, setID: 0x4},
	11:  {cat: 0x22, setID: 0x5},
	12:  {cat: 0x22, setID: 0x1},
	13:  {cat: 0x22, setID: 0x1},
	14:  {cat: 0x7, setID: 0x4},
	15:  {cat: 0x92, setID: 0x3},
	16:  {cat: 0xf, setID: 0x6},
	17:  {cat: 0x1f, setID: 0x7},
	18:  {cat: 0x82, setID: 0x3},
	19:  {cat: 0x92, setID: 0x3},
	20:  {cat: 0xf, setID: 0x6},
	21:  {cat: 0x62, setID: 0x3},
	22:  {cat: 0x4a, setID: 0x6},
	23:  {cat: 0x7, setID: 0x8},
	24:  {cat: 0x62, setID: 0x3},
	25:  {cat: 0x1f, setID: 0x9},
	26:  {cat: 0x62, setID: 0x3},
	27:  {cat: 0x5f, setID: 0x9},
	28:  {cat: 0x72, setID: 0x3},
	29:  {cat: 0x29, setID: 0xa},
	30:  {cat: 0x29, setID: 0xb},
	31:  {cat: 0x4f, setID: 0xb},
	32:  {cat: 0x61, setID: 0x2},
	33:  {cat: 0x2f, setID: 0x6},
	34:  {cat: 0x3a, setID: 0x7},
	35:  {cat: 0x4f, setID: 0x6},
	36:  {cat: 0x5f, setID: 0x7},
	37:  {cat: 0x62, setID: 0x2},
	38:  {cat: 0x4f, setID: 0x6},
	39:  {cat: 0x72, setID: 0x2},
	40:  {cat: 0x21, setID: 0x3},
	41:  {cat: 0x7, setID: 0x4},
	42:  {cat: 0x32, setID: 0x3},
	43:  {cat: 0x21, setID: 0x3},
	44:  {cat: 0x22, setID: 0x1},
	45:  {cat: 0x22, setID: 0x1},
	46:  {cat: 0x23, setID: 0x2},
	47:  {cat: 0x2, setID: 0x3},
	48:  {cat: 0x22, setID: 0x1},
	49:  {cat: 0x24, setID: 0xc},
	50:  {cat: 0x7, setID: 0x1},
	51:  {cat: 0x62, setID: 0x3},
	52:  {cat: 0x74, setID: 0x3},
	53:  {cat: 0x24, setID: 0x3},
	54:  {cat: 0x2f, setID: 0xd},
	55:  {cat: 0x34, setID: 0x1},
	56:  {cat: 0xf, setID: 0x6},
	57:  {cat: 0x1f, setID: 0x7},
	58:  {cat: 0x62, setID: 0x3},
	59:  {cat: 0x4f, setID: 0x6},
	60:  {cat: 0x5a, setID: 0x7},
	61:  {cat: 0xf, setID: 0xe},
	62:  {cat: 0x1f, setID: 0xf},
	63:  {cat: 0x64, setID: 0x3},
	64:  {cat: 0x4f, setID: 0xe},
	65:  {cat: 0x5c, setID: 0xf},
	66:  {cat: 0x22, setID: 0x10},
	67:  {cat: 0x23, setID: 0x11},
	68:  {cat: 0x24, setID: 0x12},
	69:  {cat: 0xf, setID: 0x1},
	70:  {cat: 0x62, setID: 0x3},
	71:  {cat: 0xf, setID: 0x2},
	72:  {cat: 0x63, setID: 0x3},
	73:  {cat: 0xf, setID: 0x13},
	74:  {cat: 0x64, setID: 0x3},
	75:  {cat: 0x74, setID: 0x3},
	76:  {cat: 0xf, setID: 0x1},
	77:  {cat: 0x62, setID: 0x3},
	78:  {cat: 0x4a, setID: 0x1},
	79:  {cat: 0xf, setID: 0x2},
	80:  {cat: 0x63, setID: 0x3},
	81:  {cat: 0x4b, setID: 0x2},
	82:  {cat: 0xf, setID: 0x13},
	83:  {cat: 0x64, setID: 0x3},
	84:  {cat: 0x4c, setID: 0x13},
	85:  {cat: 0x7, setID: 0x1},
	86:  {cat: 0x62, setID: 0x3},
	87:  {cat: 0x7, setID: 0x2},
	88:  {cat: 0x63, setID: 0x3},
	89:  {cat: 0x2f, setID: 0xa},
	90:  {cat: 0x37, setID: 0x14},
	91:  {cat: 0x65, setID: 0x3},
	92:  {cat: 0x7, setID: 0x1},
	93:  {cat: 0x62, setID: 0x3},
	94:  {cat: 0x7, setID: 0x15},
	95:  {cat: 0x64, setID: 0x3},
	96:  {cat: 0x75, setID: 0x3},
	97:  {cat: 0x7, setID: 0x1},
	98:  {cat: 0x62, setID: 0x3},
	99:  {cat: 0xf, setID: 0xe},
	100: {cat: 0x1f, setID: 0xf},
	101: {cat: 0x64, setID: 0x3},
	102: {cat: 0xf, setID: 0x16},
	103: {cat: 0x17, setID: 0x1},
	104: {cat: 0x65, setID: 0x3},
	105: {cat: 0xf, setID: 0x17},
	106: {cat: 0x65, setID: 0x3},
	107: {cat: 0xf, setID: 0xf},
	108: {cat: 0x65, setID: 0x3},
	109: {cat: 0x2f, setID: 0x6},
	110: {cat: 0x3a, setID: 0x7},
	111: {cat: 0x2f, setID: 0xe},
	112: {cat: 0x3c, setID: 0xf},
	113: {cat: 0x2d, setID: 0xa},
	114: {cat: 0x2d, setID: 0x17},
	115: {cat: 0x2d, setID: 0x18},
	116: {cat: 0x2f, setID: 0x6},
	117: {cat: 0x3a, setID: 0xb},
	118: {cat: 0x2f, setID: 0x19},
	119: {cat: 0x3c, setID: 0xb},
	120: {cat: 0x55, setID: 0x3},
	121: {cat: 0x22, setID: 0x1},
	122: {cat: 0x24, setID: 0x3},
	123: {cat: 0x2c, setID: 0xc},
	124: {cat: 0x2d, setID: 0xb},
	125: {cat: 0xf, setID: 0x6},
	126: {cat: 0x1f, setID: 0x7},
	127: {cat: 0x62, setID: 0x3},
	128: {cat: 0xf, setID: 0xe},
	129: {cat: 0x1f, setID: 0xf},
	130: {cat: 0x64, setID: 0x3},
	131: {cat: 0xf, setID: 0xa},
	132: {cat: 0x65, setID: 0x3},
	133: {cat: 0xf, setID: 0x17},
	134: {cat: 0x65, setID: 0x3},
	135: {cat: 0xf, setID: 0x18},
	136: {cat: 0x65, setID: 0x3},
	137: {cat: 0x2f, setID: 0x6},
	138: {cat: 0x3a, setID: 0x1a},
	139: {cat: 0x2f, setID: 0x1b},
	140: {cat: 0x3b, setID: 0x1c},
	141: {cat: 0x2f, setID: 0x1d},
	142: {cat: 0x3c, setID: 0x1e},
	143: {cat: 0x37, setID: 0x3},
	144: {cat: 0xa5, setID: 0x0},
	145: {cat: 0x22, setID: 0x1},
	146: {cat: 0x23, setID: 0x2},
	147: {cat: 0x24, setID: 0x1f},
	148: {cat: 0x25, setID: 0x20},
	149: {cat: 0xf, setID: 0x6},
	150: {cat: 0x62, setID: 0x3},
	151: {cat: 0xf, setID: 0x1b},
	152: {cat: 0x63, setID: 0x3},
	153: {cat: 0xf, setID: 0x21},
	154: {cat: 0x64, setID: 0x3},
	155: {cat: 0x75, setID: 0x3},
	156: {cat: 0x21, setID: 0x3},
	157: {cat: 0x22, setID: 0x1},
	158: {cat: 0x23, setID: 0x2},
	159: {cat: 0x2c, setID: 0x22},
	160: {cat: 0x2d, setID: 0x5},
	161: {cat: 0x21, setID: 0x3},
	162: {cat: 0x22, setID: 0x1},
	163: {cat: 0x23, setID: 0x2},
	164: {cat: 0x24, setID: 0x23},
	165: {cat: 0x25, setID: 0x24},
} // Size: 356 bytes

var cardinalIndex = []uint8{ // 36 elements
	0x00, 0x00, 0x02, 0x03, 0x04, 0x06, 0x09, 0x0a,
	0x0c, 0x0d, 0x10, 0x14, 0x17, 0x1d, 0x28, 0x2b,
	0x2d, 0x2f, 0x32, 0x38, 0x42, 0x45, 0x4c, 0x55,
	0x5c, 0x61, 0x6d, 0x74, 0x79, 0x7d, 0x89, 0x91,
	0x95, 0x9c, 0xa1, 0xa6,
} // Size: 60 bytes

var cardinalLangToIndex = []uint8{ // 754 elements
	// Entry 0 - 3F
	0x00, 0x04, 0x04, 0x08, 0x08, 0x08, 0x00, 0x00,
	0x06, 0x06, 0x01, 0x01, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x01, 0x01, 0x08, 0x08, 0x04, 0x04,
	0x08, 0x00, 0x00, 0x08, 0x08, 0x00, 0x00, 0x1a,
	0x1a, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x06,
	// Entry 40 - 7F
	0x00, 0x00, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00,
	0x1e, 0x1e, 0x08, 0x08, 0x13, 0x00, 0x00, 0x13,
	0x13, 0x04, 0x04, 0x04, 0x04, 0x04, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x18,
	0x18, 0x00, 0x00, 0x22, 0x22, 0x09, 0x09, 0x09,
	0x00, 0x00, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x00, 0x00, 0x16, 0x16, 0x00, 0x00,
	0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08,
	// Entry 80 - BF
	0x08, 0x08, 0x08, 0x08, 0x08, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	// Entry C0 - FF
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	// Entry 100 - 13F
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x04, 0x04, 0x08,
	0x08, 0x00, 0x00, 0x01, 0x01, 0x01, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x04, 0x04, 0x0c, 0x0c, 0x08,
	0x08, 0x08, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	// Entry 140 - 17F
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x08, 0x08, 0x04, 0x04, 0x1f, 0x1f, 0x14,
	0x14, 0x04, 0x04, 0x08, 0x08, 0x08, 0x08, 0x01,
	0x01, 0x06, 0x00, 0x00, 0x20, 0x20, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x17, 0x17, 0x01, 0x01,
	0x13, 0x13, 0x13, 0x16, 0x16, 0x08, 0x08, 0x02,
	0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a,
	// Entry 180 - 1BF
	0x0a, 0x04, 0x04, 0x04, 0x04, 0x04, 0x10, 0x00,
	0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x00, 0x08,
	0x08, 0x02, 0x02, 0x08, 0x00, 0x00, 0x08, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x08, 0x08, 0x08, 0x08, 0x00, 0x00, 0x0f, 0x0f,
	0x08, 0x10, 0x10, 0x08, 0x08, 0x0e, 0x0e, 0x08,
	// Entry 1C0 - 1FF
	0x08, 0x08, 0x08, 0x00, 0x00, 0x06, 0x06, 0x06,
	0x06, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1b,
	0x1b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d,
	0x0d, 0x08, 0x08, 0x08, 0x00, 0x00, 0x00, 0x00,
	0x06, 0x06, 0x00, 0x00, 0x08, 0x08, 0x0b, 0x0b,
	0x08, 0x08, 0x08, 0x08, 0x01, 0x01, 0x00, 0x00,
	0x00, 0x00, 0x1c, 0x1c, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x08, 0x10, 0x10, 0x08, 0x08, 0x08,
	// Entry 200 - 23F
	0x08, 0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x00,
	0x08, 0x06, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x06,
	0x00, 0x00, 0x06, 0x06, 0x08, 0x19, 0x19, 0x0d,
	0x0d, 0x08, 0x08, 0x03, 0x03, 0x03, 0x03, 0x03,
	0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
	// Entry 240 - 27F
	0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x00, 0x00,
	0x12, 0x12, 0x12, 0x08, 0x08, 0x1d, 0x1d, 0x1d,
	0x1d, 0x1d, 0x1d, 0x1d, 0x00, 0x00, 0x08, 0x08,
	0x00, 0x00, 0x08, 0x08, 0x00, 0x00, 0x08, 0x10,
	0x10, 0x10, 0x10, 0x08, 0x08, 0x00, 0x00, 0x00,
	0x00, 0x11, 0x00, 0x00, 0x11, 0x11, 0x05, 0x05,
	0x18, 0x18, 0x15, 0x15, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	// Entry 280 - 2BF
	0x08, 0x08, 0x08, 0x08, 0x08, 0x13, 0x13, 0x13,
	0x13, 0x13, 0x13, 0x13, 0x13, 0x13, 0x13, 0x13,
	0x08, 0x08, 0x08, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x00,
	0x00, 0x06, 0x06, 0x06, 0x08, 0x08, 0x08, 0x08,
	0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x00, 0x00,
	0x07, 0x07, 0x08, 0x08, 0x1d, 0x1d, 0x04, 0x04,
	// Entry 2C0 - 2FF
	0x04, 0x08, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
	0x08, 0x08, 0x08, 0x08, 0x06, 0x08, 0x08, 0x00,
	0x08, 0x08, 0x08, 0x00, 0x00, 0x04, 0x04, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x01,
} // Size: 778 bytes

var cardinalInclusionMasks = []uint64{ // 100 elements
	// Entry 0 - 1F
	0x0000000200500419, 0x0000000000512153, 0x000000000a327105, 0x0000000ca23c7101,
	0x00000004a23c7201, 0x0000000482943001, 0x0000001482943201, 0x0000000502943001,
	0x0000000502943001, 0x0000000522943201, 0x0000000540543401, 0x00000000454128e1,
	0x000000005b02e821, 0x000000006304e821, 0x000000006304ea21, 0x0000000042842821,
	0x0000000042842a21, 0x0000000042842821, 0x0000000042842821, 0x0000000062842a21,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000000400421, 0x0000000000400061,
	// Entry 20 - 3F
	0x000000000a004021, 0x0000000022004021, 0x0000000022004221, 0x0000000002800021,
	0x0000000002800221, 0x0000000002800021, 0x0000000002800021, 0x0000000022800221,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000000400421, 0x0000000000400061,
	0x000000000a004021, 0x0000000022004021, 0x0000000022004221, 0x0000000002800021,
	0x0000000002800221, 0x0000000002800021, 0x0000000002800021, 0x0000000022800221,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	// Entry 40 - 5F
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000040400421, 0x0000000044400061,
	0x000000005a004021, 0x0000000062004021, 0x0000000062004221, 0x0000000042800021,
	0x0000000042800221, 0x0000000042800021, 0x0000000042800021, 0x0000000062800221,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000040400421, 0x0000000044400061,
	0x000000005a004021, 0x0000000062004021, 0x0000000062004221, 0x0000000042800021,
	// Entry 60 - 7F
	0x0000000042800221, 0x0000000042800021, 0x0000000042800021, 0x0000000062800221,
} // Size: 824 bytes

// Slots used for cardinal: A6 of 0xFF rules; 24 of 0xFF indexes; 37 of 64 sets

// Total table size 3804 bytes (3KiB); checksum: FFC009FC
