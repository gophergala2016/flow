# Flow

Flow is a tool that allows you to connect to a Peer-to-peer network of people
who are willing to donate spare CPU time, for you to execute expensive computations
that you could not afford by yourself. Because we believe in equality, when you
connect to Flow, you donate you spare CPU time so other people in the network
can use it as well.

The main inspiration for developing Flow are projects like
[World Community Grid](https://www.worldcommunitygrid.org/) or
[BOINC](https://boinc.berkeley.edu/). In scientific computing,
platforms like these allow you to compete with projects that need to be run in
supercomputers.

At this early stage, Flow is not able to compete with those projects, but if
you are sitting in Starbucks with your friends, it would be fun to be able to
distribute your computations.

## Installation

We suggest using the `gb` tool:

```
gb vendor restore
gb build
```

To run Flow, you use `bin/flow`.


## Interface

The Flow interface looks like

```

Flow v0.1.0

Press Ctrl+C twice to exit

flow>
```

You have the following commands:

- `inspect`: Gives you the peers Flow has detected automatically. See `eval`.
- `usage`: Returns your current CPU use (as a percentage).
- `eval [filename]`: Evaluates your code in an available machine in the network.
`eval` automatically detects peers and chooses one for you.

## Performing computations

Currently, we are using a Lisp implementation that is native to Go,
[zygomys](https://github.com/glycerine/zygomys), as the language you can run
using Flow. In almost every aspect it's normal Lisp code. Please refer to the
language for the subtle changes.

For example, you would have
```lisp
; mycode.zy
(+ 40 2) ; Or the answer to life the universe and everything
```

inside you file, and would run:

```
flow> eval mycode.zy

42 # the answer to my computation

Press Enter to continue...
```

## A little about the server you host

Other people using Flow send you a `"usage"` request to know what your CPU usage
is. If it's lower a fixed threshold, then Flow can pick you to execute their
code. It runs in the background and you would never notice it.


## Team [menteslibres.io](https://menteslibres.io)

- [Arturo Vergara](https://github.com/ArturoVM)
- [Eduardo Villaseñor](https://github.com/evalvarez12)
- [Leonardo Castro](https://github.com/LeonardoCastro)
- [Rodrigo Leal](https://github.com/rodrigolece)
