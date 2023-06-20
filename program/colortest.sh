#!/bin/bash

#* test file for multiple diff cases

echo Try passing as arguments "hello world" --color=red

go run . "hello world" --color=red

echo Try passing as arguments "1 + 1 = 2" --color=green

go run . "1 + 1 = 2" --color=green

echo Try passing as arguments "(%&) ??" --color=yellow

go run . "(%&) ??" --color=yellow

echo Try passing as arguments Try passing as arguments "HeY GuYs" --color=orange, in order to color GuYs

go run . "HeY GuYs" GuYs --color=orange

echo Try passing as arguments "RGB()" --color=blue, in order to color just the B

go run . "RGB()" B --color=blue
