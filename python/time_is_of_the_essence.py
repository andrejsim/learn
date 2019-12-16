import os
import time


def io_timer():
    timing = os.times()
    return timing.elapsed - (timing.system + timing.user)


def measure_time(f):
    def timed(*args, **kw):
        ts = time.time()
        result = f(*args, **kw)
        te = time.time()

        print("{} ({},{}) sec {}".format(f.__name__, args, kw, te-ts))

        return result
    return timed


if __name__ == "__main__":
    timing = os.times()
    print(timing.elapsed)
    print(timing.system)
    print(timing.user)

    #odd
    print(io_timer())

    #surround with timmer
    print(measure_time(io_timer()))