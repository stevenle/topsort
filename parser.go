// Copyright 2013 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package topsort

import "strings"

const EDGE_SEP = ">"
const PAIR_SEP = ","

// ParseString Parse nodes and egdes from string
func (g *Graph) ParseString(data, edgeSep, pairSep string) {
	if edgeSep == "" {
		edgeSep = EDGE_SEP
	}
	if pairSep == "" {
		pairSep = PAIR_SEP
	}
	var pairs [][2]string
	for _, pair := range strings.Split(data, pairSep) {
		pairt := strings.Split(pair, edgeSep)
		g.AddNode(pairt...)
		if len(pairt) == 2 {
			pairs = append(pairs, [2]string{pairt[0], pairt[1]})
		}
	}

	g.AddEdgeTuple(pairs...)
}

// ParseLines Parse nodes and egdes from lines
func (g *Graph) ParseLines(edgeSep, pairSep string, lineReader func()string) {
	if edgeSep == "" {
		edgeSep = EDGE_SEP
	}
	if pairSep == "" {
		pairSep = PAIR_SEP
	}
	var pairs [][2]string
	for line := lineReader(); line != ""; {
		for _, pair := range strings.Split(line, pairSep) {
			pairt := strings.Split(pair, edgeSep)
			g.AddNode(pairt...)
			if len(pairt) == 2 {
				pairs = append(pairs, [2]string{pairt[0], pairt[1]})
			}
		}
		line = lineReader()
	}
	g.AddEdgeTuple(pairs...)
}
