def to_list(file):
  f = open(file, 'rt', encoding = 'UTF-8')
  lines = f.readlines()
  return lines

w_punct = to_list('broadcast_w_punct.txt')

f = open('SubtTV_2017_01_03_pcm.list.trn', 'rt', encoding = 'UTF-8')
new_f = open('SubtTV_2017_01_03_pcm.list.punct.trn', 'wt', encoding = 'UTF-8')

while True:
  line = f.readline()
  if not line : break
  org_txt = line.split(' :: ')[1][:-1]
  first_cmp = org_txt.split(' ')[0]
  second_cmp = org_txt.split(' ')[-1]
  
  for cmp_line in w_punct:
    cmp_txt = cmp_line[:-1]
    start = cmp_txt.find(first_cmp)
    check = cmp_txt[start + len(first_cmp):].find(second_cmp)
    end = start + len(first_cmp) + check
    if (start != -1) and (check != -1):
      break
      
  print(cmptxt)