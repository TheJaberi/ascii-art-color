#!/bin/bash

#* test file for multiple diff cases

echo Try passing as arguments "left standard --align=right"

go run . left standard --align=right

echo Try passing as arguments "right standard --align=left"

go run . right standard --align=left

echo Try passing as arguments "hello shadow --align=center"

go run . hello shadow --align=center

echo Try passing as arguments ""1 Two 4" shadow --align=justify"

go run . "1 Two 4" shadow --align=justify

echo Try passing as arguments "23/32 standard --align=right"

go run . "23/32" standard --align=right

echo Try passing as arguments "ABCabc123 thinkertoy --align=right"

go run . ABCabc123 thinkertoy --align=right

echo Try passing as arguments "#$%&" thinkertoy --align=center""

go run . "#$%&" thinkertoy --align=center

echo "Try passing as arguments ""23Hello World!" standard --align=left""

go run . "23Hello World!" standard --align=left

echo Try passing as arguments ""HELLO there HOW are YOU?!" thinkertoy --align=justify"

go run . "HELLO there HOW are YOU?!" thinkertoy --align=justify

echo "Try passing as arguments a -> A b -> B c -> C shadow --align=right"

go run . "a -> A b -> B c -> C" shadow --align=right
