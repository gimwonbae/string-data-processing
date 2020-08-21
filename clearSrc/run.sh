#!/bin/bash

# directory tree
#   Dir
#     -source file
#     -reference directory
#        -JTBC_2017_0101_0000
#           -data
#              -JTBC_2017_0101_0000.tok
#              -JTBC_2017_0101_0000.text
#     -SrcDir
#        -run.sh
#        -num_punct_process.go
#        -num_punct_process_windows
#        -num_punct_process_linux

# usage
# all files must UTF-8 encoding
# source file path don't contain ../sourceFileName just sourceFileName
# reference directory path don't contain ../ref_dir_path just ref_dir_path
# ./run.sh (source) (ref) (os) (outputname) 

if [ $# -ne 3 ]; then
  echo "Wrong Usage"
  exit
fi

source=${1}
ref=${2}
os=${3}

if [ ${os} == "windows" ]; then
  #   for matching
  echo $(date) "matching start .."
  ./num_punct_process_windows -goal matching -source ${source} -ref ${ref}
  #   for checking
  echo $(date) "checking start .."
  ./num_punct_process_windows -goal checking -source ${source} -ref ${ref}
elif [ ${os} == "linux" ]; then
  echo $(date) "matching start .."
  ./num_punct_process_linuxs -goal matching -source ${source} -ref ${ref}
  echo $(date) "checking start .."
  ./num_punct_process_linuxs -goal checking -source ${source} -ref ${ref}
elif [ ${os} == "own" ]; then
  #   It needs golang env.
  echo "build golang exec file"
  go build num_punct_process.go num_punct_process
  echo $(date) "matching start .."
  ./num_punct_process -goal matching -source ${source} -ref ${ref}
  echo $(date) "checking start .."
  ./num_punct_process -goal checking -source ${source} -ref ${ref}  
else
  echo "Wrong Usage"
  exit
fi

#   for file merge
echo $(date) "file merge .."
cat ../match_w_num > ../match; cat ../match_wo_num >> ../match 

#   if you want miss file
#cat ../miss_w_num > ../miss; cat ../miss_wo_num >> ../miss

#   for file sort
echo $(date) "file sort .."
cat ../match | sort -k 1 > ../success

#   if you want sort miss file
#cat ../miss | sort -k 1 > ../miss_s

#   for removing log files
echo $(date) "remove log files .."
rm ../match* ../miss* fail*