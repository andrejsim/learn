import os
import xarray as xr
import pyarrow as pa
import pyarrow.parquet as pq


def size_mb(size):
    return size / (1024.0*1024.0)
gi

def main():
    f1 = '../data/e024b429-3fb1-4a6d-b4e6-23fe5eaadfc5'
    f2 = '../data/468cd686-0b96-4296-92ff-45f46c73b90e'
    f3 = 'dataset.parquet'

    ds1 = xr.open_dataset(f1)
    print(ds1.keys())

    ds2 = xr.open_dataset(f2)
    print(ds2.keys())

    ds = ds1.merge(ds2)
    print(ds.info())

    # dask required TODO...
    # with xr.open_mfdataset('../data/*') as ds:
    #     print(ds.keys())

    df = ds.to_dataframe()
    table = pa.Table.from_pandas(df)
    print(table.to_pandas())

    pq.write_table(table, f3)
    for f in [f1, f2, f3]:
        print("{}           {} MB".format(f, size_mb(os.path.getsize(f))))


if __name__ == '__main__':
    main()


