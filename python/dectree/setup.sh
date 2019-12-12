conda create -n ds python=3.6
conda env export  > environment.yml
conda install --file requirements.txt 
