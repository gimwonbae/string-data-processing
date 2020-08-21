#!/bin/bash

#   directory tree

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

#   usage

# all files must UTF-8 encoding
# source file path don't contain ../sourceFileName just sourceFileName
# reference directory path don't contain ../ref_dir_path just ref_dir_path

# ./run.sh (source) (ref) (os) (outputname) 

#source file format :       SubtTV_Database/2017/03/JTBC/JTBC_2017_0101_0000/JTBC_2017_0101_0000_999_000.pcm :: blah blah

if [ $# -ne 4 ]; then
  echo "Wrong Usage"
  exit
fi

source=${1}
ref=${2}
os=${3}
output=${4}

sourceLine=$(cat ../${source} | wc -l)

if [ ${os} == "windows" ]; then
  #   for matching
  echo $(date) "matching start" ${output}
  ./num_punct_process_windows -goal matching -source ${source} -ref ${ref} -output ${output}
  #   for checking
  echo $(date) "checking start" ${output}
  ./num_punct_process_windows -goal checking -source ${source} -ref ${ref} -output ${output}
elif [ ${os} == "linux" ]; then
  echo $(date) "matching start" ${output}
  ./num_punct_process_linuxs -goal matching -source ${source} -ref ${ref} -output ${output}
  echo $(date) "checking start" ${output}
  ./num_punct_process_linuxs -goal checking -source ${source} -ref ${ref} -output ${output}
elif [ ${os} == "own" ]; then
  #   It needs golang env.
  echo "build golang exec file"
  go build -o num_punct_process num_punct_process.go 
  echo $(date) "matching start" ${output}
  ./num_punct_process -goal matching -source ${source} -ref ${ref} -output ${output}
  echo $(date) "checking start" ${output}
  ./num_punct_process -goal checking -source ${source} -ref ${ref} -output ${output}
else
  echo "Wrong Usage"
  exit
fi

#   for file merge
echo $(date) "file merge ${output}.."
cat ../${output}_match_w_num > ../${output}_match; cat ../${output}_match_wo_num >> ../${output}_match 

#   if you want miss file
#cat ../miss_w_num > ../miss; cat ../miss_wo_num >> ../miss

#   for file sort
echo $(date) "file sort ${output} .."
cat ../${output}_match | sort -k 1 > ../${output}_success

#   if you want sort miss file
#cat ../miss | sort -k 1 > ../miss_s

#   for removing log files
echo $(date) "remove log files ${output} .."
rm ../${output}_match* ../${output}_miss* ../${output}_fail ../${output}_w*

successLine=$(cat ../${output}_success | wc -l)
let per=successLine/sourceLine\*100

echo "source line :" ${sourceLine}
echo "success line :" ${successLine}
echo "Processing Rate :" ${per} "%"