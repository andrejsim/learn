import os
import xarray as xr
import pyarrow as pa
import pyarrow.parquet as pq


def size_mb(size):
    return size / (1024.0*1024.0)


def main():
    f1 = '../data/e024b429-3fb1-4a6d-b4e6-23fe5eaadfc5'
    f2 = '../data/468cd686-0b96-4296-92ff-45f46c73b90e'

    fp1 = 'dataset1.parquet'
    fp2 = 'dataset2.parquet'
    fp3 = 'dataset3.parquet'

    ds1 = xr.open_dataset(f1)
    pq.write_table(pa.Table.from_pandas(ds1.to_dataframe()), fp1)

    ds2 = xr.open_dataset(f2)
    pq.write_table(pa.Table.from_pandas(ds2.to_dataframe()), fp2)

    ds = ds1.merge(ds2)
    ds.to_netcdf("dataset3.nc")

    # dask required TODO...
    # with xr.open_mfdataset('../data/*') as ds:
    #     print(ds.keys())

    df = ds.to_dataframe()
    table = pa.Table.from_pandas(df)
    print(table.to_pandas())

    pq.write_table(table, fp3)

    for f in [f1, fp1, f2, fp2,  fp3]:
        print("{}           {} MB".format(f, size_mb(os.path.getsize(f))))

    print(fp3.info())

if __name__ == '__main__':
    main()


