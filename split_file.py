f = open('SubtTV_2017_01_03_pcm.list.trn', 'rt', encoding = 'UTF-8')
lines = f.readlines()
f1 = open('SubtTV_2017_01_03_pcm.list_1.trn', 'wt', encoding = 'UTF-8')
f2 = open('SubtTV_2017_01_03_pcm.list_2.trn', 'wt', encoding = 'UTF-8')
f3 = open('SubtTV_2017_01_03_pcm.list_3.trn', 'wt', encoding = 'UTF-8')
f4 = open('SubtTV_2017_01_03_pcm.list_4.trn', 'wt', encoding = 'UTF-8')
flist = [f1, f2, f3, f4]
i = 0
j = 0
for line in lines:
  flist[j].write(line)
  if i == 636567:
    j+=1
    i=0
  i+=1
