import matplotlib.pyplot as plt
import numpy as np


def divide(arr):
    val = arr[0]
    for i in range(len(arr)):
        arr[i] /= val
        arr[i] = 1 / arr[i]
        arr[i] = round(arr[i], 10)


def plotSpeedup():
    th = [1, 2, 4, 8, 16, 32, 64, 128]
    ws = [67.27916, 34.69823, 18.48174, 9.70778, 5.56244, 3.60574, 2.70643, 2.73279]
    divide(ws)
    wb = [67.27916, 34.80481, 18.69410, 9.49527, 5.72763, 3.51495, 2.78069, 3.24832]
    divide(wb)

    default_x_ticks = range(len(th))

    xpoints = np.array(th)
    plt.title("SPEEDUP GRAPH")

    ypoints = np.array(ws)
    plt.plot(ypoints, marker="o", label="Work-Stealing")
    plt.xticks(default_x_ticks, xpoints)

    ypoints = np.array(wb)
    plt.plot(ypoints, marker="o", label="Work-Balancing", color="red")
    plt.xticks(default_x_ticks, xpoints)

    plt.legend()

    plt.xlabel("NUM OF THREADS")
    plt.ylabel("SPEEDUP")
    # plt.savefig("speedup-work-stealing.png")
    # plt.savefig("speedup-work-balancing.png")
    plt.savefig("speedup.png")


def plotBar():
    data = {"Sequntial": 67.27916, "Work-Stealing": 2.70643, "Work-Balancing": 2.78069}

    mode = list(data.keys())
    time = list(data.values())

    fig = plt.figure(figsize=(10, 5))

    # creating the bar plot
    plt.bar(mode, time, color="maroon", width=0.4)

    plt.ylabel("Time Taken (s)")
    plt.title("Num of Bodies = 20,000, Iterations = 10")
    plt.savefig("time-taken.png")


if __name__ == "__main__":
    # plotSpeedup()
    plotBar()
