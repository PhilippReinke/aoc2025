# /// script
# dependencies = [
#   "scipy",
#   "numpy",
# ]
# ///

# Usage:
# cat in.txt | uv run day10.py

import sys
import numpy as np
from scipy.optimize import linprog


def main():
    try:
        data = sys.stdin.read()
    except EOFError:
        return

    if not data:
        return

    data_trimmed = data.rstrip("\n")
    lines = data_trimmed.split("\n")

    sol2 = 0
    for line in lines:
        if not line.strip():
            continue

        _, buttons, want_joltage = parse(line)

        c = [1 for _ in range(len(buttons))]
        b_eq = np.array(want_joltage)
        A_eq = np.array(
            [[int(j in button) for button in buttons] for j in range(len(want_joltage))]
        )

        res = linprog(c=c, A_eq=A_eq, b_eq=b_eq, integrality=1, method="highs")

        if res.success:
            sol2 += int(round(res.fun))

    print(f"sol2: {sol2}")


def parse(line):
    states = []
    buttons = []
    joltage = []

    segments = line.split(" ")
    for seg in segments:
        if not seg:
            continue

        if seg.startswith("["):
            for char in seg[1:-1]:
                states.append(char == "#")

        elif seg.startswith("("):
            nums = [int(n) for n in seg[1:-1].split(",") if n]
            buttons.append(nums)

        elif seg.startswith("{"):
            nums = [int(n) for n in seg[1:-1].split(",") if n]
            joltage.extend(nums)

        else:
            print(f"unknown seg start: {seg}")
            raise ValueError("failed to parse line")

    return states, buttons, joltage


if __name__ == "__main__":
    main()
