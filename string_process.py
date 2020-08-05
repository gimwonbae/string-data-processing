def to_list(file):
  f = open(file, 'rt', encoding = 'UTF-8')
  lines = f.readlines()
  return lines

w_punct = to_list('broadcast_w_punct.txt')

f = open('SubtTV_2017_01_03_pcm.list.trn', 'rt', encoding = 'UTF-8')
new_f = open('SubtTV_2017_01_03_pcm.list.punct.trn', 'wt', encoding = 'UTF-8')
not_found = open('SubtTV_2017_not_found', 'wt', encoding = 'UTF-8')

while True:
  line = f.readline()
  if not line : break

  file_name = line.split(' :: ')[0]
  org_txt = line.split(' :: ')[1][:-1]
  first_cmp = org_txt.split(' ')[0]
  second_cmp = org_txt.split(' ')[-1]

  flag = True

  for cmp_line in w_punct:
    cmp_txt = cmp_line[:-1]
    start = cmp_txt.find(first_cmp)
    check = cmp_txt.find(second_cmp, start + len(first_cmp))
    end = check + len(second_cmp)
    if (start != -1) and (check != -1):
      new_f.write(file_name + ' :: ' + cmp_txt[start:end] + '\n')
      flag = False

  if flag:
    not_found.write(line)