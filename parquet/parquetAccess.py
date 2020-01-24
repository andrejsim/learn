import os
import pandas as pd

fp3 = 'dataset3.parquet'

pandas3 = pd.read_parquet(fp3)

print(pandas3.keys())
print(pandas3.index)

print(type(pandas3))