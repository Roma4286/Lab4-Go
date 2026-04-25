module example.com/main

go 1.20

require (
	example.com/dog v0.0.0
	example.com/cat v0.0.0
)

replace example.com/dog => ../dog
replace example.com/cat => ../cat