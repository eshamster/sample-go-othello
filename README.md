# Sample: Othello by Go

A sample Othello to learn Go.

## Installation

```bash
$ git clone https://github.com/eshamster/sample-go-othello.git
$ cd sample-go-othello
$ go build
```

## Usage

```bash
$ ./othello -player1 <player> -player2 <player>
```

Example:

```bash
$ ./othello -player1 human -player2 uct
```

In default, the following players are defined.

- `human`: It's you.
- `minimax`: It selects move according to minimax strategy (with alpha-beta pruning).
    - The search depth is 6.
- `random`: It randomly selects move from candidates in uniform probability.
- `mc`: It selects move according to simple Monte-Carlo method.
    - The simulation times is 10,000.
- `uct`: It selects move according using Monte-Carlo Tree Search with UCT.
    - The simulation times is 10,000.

They are defined in `DEF_PLAYER`.

## Author

eshamster (hamgoostar@gmail.com)

## Copyright

Copyright (c) 2018 eshamster (hamgoostar@gmail.com)

## License

Distributed under the MIT License
