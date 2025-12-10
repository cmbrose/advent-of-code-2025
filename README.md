# advent-of-code-2025

## Starting a new day or new part

To start a new day, run `./setup.sh <target>` with the day as the arg - e.g. `./setup.sh 2.1` or `./setup.sh 4.2`

This will create a new directory with initial content for test and real input and the code. 

If the `target` is part 1 (ends in `.1`) the code will be just empty boilerplate. For part 2 (ends in `.2`) the code is copied from part 1.

## Running the code

To run the code, use `./run.sh`. By default this will run with the most recently setup day (determined by `current`) on the real input (`input.txt`). The usage of `run.sh` is:

```bash
./run.sh [input] [--debug] [--target <target>]

# e.g.
./run.sh

./run.sh test # use test.txt as the input

./run.sh --debug # enable util.Debugf outputs in the code

./run.sh --target 3.2 # override the target set up `current`

./run.sh test --debug --target 3.2 # put it all together
```