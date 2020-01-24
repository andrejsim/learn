import os
import pandas as pd
import xarray as xr
import pyarrow as pa
import pyarrow.parquet as pq
from pprint import pprint
import matplotlib.pyplot as plt


fp3 = 'dataset3.parquet'

pandas3 = pd.read_parquet(fp3)

print(pandas3.keys())
print(pandas3.index)

ds = pandas3.to_xarray()

print(ds)

ds['swell_1_of_1_max_slope'].plot()
#ds['swell_1_of_1_lm'].plot()

plt.show()


print(type(pandas3))