# SNAKE TERMINUX

![Snake Game Demo Gif](https://github.com/abhishekraj272/snake-terminux/blob/main/static/snake-demo.webp?raw=true)

## Running Locally

1. Fetch the repo
```bash
git clone https://github.com/abhishekraj272/snake-terminux.git
```

2. Download the dependencies
```bash
cd snake-terminux
go mod download
```

3. Starting the game
```bash
go run main.go
```

## Instructions

- Use arrow keys to move your snake.
- Press q, esc, Ctrl+C or Ctrl+D to quit the game.
- Hitting the wall or snake itself means GAME OVER.
- Eat food to increase score and length of snake.

## Implementation
- Chosen ```termbox-go``` for better UI/UX.
- Ran ```termbox.PollEvent``` on a GoRoutine to detect key press and used ```Go Channels``` to share data between GoRoutine.
- Created a Game object to store all required data, and created methods for all logics.
- Created a 2d Array to represent board where, 0 means blank-space, 1 means snake and 2 means food.
- Snake moves by moving its head coords and removing its tail.
- When snake eats food, its tail is increased.

## Dependency
1. [termbox-go](https://github.com/nsf/termbox-go)

## Tested On
- Ubuntu 20.04 LTS

> Time taken on this project 4-5hrs Max

## Refrences
- https://github.com/jjinux/gotetris