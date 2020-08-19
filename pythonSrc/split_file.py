f = open('SubtTV_2017_01_03_pcm.list.trn', 'rt', encoding = 'UTF-8')
lines = f.readlines()
flist = []
for i in range(80):
  ff = open('SubtTV_2017_01_03_pcm.list_' + str(i+1) + '.trn', 'wt', encoding = 'UTF-8')
  flist.append(ff)

i = 0
j = 0
for line in lines:
  flist[j].write(line)
  if i == 31829:
    j+=1
    i=0
  i+=1