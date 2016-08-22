package main

import ( 
	"fmt"
	"math"
	"math/rand"
	"time"
)

const size = 5
const traits = 5
const max_score = 15.0

var nations = []string { "german", "british", "swedish", "danish", "norsh" }
var color = []string { "red", "green", "white",  "yellow", "blue" }
var pet = []string { "dog", "bird", "cat", "horse", "fish" }
var drink = []string { "tea", "coffee", "milk", "beer", "water" }
var cigaret = []string { "pall mall", "dunhill", "winfield", "rothmanns", "marlboro" }

type Solution struct {
	nations []string	
	color []string	
	pet []string	
	drink []string	
	cigaret []string	
}

func random_permutation(src []string) (dest []string) {
	dest = make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		 dest[v] = src[i]
	}
	return dest
}

func analyze(s Solution) float64{
	var ctr float64 = 0

	for i := 0 ; i < size; i++ {
		left := i > 0
		right := i < (size -1 );
	
		if s.nations[i] == "british" && s.color[i] == "red" {
			ctr++
		}
		if s.nations[i] == "swedish" && s.pet[i] == "dog" {
			ctr++
		}
		if s.nations[i] == "danish" && s.drink[i] == "tea" {
			ctr++
		}
		if s.color[i] == "green" && right && s.color[i+1] == "white" {
			ctr++
		}
		if s.color[i] == "green" && s.drink[i] == "coffee" { 
			ctr++
		}
		if s.cigaret[i] == "pall mall" && s.pet[i] == "bird" { 
			ctr++
		}
		if i == 2  && s.drink[i] == "milk" { 
			ctr++
		}
		if s.color[i] == "yellow"  && s.cigaret[i] == "dunhill" { 
			ctr++
		}
		if i == 0  && s.nations[i] == "norsh" { 
			ctr++
		}
		if s.cigaret[i] == "marlboro" && ((left  && s.pet[i-1] == "cat") || (right &&  s.pet[i+1] == "cat")) { 
			ctr++
		}
		if s.pet[i] == "horse" && ((left  && s.cigaret[i-1] == "dunhill") || (right &&  s.cigaret[i+1] == "dunhill")) { 
			ctr++
		}
		if s.cigaret[i] == "winfield"  && s.drink[i] == "beer" { 
			ctr++
		}
		if s.nations[i] == "norsh"  && ((left  && s.color[i-1] == "blue") || (right &&  s.color[i+1] == "blue")) { 
			ctr++
		}
		if s.nations[i] == "german"  && s.cigaret[i] == "rothmanns" { 
			ctr++
		}
		if s.cigaret[i] == "marlboro" && ((left  && s.drink[i-1] == "water") || (right &&  s.drink[i+1] == "water")) { 
			ctr++
		}
	}
	return (max_score - ctr)
}

func mutate(s Solution) Solution {
	var result = Solution {}
	result.nations = make([]string, len(s.nations))
	result.color = make([]string, len(s.color))
	result.pet= make([]string, len(s.pet))
	result.drink= make([]string, len(s.drink))
	result.cigaret= make([]string, len(s.cigaret))
	
	copy(result.nations,	s.nations)
	copy(result.color, s.color)
	copy(result.pet, s.pet)
	copy(result.drink, s.drink)
	copy(result.cigaret, s.cigaret)
		
	var perm = rand.Perm(size)
	var idx1 = perm[0]
	var idx2 = perm[1]
	switch trait := rand.Intn(traits); trait {
		case 0:
			result.nations[idx1] = s.nations[idx2] 
			result.nations[idx2] = s.nations[idx1] 
		case 1:
			result.color[idx1] = s.color[idx2] 
			result.color[idx2] = s.color[idx1] 
		case 2:
			result.pet[idx1] = s.pet[idx2] 
			result.pet[idx2] = s.pet[idx1] 
		case 3:
			result.drink[idx1] = s.drink[idx2] 
			result.drink[idx2] = s.drink[idx1] 
		case 4:
			result.cigaret[idx1] = s.cigaret[idx2] 
			result.cigaret[idx2] = s.cigaret[idx1] 
	}
	return result;
}

func main() {
	s := Solution { 
		nations,
		color,
		pet,
		drink,
		cigaret, 
	}
	rand.Seed(time.Now().UTC().UnixNano())
	const delta = 0.1
	const k = 500	
	const reset_after = 100
	current_score := analyze(s)
	last_score := current_score
	resetCounter := reset_after
	t := 1.0
	for current_score > 0 {
		for i := 0 ; i < k; i++ {
			new_solution := mutate(s)
			new_score := analyze(new_solution)
			delta :=  new_score - current_score 
			if delta <= 0.0 || rand.Float64() < math.Exp(-delta / t) {
				s, current_score = new_solution, new_score
			} 
		}
		if t >= 2*delta{
			t = t - delta
		}
		if last_score != current_score {
			resetCounter = reset_after
			last_score = current_score
		} else {
			resetCounter--
			if(resetCounter == 0) {
				t = 1.0
				resetCounter = reset_after
			}
		}
	}
	fmt.Println(s)
}
